// Code generated by protoc-gen-go. DO NOT EDIT.
// source: junos-xmlproxyd_junos-rsvp-interface.proto

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

type JunosRsvpInterface struct {
	RsvpInterfaceInformation *JunosRsvpInterfaceRsvpInterfaceInformationType `protobuf:"bytes,151,opt,name=rsvp_interface_information,json=rsvpInterfaceInformation" json:"rsvp_interface_information,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                                        `json:"-"`
	XXX_unrecognized         []byte                                          `json:"-"`
	XXX_sizecache            int32                                           `json:"-"`
}

func (m *JunosRsvpInterface) Reset()         { *m = JunosRsvpInterface{} }
func (m *JunosRsvpInterface) String() string { return proto.CompactTextString(m) }
func (*JunosRsvpInterface) ProtoMessage()    {}
func (*JunosRsvpInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4225f46c0f8328, []int{0}
}

func (m *JunosRsvpInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosRsvpInterface.Unmarshal(m, b)
}
func (m *JunosRsvpInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosRsvpInterface.Marshal(b, m, deterministic)
}
func (m *JunosRsvpInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosRsvpInterface.Merge(m, src)
}
func (m *JunosRsvpInterface) XXX_Size() int {
	return xxx_messageInfo_JunosRsvpInterface.Size(m)
}
func (m *JunosRsvpInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosRsvpInterface.DiscardUnknown(m)
}

var xxx_messageInfo_JunosRsvpInterface proto.InternalMessageInfo

func (m *JunosRsvpInterface) GetRsvpInterfaceInformation() *JunosRsvpInterfaceRsvpInterfaceInformationType {
	if m != nil {
		return m.RsvpInterfaceInformation
	}
	return nil
}

type JunosRsvpInterfaceRsvpInterfaceInformationType struct {
	ActiveCount          *uint32                                                            `protobuf:"varint,51,opt,name=active_count,json=activeCount" json:"active_count,omitempty"`
	RsvpInterface        []*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList `protobuf:"bytes,151,rep,name=rsvp_interface,json=rsvpInterface" json:"rsvp_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                           `json:"-"`
	XXX_unrecognized     []byte                                                             `json:"-"`
	XXX_sizecache        int32                                                              `json:"-"`
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) Reset() {
	*m = JunosRsvpInterfaceRsvpInterfaceInformationType{}
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationType) ProtoMessage() {}
func (*JunosRsvpInterfaceRsvpInterfaceInformationType) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4225f46c0f8328, []int{0, 0}
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationType.Unmarshal(m, b)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationType.Marshal(b, m, deterministic)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationType.Merge(m, src)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) XXX_Size() int {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationType.Size(m)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationType proto.InternalMessageInfo

func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) GetActiveCount() uint32 {
	if m != nil && m.ActiveCount != nil {
		return *m.ActiveCount
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationType) GetRsvpInterface() []*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList {
	if m != nil {
		return m.RsvpInterface
	}
	return nil
}

type JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList struct {
	InterfaceName        *string                                                                                 `protobuf:"bytes,51,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	Index                *uint32                                                                                 `protobuf:"varint,52,opt,name=index" json:"index,omitempty"`
	RsvpStatus           *string                                                                                 `protobuf:"bytes,53,opt,name=rsvp_status,json=rsvpStatus" json:"rsvp_status,omitempty"`
	AuthenticationFlag   *string                                                                                 `protobuf:"bytes,54,opt,name=authentication_flag,json=authenticationFlag" json:"authentication_flag,omitempty"`
	AggregateFlag        *string                                                                                 `protobuf:"bytes,55,opt,name=aggregate_flag,json=aggregateFlag" json:"aggregate_flag,omitempty"`
	AckFlag              *string                                                                                 `protobuf:"bytes,56,opt,name=ack_flag,json=ackFlag" json:"ack_flag,omitempty"`
	ProtectFlag          *string                                                                                 `protobuf:"bytes,57,opt,name=protect_flag,json=protectFlag" json:"protect_flag,omitempty"`
	HelloInterval        *uint32                                                                                 `protobuf:"varint,58,opt,name=hello_interval,json=helloInterval" json:"hello_interval,omitempty"`
	InterfaceAddress     *string                                                                                 `protobuf:"bytes,59,opt,name=interface_address,json=interfaceAddress" json:"interface_address,omitempty"`
	MessageStatistics    []*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList `protobuf:"bytes,151,rep,name=message_statistics,json=messageStatistics" json:"message_statistics,omitempty"`
	RsvpTelink           *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType          `protobuf:"bytes,152,opt,name=rsvp_telink,json=rsvpTelink" json:"rsvp_telink,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                `json:"-"`
	XXX_unrecognized     []byte                                                                                  `json:"-"`
	XXX_sizecache        int32                                                                                   `json:"-"`
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) Reset() {
	*m = JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList{}
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) ProtoMessage() {}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4225f46c0f8328, []int{0, 0, 0}
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList.Unmarshal(m, b)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList.Marshal(b, m, deterministic)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList.Merge(m, src)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) XXX_Size() int {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList.Size(m)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList proto.InternalMessageInfo

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetInterfaceName() string {
	if m != nil && m.InterfaceName != nil {
		return *m.InterfaceName
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetRsvpStatus() string {
	if m != nil && m.RsvpStatus != nil {
		return *m.RsvpStatus
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetAuthenticationFlag() string {
	if m != nil && m.AuthenticationFlag != nil {
		return *m.AuthenticationFlag
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetAggregateFlag() string {
	if m != nil && m.AggregateFlag != nil {
		return *m.AggregateFlag
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetAckFlag() string {
	if m != nil && m.AckFlag != nil {
		return *m.AckFlag
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetProtectFlag() string {
	if m != nil && m.ProtectFlag != nil {
		return *m.ProtectFlag
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetHelloInterval() uint32 {
	if m != nil && m.HelloInterval != nil {
		return *m.HelloInterval
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetInterfaceAddress() string {
	if m != nil && m.InterfaceAddress != nil {
		return *m.InterfaceAddress
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetMessageStatistics() []*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList {
	if m != nil {
		return m.MessageStatistics
	}
	return nil
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList) GetRsvpTelink() *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType {
	if m != nil {
		return m.RsvpTelink
	}
	return nil
}

type JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList struct {
	RsvpMessage               *string  `protobuf:"bytes,51,opt,name=rsvp_message,json=rsvpMessage" json:"rsvp_message,omitempty"`
	MessagesSent              *uint32  `protobuf:"varint,52,opt,name=messages_sent,json=messagesSent" json:"messages_sent,omitempty"`
	MessagesReceived          *uint32  `protobuf:"varint,53,opt,name=messages_received,json=messagesReceived" json:"messages_received,omitempty"`
	MessagesSent_5Seconds     *uint32  `protobuf:"varint,54,opt,name=messages_sent_5seconds,json=messagesSent5seconds" json:"messages_sent_5seconds,omitempty"`
	MessagesReceived_5Seconds *uint32  `protobuf:"varint,55,opt,name=messages_received_5seconds,json=messagesReceived5seconds" json:"messages_received_5seconds,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) Reset() {
	*m = JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList{}
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) ProtoMessage() {
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4225f46c0f8328, []int{0, 0, 0, 0}
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList.Unmarshal(m, b)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList.Marshal(b, m, deterministic)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList.Merge(m, src)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) XXX_Size() int {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList.Size(m)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList proto.InternalMessageInfo

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) GetRsvpMessage() string {
	if m != nil && m.RsvpMessage != nil {
		return *m.RsvpMessage
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) GetMessagesSent() uint32 {
	if m != nil && m.MessagesSent != nil {
		return *m.MessagesSent
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) GetMessagesReceived() uint32 {
	if m != nil && m.MessagesReceived != nil {
		return *m.MessagesReceived
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) GetMessagesSent_5Seconds() uint32 {
	if m != nil && m.MessagesSent_5Seconds != nil {
		return *m.MessagesSent_5Seconds
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList) GetMessagesReceived_5Seconds() uint32 {
	if m != nil && m.MessagesReceived_5Seconds != nil {
		return *m.MessagesReceived_5Seconds
	}
	return 0
}

type JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType struct {
	ActiveReservation    *uint32                                                                                               `protobuf:"varint,51,opt,name=active_reservation,json=activeReservation" json:"active_reservation,omitempty"`
	PreemptionCount      *uint32                                                                                               `protobuf:"varint,52,opt,name=preemption_count,json=preemptionCount" json:"preemption_count,omitempty"`
	UpdateThreshold      *uint32                                                                                               `protobuf:"varint,53,opt,name=update_threshold,json=updateThreshold" json:"update_threshold,omitempty"`
	Subscription         *uint32                                                                                               `protobuf:"varint,54,opt,name=subscription" json:"subscription,omitempty"`
	StaticBandwidth      *string                                                                                               `protobuf:"bytes,55,opt,name=static_bandwidth,json=staticBandwidth" json:"static_bandwidth,omitempty"`
	AvailableBandwidth   *string                                                                                               `protobuf:"bytes,56,opt,name=available_bandwidth,json=availableBandwidth" json:"available_bandwidth,omitempty"`
	ReservedBandwidth    []*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList `protobuf:"bytes,151,rep,name=reserved_bandwidth,json=reservedBandwidth" json:"reserved_bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                              `json:"-"`
	XXX_unrecognized     []byte                                                                                                `json:"-"`
	XXX_sizecache        int32                                                                                                 `json:"-"`
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) Reset() {
	*m = JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType{}
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) String() string {
	return proto.CompactTextString(m)
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) ProtoMessage() {}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4225f46c0f8328, []int{0, 0, 0, 1}
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType.Unmarshal(m, b)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType.Marshal(b, m, deterministic)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType.Merge(m, src)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) XXX_Size() int {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType.Size(m)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType.DiscardUnknown(m)
}

var xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType proto.InternalMessageInfo

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetActiveReservation() uint32 {
	if m != nil && m.ActiveReservation != nil {
		return *m.ActiveReservation
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetPreemptionCount() uint32 {
	if m != nil && m.PreemptionCount != nil {
		return *m.PreemptionCount
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetUpdateThreshold() uint32 {
	if m != nil && m.UpdateThreshold != nil {
		return *m.UpdateThreshold
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetSubscription() uint32 {
	if m != nil && m.Subscription != nil {
		return *m.Subscription
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetStaticBandwidth() string {
	if m != nil && m.StaticBandwidth != nil {
		return *m.StaticBandwidth
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetAvailableBandwidth() string {
	if m != nil && m.AvailableBandwidth != nil {
		return *m.AvailableBandwidth
	}
	return ""
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType) GetReservedBandwidth() []*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList {
	if m != nil {
		return m.ReservedBandwidth
	}
	return nil
}

type JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList struct {
	BandwidthPriority      *uint32  `protobuf:"varint,51,opt,name=bandwidth_priority,json=bandwidthPriority" json:"bandwidth_priority,omitempty"`
	TotalReservedBandwidth *string  `protobuf:"bytes,52,opt,name=total_reserved_bandwidth,json=totalReservedBandwidth" json:"total_reserved_bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) Reset() {
	*m = JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList{}
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) String() string {
	return proto.CompactTextString(m)
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) ProtoMessage() {
}
func (*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) Descriptor() ([]byte, []int) {
	return fileDescriptor_df4225f46c0f8328, []int{0, 0, 0, 1, 0}
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList.Unmarshal(m, b)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList.Marshal(b, m, deterministic)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList.Merge(m, src)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) XXX_Size() int {
	return xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList.Size(m)
}
func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) XXX_DiscardUnknown() {
	xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList.DiscardUnknown(m)
}

var xxx_messageInfo_JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList proto.InternalMessageInfo

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) GetBandwidthPriority() uint32 {
	if m != nil && m.BandwidthPriority != nil {
		return *m.BandwidthPriority
	}
	return 0
}

func (m *JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList) GetTotalReservedBandwidth() string {
	if m != nil && m.TotalReservedBandwidth != nil {
		return *m.TotalReservedBandwidth
	}
	return ""
}

var E_JnprJunosRsvpInterfaceExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*JunosRsvpInterface)(nil),
	Field:         46,
	Name:          "jnpr_junos_rsvp_interface_ext",
	Tag:           "bytes,46,opt,name=jnpr_junos_rsvp_interface_ext",
	Filename:      "junos-xmlproxyd_junos-rsvp-interface.proto",
}

func init() {
	proto.RegisterType((*JunosRsvpInterface)(nil), "junos_rsvp_interface")
	proto.RegisterType((*JunosRsvpInterfaceRsvpInterfaceInformationType)(nil), "junos_rsvp_interface.rsvp_interface_information_type")
	proto.RegisterType((*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceList)(nil), "junos_rsvp_interface.rsvp_interface_information_type.rsvp_interface_list")
	proto.RegisterType((*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListMessageStatisticsList)(nil), "junos_rsvp_interface.rsvp_interface_information_type.rsvp_interface_list.message_statistics_list")
	proto.RegisterType((*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkType)(nil), "junos_rsvp_interface.rsvp_interface_information_type.rsvp_interface_list.rsvp_telink_type")
	proto.RegisterType((*JunosRsvpInterfaceRsvpInterfaceInformationTypeRsvpInterfaceListRsvpTelinkTypeReservedBandwidthList)(nil), "junos_rsvp_interface.rsvp_interface_information_type.rsvp_interface_list.rsvp_telink_type.reserved_bandwidth_list")
	proto.RegisterExtension(E_JnprJunosRsvpInterfaceExt)
}

func init() {
	proto.RegisterFile("junos-xmlproxyd_junos-rsvp-interface.proto", fileDescriptor_df4225f46c0f8328)
}

var fileDescriptor_df4225f46c0f8328 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x6f, 0xd3, 0x4a,
	0x14, 0x56, 0x7a, 0x6f, 0x75, 0x7b, 0x4f, 0xe2, 0x34, 0x99, 0x96, 0xd6, 0xb5, 0x84, 0xfa, 0x40,
	0x48, 0x29, 0xd0, 0x20, 0x95, 0x94, 0x96, 0xc2, 0x02, 0x8a, 0x40, 0x4a, 0x25, 0x2a, 0xe4, 0x76,
	0xc5, 0xc6, 0x9a, 0xda, 0xa7, 0xc9, 0x34, 0x8e, 0x6d, 0x66, 0x26, 0x69, 0x22, 0xb1, 0x62, 0x83,
	0xba, 0x64, 0x03, 0xac, 0xf9, 0x35, 0x6c, 0xf8, 0x4f, 0xc8, 0x33, 0x7e, 0x24, 0x69, 0x2a, 0x24,
	0x54, 0x96, 0xfe, 0x1e, 0x67, 0xce, 0x9c, 0xc7, 0x18, 0xee, 0x9d, 0xf7, 0x82, 0x50, 0x6c, 0x0d,
	0xba, 0x7e, 0xc4, 0xc3, 0xc1, 0xd0, 0x73, 0xf4, 0x37, 0x17, 0xfd, 0x68, 0x8b, 0x05, 0x12, 0xf9,
	0x19, 0x75, 0xb1, 0x1e, 0xf1, 0x50, 0x86, 0xd6, 0x82, 0x44, 0x1f, 0xbb, 0x28, 0xf9, 0xd0, 0x91,
	0x61, 0xa4, 0xc1, 0x8d, 0x2f, 0x65, 0x58, 0x54, 0x1e, 0x27, 0xf6, 0x38, 0x99, 0x87, 0x48, 0xb0,
	0xc6, 0x11, 0x87, 0x05, 0x67, 0x21, 0xef, 0x52, 0xc9, 0xc2, 0xc0, 0xfc, 0x5a, 0x58, 0x2b, 0xd4,
	0x8a, 0xdb, 0x3b, 0xf5, 0x69, 0xde, 0xfa, 0xf5, 0x46, 0x47, 0x0e, 0x23, 0xb4, 0xcd, 0x58, 0xd0,
	0x4c, 0xf9, 0x66, 0x4e, 0x5b, 0x9f, 0x0c, 0x58, 0xfd, 0x8d, 0x9b, 0xac, 0x43, 0x89, 0xba, 0x92,
	0xf5, 0xd1, 0x71, 0xc3, 0x5e, 0x20, 0xcd, 0x47, 0x6b, 0x85, 0x9a, 0x61, 0x17, 0x35, 0xf6, 0x32,
	0x86, 0xc8, 0x7b, 0x28, 0x8f, 0x47, 0x89, 0x13, 0xfe, 0xa7, 0x56, 0xdc, 0x6e, 0xfe, 0x51, 0xc2,
	0x93, 0xbc, 0xcf, 0x84, 0xb4, 0x8d, 0xb1, 0x4b, 0x58, 0x3f, 0x8b, 0xb0, 0x30, 0x45, 0x46, 0x1e,
	0x40, 0x39, 0x47, 0x02, 0xda, 0x45, 0x95, 0xef, 0xff, 0x07, 0xb3, 0x1f, 0x9f, 0xcf, 0xcc, 0x15,
	0x6c, 0x23, 0x23, 0x8f, 0x68, 0x17, 0xc9, 0x22, 0xcc, 0xb2, 0xc0, 0xc3, 0x81, 0xd9, 0x50, 0x97,
	0xd2, 0x1f, 0x64, 0x15, 0x8a, 0x2a, 0xb4, 0x90, 0x54, 0xf6, 0x84, 0xb9, 0x13, 0x07, 0xb0, 0x21,
	0x86, 0x8e, 0x15, 0x42, 0x1e, 0xc2, 0x02, 0xed, 0xc9, 0x36, 0x06, 0x92, 0xb9, 0x3a, 0xed, 0x33,
	0x9f, 0xb6, 0xcc, 0xc7, 0x4a, 0x48, 0xc6, 0xa9, 0xd7, 0x3e, 0x6d, 0x91, 0xbb, 0x50, 0xa6, 0xad,
	0x16, 0xc7, 0x16, 0x95, 0xa8, 0xb5, 0xbb, 0x4a, 0x6b, 0x64, 0xa8, 0x92, 0xad, 0xc0, 0x1c, 0x75,
	0x3b, 0x5a, 0xb0, 0xa7, 0x04, 0xff, 0x51, 0xb7, 0xa3, 0xa8, 0x75, 0x28, 0xc5, 0x13, 0x84, 0xae,
	0xd4, 0xf4, 0x13, 0x45, 0x17, 0x13, 0x2c, 0x3d, 0xa4, 0x8d, 0xbe, 0x1f, 0xea, 0x92, 0xf4, 0xa9,
	0x6f, 0xee, 0xab, 0x5b, 0x19, 0x0a, 0x6d, 0x26, 0x20, 0xb9, 0x0f, 0xd5, 0xbc, 0x42, 0xd4, 0xf3,
	0x38, 0x0a, 0x61, 0x3e, 0x55, 0xe1, 0x2a, 0x19, 0xf1, 0x42, 0xe3, 0xe4, 0x73, 0x01, 0x48, 0x17,
	0x85, 0xa0, 0x2d, 0x54, 0xe5, 0x60, 0x42, 0x32, 0x57, 0xa4, 0xed, 0xa5, 0x37, 0xd6, 0xde, 0xfa,
	0xd5, 0x43, 0x74, 0xdb, 0xab, 0x09, 0x71, 0x9c, 0xe1, 0xe4, 0x43, 0xd2, 0x1e, 0x89, 0x3e, 0x0b,
	0x3a, 0xe6, 0x37, 0xbd, 0x1b, 0xef, 0x6e, 0x2e, 0x97, 0x91, 0xe8, 0x7a, 0x81, 0x54, 0xef, 0x4f,
	0x14, 0x60, 0x5d, 0xce, 0xc0, 0xf2, 0x35, 0xc9, 0x92, 0x1a, 0x94, 0x94, 0x37, 0xe1, 0xc7, 0x47,
	0x4f, 0x25, 0xfd, 0x46, 0x33, 0xe4, 0x0e, 0x18, 0x89, 0x48, 0x38, 0x02, 0x03, 0x99, 0x0c, 0x60,
	0x29, 0x05, 0x8f, 0x31, 0x90, 0x71, 0xa7, 0x32, 0x11, 0x47, 0x17, 0x59, 0x1f, 0x3d, 0x35, 0x8d,
	0x86, 0x5d, 0x49, 0x09, 0x3b, 0xc1, 0x49, 0x03, 0x96, 0xc6, 0x22, 0x3a, 0x3b, 0x02, 0xdd, 0x30,
	0xf0, 0x84, 0x1a, 0x4b, 0xc3, 0x5e, 0x1c, 0x0d, 0x9d, 0x72, 0xe4, 0x19, 0x58, 0x57, 0x8e, 0xc8,
	0x9d, 0xbb, 0xca, 0x69, 0x4e, 0x9e, 0x95, 0xf2, 0xd6, 0x8f, 0x7f, 0xa1, 0x32, 0x59, 0x2c, 0xb2,
	0x05, 0x24, 0x79, 0x2f, 0x38, 0x8a, 0x78, 0xe4, 0xd4, 0x0b, 0xa6, 0x5f, 0x8d, 0xaa, 0x66, 0xec,
	0x9c, 0x20, 0x9b, 0x50, 0x89, 0x38, 0x62, 0x37, 0x52, 0x3d, 0xd1, 0x4f, 0x8c, 0x2e, 0xc6, 0x7c,
	0x8e, 0xeb, 0x67, 0x66, 0x13, 0x2a, 0xbd, 0xc8, 0x8b, 0x57, 0x48, 0xb6, 0x39, 0x8a, 0x76, 0xe8,
	0xa7, 0xe5, 0x98, 0xd7, 0xf8, 0x49, 0x0a, 0x93, 0x0d, 0x28, 0x89, 0xde, 0xa9, 0x70, 0x39, 0x53,
	0xfe, 0xa4, 0x06, 0x63, 0x58, 0x1c, 0x4e, 0x35, 0xd0, 0x75, 0x4e, 0x69, 0xe0, 0x5d, 0x30, 0x4f,
	0xb6, 0x93, 0xb5, 0x9c, 0xd7, 0xf8, 0x41, 0x0a, 0xab, 0x85, 0xef, 0x53, 0xe6, 0xd3, 0x53, 0x1f,
	0x47, 0xd4, 0x7b, 0xc9, 0xc2, 0xa7, 0x54, 0x6e, 0xf8, 0x5e, 0x00, 0xa2, 0xaf, 0x8f, 0xde, 0x88,
	0x21, 0xd9, 0x1b, 0xfe, 0xf7, 0x66, 0xb5, 0x7e, 0xf5, 0xd4, 0x64, 0x91, 0x52, 0x22, 0x4b, 0xd2,
	0xba, 0x2c, 0xc0, 0xf2, 0x35, 0x72, 0xd2, 0x00, 0x92, 0x23, 0x11, 0x67, 0x21, 0x67, 0x72, 0xa8,
	0xbb, 0x98, 0x0e, 0x74, 0x35, 0x13, 0xbc, 0x4d, 0x78, 0xb2, 0x07, 0xa6, 0x0c, 0x25, 0xf5, 0x9d,
	0x29, 0x77, 0x6f, 0xa8, 0x62, 0x2d, 0x29, 0xde, 0x9e, 0xcc, 0x65, 0x3f, 0x82, 0xdb, 0xe7, 0x41,
	0xc4, 0x9d, 0x69, 0x85, 0x71, 0x70, 0x20, 0xc9, 0x72, 0xfd, 0xb0, 0x17, 0xb0, 0x08, 0xf9, 0x11,
	0xca, 0x8b, 0x90, 0x77, 0xe2, 0x39, 0x16, 0x21, 0x17, 0x66, 0x5d, 0xad, 0xff, 0xad, 0xa9, 0x25,
	0xb5, 0x57, 0xe2, 0xa0, 0x87, 0x31, 0x63, 0x8f, 0xfe, 0x3e, 0x5e, 0x0d, 0xe4, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd4, 0xbb, 0x9c, 0xd8, 0xcc, 0x07, 0x00, 0x00,
}
