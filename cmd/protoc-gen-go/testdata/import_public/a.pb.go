// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/a.proto

package import_public

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	sub "github.com/golang/protobuf/v2/cmd/protoc-gen-go/testdata/import_public/sub"
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

const Default_M_S = sub.Default_M_S

// M from public import import_public/sub/a.proto
type M = sub.M
type M_OneofInt32 = sub.M_OneofInt32
type M_OneofInt64 = sub.M_OneofInt64

// M_Submessage from public import import_public/sub/a.proto
type M_Submessage = sub.M_Submessage
type M_Submessage_SubmessageOneofInt32 = sub.M_Submessage_SubmessageOneofInt32
type M_Submessage_SubmessageOneofInt64 = sub.M_Submessage_SubmessageOneofInt64

// E from public import import_public/sub/a.proto
type E = sub.E

var E_name = sub.E_name
var E_value = sub.E_value

const E_ZERO = E(sub.E_ZERO)

// M_Subenum from public import import_public/sub/a.proto
type M_Subenum = sub.M_Subenum

var M_Subenum_name = sub.M_Subenum_name
var M_Subenum_value = sub.M_Subenum_value

const M_M_ZERO = M_Subenum(sub.M_M_ZERO)

// M_Submessage_Submessage_Subenum from public import import_public/sub/a.proto
type M_Submessage_Submessage_Subenum = sub.M_Submessage_Submessage_Subenum

var M_Submessage_Submessage_Subenum_name = sub.M_Submessage_Submessage_Subenum_name
var M_Submessage_Submessage_Subenum_value = sub.M_Submessage_Submessage_Subenum_value

const M_Submessage_M_SUBMESSAGE_ZERO = M_Submessage_Submessage_Subenum(sub.M_Submessage_M_SUBMESSAGE_ZERO)

var E_ExtensionField = sub.E_ExtensionField

type Public struct {
	M                    *sub.M   `protobuf:"bytes,1,opt,name=m" json:"m,omitempty"`
	E                    *sub.E   `protobuf:"varint,2,opt,name=e,enum=goproto.protoc.import_public.sub.E" json:"e,omitempty"`
	Local                *Local   `protobuf:"bytes,3,opt,name=local" json:"local,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Public) Reset()         { *m = Public{} }
func (m *Public) String() string { return proto.CompactTextString(m) }
func (*Public) ProtoMessage()    {}
func (*Public) Descriptor() ([]byte, []int) {
	return fileDescriptor_73b7577c95fa6b70, []int{0}
}

func (m *Public) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Public.Unmarshal(m, b)
}
func (m *Public) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Public.Marshal(b, m, deterministic)
}
func (m *Public) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Public.Merge(m, src)
}
func (m *Public) XXX_Size() int {
	return xxx_messageInfo_Public.Size(m)
}
func (m *Public) XXX_DiscardUnknown() {
	xxx_messageInfo_Public.DiscardUnknown(m)
}

var xxx_messageInfo_Public proto.InternalMessageInfo

func (m *Public) GetM() *sub.M {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Public) GetE() sub.E {
	if m != nil && m.E != nil {
		return *m.E
	}
	return sub.E_ZERO
}

func (m *Public) GetLocal() *Local {
	if m != nil {
		return m.Local
	}
	return nil
}

func init() {
	proto.RegisterType((*Public)(nil), "goproto.protoc.import_public.Public")
}

func init() { proto.RegisterFile("import_public/a.proto", fileDescriptor_73b7577c95fa6b70) }

var fileDescriptor_73b7577c95fa6b70 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x4f, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x49, 0xcf, 0x07, 0x33, 0x20, 0xdc, 0x64, 0x3d, 0x14, 0x55, 0x52, 0x92, 0xa8,
	0x9a, 0x8a, 0x4b, 0x93, 0x60, 0x1a, 0xa5, 0xd0, 0xcc, 0x4b, 0x82, 0x08, 0x2b, 0xad, 0x64, 0xe4,
	0x62, 0x0b, 0x00, 0x0b, 0x09, 0x19, 0x72, 0x31, 0xe6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b,
	0x29, 0xeb, 0xe1, 0xb3, 0x46, 0xaf, 0xb8, 0x34, 0x49, 0xcf, 0x37, 0x88, 0x31, 0x17, 0xa4, 0x25,
	0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x8f, 0x18, 0x2d, 0xae, 0x41, 0x8c, 0xa9, 0x42, 0x96, 0x5c,
	0xac, 0x39, 0xf9, 0xc9, 0x89, 0x39, 0x12, 0xcc, 0xc4, 0xd8, 0xe4, 0x03, 0x52, 0x1a, 0x04, 0xd1,
	0xe1, 0xe4, 0x11, 0xe5, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x0f, 0xd6, 0x96, 0x54, 0x9a, 0xa6, 0x5f, 0x66, 0xa4, 0x9f,
	0x9c, 0x9b, 0x02, 0xe1, 0x27, 0xeb, 0xa6, 0xa7, 0xe6, 0xe9, 0xa6, 0xe7, 0xeb, 0x97, 0xa4, 0x16,
	0x97, 0xa4, 0x24, 0x96, 0x24, 0xea, 0xa3, 0x18, 0x1b, 0xc0, 0x10, 0xc0, 0x08, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0x3e, 0xda, 0xad, 0x61, 0x01, 0x00, 0x00,
}
