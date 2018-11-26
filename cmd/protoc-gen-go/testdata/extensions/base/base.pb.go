// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extensions/base/base.proto

package base

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

type BaseMessage struct {
	Field                        *string  `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *BaseMessage) Reset()         { *m = BaseMessage{} }
func (m *BaseMessage) String() string { return proto.CompactTextString(m) }
func (*BaseMessage) ProtoMessage()    {}
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_aebb28f8d5a04466, []int{0}
}

var extRange_BaseMessage = []proto.ExtensionRange{
	{Start: 4, End: 9},
	{Start: 16, End: 536870911},
}

func (*BaseMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_BaseMessage
}

func (m *BaseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMessage.Unmarshal(m, b)
}
func (m *BaseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMessage.Marshal(b, m, deterministic)
}
func (m *BaseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMessage.Merge(m, src)
}
func (m *BaseMessage) XXX_Size() int {
	return xxx_messageInfo_BaseMessage.Size(m)
}
func (m *BaseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMessage proto.InternalMessageInfo

func (m *BaseMessage) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

type MessageSetWireFormatMessage struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `protobuf_messageset:"1" json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *MessageSetWireFormatMessage) Reset()         { *m = MessageSetWireFormatMessage{} }
func (m *MessageSetWireFormatMessage) String() string { return proto.CompactTextString(m) }
func (*MessageSetWireFormatMessage) ProtoMessage()    {}
func (*MessageSetWireFormatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_aebb28f8d5a04466, []int{1}
}

func (m *MessageSetWireFormatMessage) MarshalJSON() ([]byte, error) {
	return proto.MarshalMessageSetJSON(&m.XXX_InternalExtensions)
}
func (m *MessageSetWireFormatMessage) UnmarshalJSON(buf []byte) error {
	return proto.UnmarshalMessageSetJSON(buf, &m.XXX_InternalExtensions)
}

var extRange_MessageSetWireFormatMessage = []proto.ExtensionRange{
	{Start: 100, End: 2147483646},
}

func (*MessageSetWireFormatMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MessageSetWireFormatMessage
}

func (m *MessageSetWireFormatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSetWireFormatMessage.Unmarshal(m, b)
}
func (m *MessageSetWireFormatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSetWireFormatMessage.Marshal(b, m, deterministic)
}
func (m *MessageSetWireFormatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSetWireFormatMessage.Merge(m, src)
}
func (m *MessageSetWireFormatMessage) XXX_Size() int {
	return xxx_messageInfo_MessageSetWireFormatMessage.Size(m)
}
func (m *MessageSetWireFormatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSetWireFormatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSetWireFormatMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseMessage)(nil), "goproto.protoc.extension.base.BaseMessage")
	proto.RegisterType((*MessageSetWireFormatMessage)(nil), "goproto.protoc.extension.base.MessageSetWireFormatMessage")
}

func init() { proto.RegisterFile("extensions/base/base.proto", fileDescriptor_aebb28f8d5a04466) }

var fileDescriptor_aebb28f8d5a04466 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xad, 0x28, 0x49,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0x05, 0x13, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xb2, 0xe9, 0xf9, 0x60, 0x06, 0x84, 0x9b, 0xac, 0x07, 0x57, 0xaa, 0x07,
	0x52, 0xa4, 0x64, 0xcc, 0xc5, 0xed, 0x94, 0x58, 0x9c, 0xea, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e,
	0x2a, 0x24, 0xc2, 0xc5, 0x9a, 0x96, 0x99, 0x9a, 0x93, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0xe1, 0x68, 0xb1, 0x70, 0xb0, 0x08, 0x70, 0x69, 0x71, 0x70, 0x08, 0x08, 0x34, 0x34, 0x34,
	0x34, 0x30, 0x29, 0x69, 0x73, 0x49, 0x43, 0x35, 0x04, 0xa7, 0x96, 0x84, 0x67, 0x16, 0xa5, 0xba,
	0xe5, 0x17, 0xe5, 0x26, 0x96, 0x40, 0xc5, 0xb4, 0x38, 0x38, 0x52, 0x04, 0xfe, 0xff, 0xff, 0xff,
	0x9f, 0xdd, 0x8a, 0x89, 0x83, 0xd1, 0xc9, 0x2b, 0xca, 0x23, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x1f, 0xec, 0x98, 0xa4, 0xd2,
	0x34, 0xfd, 0x32, 0x23, 0xfd, 0xe4, 0xdc, 0x14, 0x08, 0x3f, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0x37,
	0x3d, 0x5f, 0xbf, 0x24, 0xb5, 0xb8, 0x24, 0x25, 0xb1, 0x24, 0x51, 0x1f, 0xcd, 0x5f, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x75, 0xde, 0x5a, 0xe9, 0x00, 0x00, 0x00,
}
