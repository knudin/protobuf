// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_import_all.proto

package imports

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	fmt1 "github.com/golang/protobuf/protoc-gen-go/testdata/imports/fmt"
	test_a_1 "github.com/golang/protobuf/protoc-gen-go/testdata/imports/test_a_1"
	_ "github.com/golang/protobuf/protoc-gen-go/testdata/imports/test_a_2"
	test_b_1 "github.com/golang/protobuf/protoc-gen-go/testdata/imports/test_b_1"
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

type All struct {
	Am1                  *test_a_1.M1 `protobuf:"bytes,1,opt,name=am1,proto3" json:"am1,omitempty"`
	Am2                  *test_a_1.M2 `protobuf:"bytes,2,opt,name=am2,proto3" json:"am2,omitempty"`
	Bm1                  *test_b_1.M1 `protobuf:"bytes,5,opt,name=bm1,proto3" json:"bm1,omitempty"`
	Bm2                  *test_b_1.M2 `protobuf:"bytes,6,opt,name=bm2,proto3" json:"bm2,omitempty"`
	Fmt                  *fmt1.M      `protobuf:"bytes,7,opt,name=fmt,proto3" json:"fmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_324466f0afc16f77, []int{0}
}

func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (m *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(m, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetAm1() *test_a_1.M1 {
	if m != nil {
		return m.Am1
	}
	return nil
}

func (m *All) GetAm2() *test_a_1.M2 {
	if m != nil {
		return m.Am2
	}
	return nil
}

func (m *All) GetBm1() *test_b_1.M1 {
	if m != nil {
		return m.Bm1
	}
	return nil
}

func (m *All) GetBm2() *test_b_1.M2 {
	if m != nil {
		return m.Bm2
	}
	return nil
}

func (m *All) GetFmt() *fmt1.M {
	if m != nil {
		return m.Fmt
	}
	return nil
}

func init() {
	proto.RegisterType((*All)(nil), "test.All")
}

func init() { proto.RegisterFile("imports/test_import_all.proto", fileDescriptor_324466f0afc16f77) }

var fileDescriptor_324466f0afc16f77 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x85, 0x40, 0x90, 0xcc, 0x82, 0xc2, 0x62, 0x10, 0x48, 0xa8, 0x13, 0x4b, 0x6d,
	0xf9, 0xca, 0x82, 0x98, 0x60, 0xef, 0xd2, 0x91, 0x25, 0xf2, 0x95, 0xda, 0x54, 0xf2, 0xd5, 0x51,
	0x7a, 0x7d, 0x22, 0x5e, 0x14, 0xd9, 0x06, 0x09, 0x68, 0xb2, 0x25, 0xff, 0xf7, 0xfb, 0x4e, 0x27,
	0xee, 0xb6, 0xd4, 0xc7, 0x81, 0xf7, 0x9a, 0x37, 0x7b, 0xee, 0xca, 0x4f, 0x67, 0x43, 0x50, 0xfd,
	0x10, 0x39, 0xb6, 0xa7, 0x29, 0xbe, 0xb9, 0xfe, 0x53, 0xb2, 0x9d, 0xd1, 0x64, 0x4a, 0x61, 0x8c,
	0x60, 0x82, 0x40, 0xd3, 0x62, 0x9a, 0x1e, 0x47, 0x09, 0xa7, 0x77, 0xe1, 0xef, 0x5d, 0x57, 0x3f,
	0xe4, 0x88, 0x35, 0x95, 0x70, 0xf6, 0x59, 0x89, 0xfa, 0x25, 0x84, 0xf6, 0x56, 0xd4, 0x96, 0x8c,
	0xac, 0xee, 0xab, 0x87, 0x0b, 0x10, 0x2a, 0xbd, 0x56, 0x56, 0x2d, 0xcd, 0x2a, 0xc5, 0x45, 0x41,
	0x9e, 0xfc, 0x53, 0x48, 0x0a, 0xed, 0x4c, 0xd4, 0x48, 0x46, 0x9e, 0x65, 0xbd, 0x2c, 0x8a, 0xaa,
	0xb7, 0x03, 0x9b, 0x3c, 0x01, 0xc9, 0x94, 0x0e, 0xc8, 0xe6, 0xb8, 0x03, 0x79, 0x0e, 0x12, 0xb4,
	0x52, 0xd4, 0x8e, 0x58, 0x9e, 0xe7, 0x4e, 0xa3, 0x1c, 0xb1, 0x5a, 0xae, 0x52, 0xf4, 0xfa, 0xfc,
	0xf6, 0xe4, 0xb7, 0xfc, 0x71, 0x40, 0xb5, 0x8e, 0xa4, 0x7d, 0x0c, 0x76, 0xe7, 0x75, 0x3e, 0x00,
	0x0f, 0xae, 0x7c, 0xac, 0xe7, 0x7e, 0xb3, 0x9b, 0xfb, 0x98, 0x0f, 0x7f, 0xb7, 0x6c, 0xf5, 0xf7,
	0xb9, 0xd8, 0x64, 0x5f, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x88, 0x0e, 0xe2, 0x9f, 0xc7, 0x01,
	0x00, 0x00,
}
