// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ClusterMessage struct {
	Port                 int32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Shard                int32    `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterMessage) Reset()         { *m = ClusterMessage{} }
func (m *ClusterMessage) String() string { return proto.CompactTextString(m) }
func (*ClusterMessage) ProtoMessage()    {}
func (*ClusterMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{0}
}

func (m *ClusterMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterMessage.Unmarshal(m, b)
}
func (m *ClusterMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterMessage.Marshal(b, m, deterministic)
}
func (m *ClusterMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterMessage.Merge(m, src)
}
func (m *ClusterMessage) XXX_Size() int {
	return xxx_messageInfo_ClusterMessage.Size(m)
}
func (m *ClusterMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterMessage proto.InternalMessageInfo

func (m *ClusterMessage) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ClusterMessage) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

type ClusterResponse struct {
	Port                 int32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Shard                int32    `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterResponse) Reset()         { *m = ClusterResponse{} }
func (m *ClusterResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterResponse) ProtoMessage()    {}
func (*ClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cfb3b8ec240c376, []int{1}
}

func (m *ClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResponse.Unmarshal(m, b)
}
func (m *ClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResponse.Marshal(b, m, deterministic)
}
func (m *ClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResponse.Merge(m, src)
}
func (m *ClusterResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterResponse.Size(m)
}
func (m *ClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResponse proto.InternalMessageInfo

func (m *ClusterResponse) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ClusterResponse) GetShard() int32 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterMessage)(nil), "messages.ClusterMessage")
	proto.RegisterType((*ClusterResponse)(nil), "messages.ClusterResponse")
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_3cfb3b8ec240c376) }

var fileDescriptor_3cfb3b8ec240c376 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x29, 0x2d,
	0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x4f, 0x2d, 0x56, 0xb2, 0xe2, 0xe2, 0x73, 0x86, 0x48, 0xf9, 0x42, 0x84, 0x84, 0x84, 0xb8,
	0x58, 0x0a, 0xf2, 0x8b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11,
	0x2e, 0xd6, 0xe2, 0x8c, 0xc4, 0xa2, 0x14, 0x09, 0x26, 0xb0, 0x20, 0x84, 0xa3, 0x64, 0xcd, 0xc5,
	0x0f, 0xd5, 0x1b, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x4c, 0x82, 0xe6, 0x24, 0x36, 0xb0, 0x4b,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xcb, 0x5f, 0x8b, 0x9a, 0x00, 0x00, 0x00,
}
