// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pfe_mpls_sr_sid_oc.proto

package telemetry

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

type MplsPfeMplsSrSid struct {
	SignalingProtocols   *MplsPfeMplsSrSidSignalingProtocolsType `protobuf:"bytes,151,opt,name=signaling_protocols,json=signalingProtocols" json:"signaling_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *MplsPfeMplsSrSid) Reset()         { *m = MplsPfeMplsSrSid{} }
func (m *MplsPfeMplsSrSid) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSrSid) ProtoMessage()    {}
func (*MplsPfeMplsSrSid) Descriptor() ([]byte, []int) {
	return fileDescriptor_596ce4231baccfbb, []int{0}
}

func (m *MplsPfeMplsSrSid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSid.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSid.Marshal(b, m, deterministic)
}
func (m *MplsPfeMplsSrSid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSid.Merge(m, src)
}
func (m *MplsPfeMplsSrSid) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSid.Size(m)
}
func (m *MplsPfeMplsSrSid) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSid.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSid proto.InternalMessageInfo

func (m *MplsPfeMplsSrSid) GetSignalingProtocols() *MplsPfeMplsSrSidSignalingProtocolsType {
	if m != nil {
		return m.SignalingProtocols
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsType struct {
	SegmentRouting       *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType `protobuf:"bytes,151,opt,name=segment_routing,json=segmentRouting" json:"segment_routing,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) String() string { return proto.CompactTextString(m) }
func (*MplsPfeMplsSrSidSignalingProtocolsType) ProtoMessage()    {}
func (*MplsPfeMplsSrSidSignalingProtocolsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_596ce4231baccfbb, []int{0, 0}
}

func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Marshal(b, m, deterministic)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Merge(m, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsType proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsType) GetSegmentRouting() *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType {
	if m != nil {
		return m.SegmentRouting
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType struct {
	AggregateSidCounters *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType `protobuf:"bytes,151,opt,name=aggregate_sid_counters,json=aggregateSidCounters" json:"aggregate_sid_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                          `json:"-"`
	XXX_unrecognized     []byte                                                                            `json:"-"`
	XXX_sizecache        int32                                                                             `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) ProtoMessage() {}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) Descriptor() ([]byte, []int) {
	return fileDescriptor_596ce4231baccfbb, []int{0, 0, 0}
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Marshal(b, m, deterministic)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Merge(m, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType) GetAggregateSidCounters() *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType {
	if m != nil {
		return m.AggregateSidCounters
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType struct {
	AggregateSidCounter  []*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList `protobuf:"bytes,151,rep,name=aggregate_sid_counter,json=aggregateSidCounter" json:"aggregate_sid_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                                   `json:"-"`
	XXX_unrecognized     []byte                                                                                                     `json:"-"`
	XXX_sizecache        int32                                                                                                      `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) ProtoMessage() {
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) Descriptor() ([]byte, []int) {
	return fileDescriptor_596ce4231baccfbb, []int{0, 0, 0, 0}
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Marshal(b, m, deterministic)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Merge(m, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType) GetAggregateSidCounter() []*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList {
	if m != nil {
		return m.AggregateSidCounter
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList struct {
	State                *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                                          `json:"-"`
	XXX_unrecognized     []byte                                                                                                            `json:"-"`
	XXX_sizecache        int32                                                                                                             `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) ProtoMessage() {
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) Descriptor() ([]byte, []int) {
	return fileDescriptor_596ce4231baccfbb, []int{0, 0, 0, 0, 0}
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Marshal(b, m, deterministic)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Merge(m, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList proto.InternalMessageInfo

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList) GetState() *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) Reset() {
	*m = MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType{}
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) ProtoMessage() {
}
func (*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_596ce4231baccfbb, []int{0, 0, 0, 0, 0, 0}
}

func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Unmarshal(m, b)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Marshal(b, m, deterministic)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Merge(m, src)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_Size() int {
	return xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.Size(m)
}
func (m *MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType proto.InternalMessageInfo

var E_JnprMplsPfeMplsSrSidExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*MplsPfeMplsSrSid)(nil),
	Field:         75,
	Name:          "jnpr_mpls_pfe_mpls_sr_sid_ext",
	Tag:           "bytes,75,opt,name=jnpr_mpls_pfe_mpls_sr_sid_ext",
	Filename:      "pfe_mpls_sr_sid_oc.proto",
}

func init() {
	proto.RegisterType((*MplsPfeMplsSrSid)(nil), "mpls_pfe_mpls_sr_sid")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type.aggregate_sid_counters_type")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterList)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type.aggregate_sid_counters_type.aggregate_sid_counter_list")
	proto.RegisterType((*MplsPfeMplsSrSidSignalingProtocolsTypeSegmentRoutingTypeAggregateSidCountersTypeAggregateSidCounterListStateType)(nil), "mpls_pfe_mpls_sr_sid.signaling_protocols_type.segment_routing_type.aggregate_sid_counters_type.aggregate_sid_counter_list.state_type")
	proto.RegisterExtension(E_JnprMplsPfeMplsSrSidExt)
}

func init() { proto.RegisterFile("pfe_mpls_sr_sid_oc.proto", fileDescriptor_596ce4231baccfbb) }

var fileDescriptor_596ce4231baccfbb = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x39, 0x62, 0x2c, 0x26, 0xa2, 0x70, 0x49, 0xcc, 0x71, 0x22, 0x88, 0x55, 0xaa, 0x2d,
	0x52, 0xda, 0x06, 0x1b, 0x45, 0x09, 0x77, 0x0f, 0xb0, 0x1e, 0xc9, 0x64, 0x59, 0xbd, 0xdb, 0x5d,
	0x76, 0x27, 0x98, 0xd8, 0xfb, 0x0a, 0xda, 0xe8, 0x33, 0xf8, 0x1e, 0x16, 0x76, 0x36, 0x3e, 0x8d,
	0xe4, 0x2e, 0x89, 0x10, 0x36, 0x82, 0xa0, 0x58, 0x1d, 0xfc, 0xff, 0xfc, 0x3b, 0xdf, 0x3f, 0x07,
	0x91, 0x19, 0x23, 0x2f, 0x4c, 0xee, 0xb8, 0xb3, 0xdc, 0xc9, 0x11, 0xd7, 0x43, 0x66, 0xac, 0x26,
	0x1d, 0x37, 0x09, 0x73, 0x2c, 0x90, 0xec, 0x8c, 0x93, 0x36, 0x95, 0x78, 0xfc, 0xbe, 0x0d, 0xad,
	0x72, 0x7a, 0x2d, 0x16, 0x5e, 0x41, 0xd3, 0x49, 0xa1, 0xb2, 0x5c, 0x2a, 0xc1, 0xcb, 0xd9, 0xa1,
	0xce, 0x5d, 0xf4, 0x18, 0x1c, 0x05, 0xdd, 0x46, 0x8f, 0x31, 0x5f, 0x88, 0x79, 0x12, 0x9c, 0x66,
	0x06, 0x93, 0x70, 0xe5, 0x0c, 0x96, 0x46, 0xfc, 0x51, 0x87, 0x68, 0x53, 0x20, 0x2c, 0x60, 0xcf,
	0xa1, 0x28, 0x50, 0x11, 0xb7, 0x7a, 0x42, 0x52, 0x89, 0xe5, 0xea, 0xfe, 0xcf, 0x56, 0xb3, 0xb5,
	0x67, 0x2a, 0x9e, 0xdd, 0x85, 0x9a, 0x54, 0x62, 0xfc, 0xb6, 0x05, 0x2d, 0xdf, 0x60, 0xf8, 0x10,
	0xc0, 0x7e, 0x26, 0x84, 0x45, 0x91, 0x11, 0x96, 0xf7, 0x1c, 0xea, 0x89, 0x22, 0xb4, 0xab, 0x53,
	0xf0, 0x5f, 0xe0, 0x61, 0xfe, 0x15, 0x15, 0x6b, 0x6b, 0x65, 0xa6, 0x72, 0xd4, 0x5f, 0x58, 0xf1,
	0x53, 0x0d, 0x0e, 0xbe, 0x49, 0x85, 0x2f, 0x01, 0xb4, 0xbd, 0xfe, 0x9c, 0xbb, 0xd6, 0x6d, 0xf4,
	0xee, 0xfe, 0x98, 0xdb, 0xef, 0xf1, 0x5c, 0x3a, 0x4a, 0x9a, 0x9e, 0x4a, 0xf1, 0x6b, 0x00, 0xf1,
	0xe6, 0x4c, 0xf8, 0x1c, 0x40, 0xdd, 0x51, 0x46, 0xb8, 0x3c, 0xfc, 0x7d, 0xf0, 0x7f, 0x0d, 0x58,
	0x49, 0x52, 0xfd, 0x9f, 0x8a, 0x2a, 0xde, 0x01, 0xf8, 0x12, 0x4f, 0x14, 0x1c, 0x5e, 0x2b, 0x63,
	0xb9, 0x8f, 0x90, 0xe3, 0x94, 0xc2, 0x0e, 0x3b, 0x9b, 0x28, 0x69, 0xd0, 0x5e, 0x22, 0xdd, 0x6a,
	0x7b, 0xe3, 0x52, 0x54, 0x4e, 0x5b, 0x17, 0x9d, 0x97, 0xe5, 0xda, 0xde, 0x6e, 0x49, 0x67, 0xfe,
	0xe8, 0x85, 0xc9, 0xdd, 0x60, 0x8c, 0xf3, 0x4f, 0x6a, 0x53, 0x39, 0x3a, 0x9d, 0xd2, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x41, 0x6c, 0xd9, 0xc3, 0xf7, 0x03, 0x00, 0x00,
}
