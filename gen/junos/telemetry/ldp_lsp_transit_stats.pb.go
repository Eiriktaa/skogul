// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ldp_lsp_transit_stats.proto

package telemetry

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
// Top-level message
//
type LdpLspTransitStats struct {
	// List of LDP LSP stats record
	LdpLspTransitStatsRecords []*LdpLspTransitRecord `protobuf:"bytes,1,rep,name=ldp_lsp_transit_stats_records,json=ldpLspTransitStatsRecords" json:"ldp_lsp_transit_stats_records,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}               `json:"-"`
	XXX_unrecognized          []byte                 `json:"-"`
	XXX_sizecache             int32                  `json:"-"`
}

func (m *LdpLspTransitStats) Reset()         { *m = LdpLspTransitStats{} }
func (m *LdpLspTransitStats) String() string { return proto.CompactTextString(m) }
func (*LdpLspTransitStats) ProtoMessage()    {}
func (*LdpLspTransitStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b688fa85cbb319, []int{0}
}
func (m *LdpLspTransitStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpLspTransitStats.Unmarshal(m, b)
}
func (m *LdpLspTransitStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpLspTransitStats.Marshal(b, m, deterministic)
}
func (m *LdpLspTransitStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpLspTransitStats.Merge(m, src)
}
func (m *LdpLspTransitStats) XXX_Size() int {
	return xxx_messageInfo_LdpLspTransitStats.Size(m)
}
func (m *LdpLspTransitStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpLspTransitStats.DiscardUnknown(m)
}

var xxx_messageInfo_LdpLspTransitStats proto.InternalMessageInfo

func (m *LdpLspTransitStats) GetLdpLspTransitStatsRecords() []*LdpLspTransitRecord {
	if m != nil {
		return m.LdpLspTransitStatsRecords
	}
	return nil
}

//
// LDP LSP Transit statistics record
//
type LdpLspTransitRecord struct {
	// IP prefix
	IpPrefix *string `protobuf:"bytes,1,req,name=ip_prefix,json=ipPrefix" json:"ip_prefix,omitempty"`
	// Instance Identifier for cases when RPD creates multiple instances
	InstanceIdentifier *uint32 `protobuf:"varint,2,opt,name=instance_identifier,json=instanceIdentifier" json:"instance_identifier,omitempty"`
	// Name of the counter.
	CounterName *string `protobuf:"bytes,3,opt,name=counter_name,json=counterName" json:"counter_name,omitempty"`
	// Statistics
	Stats                *LabelDistributionProtocolLspTransitStats `protobuf:"bytes,4,opt,name=stats" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *LdpLspTransitRecord) Reset()         { *m = LdpLspTransitRecord{} }
func (m *LdpLspTransitRecord) String() string { return proto.CompactTextString(m) }
func (*LdpLspTransitRecord) ProtoMessage()    {}
func (*LdpLspTransitRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b688fa85cbb319, []int{1}
}
func (m *LdpLspTransitRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpLspTransitRecord.Unmarshal(m, b)
}
func (m *LdpLspTransitRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpLspTransitRecord.Marshal(b, m, deterministic)
}
func (m *LdpLspTransitRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpLspTransitRecord.Merge(m, src)
}
func (m *LdpLspTransitRecord) XXX_Size() int {
	return xxx_messageInfo_LdpLspTransitRecord.Size(m)
}
func (m *LdpLspTransitRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpLspTransitRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LdpLspTransitRecord proto.InternalMessageInfo

func (m *LdpLspTransitRecord) GetIpPrefix() string {
	if m != nil && m.IpPrefix != nil {
		return *m.IpPrefix
	}
	return ""
}

func (m *LdpLspTransitRecord) GetInstanceIdentifier() uint32 {
	if m != nil && m.InstanceIdentifier != nil {
		return *m.InstanceIdentifier
	}
	return 0
}

func (m *LdpLspTransitRecord) GetCounterName() string {
	if m != nil && m.CounterName != nil {
		return *m.CounterName
	}
	return ""
}

func (m *LdpLspTransitRecord) GetStats() *LabelDistributionProtocolLspTransitStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type LabelDistributionProtocolLspTransitStats struct {
	// Packet and Byte statistics
	Packets *uint64 `protobuf:"varint,1,opt,name=packets" json:"packets,omitempty"`
	Bytes   *uint64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	// Rates of the above counters
	PacketRate           *uint64  `protobuf:"varint,3,opt,name=packet_rate,json=packetRate" json:"packet_rate,omitempty"`
	ByteRate             *uint64  `protobuf:"varint,4,opt,name=byte_rate,json=byteRate" json:"byte_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelDistributionProtocolLspTransitStats) Reset() {
	*m = LabelDistributionProtocolLspTransitStats{}
}
func (m *LabelDistributionProtocolLspTransitStats) String() string { return proto.CompactTextString(m) }
func (*LabelDistributionProtocolLspTransitStats) ProtoMessage()    {}
func (*LabelDistributionProtocolLspTransitStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b688fa85cbb319, []int{2}
}
func (m *LabelDistributionProtocolLspTransitStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelDistributionProtocolLspTransitStats.Unmarshal(m, b)
}
func (m *LabelDistributionProtocolLspTransitStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelDistributionProtocolLspTransitStats.Marshal(b, m, deterministic)
}
func (m *LabelDistributionProtocolLspTransitStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelDistributionProtocolLspTransitStats.Merge(m, src)
}
func (m *LabelDistributionProtocolLspTransitStats) XXX_Size() int {
	return xxx_messageInfo_LabelDistributionProtocolLspTransitStats.Size(m)
}
func (m *LabelDistributionProtocolLspTransitStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelDistributionProtocolLspTransitStats.DiscardUnknown(m)
}

var xxx_messageInfo_LabelDistributionProtocolLspTransitStats proto.InternalMessageInfo

func (m *LabelDistributionProtocolLspTransitStats) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *LabelDistributionProtocolLspTransitStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *LabelDistributionProtocolLspTransitStats) GetPacketRate() uint64 {
	if m != nil && m.PacketRate != nil {
		return *m.PacketRate
	}
	return 0
}

func (m *LabelDistributionProtocolLspTransitStats) GetByteRate() uint64 {
	if m != nil && m.ByteRate != nil {
		return *m.ByteRate
	}
	return 0
}

var E_JnprLdpLspTransitStatsExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*LdpLspTransitStats)(nil),
	Field:         154,
	Name:          "jnpr_ldp_lsp_transit_stats_ext",
	Tag:           "bytes,154,opt,name=jnpr_ldp_lsp_transit_stats_ext",
	Filename:      "ldp_lsp_transit_stats.proto",
}

func init() {
	proto.RegisterType((*LdpLspTransitStats)(nil), "LdpLspTransitStats")
	proto.RegisterType((*LdpLspTransitRecord)(nil), "LdpLspTransitRecord")
	proto.RegisterType((*LabelDistributionProtocolLspTransitStats)(nil), "LabelDistributionProtocolLspTransitStats")
	proto.RegisterExtension(E_JnprLdpLspTransitStatsExt)
}

func init() { proto.RegisterFile("ldp_lsp_transit_stats.proto", fileDescriptor_a6b688fa85cbb319) }

var fileDescriptor_a6b688fa85cbb319 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xe5, 0xfe, 0x88, 0xd6, 0x81, 0x8d, 0x83, 0x44, 0x4a, 0x05, 0x44, 0x59, 0xa0, 0xb0,
	0xc9, 0xa2, 0x0b, 0x16, 0x6c, 0x40, 0x08, 0x16, 0xa0, 0xaa, 0xaa, 0x52, 0xc4, 0xd6, 0x72, 0x93,
	0x5b, 0xc9, 0xd3, 0xd4, 0xb6, 0xec, 0x5b, 0x4d, 0xbb, 0x9d, 0xc7, 0x98, 0xe7, 0x98, 0xc7, 0x99,
	0x87, 0x19, 0x39, 0x9e, 0x76, 0x7e, 0xda, 0xc5, 0x6c, 0xcf, 0xf9, 0xce, 0x55, 0xce, 0x89, 0xe9,
	0xb8, 0xa9, 0x0d, 0x6f, 0x9c, 0xe1, 0x68, 0x85, 0x72, 0x12, 0xb9, 0x43, 0x81, 0xae, 0x30, 0x56,
	0xa3, 0x7e, 0x1f, 0x23, 0x34, 0xb0, 0x01, 0xb4, 0x7b, 0x8e, 0xda, 0x04, 0x31, 0x6b, 0x28, 0x9b,
	0xd6, 0x66, 0xea, 0xcc, 0xbf, 0x90, 0x58, 0xf8, 0x00, 0xfb, 0x4f, 0x3f, 0x9c, 0xbd, 0xc4, 0x2d,
	0x54, 0xda, 0xd6, 0x2e, 0x21, 0x69, 0x37, 0x8f, 0x26, 0x6f, 0x8b, 0x27, 0xd9, 0xb2, 0x35, 0xcb,
	0x51, 0x73, 0x72, 0x30, 0x38, 0x2e, 0xbb, 0x25, 0x34, 0x3e, 0x13, 0x61, 0x19, 0x1d, 0x4a, 0xc3,
	0x8d, 0x85, 0x95, 0xdc, 0x25, 0x24, 0xed, 0xe4, 0xc3, 0x9f, 0xfd, 0xab, 0x1f, 0x9d, 0x01, 0x29,
	0x07, 0xd2, 0xcc, 0x5b, 0x99, 0x7d, 0xa5, 0xb1, 0x54, 0x0e, 0x85, 0xaa, 0x80, 0xcb, 0x1a, 0x14,
	0xca, 0x95, 0x04, 0x9b, 0x74, 0x52, 0x92, 0xbf, 0x39, 0xd0, 0xec, 0x40, 0xfc, 0x39, 0x02, 0x2c,
	0xa7, 0xaf, 0x2b, 0xbd, 0x55, 0x08, 0x96, 0x2b, 0xb1, 0x81, 0xa4, 0x9b, 0x92, 0x87, 0xf3, 0xd1,
	0xbd, 0x35, 0x13, 0x1b, 0x60, 0xdf, 0x69, 0xbf, 0x6d, 0x99, 0xf4, 0x52, 0x92, 0x47, 0x93, 0x2f,
	0xc5, 0x54, 0x2c, 0xa1, 0xf9, 0x25, 0x1d, 0x5a, 0xb9, 0xdc, 0xa2, 0xd4, 0x6a, 0xee, 0x37, 0xab,
	0x74, 0xf3, 0xbc, 0x5e, 0xc8, 0x65, 0x37, 0x84, 0xe6, 0x2f, 0xcd, 0xb0, 0x4f, 0xf4, 0x95, 0x11,
	0xd5, 0x1a, 0xd0, 0xaf, 0x49, 0xf2, 0x5e, 0xfb, 0x49, 0x09, 0x29, 0x0f, 0x2a, 0x1b, 0xd3, 0xfe,
	0x72, 0x8f, 0xe0, 0xda, 0x8a, 0x47, 0x3b, 0x68, 0xec, 0x33, 0x8d, 0x02, 0xc7, 0xad, 0xc0, 0x50,
	0x2a, 0x20, 0x29, 0x29, 0x69, 0x70, 0x4a, 0x81, 0xe0, 0x97, 0xf5, 0x81, 0x40, 0xf5, 0x1e, 0x53,
	0x03, 0xaf, 0x7b, 0xe6, 0x9b, 0xa1, 0x1f, 0x2f, 0x94, 0xb1, 0xfc, 0xfc, 0x2f, 0x87, 0x1d, 0xb2,
	0x77, 0xc5, 0xdf, 0xad, 0x92, 0x06, 0xec, 0x0c, 0xf0, 0x52, 0xdb, 0xb5, 0x5b, 0x80, 0x72, 0xda,
	0xba, 0xe4, 0x9a, 0xb4, 0x53, 0xc5, 0xc5, 0xe9, 0x23, 0x2a, 0x47, 0xfe, 0xe8, 0xa9, 0xfe, 0x7b,
	0x87, 0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xf9, 0x28, 0xf8, 0xa8, 0x02, 0x00, 0x00,
}