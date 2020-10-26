// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/topic.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicRequest) Reset()         { *m = CreateTopicRequest{} }
func (m *CreateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTopicRequest) ProtoMessage()    {}
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{0}
}

func (m *CreateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicRequest.Unmarshal(m, b)
}
func (m *CreateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicRequest.Marshal(b, m, deterministic)
}
func (m *CreateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicRequest.Merge(m, src)
}
func (m *CreateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTopicRequest.Size(m)
}
func (m *CreateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicRequest proto.InternalMessageInfo

type CreateTopicResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicResponse) Reset()         { *m = CreateTopicResponse{} }
func (m *CreateTopicResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTopicResponse) ProtoMessage()    {}
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{1}
}

func (m *CreateTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicResponse.Unmarshal(m, b)
}
func (m *CreateTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicResponse.Marshal(b, m, deterministic)
}
func (m *CreateTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicResponse.Merge(m, src)
}
func (m *CreateTopicResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTopicResponse.Size(m)
}
func (m *CreateTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicResponse proto.InternalMessageInfo

type ListTopicsRequest struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsRequest) Reset()         { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()    {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{2}
}

func (m *ListTopicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsRequest.Unmarshal(m, b)
}
func (m *ListTopicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsRequest.Marshal(b, m, deterministic)
}
func (m *ListTopicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsRequest.Merge(m, src)
}
func (m *ListTopicsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTopicsRequest.Size(m)
}
func (m *ListTopicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsRequest proto.InternalMessageInfo

func (m *ListTopicsRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type ListTopicsResponse struct {
	Topics               []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsResponse) Reset()         { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()    {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{3}
}

func (m *ListTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsResponse.Unmarshal(m, b)
}
func (m *ListTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsResponse.Marshal(b, m, deterministic)
}
func (m *ListTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsResponse.Merge(m, src)
}
func (m *ListTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopicsResponse.Size(m)
}
func (m *ListTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsResponse proto.InternalMessageInfo

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type GetTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicRequest) Reset()         { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()    {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{4}
}

func (m *GetTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicRequest.Unmarshal(m, b)
}
func (m *GetTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicRequest.Merge(m, src)
}
func (m *GetTopicRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicRequest.Size(m)
}
func (m *GetTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicRequest proto.InternalMessageInfo

type DeleteTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTopicRequest) Reset()         { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()    {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{5}
}

func (m *DeleteTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTopicRequest.Unmarshal(m, b)
}
func (m *DeleteTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTopicRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTopicRequest.Merge(m, src)
}
func (m *DeleteTopicRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTopicRequest.Size(m)
}
func (m *DeleteTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTopicRequest proto.InternalMessageInfo

type UpdateTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTopicRequest) Reset()         { *m = UpdateTopicRequest{} }
func (m *UpdateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTopicRequest) ProtoMessage()    {}
func (*UpdateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{6}
}

func (m *UpdateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTopicRequest.Unmarshal(m, b)
}
func (m *UpdateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTopicRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTopicRequest.Merge(m, src)
}
func (m *UpdateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTopicRequest.Size(m)
}
func (m *UpdateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTopicRequest proto.InternalMessageInfo

type TopicDetail struct {
	Partitions           []*Partition   `protobuf:"bytes,4,rep,name=partitions,proto3" json:"partitions,omitempty"`
	Config               []*TopicConfig `protobuf:"bytes,5,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TopicDetail) Reset()         { *m = TopicDetail{} }
func (m *TopicDetail) String() string { return proto.CompactTextString(m) }
func (*TopicDetail) ProtoMessage()    {}
func (*TopicDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{7}
}

func (m *TopicDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicDetail.Unmarshal(m, b)
}
func (m *TopicDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicDetail.Marshal(b, m, deterministic)
}
func (m *TopicDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicDetail.Merge(m, src)
}
func (m *TopicDetail) XXX_Size() int {
	return xxx_messageInfo_TopicDetail.Size(m)
}
func (m *TopicDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TopicDetail proto.InternalMessageInfo

func (m *TopicDetail) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *TopicDetail) GetConfig() []*TopicConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type Topic struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumReplicas          int32        `protobuf:"varint,2,opt,name=num_replicas,json=numReplicas,proto3" json:"num_replicas,omitempty"`
	NumPartitions        int32        `protobuf:"varint,3,opt,name=num_partitions,json=numPartitions,proto3" json:"num_partitions,omitempty"`
	TotalHighWatermarks  int64        `protobuf:"varint,6,opt,name=total_high_watermarks,json=totalHighWatermarks,proto3" json:"total_high_watermarks,omitempty"`
	TopicDetail          *TopicDetail `protobuf:"bytes,7,opt,name=topic_detail,json=topicDetail,proto3" json:"topic_detail,omitempty"`
	LogDirBytes          int64        `protobuf:"varint,8,opt,name=log_dir_bytes,json=logDirBytes,proto3" json:"log_dir_bytes,omitempty"`
	Messages             int64        `protobuf:"varint,9,opt,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{8}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Topic) GetNumReplicas() int32 {
	if m != nil {
		return m.NumReplicas
	}
	return 0
}

func (m *Topic) GetNumPartitions() int32 {
	if m != nil {
		return m.NumPartitions
	}
	return 0
}

func (m *Topic) GetTotalHighWatermarks() int64 {
	if m != nil {
		return m.TotalHighWatermarks
	}
	return 0
}

func (m *Topic) GetTopicDetail() *TopicDetail {
	if m != nil {
		return m.TopicDetail
	}
	return nil
}

func (m *Topic) GetLogDirBytes() int64 {
	if m != nil {
		return m.LogDirBytes
	}
	return 0
}

func (m *Topic) GetMessages() int64 {
	if m != nil {
		return m.Messages
	}
	return 0
}

type TopicConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ReadOnly             bool     `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	Sensitive            bool     `protobuf:"varint,4,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicConfig) Reset()         { *m = TopicConfig{} }
func (m *TopicConfig) String() string { return proto.CompactTextString(m) }
func (*TopicConfig) ProtoMessage()    {}
func (*TopicConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{9}
}

func (m *TopicConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicConfig.Unmarshal(m, b)
}
func (m *TopicConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicConfig.Marshal(b, m, deterministic)
}
func (m *TopicConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicConfig.Merge(m, src)
}
func (m *TopicConfig) XXX_Size() int {
	return xxx_messageInfo_TopicConfig.Size(m)
}
func (m *TopicConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TopicConfig proto.InternalMessageInfo

func (m *TopicConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopicConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TopicConfig) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *TopicConfig) GetSensitive() bool {
	if m != nil {
		return m.Sensitive
	}
	return false
}

type Partition struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	HighWatermark        int64    `protobuf:"varint,3,opt,name=high_watermark,json=highWatermark,proto3" json:"high_watermark,omitempty"`
	Leader               int32    `protobuf:"varint,4,opt,name=leader,proto3" json:"leader,omitempty"`
	Replicas             []int32  `protobuf:"varint,5,rep,packed,name=replicas,proto3" json:"replicas,omitempty"`
	Isr                  []int32  `protobuf:"varint,6,rep,packed,name=isr,proto3" json:"isr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{10}
}

func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (m *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(m, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Partition) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Partition) GetHighWatermark() int64 {
	if m != nil {
		return m.HighWatermark
	}
	return 0
}

func (m *Partition) GetLeader() int32 {
	if m != nil {
		return m.Leader
	}
	return 0
}

func (m *Partition) GetReplicas() []int32 {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *Partition) GetIsr() []int32 {
	if m != nil {
		return m.Isr
	}
	return nil
}

type GetHighWatermarksRequest struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Topics               []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHighWatermarksRequest) Reset()         { *m = GetHighWatermarksRequest{} }
func (m *GetHighWatermarksRequest) String() string { return proto.CompactTextString(m) }
func (*GetHighWatermarksRequest) ProtoMessage()    {}
func (*GetHighWatermarksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{11}
}

func (m *GetHighWatermarksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHighWatermarksRequest.Unmarshal(m, b)
}
func (m *GetHighWatermarksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHighWatermarksRequest.Marshal(b, m, deterministic)
}
func (m *GetHighWatermarksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHighWatermarksRequest.Merge(m, src)
}
func (m *GetHighWatermarksRequest) XXX_Size() int {
	return xxx_messageInfo_GetHighWatermarksRequest.Size(m)
}
func (m *GetHighWatermarksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHighWatermarksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHighWatermarksRequest proto.InternalMessageInfo

func (m *GetHighWatermarksRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *GetHighWatermarksRequest) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type TopicHighWatermarks struct {
	HighWatermarks       map[int32]int64 `protobuf:"bytes,1,rep,name=high_watermarks,json=highWatermarks,proto3" json:"high_watermarks,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TopicHighWatermarks) Reset()         { *m = TopicHighWatermarks{} }
func (m *TopicHighWatermarks) String() string { return proto.CompactTextString(m) }
func (*TopicHighWatermarks) ProtoMessage()    {}
func (*TopicHighWatermarks) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{12}
}

func (m *TopicHighWatermarks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicHighWatermarks.Unmarshal(m, b)
}
func (m *TopicHighWatermarks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicHighWatermarks.Marshal(b, m, deterministic)
}
func (m *TopicHighWatermarks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicHighWatermarks.Merge(m, src)
}
func (m *TopicHighWatermarks) XXX_Size() int {
	return xxx_messageInfo_TopicHighWatermarks.Size(m)
}
func (m *TopicHighWatermarks) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicHighWatermarks.DiscardUnknown(m)
}

var xxx_messageInfo_TopicHighWatermarks proto.InternalMessageInfo

func (m *TopicHighWatermarks) GetHighWatermarks() map[int32]int64 {
	if m != nil {
		return m.HighWatermarks
	}
	return nil
}

type GetHighWatermarksResponse struct {
	HighWatermarks       map[string]*TopicHighWatermarks `protobuf:"bytes,1,rep,name=high_watermarks,json=highWatermarks,proto3" json:"high_watermarks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetHighWatermarksResponse) Reset()         { *m = GetHighWatermarksResponse{} }
func (m *GetHighWatermarksResponse) String() string { return proto.CompactTextString(m) }
func (*GetHighWatermarksResponse) ProtoMessage()    {}
func (*GetHighWatermarksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{13}
}

func (m *GetHighWatermarksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHighWatermarksResponse.Unmarshal(m, b)
}
func (m *GetHighWatermarksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHighWatermarksResponse.Marshal(b, m, deterministic)
}
func (m *GetHighWatermarksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHighWatermarksResponse.Merge(m, src)
}
func (m *GetHighWatermarksResponse) XXX_Size() int {
	return xxx_messageInfo_GetHighWatermarksResponse.Size(m)
}
func (m *GetHighWatermarksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHighWatermarksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHighWatermarksResponse proto.InternalMessageInfo

func (m *GetHighWatermarksResponse) GetHighWatermarks() map[string]*TopicHighWatermarks {
	if m != nil {
		return m.HighWatermarks
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTopicRequest)(nil), "kaf.api.CreateTopicRequest")
	proto.RegisterType((*CreateTopicResponse)(nil), "kaf.api.CreateTopicResponse")
	proto.RegisterType((*ListTopicsRequest)(nil), "kaf.api.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "kaf.api.ListTopicsResponse")
	proto.RegisterType((*GetTopicRequest)(nil), "kaf.api.GetTopicRequest")
	proto.RegisterType((*DeleteTopicRequest)(nil), "kaf.api.DeleteTopicRequest")
	proto.RegisterType((*UpdateTopicRequest)(nil), "kaf.api.UpdateTopicRequest")
	proto.RegisterType((*TopicDetail)(nil), "kaf.api.TopicDetail")
	proto.RegisterType((*Topic)(nil), "kaf.api.Topic")
	proto.RegisterType((*TopicConfig)(nil), "kaf.api.TopicConfig")
	proto.RegisterType((*Partition)(nil), "kaf.api.Partition")
	proto.RegisterType((*GetHighWatermarksRequest)(nil), "kaf.api.GetHighWatermarksRequest")
	proto.RegisterType((*TopicHighWatermarks)(nil), "kaf.api.TopicHighWatermarks")
	proto.RegisterMapType((map[int32]int64)(nil), "kaf.api.TopicHighWatermarks.HighWatermarksEntry")
	proto.RegisterType((*GetHighWatermarksResponse)(nil), "kaf.api.GetHighWatermarksResponse")
	proto.RegisterMapType((map[string]*TopicHighWatermarks)(nil), "kaf.api.GetHighWatermarksResponse.HighWatermarksEntry")
}

func init() { proto.RegisterFile("api/topic.proto", fileDescriptor_577e618acc61581b) }

var fileDescriptor_577e618acc61581b = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x86, 0xac, 0xc8, 0xb1, 0x46, 0xf9, 0xd9, 0xd0, 0x4e, 0xa0, 0xb5, 0x73, 0x70, 0x04, 0xec,
	0xc2, 0x87, 0x5d, 0x65, 0xe1, 0x2d, 0xda, 0x22, 0xe8, 0xa5, 0xf9, 0x41, 0x7a, 0x08, 0xd0, 0x80,
	0x6d, 0x51, 0x34, 0x17, 0x81, 0xb6, 0x69, 0x99, 0x88, 0xfe, 0x2a, 0x52, 0x29, 0xfc, 0x1a, 0x7d,
	0x85, 0x5e, 0x7a, 0xef, 0xf3, 0xf4, 0x5d, 0x0a, 0x51, 0xb4, 0x22, 0xd9, 0x4e, 0xd2, 0x1b, 0xe7,
	0x8f, 0x33, 0xf3, 0xcd, 0xf0, 0x23, 0xec, 0x92, 0x84, 0x1d, 0x8b, 0x38, 0x61, 0x63, 0x37, 0x49,
	0x63, 0x11, 0xa3, 0xcd, 0x5b, 0x32, 0x75, 0x49, 0xc2, 0xba, 0x3d, 0x3f, 0x8e, 0xfd, 0x80, 0x1e,
	0x4b, 0xf5, 0x28, 0x9b, 0x1e, 0xd3, 0x30, 0x11, 0xf3, 0xc2, 0xcb, 0xe9, 0x00, 0x3a, 0x4b, 0x29,
	0x11, 0xf4, 0x7d, 0x1e, 0x8a, 0xe9, 0xe7, 0x8c, 0x72, 0xe1, 0xec, 0x43, 0xbb, 0xa6, 0xe5, 0x49,
	0x1c, 0x71, 0xea, 0xfc, 0x0b, 0x7b, 0x57, 0x8c, 0x0b, 0xa9, 0xe4, 0xca, 0x17, 0xd9, 0xb0, 0x39,
	0x0e, 0x32, 0x2e, 0x68, 0x6a, 0x6b, 0x7d, 0x6d, 0x60, 0xe2, 0x85, 0xe8, 0xbc, 0x02, 0x54, 0x75,
	0x2f, 0x2e, 0x41, 0x7f, 0x43, 0x53, 0x96, 0xc9, 0x6d, 0xad, 0xaf, 0x0f, 0xac, 0xe1, 0x8e, 0xab,
	0x0a, 0x75, 0x8b, 0x64, 0xca, 0xea, 0xec, 0xc1, 0xee, 0x25, 0x15, 0xb5, 0xb2, 0x3a, 0x80, 0xce,
	0x69, 0x40, 0x97, 0x8a, 0xed, 0x00, 0xfa, 0x90, 0x4c, 0x96, 0x5b, 0x88, 0xc1, 0x92, 0xf2, 0x39,
	0x15, 0x84, 0x05, 0x68, 0x08, 0x90, 0x90, 0x54, 0x30, 0xc1, 0xe2, 0x88, 0xdb, 0x1b, 0x32, 0x33,
	0x2a, 0x33, 0x5f, 0x2f, 0x4c, 0xb8, 0xe2, 0x85, 0xfe, 0x81, 0xe6, 0x38, 0x8e, 0xa6, 0xcc, 0xb7,
	0x0d, 0xe9, 0xdf, 0xa9, 0x57, 0x7a, 0x26, 0x6d, 0x58, 0xf9, 0x38, 0x5f, 0x1b, 0x60, 0x48, 0x3d,
	0x42, 0xb0, 0x11, 0x91, 0x90, 0x2a, 0x38, 0xe4, 0x19, 0x1d, 0xc1, 0x56, 0x94, 0x85, 0x5e, 0x4a,
	0x93, 0x80, 0x8d, 0x09, 0xb7, 0x1b, 0x7d, 0x6d, 0x60, 0x60, 0x2b, 0xca, 0x42, 0xac, 0x54, 0xe8,
	0x2f, 0xd8, 0xc9, 0x5d, 0x2a, 0x65, 0xea, 0xd2, 0x69, 0x3b, 0xca, 0xc2, 0xeb, 0xfb, 0xaa, 0x86,
	0xb0, 0x2f, 0x62, 0x41, 0x02, 0x6f, 0xc6, 0xfc, 0x99, 0xf7, 0x85, 0x08, 0x9a, 0x86, 0x24, 0xbd,
	0xe5, 0x76, 0xb3, 0xaf, 0x0d, 0x74, 0xdc, 0x96, 0xc6, 0x37, 0xcc, 0x9f, 0x7d, 0x2c, 0x4d, 0xe8,
	0x05, 0x6c, 0x49, 0x54, 0xbd, 0x89, 0x44, 0xc3, 0xde, 0xec, 0x6b, 0xab, 0xfd, 0x14, 0x48, 0x61,
	0x4b, 0x54, 0x60, 0x73, 0x60, 0x3b, 0x88, 0x7d, 0x6f, 0xc2, 0x52, 0x6f, 0x34, 0x17, 0x94, 0xdb,
	0x2d, 0x99, 0xc4, 0x0a, 0x62, 0xff, 0x9c, 0xa5, 0xa7, 0xb9, 0x0a, 0x75, 0xa1, 0x15, 0x52, 0xce,
	0x89, 0x4f, 0xb9, 0x6d, 0x4a, 0x73, 0x29, 0x3b, 0xa9, 0x9a, 0x42, 0x81, 0xd5, 0x5a, 0x64, 0x3a,
	0x60, 0xdc, 0x91, 0x20, 0xa3, 0x12, 0x12, 0x13, 0x17, 0x02, 0xea, 0x81, 0x99, 0x52, 0x32, 0xf1,
	0xe2, 0x28, 0x98, 0x4b, 0x1c, 0x5a, 0xb8, 0x95, 0x2b, 0xde, 0x46, 0xc1, 0x1c, 0x1d, 0x82, 0xc9,
	0x69, 0xc4, 0x99, 0x60, 0x77, 0xd4, 0xde, 0x90, 0xc6, 0x7b, 0x85, 0xf3, 0x5d, 0x03, 0xb3, 0xc4,
	0x0b, 0x1d, 0x40, 0x33, 0xca, 0xc2, 0x91, 0xda, 0x4e, 0x03, 0x2b, 0x29, 0xd7, 0xc7, 0xd3, 0x29,
	0xa7, 0x42, 0xe6, 0xd5, 0xb1, 0x92, 0xf2, 0x29, 0xd4, 0x81, 0x95, 0xd9, 0x75, 0xbc, 0x3d, 0xab,
	0x42, 0x9a, 0x87, 0x07, 0x94, 0x4c, 0x68, 0x2a, 0xf3, 0x1b, 0x58, 0x49, 0x39, 0x18, 0xe5, 0x8c,
	0xf3, 0xad, 0x31, 0x70, 0x29, 0xa3, 0x3f, 0x40, 0x67, 0x3c, 0xb5, 0x9b, 0x52, 0x9d, 0x1f, 0x9d,
	0x2b, 0xb0, 0x2f, 0xa9, 0xa8, 0x0f, 0xeb, 0xc9, 0x77, 0x95, 0xe7, 0x56, 0x2f, 0xa8, 0xd1, 0xd7,
	0x07, 0x66, 0xf9, 0x62, 0x7e, 0x68, 0xd0, 0x96, 0x68, 0x2f, 0x4d, 0xff, 0x13, 0xec, 0x2e, 0xef,
	0x4a, 0xf1, 0xf4, 0xfe, 0xab, 0x2f, 0x40, 0x3d, 0xcc, 0xad, 0x8b, 0x17, 0x91, 0x48, 0xe7, 0x78,
	0xa7, 0x86, 0x02, 0xef, 0xbe, 0x86, 0xf6, 0x1a, 0xb7, 0xbc, 0xd3, 0x5b, 0x3a, 0x57, 0x88, 0xe7,
	0xc7, 0xfa, 0x94, 0x75, 0x35, 0xe5, 0x93, 0xc6, 0x4b, 0xcd, 0xf9, 0xa9, 0xc1, 0x9f, 0x6b, 0x40,
	0x50, 0x6c, 0xe1, 0x3d, 0x54, 0xfb, 0xf3, 0xb2, 0xf6, 0x07, 0x83, 0x7f, 0xab, 0x03, 0xef, 0xc9,
	0x0e, 0xcc, 0xa2, 0x83, 0x61, 0xb5, 0x03, 0x6b, 0x78, 0xf8, 0x18, 0x76, 0x95, 0xfe, 0x86, 0xdf,
	0x74, 0xd8, 0x92, 0x2e, 0xef, 0x68, 0x7a, 0xc7, 0xc6, 0x14, 0x9d, 0x80, 0x55, 0x21, 0x57, 0xd4,
	0x2b, 0x2f, 0x5a, 0x25, 0xe2, 0xee, 0x12, 0x39, 0xa2, 0x67, 0xd0, 0x5a, 0x90, 0x22, 0xb2, 0xab,
	0x08, 0x3c, 0x1a, 0x75, 0x02, 0x56, 0x85, 0x21, 0x2b, 0x19, 0x57, 0x79, 0x73, 0x25, 0xf6, 0x02,
	0xe0, 0x9e, 0xc4, 0x51, 0xb7, 0xb4, 0xae, 0x7c, 0x04, 0xdd, 0xde, 0x5a, 0x9b, 0x9a, 0xe3, 0x29,
	0x58, 0x15, 0xea, 0xae, 0x94, 0xb0, 0x4a, 0xe8, 0xdd, 0x03, 0xb7, 0xf8, 0xb1, 0xdc, 0xc5, 0x8f,
	0xe5, 0x5e, 0xe4, 0x3f, 0x16, 0xba, 0x81, 0xbd, 0x95, 0x59, 0xa3, 0xa3, 0xc7, 0xf6, 0xa0, 0xb8,
	0xcf, 0x79, 0x7a, 0x55, 0x4e, 0x8d, 0x1b, 0x9d, 0x24, 0x6c, 0xd4, 0x94, 0x29, 0xff, 0xff, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0xb9, 0x29, 0xd1, 0x4e, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetHighWatermarks(ctx context.Context, in *GetHighWatermarksRequest, opts ...grpc.CallOption) (*GetHighWatermarksResponse, error)
}

type topicServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopicServiceClient(cc *grpc.ClientConn) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/GetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/UpdateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetHighWatermarks(ctx context.Context, in *GetHighWatermarksRequest, opts ...grpc.CallOption) (*GetHighWatermarksResponse, error) {
	out := new(GetHighWatermarksResponse)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/GetHighWatermarks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*Topic, error)
	GetTopic(context.Context, *GetTopicRequest) (*Topic, error)
	UpdateTopic(context.Context, *UpdateTopicRequest) (*Topic, error)
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*empty.Empty, error)
	GetHighWatermarks(context.Context, *GetHighWatermarksRequest) (*GetHighWatermarksResponse, error)
}

func RegisterTopicServiceServer(s *grpc.Server, srv TopicServiceServer) {
	s.RegisterService(&_TopicService_serviceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/GetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetHighWatermarks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHighWatermarksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetHighWatermarks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/GetHighWatermarks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetHighWatermarks(ctx, req.(*GetHighWatermarksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kaf.api.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicService_GetTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _TopicService_UpdateTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _TopicService_ListTopics_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
		{
			MethodName: "GetHighWatermarks",
			Handler:    _TopicService_GetHighWatermarks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/topic.proto",
}
