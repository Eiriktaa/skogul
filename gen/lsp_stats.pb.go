// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lsp_stats.proto

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
type LspStats struct {
	// List of LSP statistics records
	LspStatsRecords      []*LspStatsRecord `protobuf:"bytes,1,rep,name=lsp_stats_records,json=lspStatsRecords" json:"lsp_stats_records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LspStats) Reset()         { *m = LspStats{} }
func (m *LspStats) String() string { return proto.CompactTextString(m) }
func (*LspStats) ProtoMessage()    {}
func (*LspStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc5e73ce6e34840, []int{0}
}

func (m *LspStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LspStats.Unmarshal(m, b)
}
func (m *LspStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LspStats.Marshal(b, m, deterministic)
}
func (m *LspStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LspStats.Merge(m, src)
}
func (m *LspStats) XXX_Size() int {
	return xxx_messageInfo_LspStats.Size(m)
}
func (m *LspStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LspStats.DiscardUnknown(m)
}

var xxx_messageInfo_LspStats proto.InternalMessageInfo

func (m *LspStats) GetLspStatsRecords() []*LspStatsRecord {
	if m != nil {
		return m.LspStatsRecords
	}
	return nil
}

//
// LSP statistics record
//
type LspStatsRecord struct {
	// Name of the LSP
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Instance Identifier for cases when RPD creates multiple instances
	InstanceIdentifier *uint32 `protobuf:"varint,2,req,name=instance_identifier,json=instanceIdentifier" json:"instance_identifier,omitempty"`
	// Name of the counter. This is useful when an LSP has multiple counters.
	// When an LSP is resignalled, it is possible that a new counter is
	// created in the hardware.
	CounterName *string `protobuf:"bytes,3,req,name=counter_name,json=counterName" json:"counter_name,omitempty"`
	// The total number of packets
	Packets *uint64 `protobuf:"varint,4,opt,name=packets" json:"packets,omitempty"`
	// The total number of bytes
	Bytes *uint64 `protobuf:"varint,5,opt,name=bytes" json:"bytes,omitempty"`
	// Packet rate computed over the most recent 3 second interval
	PacketRate *uint64 `protobuf:"varint,6,opt,name=packet_rate,json=packetRate" json:"packet_rate,omitempty"`
	// Byte rate computed over the most recent 3 second interval
	ByteRate             *uint64  `protobuf:"varint,7,opt,name=byte_rate,json=byteRate" json:"byte_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LspStatsRecord) Reset()         { *m = LspStatsRecord{} }
func (m *LspStatsRecord) String() string { return proto.CompactTextString(m) }
func (*LspStatsRecord) ProtoMessage()    {}
func (*LspStatsRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_edc5e73ce6e34840, []int{1}
}

func (m *LspStatsRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LspStatsRecord.Unmarshal(m, b)
}
func (m *LspStatsRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LspStatsRecord.Marshal(b, m, deterministic)
}
func (m *LspStatsRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LspStatsRecord.Merge(m, src)
}
func (m *LspStatsRecord) XXX_Size() int {
	return xxx_messageInfo_LspStatsRecord.Size(m)
}
func (m *LspStatsRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LspStatsRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LspStatsRecord proto.InternalMessageInfo

func (m *LspStatsRecord) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LspStatsRecord) GetInstanceIdentifier() uint32 {
	if m != nil && m.InstanceIdentifier != nil {
		return *m.InstanceIdentifier
	}
	return 0
}

func (m *LspStatsRecord) GetCounterName() string {
	if m != nil && m.CounterName != nil {
		return *m.CounterName
	}
	return ""
}

func (m *LspStatsRecord) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *LspStatsRecord) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *LspStatsRecord) GetPacketRate() uint64 {
	if m != nil && m.PacketRate != nil {
		return *m.PacketRate
	}
	return 0
}

func (m *LspStatsRecord) GetByteRate() uint64 {
	if m != nil && m.ByteRate != nil {
		return *m.ByteRate
	}
	return 0
}

var E_JnprLspStatisticsExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*LspStats)(nil),
	Field:         5,
	Name:          "jnpr_lsp_statistics_ext",
	Tag:           "bytes,5,opt,name=jnpr_lsp_statistics_ext",
	Filename:      "lsp_stats.proto",
}

func init() {
	proto.RegisterType((*LspStats)(nil), "LspStats")
	proto.RegisterType((*LspStatsRecord)(nil), "LspStatsRecord")
	proto.RegisterExtension(E_JnprLspStatisticsExt)
}

func init() { proto.RegisterFile("lsp_stats.proto", fileDescriptor_edc5e73ce6e34840) }

var fileDescriptor_edc5e73ce6e34840 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4f, 0x02, 0x31,
	0x14, 0xc6, 0x73, 0x07, 0x08, 0xf4, 0x54, 0x62, 0x31, 0xa1, 0xca, 0xc0, 0x85, 0xe9, 0x26, 0x06,
	0x06, 0x07, 0x5d, 0x8c, 0x89, 0x31, 0x1a, 0xc3, 0x70, 0x4c, 0x4e, 0xcd, 0x79, 0x3c, 0x93, 0x0a,
	0xb4, 0x4d, 0xdf, 0x23, 0xc2, 0xea, 0xe6, 0x7f, 0x6d, 0x7a, 0xe5, 0x10, 0xc7, 0xf7, 0xfb, 0x7e,
	0xfd, 0xfa, 0xf2, 0x58, 0x6f, 0x85, 0x56, 0x22, 0x15, 0x84, 0x13, 0xeb, 0x0c, 0x99, 0xeb, 0x3e,
	0xc1, 0x0a, 0xd6, 0x40, 0x6e, 0x27, 0xc9, 0xd8, 0x00, 0xc7, 0x4f, 0xac, 0xf3, 0x8a, 0x76, 0xee,
	0x35, 0x7e, 0xc7, 0x2e, 0x0e, 0x6f, 0xa4, 0x83, 0xd2, 0xb8, 0x05, 0x8a, 0x28, 0x6d, 0x64, 0xc9,
	0xb4, 0x37, 0xa9, 0xad, 0xbc, 0xe2, 0xb9, 0x6f, 0x3f, 0x9a, 0x71, 0xfc, 0x13, 0xb3, 0xf3, 0xff,
	0x0e, 0xbf, 0x62, 0x4d, 0x5d, 0xac, 0x41, 0x44, 0x69, 0x9c, 0x75, 0x1f, 0x5a, 0xdf, 0xf7, 0x71,
	0x27, 0xca, 0x2b, 0xc4, 0x6f, 0x58, 0x5f, 0x69, 0xa4, 0x42, 0x97, 0x20, 0xd5, 0x02, 0x34, 0xa9,
	0x0f, 0x05, 0x4e, 0xc4, 0x69, 0x9c, 0x9d, 0xd5, 0x26, 0xaf, 0x8d, 0xe7, 0x83, 0xc0, 0x33, 0x76,
	0x5a, 0x9a, 0x8d, 0x26, 0x70, 0xb2, 0xaa, 0x6e, 0x1c, 0x57, 0x27, 0xfb, 0x68, 0xe6, 0x7f, 0x18,
	0xb1, 0xb6, 0x2d, 0xca, 0x25, 0x10, 0x8a, 0x66, 0x1a, 0x65, 0xcd, 0x4a, 0x12, 0x51, 0x5e, 0x53,
	0x3e, 0x64, 0xad, 0xf7, 0x1d, 0x01, 0x8a, 0xd6, 0x71, 0x1c, 0x18, 0x1f, 0xb1, 0x24, 0x78, 0xd2,
	0x15, 0x04, 0xe2, 0xc4, 0x2b, 0x39, 0x0b, 0x28, 0x2f, 0x08, 0xf8, 0x90, 0x75, 0xbd, 0x19, 0xe2,
	0x76, 0x15, 0x77, 0x3c, 0xf0, 0xe1, 0xed, 0x1b, 0x1b, 0x7c, 0x6a, 0xeb, 0x64, 0x7d, 0x4d, 0x85,
	0xa4, 0x4a, 0x94, 0xb0, 0x25, 0x3e, 0x98, 0xbc, 0x6c, 0xb4, 0xb2, 0xe0, 0x66, 0x40, 0x5f, 0xc6,
	0x2d, 0x71, 0x0e, 0x1a, 0x8d, 0x0b, 0x5b, 0x24, 0xd3, 0xee, 0xdf, 0x9d, 0x2f, 0x7d, 0xc5, 0x7e,
	0x0a, 0x05, 0x8f, 0x5b, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x6c, 0xfb, 0xd8, 0xd6, 0x01,
	0x00, 0x00,
}
