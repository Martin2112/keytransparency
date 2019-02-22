// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package entry

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/google/keytransparency/core/mutator"
	"github.com/google/tink/go/keyset"
	"github.com/google/tink/go/signature"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	pb "github.com/google/keytransparency/core/api/v1/keytransparency_go_proto"
	tinkpb "github.com/google/tink/proto/tink_go_proto"
)

// MapLogItemFn maps elements from *mutator.LogMessage to KV<index, *pb.EntryUpdate>.
func MapLogItemFn(m *mutator.LogMessage, emit func(index []byte, mutation *pb.EntryUpdate)) error {
	var entry pb.Entry
	if err := proto.Unmarshal(m.Mutation.Entry, &entry); err != nil {
		return err
	}
	emit(entry.Index, &pb.EntryUpdate{
		Mutation:  m.Mutation,
		Committed: m.ExtraData,
	})
	return nil
}

// IsValidEntry checks the internal consistency and correctness of a mutation.
func IsValidEntry(signedEntry *pb.SignedEntry) error {
	// Ensure that the mutation size is within bounds.
	if got, want := proto.Size(signedEntry), mutator.MaxMutationSize; got > want {
		glog.Warningf("mutation is %v bytes, want <= %v", got, want)
		return mutator.ErrSize
	}

	newEntry := pb.Entry{}
	if err := proto.Unmarshal(signedEntry.GetEntry(), &newEntry); err != nil {
		return fmt.Errorf("proto.Unmarshal(): %v", err)
	}

	ks := newEntry.GetAuthorizedKeys()
	return verifyKeys(ks, signedEntry.GetEntry(), signedEntry.GetSignatures())
}

// MutateFn verifies that newSignedEntry is a valid mutation for oldSignedEntry and returns the
// application of newSignedEntry to oldSignedEntry.
func MutateFn(oldSignedEntry, newSignedEntry *pb.SignedEntry) (*pb.SignedEntry, error) {
	if err := IsValidEntry(newSignedEntry); err != nil {
		return nil, err
	}

	newEntry := pb.Entry{}
	if err := proto.Unmarshal(newSignedEntry.GetEntry(), &newEntry); err != nil {
		return nil, fmt.Errorf("proto.Unmarshal(): %v", err)
	}
	oldEntry := pb.Entry{}
	// oldSignedEntry may be nil, resulting an empty oldEntry struct.
	if err := proto.Unmarshal(oldSignedEntry.GetEntry(), &oldEntry); err != nil {
		return nil, fmt.Errorf("proto.Unmarshal(): %v", err)
	}

	// Verify pointer to previous data.  The very first entry will have
	// oldSignedEntry=nil, so its hash is the Sha256 value of nil.
	prevEntryHash := sha256.Sum256(oldSignedEntry.GetEntry())
	if got, want := prevEntryHash[:], newEntry.GetPrevious(); !bytes.Equal(got, want) {
		// Check if this mutation is a replay.
		if bytes.Equal(oldSignedEntry.GetEntry(), newSignedEntry.Entry) {
			glog.Warningf("mutation is a replay of an old one")
			return nil, mutator.ErrReplay
		}
		glog.Warningf("previous entry hash: %x, want %x", got, want)
		return nil, mutator.ErrPreviousHash
	}

	if oldSignedEntry == nil {
		// Skip verificaion checks if there is no previous oldSignedEntry.
		return newSignedEntry, nil
	}

	ks := oldEntry.GetAuthorizedKeys()
	if err := verifyKeys(ks, newSignedEntry.Entry, newSignedEntry.GetSignatures()); err != nil {
		return nil, err
	}

	return newSignedEntry, nil
}

// verifyKeys verifies both old and new authorized keys based on the following
// criteria:
//   1. At least one signature with a key in the entry should exist.
//   2. Signatures with no matching keys are simply ignored.
func verifyKeys(ks *tinkpb.Keyset, data []byte, sigs [][]byte) error {
	if ks == nil {
		return errors.New("entry: nil keyset")
	}
	handle, err := keyset.NewHandleWithNoSecrets(ks)
	if err != nil {
		return fmt.Errorf("tink.KeysetHanldeWithNoSecret(new): %v", err)
	}

	verifier, err := signature.NewVerifier(handle)
	if err != nil {
		return fmt.Errorf("signature.NewVerifier(%v): %v", ks, err)
	}

	for _, sig := range sigs {
		if err := verifier.Verify(sig, data); err == nil {
			return nil
		}
	}
	return mutator.ErrUnauthorized
}
