// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logical_port.proto

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
type LogicalPort struct {
	InterfaceInfo        []*LogicalInterfaceInfo `protobuf:"bytes,1,rep,name=interface_info,json=interfaceInfo" json:"interface_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LogicalPort) Reset()         { *m = LogicalPort{} }
func (m *LogicalPort) String() string { return proto.CompactTextString(m) }
func (*LogicalPort) ProtoMessage()    {}
func (*LogicalPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{0}
}

func (m *LogicalPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPort.Unmarshal(m, b)
}
func (m *LogicalPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPort.Marshal(b, m, deterministic)
}
func (m *LogicalPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPort.Merge(m, src)
}
func (m *LogicalPort) XXX_Size() int {
	return xxx_messageInfo_LogicalPort.Size(m)
}
func (m *LogicalPort) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPort.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPort proto.InternalMessageInfo

func (m *LogicalPort) GetInterfaceInfo() []*LogicalInterfaceInfo {
	if m != nil {
		return m.InterfaceInfo
	}
	return nil
}

//
// Logical Interaface information
//
type LogicalInterfaceInfo struct {
	// Logical interface name (e.g. xe-0/0/0.0)
	IfName *string `protobuf:"bytes,1,req,name=if_name,json=ifName" json:"if_name,omitempty"`
	// Time reset
	InitTime *uint64 `protobuf:"varint,2,req,name=init_time,json=initTime" json:"init_time,omitempty"`
	// Global Index
	SnmpIfIndex *uint32 `protobuf:"varint,3,opt,name=snmp_if_index,json=snmpIfIndex" json:"snmp_if_index,omitempty"`
	// Name of the aggregate bundle
	ParentAeName *string `protobuf:"bytes,4,opt,name=parent_ae_name,json=parentAeName" json:"parent_ae_name,omitempty"`
	// Inbound traffic statistics
	IngressStats *IngressInterfaceStats `protobuf:"bytes,5,opt,name=ingress_stats,json=ingressStats" json:"ingress_stats,omitempty"`
	// Outbound traffic statistics
	EgressStats *EgressInterfaceStats `protobuf:"bytes,6,opt,name=egress_stats,json=egressStats" json:"egress_stats,omitempty"`
	// Link state UP\DOWN etc.
	OpState *OperationalState `protobuf:"bytes,7,opt,name=op_state,json=opState" json:"op_state,omitempty"`
	// administrative status, i.e.. enabled/disabled
	AdministractiveStatus *string `protobuf:"bytes,8,opt,name=administractive_status,json=administractiveStatus" json:"administractive_status,omitempty"`
	// Description of the interface
	Description *string `protobuf:"bytes,9,opt,name=description" json:"description,omitempty"`
	// This corresponds to the ifLastChange object in the standard interface MIB
	LastChange *uint32 `protobuf:"varint,10,opt,name=last_change,json=lastChange" json:"last_change,omitempty"`
	// This corresponds to the ifHighSpeed object in the standard interface MIB
	HighSpeed *uint32 `protobuf:"varint,11,opt,name=high_speed,json=highSpeed" json:"high_speed,omitempty"`
	// Ingress queue information
	IngressQueueInfo []*LogicalInterfaceQueueStats `protobuf:"bytes,12,rep,name=ingress_queue_info,json=ingressQueueInfo" json:"ingress_queue_info,omitempty"`
	// Egress queue information
	EgressQueueInfo      []*LogicalInterfaceQueueStats `protobuf:"bytes,13,rep,name=egress_queue_info,json=egressQueueInfo" json:"egress_queue_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *LogicalInterfaceInfo) Reset()         { *m = LogicalInterfaceInfo{} }
func (m *LogicalInterfaceInfo) String() string { return proto.CompactTextString(m) }
func (*LogicalInterfaceInfo) ProtoMessage()    {}
func (*LogicalInterfaceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{1}
}

func (m *LogicalInterfaceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalInterfaceInfo.Unmarshal(m, b)
}
func (m *LogicalInterfaceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalInterfaceInfo.Marshal(b, m, deterministic)
}
func (m *LogicalInterfaceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalInterfaceInfo.Merge(m, src)
}
func (m *LogicalInterfaceInfo) XXX_Size() int {
	return xxx_messageInfo_LogicalInterfaceInfo.Size(m)
}
func (m *LogicalInterfaceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalInterfaceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalInterfaceInfo proto.InternalMessageInfo

func (m *LogicalInterfaceInfo) GetIfName() string {
	if m != nil && m.IfName != nil {
		return *m.IfName
	}
	return ""
}

func (m *LogicalInterfaceInfo) GetInitTime() uint64 {
	if m != nil && m.InitTime != nil {
		return *m.InitTime
	}
	return 0
}

func (m *LogicalInterfaceInfo) GetSnmpIfIndex() uint32 {
	if m != nil && m.SnmpIfIndex != nil {
		return *m.SnmpIfIndex
	}
	return 0
}

func (m *LogicalInterfaceInfo) GetParentAeName() string {
	if m != nil && m.ParentAeName != nil {
		return *m.ParentAeName
	}
	return ""
}

func (m *LogicalInterfaceInfo) GetIngressStats() *IngressInterfaceStats {
	if m != nil {
		return m.IngressStats
	}
	return nil
}

func (m *LogicalInterfaceInfo) GetEgressStats() *EgressInterfaceStats {
	if m != nil {
		return m.EgressStats
	}
	return nil
}

func (m *LogicalInterfaceInfo) GetOpState() *OperationalState {
	if m != nil {
		return m.OpState
	}
	return nil
}

func (m *LogicalInterfaceInfo) GetAdministractiveStatus() string {
	if m != nil && m.AdministractiveStatus != nil {
		return *m.AdministractiveStatus
	}
	return ""
}

func (m *LogicalInterfaceInfo) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *LogicalInterfaceInfo) GetLastChange() uint32 {
	if m != nil && m.LastChange != nil {
		return *m.LastChange
	}
	return 0
}

func (m *LogicalInterfaceInfo) GetHighSpeed() uint32 {
	if m != nil && m.HighSpeed != nil {
		return *m.HighSpeed
	}
	return 0
}

func (m *LogicalInterfaceInfo) GetIngressQueueInfo() []*LogicalInterfaceQueueStats {
	if m != nil {
		return m.IngressQueueInfo
	}
	return nil
}

func (m *LogicalInterfaceInfo) GetEgressQueueInfo() []*LogicalInterfaceQueueStats {
	if m != nil {
		return m.EgressQueueInfo
	}
	return nil
}

//
//  Interface inbound/Ingress traffic statistics
//
type IngressInterfaceStats struct {
	// Count of packets
	IfPackets *uint64 `protobuf:"varint,1,req,name=if_packets,json=ifPackets" json:"if_packets,omitempty"`
	// Count of bytes
	IfOctets *uint64 `protobuf:"varint,2,req,name=if_octets,json=ifOctets" json:"if_octets,omitempty"`
	// Count of unicast packets
	IfUcastPackets *uint64 `protobuf:"varint,3,opt,name=if_ucast_packets,json=ifUcastPackets" json:"if_ucast_packets,omitempty"`
	// Count of multicast packets
	IfMcastPackets       *uint64                      `protobuf:"varint,4,opt,name=if_mcast_packets,json=ifMcastPackets" json:"if_mcast_packets,omitempty"`
	IfFcStats            []*ForwardingClassAccounting `protobuf:"bytes,5,rep,name=if_fc_stats,json=ifFcStats" json:"if_fc_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *IngressInterfaceStats) Reset()         { *m = IngressInterfaceStats{} }
func (m *IngressInterfaceStats) String() string { return proto.CompactTextString(m) }
func (*IngressInterfaceStats) ProtoMessage()    {}
func (*IngressInterfaceStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{2}
}

func (m *IngressInterfaceStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressInterfaceStats.Unmarshal(m, b)
}
func (m *IngressInterfaceStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressInterfaceStats.Marshal(b, m, deterministic)
}
func (m *IngressInterfaceStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressInterfaceStats.Merge(m, src)
}
func (m *IngressInterfaceStats) XXX_Size() int {
	return xxx_messageInfo_IngressInterfaceStats.Size(m)
}
func (m *IngressInterfaceStats) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressInterfaceStats.DiscardUnknown(m)
}

var xxx_messageInfo_IngressInterfaceStats proto.InternalMessageInfo

func (m *IngressInterfaceStats) GetIfPackets() uint64 {
	if m != nil && m.IfPackets != nil {
		return *m.IfPackets
	}
	return 0
}

func (m *IngressInterfaceStats) GetIfOctets() uint64 {
	if m != nil && m.IfOctets != nil {
		return *m.IfOctets
	}
	return 0
}

func (m *IngressInterfaceStats) GetIfUcastPackets() uint64 {
	if m != nil && m.IfUcastPackets != nil {
		return *m.IfUcastPackets
	}
	return 0
}

func (m *IngressInterfaceStats) GetIfMcastPackets() uint64 {
	if m != nil && m.IfMcastPackets != nil {
		return *m.IfMcastPackets
	}
	return 0
}

func (m *IngressInterfaceStats) GetIfFcStats() []*ForwardingClassAccounting {
	if m != nil {
		return m.IfFcStats
	}
	return nil
}

//
//  Interface outbound/Egress traffic statistics
//
type EgressInterfaceStats struct {
	// Count of packets
	IfPackets *uint64 `protobuf:"varint,1,req,name=if_packets,json=ifPackets" json:"if_packets,omitempty"`
	// Count of bytes
	IfOctets *uint64 `protobuf:"varint,2,req,name=if_octets,json=ifOctets" json:"if_octets,omitempty"`
	// Count of unicast packets
	IfUcastPackets *uint64 `protobuf:"varint,3,opt,name=if_ucast_packets,json=ifUcastPackets" json:"if_ucast_packets,omitempty"`
	// Count of multicast packets
	IfMcastPackets       *uint64  `protobuf:"varint,4,opt,name=if_mcast_packets,json=ifMcastPackets" json:"if_mcast_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EgressInterfaceStats) Reset()         { *m = EgressInterfaceStats{} }
func (m *EgressInterfaceStats) String() string { return proto.CompactTextString(m) }
func (*EgressInterfaceStats) ProtoMessage()    {}
func (*EgressInterfaceStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{3}
}

func (m *EgressInterfaceStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EgressInterfaceStats.Unmarshal(m, b)
}
func (m *EgressInterfaceStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EgressInterfaceStats.Marshal(b, m, deterministic)
}
func (m *EgressInterfaceStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EgressInterfaceStats.Merge(m, src)
}
func (m *EgressInterfaceStats) XXX_Size() int {
	return xxx_messageInfo_EgressInterfaceStats.Size(m)
}
func (m *EgressInterfaceStats) XXX_DiscardUnknown() {
	xxx_messageInfo_EgressInterfaceStats.DiscardUnknown(m)
}

var xxx_messageInfo_EgressInterfaceStats proto.InternalMessageInfo

func (m *EgressInterfaceStats) GetIfPackets() uint64 {
	if m != nil && m.IfPackets != nil {
		return *m.IfPackets
	}
	return 0
}

func (m *EgressInterfaceStats) GetIfOctets() uint64 {
	if m != nil && m.IfOctets != nil {
		return *m.IfOctets
	}
	return 0
}

func (m *EgressInterfaceStats) GetIfUcastPackets() uint64 {
	if m != nil && m.IfUcastPackets != nil {
		return *m.IfUcastPackets
	}
	return 0
}

func (m *EgressInterfaceStats) GetIfMcastPackets() uint64 {
	if m != nil && m.IfMcastPackets != nil {
		return *m.IfMcastPackets
	}
	return 0
}

//
//  Interface operational State details
//
type OperationalState struct {
	// If the link is up/down
	OperationalStatus    *string  `protobuf:"bytes,1,opt,name=operational_status,json=operationalStatus" json:"operational_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationalState) Reset()         { *m = OperationalState{} }
func (m *OperationalState) String() string { return proto.CompactTextString(m) }
func (*OperationalState) ProtoMessage()    {}
func (*OperationalState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{4}
}

func (m *OperationalState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationalState.Unmarshal(m, b)
}
func (m *OperationalState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationalState.Marshal(b, m, deterministic)
}
func (m *OperationalState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationalState.Merge(m, src)
}
func (m *OperationalState) XXX_Size() int {
	return xxx_messageInfo_OperationalState.Size(m)
}
func (m *OperationalState) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationalState.DiscardUnknown(m)
}

var xxx_messageInfo_OperationalState proto.InternalMessageInfo

func (m *OperationalState) GetOperationalStatus() string {
	if m != nil && m.OperationalStatus != nil {
		return *m.OperationalStatus
	}
	return ""
}

//
//  Interface forwarding class accounting
//
type ForwardingClassAccounting struct {
	// Interface protocol
	IfFamily *string `protobuf:"bytes,1,opt,name=if_family,json=ifFamily" json:"if_family,omitempty"`
	// Forwarding class number
	FcNumber *uint32 `protobuf:"varint,2,opt,name=fc_number,json=fcNumber" json:"fc_number,omitempty"`
	// Count of packets
	IfPackets *uint64 `protobuf:"varint,3,opt,name=if_packets,json=ifPackets" json:"if_packets,omitempty"`
	// Count of bytes
	IfOctets             *uint64  `protobuf:"varint,4,opt,name=if_octets,json=ifOctets" json:"if_octets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwardingClassAccounting) Reset()         { *m = ForwardingClassAccounting{} }
func (m *ForwardingClassAccounting) String() string { return proto.CompactTextString(m) }
func (*ForwardingClassAccounting) ProtoMessage()    {}
func (*ForwardingClassAccounting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{5}
}

func (m *ForwardingClassAccounting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardingClassAccounting.Unmarshal(m, b)
}
func (m *ForwardingClassAccounting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardingClassAccounting.Marshal(b, m, deterministic)
}
func (m *ForwardingClassAccounting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardingClassAccounting.Merge(m, src)
}
func (m *ForwardingClassAccounting) XXX_Size() int {
	return xxx_messageInfo_ForwardingClassAccounting.Size(m)
}
func (m *ForwardingClassAccounting) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardingClassAccounting.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardingClassAccounting proto.InternalMessageInfo

func (m *ForwardingClassAccounting) GetIfFamily() string {
	if m != nil && m.IfFamily != nil {
		return *m.IfFamily
	}
	return ""
}

func (m *ForwardingClassAccounting) GetFcNumber() uint32 {
	if m != nil && m.FcNumber != nil {
		return *m.FcNumber
	}
	return 0
}

func (m *ForwardingClassAccounting) GetIfPackets() uint64 {
	if m != nil && m.IfPackets != nil {
		return *m.IfPackets
	}
	return 0
}

func (m *ForwardingClassAccounting) GetIfOctets() uint64 {
	if m != nil && m.IfOctets != nil {
		return *m.IfOctets
	}
	return 0
}

//
// Interface queue statistics
//
type LogicalInterfaceQueueStats struct {
	// Queue number
	QueueNumber *uint32 `protobuf:"varint,1,opt,name=queue_number,json=queueNumber" json:"queue_number,omitempty"`
	// The total number of packets that have been added to this queue
	Packets *uint64 `protobuf:"varint,2,opt,name=packets" json:"packets,omitempty"`
	// The total number of bytes that have been added to this queue
	Bytes *uint64 `protobuf:"varint,3,opt,name=bytes" json:"bytes,omitempty"`
	// The total number of tail dropped packets
	TailDropPackets *uint64 `protobuf:"varint,4,opt,name=tail_drop_packets,json=tailDropPackets" json:"tail_drop_packets,omitempty"`
	// The total number of rate-limited packets
	RateLimitDropPackets *uint64 `protobuf:"varint,5,opt,name=rate_limit_drop_packets,json=rateLimitDropPackets" json:"rate_limit_drop_packets,omitempty"`
	// The total number of rate-limited bytes
	RateLimitDropBytes *uint64 `protobuf:"varint,6,opt,name=rate_limit_drop_bytes,json=rateLimitDropBytes" json:"rate_limit_drop_bytes,omitempty"`
	// The total number of red-dropped packets
	RedDropPackets *uint64 `protobuf:"varint,7,opt,name=red_drop_packets,json=redDropPackets" json:"red_drop_packets,omitempty"`
	// The total number of red-dropped bytes
	RedDropBytes *uint64 `protobuf:"varint,8,opt,name=red_drop_bytes,json=redDropBytes" json:"red_drop_bytes,omitempty"`
	// Average queue depth, in packets
	AverageBufferOccupancy *uint64 `protobuf:"varint,9,opt,name=average_buffer_occupancy,json=averageBufferOccupancy" json:"average_buffer_occupancy,omitempty"`
	// Current queue depth, in packets
	CurrentBufferOccupancy *uint64 `protobuf:"varint,10,opt,name=current_buffer_occupancy,json=currentBufferOccupancy" json:"current_buffer_occupancy,omitempty"`
	// The max measured queue depth, in packets, across all measurements since boot
	PeakBufferOccupancy *uint64 `protobuf:"varint,11,opt,name=peak_buffer_occupancy,json=peakBufferOccupancy" json:"peak_buffer_occupancy,omitempty"`
	// Allocated buffer size
	AllocatedBufferSize  *uint64  `protobuf:"varint,12,opt,name=allocated_buffer_size,json=allocatedBufferSize" json:"allocated_buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogicalInterfaceQueueStats) Reset()         { *m = LogicalInterfaceQueueStats{} }
func (m *LogicalInterfaceQueueStats) String() string { return proto.CompactTextString(m) }
func (*LogicalInterfaceQueueStats) ProtoMessage()    {}
func (*LogicalInterfaceQueueStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed53654dcd9b9a05, []int{6}
}

func (m *LogicalInterfaceQueueStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalInterfaceQueueStats.Unmarshal(m, b)
}
func (m *LogicalInterfaceQueueStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalInterfaceQueueStats.Marshal(b, m, deterministic)
}
func (m *LogicalInterfaceQueueStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalInterfaceQueueStats.Merge(m, src)
}
func (m *LogicalInterfaceQueueStats) XXX_Size() int {
	return xxx_messageInfo_LogicalInterfaceQueueStats.Size(m)
}
func (m *LogicalInterfaceQueueStats) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalInterfaceQueueStats.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalInterfaceQueueStats proto.InternalMessageInfo

func (m *LogicalInterfaceQueueStats) GetQueueNumber() uint32 {
	if m != nil && m.QueueNumber != nil {
		return *m.QueueNumber
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetTailDropPackets() uint64 {
	if m != nil && m.TailDropPackets != nil {
		return *m.TailDropPackets
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetRateLimitDropPackets() uint64 {
	if m != nil && m.RateLimitDropPackets != nil {
		return *m.RateLimitDropPackets
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetRateLimitDropBytes() uint64 {
	if m != nil && m.RateLimitDropBytes != nil {
		return *m.RateLimitDropBytes
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetRedDropPackets() uint64 {
	if m != nil && m.RedDropPackets != nil {
		return *m.RedDropPackets
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetRedDropBytes() uint64 {
	if m != nil && m.RedDropBytes != nil {
		return *m.RedDropBytes
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetAverageBufferOccupancy() uint64 {
	if m != nil && m.AverageBufferOccupancy != nil {
		return *m.AverageBufferOccupancy
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetCurrentBufferOccupancy() uint64 {
	if m != nil && m.CurrentBufferOccupancy != nil {
		return *m.CurrentBufferOccupancy
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetPeakBufferOccupancy() uint64 {
	if m != nil && m.PeakBufferOccupancy != nil {
		return *m.PeakBufferOccupancy
	}
	return 0
}

func (m *LogicalInterfaceQueueStats) GetAllocatedBufferSize() uint64 {
	if m != nil && m.AllocatedBufferSize != nil {
		return *m.AllocatedBufferSize
	}
	return 0
}

var E_JnprLogicalInterfaceExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*LogicalPort)(nil),
	Field:         7,
	Name:          "jnprLogicalInterfaceExt",
	Tag:           "bytes,7,opt,name=jnprLogicalInterfaceExt",
	Filename:      "logical_port.proto",
}

func init() {
	proto.RegisterType((*LogicalPort)(nil), "LogicalPort")
	proto.RegisterType((*LogicalInterfaceInfo)(nil), "LogicalInterfaceInfo")
	proto.RegisterType((*IngressInterfaceStats)(nil), "IngressInterfaceStats")
	proto.RegisterType((*EgressInterfaceStats)(nil), "EgressInterfaceStats")
	proto.RegisterType((*OperationalState)(nil), "OperationalState")
	proto.RegisterType((*ForwardingClassAccounting)(nil), "ForwardingClassAccounting")
	proto.RegisterType((*LogicalInterfaceQueueStats)(nil), "logicalInterfaceQueueStats")
	proto.RegisterExtension(E_JnprLogicalInterfaceExt)
}

func init() { proto.RegisterFile("logical_port.proto", fileDescriptor_ed53654dcd9b9a05) }

var fileDescriptor_ed53654dcd9b9a05 = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x5d, 0x73, 0xdb, 0x44,
	0x14, 0x1d, 0xb9, 0x75, 0x62, 0x5f, 0xd9, 0xa9, 0xb3, 0x8d, 0x13, 0x91, 0x0e, 0xd4, 0xe3, 0xe1,
	0xc1, 0x4c, 0xc1, 0x1d, 0x3a, 0xc3, 0x4c, 0x29, 0x9d, 0x81, 0xa4, 0x24, 0x8c, 0xa1, 0x24, 0x45,
	0x86, 0xe7, 0x9d, 0x8d, 0x7c, 0xd7, 0x59, 0x22, 0xed, 0x8a, 0xd5, 0xaa, 0x6d, 0xfa, 0xc8, 0xc0,
	0xcf, 0xe1, 0x4f, 0xf0, 0xc7, 0x60, 0x76, 0x25, 0xb9, 0x92, 0xe3, 0x0c, 0xcf, 0xbc, 0x59, 0xe7,
	0x9e, 0x73, 0xee, 0x87, 0xae, 0x76, 0x0d, 0x24, 0x56, 0x4b, 0x11, 0xb1, 0x98, 0xa6, 0x4a, 0x9b,
	0x69, 0xaa, 0x95, 0x51, 0x87, 0xf7, 0x0d, 0xc6, 0x98, 0xa0, 0xd1, 0xd7, 0xd4, 0xa8, 0xb4, 0x00,
	0xc7, 0x3f, 0x80, 0xff, 0xb2, 0xa0, 0xbe, 0x52, 0xda, 0x90, 0xe7, 0xb0, 0x23, 0xa4, 0x41, 0xcd,
	0x59, 0x84, 0x54, 0x48, 0xae, 0x02, 0x6f, 0x74, 0x67, 0xe2, 0x3f, 0x19, 0x4e, 0x4b, 0xd6, 0xac,
	0x8a, 0xce, 0x24, 0x57, 0x61, 0x5f, 0xd4, 0x1f, 0xc7, 0x7f, 0xb4, 0x61, 0x6f, 0x13, 0x8f, 0x7c,
	0x04, 0xdb, 0x82, 0x53, 0xc9, 0x12, 0x0c, 0xbc, 0x51, 0x6b, 0xd2, 0x3d, 0x6e, 0xff, 0xfe, 0x4d,
	0xab, 0xe3, 0x85, 0x5b, 0x82, 0x9f, 0xb1, 0x04, 0xc9, 0x18, 0xba, 0x42, 0x0a, 0x43, 0x8d, 0x48,
	0x30, 0x68, 0x8d, 0x5a, 0x93, 0xbb, 0x8e, 0x31, 0xf0, 0xc2, 0x8e, 0xc5, 0x7f, 0x16, 0x09, 0x92,
	0x4f, 0xa0, 0x9f, 0xc9, 0x24, 0xa5, 0x82, 0x53, 0x21, 0x17, 0xf8, 0x36, 0xb8, 0x33, 0xf2, 0x26,
	0xfd, 0xca, 0xc9, 0xb7, 0xb1, 0x19, 0x9f, 0xd9, 0x08, 0x79, 0x04, 0x3b, 0x29, 0xd3, 0x28, 0x0d,
	0x65, 0x58, 0x64, 0xbd, 0x3b, 0xf2, 0xde, 0x67, 0xed, 0x15, 0xc1, 0x23, 0x74, 0xb9, 0xbf, 0x82,
	0xbe, 0x90, 0x4b, 0x8d, 0x59, 0x46, 0x33, 0xc3, 0x4c, 0x16, 0xb4, 0x47, 0xde, 0xc4, 0x7f, 0xb2,
	0x3f, 0x9d, 0x15, 0xe8, 0xaa, 0x93, 0xb9, 0x8d, 0x86, 0xbd, 0x92, 0xec, 0x9e, 0xc8, 0x53, 0xe8,
	0x61, 0x5d, 0xbb, 0xe5, 0xb4, 0xc3, 0xe9, 0xc9, 0x26, 0xa9, 0x8f, 0x35, 0xe5, 0xa7, 0xd0, 0x51,
	0xa9, 0x53, 0x61, 0xb0, 0xed, 0x54, 0xbb, 0xd3, 0xf3, 0x14, 0x35, 0x33, 0x42, 0x49, 0x16, 0x5b,
	0x12, 0x86, 0xdb, 0x2a, 0x75, 0x3f, 0xc8, 0x17, 0xb0, 0xcf, 0x16, 0x89, 0x90, 0x22, 0x33, 0x9a,
	0x45, 0x46, 0xbc, 0x46, 0x27, 0xcd, 0xb3, 0xa0, 0x63, 0x3b, 0x0b, 0x87, 0x6b, 0xd1, 0xb9, 0x0b,
	0x92, 0x11, 0xf8, 0x0b, 0xcc, 0x22, 0x2d, 0x52, 0xeb, 0x1a, 0x74, 0x1d, 0xb7, 0x0e, 0x91, 0x87,
	0xe0, 0xc7, 0x2c, 0x33, 0x34, 0xba, 0x64, 0x72, 0x89, 0x01, 0xd8, 0x99, 0x86, 0x60, 0xa1, 0x17,
	0x0e, 0x21, 0x1f, 0x02, 0x5c, 0x8a, 0xe5, 0x25, 0xcd, 0x52, 0xc4, 0x45, 0xe0, 0xbb, 0x78, 0xd7,
	0x22, 0x73, 0x0b, 0x90, 0x19, 0x90, 0x6a, 0x7a, 0xbf, 0xe5, 0x98, 0x97, 0x4b, 0xd3, 0x73, 0x4b,
	0xf3, 0x60, 0x1a, 0xaf, 0x2d, 0xc3, 0x4f, 0x96, 0x52, 0x0c, 0x63, 0x50, 0xca, 0x1c, 0xe4, 0x96,
	0xe4, 0x3b, 0xd8, 0xc5, 0x1b, 0x4e, 0xfd, 0xff, 0x76, 0xba, 0x87, 0x4d, 0xa3, 0xf1, 0x3f, 0x1e,
	0x0c, 0x37, 0xbe, 0x3c, 0xf2, 0x31, 0x80, 0xe0, 0x34, 0x65, 0xd1, 0x15, 0x9a, 0xcc, 0xad, 0x62,
	0xb1, 0x68, 0x81, 0x17, 0x76, 0x05, 0x7f, 0x55, 0xe0, 0x6e, 0x1b, 0x39, 0x55, 0x91, 0xb1, 0xa4,
	0x56, 0x9d, 0xd4, 0x11, 0xfc, 0xdc, 0xc1, 0xe4, 0x31, 0x0c, 0x04, 0xa7, 0x79, 0x64, 0x67, 0x57,
	0xf9, 0xd9, 0x85, 0x5c, 0x51, 0x77, 0x04, 0xff, 0xc5, 0x46, 0x2b, 0xd3, 0x42, 0x90, 0x34, 0x04,
	0x77, 0xd7, 0x04, 0x3f, 0xd6, 0x05, 0xcf, 0xc0, 0x17, 0x9c, 0xf2, 0x68, 0xb5, 0x95, 0x76, 0x10,
	0x87, 0xd3, 0x53, 0xa5, 0xdf, 0x30, 0xbd, 0x10, 0x72, 0xf9, 0x22, 0x66, 0x59, 0x76, 0x14, 0x45,
	0x2a, 0x97, 0x46, 0xc8, 0xa5, 0xed, 0xe0, 0x34, 0x72, 0x7d, 0x8e, 0xff, 0xf6, 0x60, 0xef, 0xe4,
	0x7f, 0x3e, 0x80, 0xf1, 0x11, 0x0c, 0xd6, 0x3f, 0x08, 0xf2, 0x19, 0x10, 0xf5, 0x1e, 0xab, 0xbe,
	0x01, 0xcf, 0xed, 0xf5, 0xae, 0x6a, 0xb2, 0xf3, 0x6c, 0xfc, 0x97, 0x07, 0x1f, 0xdc, 0x3a, 0xb0,
	0xb2, 0x4d, 0xce, 0x12, 0x11, 0x5f, 0x17, 0x1e, 0xd5, 0x09, 0xd1, 0x11, 0xfc, 0xd4, 0xc1, 0x96,
	0xc3, 0x23, 0x2a, 0xf3, 0xe4, 0x02, 0x75, 0xd0, 0xaa, 0x9f, 0x38, 0x1d, 0x1e, 0x9d, 0x39, 0x78,
	0x6d, 0xa8, 0x8d, 0x21, 0xdc, 0x36, 0xd4, 0x46, 0xe3, 0xab, 0xa1, 0x8e, 0xff, 0x6c, 0xc3, 0xe1,
	0xed, 0x9b, 0x4e, 0x26, 0xd0, 0x2b, 0x3e, 0x8d, 0xb2, 0x1e, 0xaf, 0x71, 0x02, 0xba, 0x50, 0x59,
	0xd2, 0x43, 0xd8, 0xae, 0xea, 0x69, 0xd5, 0x53, 0x55, 0x28, 0x79, 0x00, 0xed, 0x8b, 0x6b, 0x83,
	0x6b, 0xe5, 0x16, 0x18, 0xf9, 0x1c, 0x76, 0x0d, 0x13, 0x31, 0x5d, 0x68, 0x95, 0x6e, 0x7e, 0x57,
	0xf7, 0x6c, 0xfc, 0x5b, 0xad, 0xd2, 0xaa, 0xbb, 0xe7, 0x70, 0xa0, 0x99, 0x41, 0x1a, 0x8b, 0x44,
	0x98, 0xa6, 0xb0, 0x5d, 0x17, 0xee, 0x59, 0xd6, 0x4b, 0x4b, 0xaa, 0xab, 0x9f, 0xc2, 0x70, 0x5d,
	0x5d, 0x54, 0xb7, 0x55, 0xd7, 0x92, 0x86, 0xf6, 0xd8, 0x95, 0xfa, 0x18, 0x06, 0x1a, 0x17, 0xcd,
	0x84, 0xdb, 0x8d, 0xad, 0xd2, 0xb8, 0xa8, 0xa7, 0x7a, 0x04, 0x3b, 0x2b, 0x41, 0x91, 0xa3, 0x53,
	0xa7, 0xf7, 0x4a, 0x7a, 0xe1, 0xfe, 0x35, 0x04, 0xec, 0x35, 0x6a, 0xb6, 0x44, 0x7a, 0x91, 0x73,
	0x8e, 0x9a, 0xaa, 0x28, 0xca, 0x53, 0x26, 0xa3, 0x6b, 0x77, 0x98, 0x16, 0xb2, 0x91, 0x17, 0xee,
	0x97, 0xb4, 0x63, 0xc7, 0x3a, 0xaf, 0x48, 0xd6, 0x20, 0xca, 0xb5, 0xbb, 0x8a, 0x6e, 0x18, 0x40,
	0xc3, 0xa0, 0xa4, 0xad, 0x1b, 0x7c, 0x09, 0xc3, 0x14, 0xd9, 0xd5, 0x4d, 0xb5, 0x5f, 0x57, 0xdf,
	0xb7, 0x9c, 0x0d, 0x52, 0x16, 0xc7, 0x2a, 0x62, 0x06, 0x17, 0x95, 0x3e, 0x13, 0xef, 0x30, 0xe8,
	0x35, 0xa4, 0x2b, 0x4e, 0xa1, 0x9f, 0x8b, 0x77, 0xf8, 0x8c, 0xc2, 0xc1, 0xaf, 0x32, 0xd5, 0xeb,
	0x77, 0xf9, 0xc9, 0x5b, 0x43, 0x0e, 0xa6, 0xdf, 0xe7, 0x52, 0xa4, 0xa8, 0xcf, 0xd0, 0xbc, 0x51,
	0xfa, 0x2a, 0x9b, 0xa3, 0xcc, 0x94, 0xce, 0xca, 0x4b, 0xac, 0x37, 0xad, 0xfd, 0x9d, 0x08, 0x6f,
	0x73, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x48, 0x74, 0x81, 0xa0, 0x08, 0x00, 0x00,
}
