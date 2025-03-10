// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sequencer_api.proto

// Key Transparency Sequencer
//
// The Key Transparency Sequencer API supplies an api for applying mutations to the current
// state of the map.

package sequencer_go_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MapMetadata struct {
	// sources is a list of log sources that were used to construct this map revision.
	Sources              []*MapMetadata_SourceSlice `protobuf:"bytes,2,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MapMetadata) Reset()         { *m = MapMetadata{} }
func (m *MapMetadata) String() string { return proto.CompactTextString(m) }
func (*MapMetadata) ProtoMessage()    {}
func (*MapMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{0}
}

func (m *MapMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapMetadata.Unmarshal(m, b)
}
func (m *MapMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapMetadata.Marshal(b, m, deterministic)
}
func (m *MapMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapMetadata.Merge(m, src)
}
func (m *MapMetadata) XXX_Size() int {
	return xxx_messageInfo_MapMetadata.Size(m)
}
func (m *MapMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MapMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MapMetadata proto.InternalMessageInfo

func (m *MapMetadata) GetSources() []*MapMetadata_SourceSlice {
	if m != nil {
		return m.Sources
	}
	return nil
}

// SourceSlice is the range of inputs that have been included in a map
// revision.
type MapMetadata_SourceSlice struct {
	// lowest_inclusive is the lowest primary key (inclusive) of the source
	// log that has been incorporated into this map revision. The primary
	// keys of logged items MUST be monotonically increasing.
	LowestInclusive int64 `protobuf:"varint,1,opt,name=lowest_inclusive,json=lowestInclusive,proto3" json:"lowest_inclusive,omitempty"`
	// highest_exclusive is the highest primary key (exclusive) of the source
	// log that has been incorporated into this map revision. The primary keys
	// of logged items MUST be monotonically increasing.
	HighestExclusive int64 `protobuf:"varint,2,opt,name=highest_exclusive,json=highestExclusive,proto3" json:"highest_exclusive,omitempty"`
	// log_id is the ID of the source log.
	LogId                int64    `protobuf:"varint,3,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapMetadata_SourceSlice) Reset()         { *m = MapMetadata_SourceSlice{} }
func (m *MapMetadata_SourceSlice) String() string { return proto.CompactTextString(m) }
func (*MapMetadata_SourceSlice) ProtoMessage()    {}
func (*MapMetadata_SourceSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{0, 0}
}

func (m *MapMetadata_SourceSlice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapMetadata_SourceSlice.Unmarshal(m, b)
}
func (m *MapMetadata_SourceSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapMetadata_SourceSlice.Marshal(b, m, deterministic)
}
func (m *MapMetadata_SourceSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapMetadata_SourceSlice.Merge(m, src)
}
func (m *MapMetadata_SourceSlice) XXX_Size() int {
	return xxx_messageInfo_MapMetadata_SourceSlice.Size(m)
}
func (m *MapMetadata_SourceSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_MapMetadata_SourceSlice.DiscardUnknown(m)
}

var xxx_messageInfo_MapMetadata_SourceSlice proto.InternalMessageInfo

func (m *MapMetadata_SourceSlice) GetLowestInclusive() int64 {
	if m != nil {
		return m.LowestInclusive
	}
	return 0
}

func (m *MapMetadata_SourceSlice) GetHighestExclusive() int64 {
	if m != nil {
		return m.HighestExclusive
	}
	return 0
}

func (m *MapMetadata_SourceSlice) GetLogId() int64 {
	if m != nil {
		return m.LogId
	}
	return 0
}

// RunBatchRequest triggers the sequencing of a batch of mutations for a
// directory, with the batch size governed by the request parameters.
type RunBatchRequest struct {
	// directory_id is the directory to run for.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// min_batch is the minimum number of items in a batch.
	// If less than min_batch items are available, nothing happens.
	// TODO(#1047): Replace with timeout so items in the log get processed
	// eventually.
	MinBatch int32 `protobuf:"varint,2,opt,name=min_batch,json=minBatch,proto3" json:"min_batch,omitempty"`
	// max_batch is the maximum number of items in a batch.
	MaxBatch int32 `protobuf:"varint,3,opt,name=max_batch,json=maxBatch,proto3" json:"max_batch,omitempty"`
	// block until a Signed Log Root has been published which encompases all map roots.
	Block                bool     `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunBatchRequest) Reset()         { *m = RunBatchRequest{} }
func (m *RunBatchRequest) String() string { return proto.CompactTextString(m) }
func (*RunBatchRequest) ProtoMessage()    {}
func (*RunBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{1}
}

func (m *RunBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunBatchRequest.Unmarshal(m, b)
}
func (m *RunBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunBatchRequest.Marshal(b, m, deterministic)
}
func (m *RunBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunBatchRequest.Merge(m, src)
}
func (m *RunBatchRequest) XXX_Size() int {
	return xxx_messageInfo_RunBatchRequest.Size(m)
}
func (m *RunBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunBatchRequest proto.InternalMessageInfo

func (m *RunBatchRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *RunBatchRequest) GetMinBatch() int32 {
	if m != nil {
		return m.MinBatch
	}
	return 0
}

func (m *RunBatchRequest) GetMaxBatch() int32 {
	if m != nil {
		return m.MaxBatch
	}
	return 0
}

func (m *RunBatchRequest) GetBlock() bool {
	if m != nil {
		return m.Block
	}
	return false
}

// DefineRevisionRequest contains information needed to define a new revision.
type DefineRevisionsRequest struct {
	// directory_id is the directory to examine the outstanding mutations for.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// min_batch is the minimum number of items in a batch.
	// If less than min_batch items are available, nothing happens.
	// TODO(#1047): Replace with timeout so items in the log get processed
	// eventually.
	MinBatch int32 `protobuf:"varint,2,opt,name=min_batch,json=minBatch,proto3" json:"min_batch,omitempty"`
	// max_batch is the maximum number of items in a batch.
	MaxBatch             int32    `protobuf:"varint,3,opt,name=max_batch,json=maxBatch,proto3" json:"max_batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefineRevisionsRequest) Reset()         { *m = DefineRevisionsRequest{} }
func (m *DefineRevisionsRequest) String() string { return proto.CompactTextString(m) }
func (*DefineRevisionsRequest) ProtoMessage()    {}
func (*DefineRevisionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{2}
}

func (m *DefineRevisionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefineRevisionsRequest.Unmarshal(m, b)
}
func (m *DefineRevisionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefineRevisionsRequest.Marshal(b, m, deterministic)
}
func (m *DefineRevisionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefineRevisionsRequest.Merge(m, src)
}
func (m *DefineRevisionsRequest) XXX_Size() int {
	return xxx_messageInfo_DefineRevisionsRequest.Size(m)
}
func (m *DefineRevisionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DefineRevisionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DefineRevisionsRequest proto.InternalMessageInfo

func (m *DefineRevisionsRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *DefineRevisionsRequest) GetMinBatch() int32 {
	if m != nil {
		return m.MinBatch
	}
	return 0
}

func (m *DefineRevisionsRequest) GetMaxBatch() int32 {
	if m != nil {
		return m.MaxBatch
	}
	return 0
}

// DefineRevisionResponse contains information about freshly defined revisions.
type DefineRevisionsResponse struct {
	// outsanding_revisions a list of all the defined revisions which are not yet applied.
	OutstandingRevisions []int64  `protobuf:"varint,1,rep,packed,name=outstanding_revisions,json=outstandingRevisions,proto3" json:"outstanding_revisions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefineRevisionsResponse) Reset()         { *m = DefineRevisionsResponse{} }
func (m *DefineRevisionsResponse) String() string { return proto.CompactTextString(m) }
func (*DefineRevisionsResponse) ProtoMessage()    {}
func (*DefineRevisionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{3}
}

func (m *DefineRevisionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefineRevisionsResponse.Unmarshal(m, b)
}
func (m *DefineRevisionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefineRevisionsResponse.Marshal(b, m, deterministic)
}
func (m *DefineRevisionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefineRevisionsResponse.Merge(m, src)
}
func (m *DefineRevisionsResponse) XXX_Size() int {
	return xxx_messageInfo_DefineRevisionsResponse.Size(m)
}
func (m *DefineRevisionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DefineRevisionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DefineRevisionsResponse proto.InternalMessageInfo

func (m *DefineRevisionsResponse) GetOutstandingRevisions() []int64 {
	if m != nil {
		return m.OutstandingRevisions
	}
	return nil
}

// ApplyRevisionRequest contains information needed to create a new revision.
type ApplyRevisionRequest struct {
	// directory_id is the directory to apply the mutations to.
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// revision is the expected revision of the new revision.
	Revision             int64    `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRevisionRequest) Reset()         { *m = ApplyRevisionRequest{} }
func (m *ApplyRevisionRequest) String() string { return proto.CompactTextString(m) }
func (*ApplyRevisionRequest) ProtoMessage()    {}
func (*ApplyRevisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{4}
}

func (m *ApplyRevisionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRevisionRequest.Unmarshal(m, b)
}
func (m *ApplyRevisionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRevisionRequest.Marshal(b, m, deterministic)
}
func (m *ApplyRevisionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRevisionRequest.Merge(m, src)
}
func (m *ApplyRevisionRequest) XXX_Size() int {
	return xxx_messageInfo_ApplyRevisionRequest.Size(m)
}
func (m *ApplyRevisionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRevisionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRevisionRequest proto.InternalMessageInfo

func (m *ApplyRevisionRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *ApplyRevisionRequest) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

// ApplyRevisionResponse contains stats about the created revision.
type ApplyRevisionResponse struct {
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// The revision this is for.
	Revision int64 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// mutations processed.
	Mutations int64 `protobuf:"varint,3,opt,name=mutations,proto3" json:"mutations,omitempty"`
	// map_leaves written.
	MapLeaves            int64    `protobuf:"varint,4,opt,name=map_leaves,json=mapLeaves,proto3" json:"map_leaves,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRevisionResponse) Reset()         { *m = ApplyRevisionResponse{} }
func (m *ApplyRevisionResponse) String() string { return proto.CompactTextString(m) }
func (*ApplyRevisionResponse) ProtoMessage()    {}
func (*ApplyRevisionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{5}
}

func (m *ApplyRevisionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRevisionResponse.Unmarshal(m, b)
}
func (m *ApplyRevisionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRevisionResponse.Marshal(b, m, deterministic)
}
func (m *ApplyRevisionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRevisionResponse.Merge(m, src)
}
func (m *ApplyRevisionResponse) XXX_Size() int {
	return xxx_messageInfo_ApplyRevisionResponse.Size(m)
}
func (m *ApplyRevisionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRevisionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRevisionResponse proto.InternalMessageInfo

func (m *ApplyRevisionResponse) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *ApplyRevisionResponse) GetRevision() int64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ApplyRevisionResponse) GetMutations() int64 {
	if m != nil {
		return m.Mutations
	}
	return 0
}

func (m *ApplyRevisionResponse) GetMapLeaves() int64 {
	if m != nil {
		return m.MapLeaves
	}
	return 0
}

// PublishRevisionsRequest copies all available SignedMapRoots into the Log of SignedMapRoots.
type PublishRevisionsRequest struct {
	DirectoryId string `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	// block until a Signed Log Root has been published which encompases all map roots.
	Block                bool     `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRevisionsRequest) Reset()         { *m = PublishRevisionsRequest{} }
func (m *PublishRevisionsRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRevisionsRequest) ProtoMessage()    {}
func (*PublishRevisionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{6}
}

func (m *PublishRevisionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRevisionsRequest.Unmarshal(m, b)
}
func (m *PublishRevisionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRevisionsRequest.Marshal(b, m, deterministic)
}
func (m *PublishRevisionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRevisionsRequest.Merge(m, src)
}
func (m *PublishRevisionsRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRevisionsRequest.Size(m)
}
func (m *PublishRevisionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRevisionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRevisionsRequest proto.InternalMessageInfo

func (m *PublishRevisionsRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *PublishRevisionsRequest) GetBlock() bool {
	if m != nil {
		return m.Block
	}
	return false
}

// PublishRevisionsResponse contains metrics about the publishing operation.
type PublishRevisionsResponse struct {
	// revisions published to the log of signed map roots.
	Revisions            []int64  `protobuf:"varint,1,rep,packed,name=revisions,proto3" json:"revisions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRevisionsResponse) Reset()         { *m = PublishRevisionsResponse{} }
func (m *PublishRevisionsResponse) String() string { return proto.CompactTextString(m) }
func (*PublishRevisionsResponse) ProtoMessage()    {}
func (*PublishRevisionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{7}
}

func (m *PublishRevisionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRevisionsResponse.Unmarshal(m, b)
}
func (m *PublishRevisionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRevisionsResponse.Marshal(b, m, deterministic)
}
func (m *PublishRevisionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRevisionsResponse.Merge(m, src)
}
func (m *PublishRevisionsResponse) XXX_Size() int {
	return xxx_messageInfo_PublishRevisionsResponse.Size(m)
}
func (m *PublishRevisionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRevisionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRevisionsResponse proto.InternalMessageInfo

func (m *PublishRevisionsResponse) GetRevisions() []int64 {
	if m != nil {
		return m.Revisions
	}
	return nil
}

// UpdateMetricsRequest is empty.
type UpdateMetricsRequest struct {
	DirectoryId          string   `protobuf:"bytes,1,opt,name=directory_id,json=directoryId,proto3" json:"directory_id,omitempty"`
	MaxUnappliedCount    int32    `protobuf:"varint,2,opt,name=max_unapplied_count,json=maxUnappliedCount,proto3" json:"max_unapplied_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMetricsRequest) Reset()         { *m = UpdateMetricsRequest{} }
func (m *UpdateMetricsRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMetricsRequest) ProtoMessage()    {}
func (*UpdateMetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{8}
}

func (m *UpdateMetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMetricsRequest.Unmarshal(m, b)
}
func (m *UpdateMetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMetricsRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMetricsRequest.Merge(m, src)
}
func (m *UpdateMetricsRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMetricsRequest.Size(m)
}
func (m *UpdateMetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMetricsRequest proto.InternalMessageInfo

func (m *UpdateMetricsRequest) GetDirectoryId() string {
	if m != nil {
		return m.DirectoryId
	}
	return ""
}

func (m *UpdateMetricsRequest) GetMaxUnappliedCount() int32 {
	if m != nil {
		return m.MaxUnappliedCount
	}
	return 0
}

// UpdateMetricsResponse is empty.
type UpdateMetricsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMetricsResponse) Reset()         { *m = UpdateMetricsResponse{} }
func (m *UpdateMetricsResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateMetricsResponse) ProtoMessage()    {}
func (*UpdateMetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5d61b2e27141ee, []int{9}
}

func (m *UpdateMetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMetricsResponse.Unmarshal(m, b)
}
func (m *UpdateMetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMetricsResponse.Marshal(b, m, deterministic)
}
func (m *UpdateMetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMetricsResponse.Merge(m, src)
}
func (m *UpdateMetricsResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateMetricsResponse.Size(m)
}
func (m *UpdateMetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMetricsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MapMetadata)(nil), "google.keytransparency.sequencer.MapMetadata")
	proto.RegisterType((*MapMetadata_SourceSlice)(nil), "google.keytransparency.sequencer.MapMetadata.SourceSlice")
	proto.RegisterType((*RunBatchRequest)(nil), "google.keytransparency.sequencer.RunBatchRequest")
	proto.RegisterType((*DefineRevisionsRequest)(nil), "google.keytransparency.sequencer.DefineRevisionsRequest")
	proto.RegisterType((*DefineRevisionsResponse)(nil), "google.keytransparency.sequencer.DefineRevisionsResponse")
	proto.RegisterType((*ApplyRevisionRequest)(nil), "google.keytransparency.sequencer.ApplyRevisionRequest")
	proto.RegisterType((*ApplyRevisionResponse)(nil), "google.keytransparency.sequencer.ApplyRevisionResponse")
	proto.RegisterType((*PublishRevisionsRequest)(nil), "google.keytransparency.sequencer.PublishRevisionsRequest")
	proto.RegisterType((*PublishRevisionsResponse)(nil), "google.keytransparency.sequencer.PublishRevisionsResponse")
	proto.RegisterType((*UpdateMetricsRequest)(nil), "google.keytransparency.sequencer.UpdateMetricsRequest")
	proto.RegisterType((*UpdateMetricsResponse)(nil), "google.keytransparency.sequencer.UpdateMetricsResponse")
}

func init() { proto.RegisterFile("sequencer_api.proto", fileDescriptor_0a5d61b2e27141ee) }

var fileDescriptor_0a5d61b2e27141ee = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0xa6, 0x24, 0x13, 0x50, 0xd3, 0x6d, 0xd2, 0x5a, 0x6e, 0x91, 0x82, 0x4f, 0x41,
	0x48, 0x8e, 0x68, 0x25, 0x68, 0xb9, 0xd1, 0xd2, 0x43, 0x81, 0x22, 0xe4, 0x90, 0x0b, 0x17, 0x6b,
	0x63, 0x6f, 0x9d, 0x55, 0xed, 0xdd, 0xc5, 0xbb, 0x2e, 0x89, 0xc4, 0x81, 0x03, 0x42, 0x42, 0xe2,
	0xc6, 0x7f, 0xe4, 0x77, 0x20, 0x7f, 0xa5, 0xa9, 0x1b, 0x94, 0xb6, 0x48, 0x9c, 0x5a, 0xcf, 0x9b,
	0xf7, 0x66, 0x76, 0xe7, 0xed, 0x04, 0x36, 0x24, 0xf9, 0x14, 0x13, 0xe6, 0x92, 0xc8, 0xc1, 0x82,
	0x5a, 0x22, 0xe2, 0x8a, 0xa3, 0xae, 0xcf, 0xb9, 0x1f, 0x10, 0xeb, 0x9c, 0x4c, 0x55, 0x84, 0x99,
	0x14, 0x38, 0x22, 0xcc, 0x9d, 0x5a, 0xb3, 0x5c, 0x63, 0x3b, 0xcb, 0xe8, 0xa7, 0xf9, 0xa3, 0xf8,
	0xac, 0x4f, 0x42, 0xa1, 0xa6, 0x19, 0xdd, 0xfc, 0xad, 0x41, 0xf3, 0x14, 0x8b, 0x53, 0xa2, 0xb0,
	0x87, 0x15, 0x46, 0x03, 0xb8, 0x27, 0x79, 0x1c, 0xb9, 0x44, 0xea, 0x95, 0x6e, 0xb5, 0xd7, 0xdc,
	0x3d, 0xb0, 0x96, 0x15, 0xb0, 0xe6, 0xf8, 0xd6, 0x20, 0x25, 0x0f, 0x02, 0xea, 0x12, 0xbb, 0x50,
	0x32, 0xbe, 0x40, 0x73, 0x2e, 0x8e, 0x1e, 0x43, 0x2b, 0xe0, 0x9f, 0x89, 0x54, 0x0e, 0x65, 0x6e,
	0x10, 0x4b, 0x7a, 0x41, 0x74, 0xad, 0xab, 0xf5, 0xaa, 0xf6, 0x5a, 0x16, 0x3f, 0x29, 0xc2, 0xe8,
	0x09, 0xac, 0x8f, 0xa9, 0x3f, 0x4e, 0x72, 0xc9, 0xa4, 0xc8, 0xad, 0xa4, 0xb9, 0xad, 0x1c, 0x38,
	0x2e, 0xe2, 0xa8, 0x03, 0xab, 0x01, 0xf7, 0x1d, 0xea, 0xe9, 0xd5, 0x34, 0xa3, 0x16, 0x70, 0xff,
	0xc4, 0x7b, 0xbd, 0x52, 0xd7, 0x5a, 0x15, 0xf3, 0x9b, 0x06, 0x6b, 0x76, 0xcc, 0x0e, 0xb1, 0x72,
	0xc7, 0x76, 0xd2, 0xba, 0x54, 0xe8, 0x11, 0xdc, 0xf7, 0x68, 0x44, 0x5c, 0xc5, 0xa3, 0x69, 0x42,
	0x4b, 0x9a, 0x68, 0xd8, 0xcd, 0x59, 0xec, 0xc4, 0x43, 0xdb, 0xd0, 0x08, 0x29, 0x73, 0x46, 0x09,
	0x2d, 0x2d, 0x5c, 0xb3, 0xeb, 0x21, 0xcd, 0x64, 0x52, 0x10, 0x4f, 0x72, 0xb0, 0x9a, 0x83, 0x78,
	0x92, 0x81, 0x6d, 0xa8, 0x8d, 0x02, 0xee, 0x9e, 0xeb, 0x2b, 0x5d, 0xad, 0x57, 0xb7, 0xb3, 0x0f,
	0x33, 0x86, 0xcd, 0x57, 0xe4, 0x8c, 0x32, 0x62, 0x93, 0x0b, 0x2a, 0x29, 0x67, 0xf2, 0x7f, 0x34,
	0x63, 0xbe, 0x83, 0xad, 0x6b, 0x65, 0xa5, 0xe0, 0x4c, 0x12, 0xb4, 0x07, 0x1d, 0x1e, 0x2b, 0xa9,
	0x30, 0xf3, 0x28, 0xf3, 0x9d, 0xa8, 0x48, 0xd0, 0xb5, 0x6e, 0xb5, 0x57, 0xb5, 0xdb, 0x73, 0xe0,
	0x8c, 0x6c, 0x0e, 0xa1, 0xfd, 0x52, 0x88, 0x60, 0x5a, 0x44, 0x6e, 0x71, 0x08, 0x03, 0xea, 0x45,
	0x8d, 0x7c, 0x92, 0xb3, 0x6f, 0xf3, 0x97, 0x06, 0x9d, 0x92, 0x6e, 0xde, 0xe5, 0xbf, 0x09, 0xa3,
	0x1d, 0x68, 0x84, 0xb1, 0xc2, 0x2a, 0x3d, 0x58, 0xe6, 0x8e, 0xcb, 0x00, 0x7a, 0x08, 0x10, 0x62,
	0xe1, 0x04, 0x04, 0x5f, 0x10, 0x99, 0xce, 0x2b, 0x81, 0xb1, 0x78, 0x9b, 0x06, 0x4c, 0x1b, 0xb6,
	0xde, 0xc7, 0xa3, 0x80, 0xca, 0xf1, 0x5d, 0x86, 0x36, 0xf3, 0x41, 0x65, 0xde, 0x07, 0xfb, 0xa0,
	0x5f, 0xd7, 0xcc, 0xcf, 0xba, 0x03, 0x8d, 0xf2, 0x14, 0x2e, 0x03, 0x26, 0x85, 0xf6, 0x50, 0x78,
	0x58, 0x91, 0x53, 0xa2, 0x22, 0xea, 0xde, 0xa6, 0x15, 0x0b, 0x36, 0x12, 0x8b, 0xc4, 0x0c, 0x0b,
	0x11, 0x50, 0xe2, 0x39, 0x2e, 0x8f, 0x99, 0xca, 0x9d, 0xb4, 0x1e, 0xe2, 0xc9, 0xb0, 0x40, 0x8e,
	0x12, 0xc0, 0xdc, 0x82, 0x4e, 0xa9, 0x54, 0xd6, 0xe1, 0xee, 0xcf, 0x1a, 0xe8, 0x6f, 0xc8, 0xf4,
	0xc3, 0xdc, 0x3e, 0x18, 0x14, 0xeb, 0x00, 0x0d, 0xa1, 0x5e, 0x3c, 0x34, 0xf4, 0x74, 0xf9, 0xf6,
	0x28, 0x3d, 0x4a, 0x63, 0xb3, 0xa0, 0x14, 0xfb, 0xca, 0x3a, 0x4e, 0xf6, 0x15, 0xfa, 0xae, 0xc1,
	0x5a, 0xc9, 0xc3, 0x68, 0x7f, 0xb9, 0xfc, 0xe2, 0xd7, 0x66, 0x1c, 0xdc, 0x81, 0x99, 0x8f, 0xe7,
	0xab, 0x06, 0x0f, 0xae, 0x98, 0x14, 0x3d, 0x5b, 0x2e, 0xb6, 0xe8, 0xb5, 0x18, 0xcf, 0x6f, 0xcd,
	0xcb, 0x5b, 0xf8, 0xa1, 0x41, 0xab, 0x6c, 0x1f, 0x74, 0x83, 0x23, 0xfd, 0xc5, 0xc6, 0xc6, 0x8b,
	0xbb, 0x50, 0xe7, 0xae, 0xe3, 0x8a, 0x4b, 0x6e, 0x72, 0x1d, 0x8b, 0x1c, 0x7c, 0x93, 0xeb, 0x58,
	0x68, 0xc7, 0xc3, 0xe3, 0x8f, 0x47, 0x3e, 0x55, 0xe3, 0x78, 0x64, 0xb9, 0x3c, 0xec, 0xe7, 0x3f,
	0x77, 0x25, 0x91, 0xbe, 0xcb, 0x23, 0xd2, 0x9f, 0x29, 0x5d, 0xfe, 0xe7, 0xf8, 0xdc, 0xc9, 0xac,
	0xb6, 0x9a, 0xfe, 0xd9, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x00, 0x12, 0xd9, 0x68, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeyTransparencySequencerClient is the client API for KeyTransparencySequencer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyTransparencySequencerClient interface {
	// RunBatch calls DefineRevisions, ApplyRevision, and PublishRevisions successively.
	RunBatch(ctx context.Context, in *RunBatchRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DefineRevision examines the outstanding items in the queue and optionally
	// writes the metadata for one or more revisions to the metadata database.
	DefineRevisions(ctx context.Context, in *DefineRevisionsRequest, opts ...grpc.CallOption) (*DefineRevisionsResponse, error)
	// ApplyRevision applies the contained mutations to the current map root.
	// If this method fails, it must be retried with the same arguments.
	ApplyRevision(ctx context.Context, in *ApplyRevisionRequest, opts ...grpc.CallOption) (*ApplyRevisionResponse, error)
	// PublishRevisions copies the MapRoots of all known map revisions into the Log
	// of MapRoots.
	PublishRevisions(ctx context.Context, in *PublishRevisionsRequest, opts ...grpc.CallOption) (*PublishRevisionsResponse, error)
	// UpdateMetrics will update various counters on the server. Call periodically.
	UpdateMetrics(ctx context.Context, in *UpdateMetricsRequest, opts ...grpc.CallOption) (*UpdateMetricsResponse, error)
}

type keyTransparencySequencerClient struct {
	cc *grpc.ClientConn
}

func NewKeyTransparencySequencerClient(cc *grpc.ClientConn) KeyTransparencySequencerClient {
	return &keyTransparencySequencerClient{cc}
}

func (c *keyTransparencySequencerClient) RunBatch(ctx context.Context, in *RunBatchRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.keytransparency.sequencer.KeyTransparencySequencer/RunBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencySequencerClient) DefineRevisions(ctx context.Context, in *DefineRevisionsRequest, opts ...grpc.CallOption) (*DefineRevisionsResponse, error) {
	out := new(DefineRevisionsResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.sequencer.KeyTransparencySequencer/DefineRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencySequencerClient) ApplyRevision(ctx context.Context, in *ApplyRevisionRequest, opts ...grpc.CallOption) (*ApplyRevisionResponse, error) {
	out := new(ApplyRevisionResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.sequencer.KeyTransparencySequencer/ApplyRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencySequencerClient) PublishRevisions(ctx context.Context, in *PublishRevisionsRequest, opts ...grpc.CallOption) (*PublishRevisionsResponse, error) {
	out := new(PublishRevisionsResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.sequencer.KeyTransparencySequencer/PublishRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyTransparencySequencerClient) UpdateMetrics(ctx context.Context, in *UpdateMetricsRequest, opts ...grpc.CallOption) (*UpdateMetricsResponse, error) {
	out := new(UpdateMetricsResponse)
	err := c.cc.Invoke(ctx, "/google.keytransparency.sequencer.KeyTransparencySequencer/UpdateMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyTransparencySequencerServer is the server API for KeyTransparencySequencer service.
type KeyTransparencySequencerServer interface {
	// RunBatch calls DefineRevisions, ApplyRevision, and PublishRevisions successively.
	RunBatch(context.Context, *RunBatchRequest) (*empty.Empty, error)
	// DefineRevision examines the outstanding items in the queue and optionally
	// writes the metadata for one or more revisions to the metadata database.
	DefineRevisions(context.Context, *DefineRevisionsRequest) (*DefineRevisionsResponse, error)
	// ApplyRevision applies the contained mutations to the current map root.
	// If this method fails, it must be retried with the same arguments.
	ApplyRevision(context.Context, *ApplyRevisionRequest) (*ApplyRevisionResponse, error)
	// PublishRevisions copies the MapRoots of all known map revisions into the Log
	// of MapRoots.
	PublishRevisions(context.Context, *PublishRevisionsRequest) (*PublishRevisionsResponse, error)
	// UpdateMetrics will update various counters on the server. Call periodically.
	UpdateMetrics(context.Context, *UpdateMetricsRequest) (*UpdateMetricsResponse, error)
}

// UnimplementedKeyTransparencySequencerServer can be embedded to have forward compatible implementations.
type UnimplementedKeyTransparencySequencerServer struct {
}

func (*UnimplementedKeyTransparencySequencerServer) RunBatch(ctx context.Context, req *RunBatchRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunBatch not implemented")
}
func (*UnimplementedKeyTransparencySequencerServer) DefineRevisions(ctx context.Context, req *DefineRevisionsRequest) (*DefineRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefineRevisions not implemented")
}
func (*UnimplementedKeyTransparencySequencerServer) ApplyRevision(ctx context.Context, req *ApplyRevisionRequest) (*ApplyRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRevision not implemented")
}
func (*UnimplementedKeyTransparencySequencerServer) PublishRevisions(ctx context.Context, req *PublishRevisionsRequest) (*PublishRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishRevisions not implemented")
}
func (*UnimplementedKeyTransparencySequencerServer) UpdateMetrics(ctx context.Context, req *UpdateMetricsRequest) (*UpdateMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetrics not implemented")
}

func RegisterKeyTransparencySequencerServer(s *grpc.Server, srv KeyTransparencySequencerServer) {
	s.RegisterService(&_KeyTransparencySequencer_serviceDesc, srv)
}

func _KeyTransparencySequencer_RunBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencySequencerServer).RunBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.sequencer.KeyTransparencySequencer/RunBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencySequencerServer).RunBatch(ctx, req.(*RunBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencySequencer_DefineRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefineRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencySequencerServer).DefineRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.sequencer.KeyTransparencySequencer/DefineRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencySequencerServer).DefineRevisions(ctx, req.(*DefineRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencySequencer_ApplyRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencySequencerServer).ApplyRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.sequencer.KeyTransparencySequencer/ApplyRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencySequencerServer).ApplyRevision(ctx, req.(*ApplyRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencySequencer_PublishRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencySequencerServer).PublishRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.sequencer.KeyTransparencySequencer/PublishRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencySequencerServer).PublishRevisions(ctx, req.(*PublishRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyTransparencySequencer_UpdateMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyTransparencySequencerServer).UpdateMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.keytransparency.sequencer.KeyTransparencySequencer/UpdateMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyTransparencySequencerServer).UpdateMetrics(ctx, req.(*UpdateMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyTransparencySequencer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.keytransparency.sequencer.KeyTransparencySequencer",
	HandlerType: (*KeyTransparencySequencerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunBatch",
			Handler:    _KeyTransparencySequencer_RunBatch_Handler,
		},
		{
			MethodName: "DefineRevisions",
			Handler:    _KeyTransparencySequencer_DefineRevisions_Handler,
		},
		{
			MethodName: "ApplyRevision",
			Handler:    _KeyTransparencySequencer_ApplyRevision_Handler,
		},
		{
			MethodName: "PublishRevisions",
			Handler:    _KeyTransparencySequencer_PublishRevisions_Handler,
		},
		{
			MethodName: "UpdateMetrics",
			Handler:    _KeyTransparencySequencer_UpdateMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sequencer_api.proto",
}
