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
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x6a, 0xea, 0x40,
	0x14, 0xc6, 0x8d, 0xf7, 0x1a, 0xef, 0x3d, 0xfe, 0x69, 0x9c, 0x8a, 0x48, 0x17, 0xad, 0xb8, 0xa9,
	0xb8, 0x88, 0x54, 0x77, 0xee, 0x62, 0x1c, 0x6a, 0x60, 0x34, 0x32, 0xb1, 0x16, 0xba, 0x09, 0x31,
	0x4e, 0x35, 0x10, 0x9c, 0x90, 0x44, 0x17, 0x7d, 0x8a, 0xd2, 0xe7, 0xe8, 0x83, 0xf4, 0xb1, 0x4a,
	0x32, 0x09, 0x68, 0xbb, 0x3c, 0xdf, 0x77, 0xbe, 0x8f, 0xc3, 0xef, 0x40, 0xcb, 0x39, 0xc6, 0xfb,
	0x81, 0x73, 0x38, 0xf0, 0xd8, 0x89, 0x3d, 0x7e, 0x88, 0xd4, 0x20, 0xe4, 0x31, 0x47, 0x7f, 0x13,
	0xfd, 0xa6, 0xb3, 0xe3, 0x7c, 0xe7, 0xb3, 0x41, 0xaa, 0x6d, 0x8e, 0xaf, 0x83, 0x2d, 0x8b, 0xdc,
	0xd0, 0x0b, 0x62, 0x1e, 0x8a, 0xbd, 0xee, 0x03, 0xc8, 0x9a, 0xeb, 0xb2, 0x28, 0x42, 0xf7, 0x50,
	0xf2, 0xd9, 0x89, 0xf9, 0x6d, 0xa9, 0x23, 0xf5, 0xea, 0xc3, 0x86, 0x9a, 0x34, 0xa8, 0xc2, 0x24,
	0x89, 0x41, 0x85, 0xdf, 0x1d, 0x41, 0x79, 0xed, 0x45, 0xde, 0xc6, 0x67, 0xa8, 0x07, 0xa5, 0xc8,
	0xe5, 0x01, 0xcb, 0x32, 0x48, 0x64, 0x32, 0xd7, 0x4a, 0x1c, 0x2a, 0x16, 0xfa, 0x6f, 0x50, 0x39,
	0xab, 0x42, 0x35, 0xf8, 0x6f, 0x2f, 0x4c, 0x9b, 0x18, 0x73, 0x63, 0xa5, 0x14, 0x50, 0x13, 0x14,
	0x62, 0x3e, 0xdb, 0x9a, 0xae, 0x63, 0xcb, 0xb2, 0x09, 0x5e, 0x63, 0xa2, 0x6c, 0x51, 0x1b, 0xae,
	0xe7, 0xc6, 0x74, 0x4a, 0xf0, 0xa5, 0xf1, 0x25, 0xa1, 0x16, 0x34, 0x66, 0xc6, 0xe3, 0xec, 0x52,
	0xff, 0x2c, 0xa2, 0x26, 0x5c, 0x59, 0x98, 0xae, 0x31, 0xb5, 0x8d, 0xc5, 0x0a, 0xd3, 0x85, 0x46,
	0x94, 0xf7, 0x3f, 0x7d, 0x1d, 0xaa, 0xe7, 0x27, 0x21, 0x05, 0xaa, 0xcb, 0xa7, 0x09, 0x31, 0x74,
	0xdb, 0xd2, 0xcd, 0x25, 0x56, 0x0a, 0x08, 0x41, 0x3d, 0x0f, 0x64, 0x9a, 0x94, 0x9c, 0xa8, 0x91,
	0x7c, 0x2c, 0x8e, 0x31, 0xc8, 0x8e, 0x00, 0x75, 0xab, 0x0a, 0xaa, 0x6a, 0x4e, 0x55, 0x9d, 0xb3,
	0x78, 0xcf, 0xb7, 0x66, 0x90, 0x3e, 0xa0, 0xfd, 0xf1, 0xaf, 0x23, 0xf5, 0x2a, 0xc3, 0xea, 0x39,
	0x40, 0x9a, 0x85, 0xc7, 0x06, 0x94, 0x4f, 0x19, 0xbc, 0xbb, 0x5f, 0x3d, 0x16, 0x0b, 0x4f, 0x9e,
	0xcb, 0x7e, 0x14, 0xd5, 0x2e, 0xa8, 0xd2, 0x3c, 0x3f, 0x91, 0x5f, 0xd2, 0x27, 0x6f, 0xe4, 0x34,
	0x3f, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x10, 0x17, 0x0d, 0xc9, 0x0b, 0x02, 0x00, 0x00,
}
