// Code generated by protoc-gen-go. DO NOT EDIT.
// source: port_exp.proto

package gen

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

//
// Top-level message
//
type PortExp struct {
	InterfaceExpStats    []*InterfaceExpInfos `protobuf:"bytes,1,rep,name=interfaceExp_stats,json=interfaceExpStats" json:"interfaceExp_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PortExp) Reset()         { *m = PortExp{} }
func (m *PortExp) String() string { return proto.CompactTextString(m) }
func (*PortExp) ProtoMessage()    {}
func (*PortExp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd3f2d611fd16ff, []int{0}
}

func (m *PortExp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortExp.Unmarshal(m, b)
}
func (m *PortExp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortExp.Marshal(b, m, deterministic)
}
func (m *PortExp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortExp.Merge(m, src)
}
func (m *PortExp) XXX_Size() int {
	return xxx_messageInfo_PortExp.Size(m)
}
func (m *PortExp) XXX_DiscardUnknown() {
	xxx_messageInfo_PortExp.DiscardUnknown(m)
}

var xxx_messageInfo_PortExp proto.InternalMessageInfo

func (m *PortExp) GetInterfaceExpStats() []*InterfaceExpInfos {
	if m != nil {
		return m.InterfaceExpStats
	}
	return nil
}

//
// Interface information
//
type InterfaceExpInfos struct {
	// Interface name, e.g., et-0/0/0
	IfName *string `protobuf:"bytes,1,req,name=if_name,json=ifName" json:"if_name,omitempty"`
	// Interface operational status
	IfOperationalStatus  *string  `protobuf:"bytes,2,opt,name=if_operational_status,json=ifOperationalStatus" json:"if_operational_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceExpInfos) Reset()         { *m = InterfaceExpInfos{} }
func (m *InterfaceExpInfos) String() string { return proto.CompactTextString(m) }
func (*InterfaceExpInfos) ProtoMessage()    {}
func (*InterfaceExpInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd3f2d611fd16ff, []int{1}
}

func (m *InterfaceExpInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceExpInfos.Unmarshal(m, b)
}
func (m *InterfaceExpInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceExpInfos.Marshal(b, m, deterministic)
}
func (m *InterfaceExpInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceExpInfos.Merge(m, src)
}
func (m *InterfaceExpInfos) XXX_Size() int {
	return xxx_messageInfo_InterfaceExpInfos.Size(m)
}
func (m *InterfaceExpInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceExpInfos.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceExpInfos proto.InternalMessageInfo

func (m *InterfaceExpInfos) GetIfName() string {
	if m != nil && m.IfName != nil {
		return *m.IfName
	}
	return ""
}

func (m *InterfaceExpInfos) GetIfOperationalStatus() string {
	if m != nil && m.IfOperationalStatus != nil {
		return *m.IfOperationalStatus
	}
	return ""
}

var E_JnprInterfaceExpExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*PortExp)(nil),
	Field:         18,
	Name:          "jnpr_interface_exp_ext",
	Tag:           "bytes,18,opt,name=jnpr_interface_exp_ext",
	Filename:      "port_exp.proto",
}

func init() {
	proto.RegisterType((*PortExp)(nil), "Port_exp")
	proto.RegisterType((*InterfaceExpInfos)(nil), "InterfaceExpInfos")
	proto.RegisterExtension(E_JnprInterfaceExpExt)
}

func init() { proto.RegisterFile("port_exp.proto", fileDescriptor_dbd3f2d611fd16ff) }

var fileDescriptor_dbd3f2d611fd16ff = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0xd9, 0x88, 0x7a, 0xd9, 0x03, 0xe1, 0x36, 0xa8, 0xc1, 0x42, 0xc2, 0x55, 0xa9, 0x52,
	0xa4, 0xb4, 0x52, 0xe1, 0x8a, 0x13, 0x3c, 0x25, 0x29, 0x2c, 0x97, 0x45, 0x66, 0x65, 0xf5, 0xb2,
	0xb3, 0xcc, 0xce, 0x61, 0x6c, 0xfd, 0xe5, 0x92, 0x23, 0x6a, 0xe0, 0xda, 0xf9, 0x3e, 0xe6, 0xbd,
	0x27, 0xcf, 0x02, 0x12, 0x6b, 0xe8, 0x43, 0x15, 0x08, 0x19, 0xaf, 0x32, 0x86, 0x2d, 0x74, 0xc0,
	0xf4, 0xa5, 0x19, 0xc7, 0xe3, 0xf2, 0x51, 0xce, 0x9e, 0x47, 0x4d, 0xdd, 0x49, 0xe5, 0x3c, 0x03,
	0x59, 0xf3, 0x0a, 0xab, 0x3e, 0xe8, 0xc8, 0x86, 0x63, 0x2e, 0x8a, 0xa3, 0x72, 0x5e, 0xab, 0x6a,
	0x3d, 0x41, 0x6b, 0x6f, 0x31, 0x36, 0x8b, 0xa9, 0xdd, 0x0e, 0xf2, 0xf2, 0x4d, 0x2e, 0x0e, 0x3c,
	0x75, 0x2d, 0x4f, 0x9d, 0xd5, 0xde, 0x74, 0x90, 0x8b, 0x22, 0x29, 0xd3, 0xfb, 0xe3, 0xef, 0xdb,
	0x64, 0x26, 0x9a, 0x13, 0x67, 0x37, 0xa6, 0x03, 0x55, 0xcb, 0x73, 0x67, 0x35, 0x06, 0x20, 0xc3,
	0x0e, 0xbd, 0xd9, 0xee, 0x93, 0x77, 0x31, 0x4f, 0x0a, 0x51, 0xa6, 0x4d, 0xe6, 0xec, 0xd3, 0x3f,
	0x6b, 0xf7, 0xe8, 0xe6, 0x45, 0x5e, 0xbc, 0xfb, 0x40, 0xfa, 0xaf, 0xc2, 0xb0, 0x40, 0x43, 0xcf,
	0xea, 0xb2, 0x7a, 0xd8, 0x79, 0x17, 0x80, 0x36, 0xc0, 0x9f, 0x48, 0x1f, 0xb1, 0x05, 0x1f, 0x91,
	0x62, 0xae, 0x0a, 0x51, 0xce, 0xeb, 0xb4, 0xfa, 0xdd, 0xdb, 0x64, 0xc3, 0x87, 0x69, 0xdd, 0x55,
	0xcf, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x47, 0xa5, 0xae, 0xab, 0x36, 0x01, 0x00, 0x00,
}
