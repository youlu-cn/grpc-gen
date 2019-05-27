// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/annotations.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type AccessLevel int32

const (
	AccessLevel__NO_LIMIT           AccessLevel = 0
	AccessLevel_LOW_ACCESS_LEVEL    AccessLevel = 100
	AccessLevel_MIDDLE_ACCESS_LEVEL AccessLevel = 200
	AccessLevel_HIGH_ACCESS_LEVEL   AccessLevel = 300
	AccessLevel_SERVER_INTERNAL     AccessLevel = 400
)

var AccessLevel_name = map[int32]string{
	0:   "_NO_LIMIT",
	100: "LOW_ACCESS_LEVEL",
	200: "MIDDLE_ACCESS_LEVEL",
	300: "HIGH_ACCESS_LEVEL",
	400: "SERVER_INTERNAL",
}

var AccessLevel_value = map[string]int32{
	"_NO_LIMIT":           0,
	"LOW_ACCESS_LEVEL":    100,
	"MIDDLE_ACCESS_LEVEL": 200,
	"HIGH_ACCESS_LEVEL":   300,
	"SERVER_INTERNAL":     400,
}

func (x AccessLevel) String() string {
	return proto.EnumName(AccessLevel_name, int32(x))
}

func (AccessLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51f13039a448e4fa, []int{0}
}

type VisibleScope int32

const (
	VisibleScope_PUBLIC_SCOPE   VisibleScope = 0
	VisibleScope_INTERNAL_SCOPE VisibleScope = 1
	VisibleScope_ALL_SCOPE      VisibleScope = 2
)

var VisibleScope_name = map[int32]string{
	0: "PUBLIC_SCOPE",
	1: "INTERNAL_SCOPE",
	2: "ALL_SCOPE",
}

var VisibleScope_value = map[string]int32{
	"PUBLIC_SCOPE":   0,
	"INTERNAL_SCOPE": 1,
	"ALL_SCOPE":      2,
}

func (x VisibleScope) String() string {
	return proto.EnumName(VisibleScope_name, int32(x))
}

func (VisibleScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51f13039a448e4fa, []int{1}
}

type Access struct {
	Level                AccessLevel `protobuf:"varint,1,opt,name=level,proto3,enum=auth.AccessLevel" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Access) Reset()         { *m = Access{} }
func (m *Access) String() string { return proto.CompactTextString(m) }
func (*Access) ProtoMessage()    {}
func (*Access) Descriptor() ([]byte, []int) {
	return fileDescriptor_51f13039a448e4fa, []int{0}
}

func (m *Access) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Access.Unmarshal(m, b)
}
func (m *Access) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Access.Marshal(b, m, deterministic)
}
func (m *Access) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Access.Merge(m, src)
}
func (m *Access) XXX_Size() int {
	return xxx_messageInfo_Access.Size(m)
}
func (m *Access) XXX_DiscardUnknown() {
	xxx_messageInfo_Access.DiscardUnknown(m)
}

var xxx_messageInfo_Access proto.InternalMessageInfo

func (m *Access) GetLevel() AccessLevel {
	if m != nil {
		return m.Level
	}
	return AccessLevel__NO_LIMIT
}

type Visible struct {
	Scope                VisibleScope `protobuf:"varint,1,opt,name=scope,proto3,enum=auth.VisibleScope" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Visible) Reset()         { *m = Visible{} }
func (m *Visible) String() string { return proto.CompactTextString(m) }
func (*Visible) ProtoMessage()    {}
func (*Visible) Descriptor() ([]byte, []int) {
	return fileDescriptor_51f13039a448e4fa, []int{1}
}

func (m *Visible) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Visible.Unmarshal(m, b)
}
func (m *Visible) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Visible.Marshal(b, m, deterministic)
}
func (m *Visible) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Visible.Merge(m, src)
}
func (m *Visible) XXX_Size() int {
	return xxx_messageInfo_Visible.Size(m)
}
func (m *Visible) XXX_DiscardUnknown() {
	xxx_messageInfo_Visible.DiscardUnknown(m)
}

var xxx_messageInfo_Visible proto.InternalMessageInfo

func (m *Visible) GetScope() VisibleScope {
	if m != nil {
		return m.Scope
	}
	return VisibleScope_PUBLIC_SCOPE
}

var E_Access = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Access)(nil),
	Field:         1042,
	Name:          "auth.access",
	Tag:           "bytes,1042,opt,name=access",
	Filename:      "auth/annotations.proto",
}

var E_Visible = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Visible)(nil),
	Field:         1042,
	Name:          "auth.visible",
	Tag:           "bytes,1042,opt,name=visible",
	Filename:      "auth/annotations.proto",
}

func init() {
	proto.RegisterEnum("auth.AccessLevel", AccessLevel_name, AccessLevel_value)
	proto.RegisterEnum("auth.VisibleScope", VisibleScope_name, VisibleScope_value)
	proto.RegisterType((*Access)(nil), "auth.Access")
	proto.RegisterType((*Visible)(nil), "auth.Visible")
	proto.RegisterExtension(E_Access)
	proto.RegisterExtension(E_Visible)
}

func init() { proto.RegisterFile("auth/annotations.proto", fileDescriptor_51f13039a448e4fa) }

var fileDescriptor_51f13039a448e4fa = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xd1, 0x8a, 0x9b, 0x40,
	0x14, 0x86, 0xd7, 0x6d, 0x37, 0xdb, 0x9e, 0xcd, 0x6e, 0xdd, 0xe9, 0x12, 0x42, 0x2f, 0xda, 0x90,
	0x9b, 0x86, 0x40, 0x94, 0x26, 0x50, 0x4a, 0x7a, 0x65, 0xcc, 0xd0, 0x08, 0x93, 0x18, 0x34, 0x4d,
	0xa1, 0x37, 0xa2, 0x93, 0xa9, 0x0a, 0xd6, 0x11, 0x1d, 0x03, 0xed, 0x53, 0x94, 0x3e, 0x47, 0x1f,
	0xa4, 0x8f, 0x55, 0x74, 0x14, 0x92, 0xee, 0x8d, 0xe8, 0xff, 0x9d, 0xff, 0xe7, 0xf0, 0x1f, 0xa1,
	0xe7, 0x97, 0x22, 0xd2, 0xfd, 0x34, 0xe5, 0xc2, 0x17, 0x31, 0x4f, 0x0b, 0x2d, 0xcb, 0xb9, 0xe0,
	0xe8, 0x69, 0xa5, 0xbf, 0x1a, 0x84, 0x9c, 0x87, 0x09, 0xd3, 0x6b, 0x2d, 0x28, 0xbf, 0xe9, 0x07,
	0x56, 0xd0, 0x3c, 0xce, 0x04, 0xcf, 0xe5, 0xdc, 0xf0, 0x1d, 0x74, 0x0c, 0x4a, 0x59, 0x51, 0xa0,
	0xb7, 0x70, 0x95, 0xb0, 0x23, 0x4b, 0xfa, 0xca, 0x40, 0x19, 0xdd, 0x4d, 0xef, 0xb5, 0x2a, 0x41,
	0x93, 0x90, 0x54, 0xc0, 0x91, 0x7c, 0x38, 0x83, 0xeb, 0x7d, 0x5c, 0xc4, 0x41, 0xc2, 0xd0, 0x08,
	0xae, 0x0a, 0xca, 0x33, 0xd6, 0x78, 0x90, 0xf4, 0x34, 0xd4, 0xad, 0x88, 0x23, 0x07, 0xc6, 0x3f,
	0xe1, 0xe6, 0x24, 0x0a, 0xdd, 0xc2, 0x73, 0x6f, 0x63, 0x7b, 0xc4, 0x5a, 0x5b, 0x3b, 0xf5, 0x02,
	0x3d, 0x80, 0x4a, 0xec, 0x2f, 0x9e, 0x61, 0x9a, 0xd8, 0x75, 0x3d, 0x82, 0xf7, 0x98, 0xa8, 0x07,
	0xd4, 0x87, 0x97, 0x6b, 0x6b, 0xb9, 0x24, 0xf8, 0x1c, 0xfc, 0x55, 0x50, 0x0f, 0xee, 0x57, 0xd6,
	0xa7, 0xd5, 0xb9, 0xfe, 0xe7, 0x12, 0x3d, 0xc0, 0x0b, 0x17, 0x3b, 0x7b, 0xec, 0x78, 0xd6, 0x66,
	0x87, 0x9d, 0x8d, 0x41, 0xd4, 0x5f, 0x4f, 0xc6, 0x26, 0x74, 0x4f, 0x57, 0x42, 0x2a, 0x74, 0xb7,
	0x9f, 0x17, 0xc4, 0x32, 0x3d, 0xd7, 0xb4, 0xb7, 0x58, 0xbd, 0x40, 0x08, 0xee, 0x5a, 0x43, 0xa3,
	0x29, 0xd5, 0x8a, 0x06, 0x69, 0x3f, 0x2f, 0xe7, 0x18, 0x3a, 0xbe, 0x2c, 0xea, 0xb5, 0x26, 0x5b,
	0xd5, 0xda, 0x56, 0xb5, 0x35, 0x13, 0x11, 0x3f, 0xd8, 0x59, 0x7d, 0x80, 0xfe, 0xef, 0x67, 0x03,
	0x65, 0x74, 0x33, 0xed, 0x9e, 0x16, 0xe8, 0x34, 0xe6, 0xb9, 0x05, 0xd7, 0xc7, 0xa6, 0xbc, 0x37,
	0x8f, 0x72, 0x5c, 0x96, 0x1f, 0x63, 0xca, 0xfe, 0x0b, 0xba, 0x3d, 0x6b, 0xd5, 0x69, 0xfd, 0x8b,
	0x0f, 0x5f, 0xdf, 0x87, 0xb1, 0x88, 0xca, 0x40, 0xa3, 0xfc, 0xbb, 0xfe, 0x83, 0x97, 0x49, 0x39,
	0xa1, 0xa9, 0x1e, 0xe6, 0x19, 0x9d, 0x84, 0x2c, 0x95, 0x47, 0xaf, 0x5f, 0x27, 0xf2, 0x27, 0x29,
	0x45, 0xf4, 0xb1, 0x7a, 0x04, 0x9d, 0x1a, 0xcd, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x86,
	0x14, 0xc4, 0x3d, 0x02, 0x00, 0x00,
}
