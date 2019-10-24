// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmerror.proto

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
// Juniper Error Item information
//
type Error struct {
	// Identifier that uniquely identifies the source of
	// the error.
	// e.g.
	//
	// junos/system/linecard/0/pcie/0/lane/0/pcie_cmerror_uncorrectable_major
	//
	Identifier *string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Name of the error
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Instance id of the associated resource
	ComponentId *uint32 `protobuf:"varint,3,opt,name=component_id,json=componentId" json:"component_id,omitempty"`
	// Fru information
	FruType *string `protobuf:"bytes,4,opt,name=fru_type,json=fruType" json:"fru_type,omitempty"`
	FruSlot *uint32 `protobuf:"varint,5,opt,name=fru_slot,json=fruSlot" json:"fru_slot,omitempty"`
	// Scope,Category,Severity
	// in which this error belong to.
	Scope    *string `protobuf:"bytes,6,opt,name=scope" json:"scope,omitempty"`
	Category *string `protobuf:"bytes,7,opt,name=category" json:"category,omitempty"`
	Severity *string `protobuf:"bytes,8,opt,name=severity" json:"severity,omitempty"`
	// Thresholds and action configured for this
	// error.
	Threshold         *uint32 `protobuf:"varint,9,opt,name=threshold" json:"threshold,omitempty"`
	Limit             *uint32 `protobuf:"varint,10,opt,name=limit" json:"limit,omitempty"`
	RaisingThreshold  *uint32 `protobuf:"varint,11,opt,name=raising_threshold,json=raisingThreshold" json:"raising_threshold,omitempty"`
	ClearingThreshold *uint32 `protobuf:"varint,12,opt,name=clearing_threshold,json=clearingThreshold" json:"clearing_threshold,omitempty"`
	Action            *uint32 `protobuf:"varint,13,opt,name=action" json:"action,omitempty"`
	// local/global/both
	ActionHandlingType *uint32 `protobuf:"varint,14,opt,name=action_handling_type,json=actionHandlingType" json:"action_handling_type,omitempty"`
	// user configured thresholds and limits for this error.
	ConfiguredThreshold   *uint32  `protobuf:"varint,15,opt,name=configured_threshold,json=configuredThreshold" json:"configured_threshold,omitempty"`
	ConfiguredLimit       *uint32  `protobuf:"varint,16,opt,name=configured_limit,json=configuredLimit" json:"configured_limit,omitempty"`
	ConfiguredAction      *uint32  `protobuf:"varint,17,opt,name=configured_action,json=configuredAction" json:"configured_action,omitempty"`
	ConfiguredClearAction *uint32  `protobuf:"varint,18,opt,name=configured_clear_action,json=configuredClearAction" json:"configured_clear_action,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_747f0735808ade43, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *Error) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Error) GetComponentId() uint32 {
	if m != nil && m.ComponentId != nil {
		return *m.ComponentId
	}
	return 0
}

func (m *Error) GetFruType() string {
	if m != nil && m.FruType != nil {
		return *m.FruType
	}
	return ""
}

func (m *Error) GetFruSlot() uint32 {
	if m != nil && m.FruSlot != nil {
		return *m.FruSlot
	}
	return 0
}

func (m *Error) GetScope() string {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return ""
}

func (m *Error) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *Error) GetSeverity() string {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return ""
}

func (m *Error) GetThreshold() uint32 {
	if m != nil && m.Threshold != nil {
		return *m.Threshold
	}
	return 0
}

func (m *Error) GetLimit() uint32 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *Error) GetRaisingThreshold() uint32 {
	if m != nil && m.RaisingThreshold != nil {
		return *m.RaisingThreshold
	}
	return 0
}

func (m *Error) GetClearingThreshold() uint32 {
	if m != nil && m.ClearingThreshold != nil {
		return *m.ClearingThreshold
	}
	return 0
}

func (m *Error) GetAction() uint32 {
	if m != nil && m.Action != nil {
		return *m.Action
	}
	return 0
}

func (m *Error) GetActionHandlingType() uint32 {
	if m != nil && m.ActionHandlingType != nil {
		return *m.ActionHandlingType
	}
	return 0
}

func (m *Error) GetConfiguredThreshold() uint32 {
	if m != nil && m.ConfiguredThreshold != nil {
		return *m.ConfiguredThreshold
	}
	return 0
}

func (m *Error) GetConfiguredLimit() uint32 {
	if m != nil && m.ConfiguredLimit != nil {
		return *m.ConfiguredLimit
	}
	return 0
}

func (m *Error) GetConfiguredAction() uint32 {
	if m != nil && m.ConfiguredAction != nil {
		return *m.ConfiguredAction
	}
	return 0
}

func (m *Error) GetConfiguredClearAction() uint32 {
	if m != nil && m.ConfiguredClearAction != nil {
		return *m.ConfiguredClearAction
	}
	return 0
}

//
// Top-level Cmerror message
//
type Cmerror struct {
	// collection of error items
	ErrorItem []*Error `protobuf:"bytes,1,rep,name=error_item,json=errorItem" json:"error_item,omitempty"`
	// last configuration change for cmerror.
	LastConfigurationChange *uint64 `protobuf:"varint,2,opt,name=last_configuration_change,json=lastConfigurationChange" json:"last_configuration_change,omitempty"`
	// This will toggle at start of every wrap cycle
	WrapCycleMarker      *bool    `protobuf:"varint,3,opt,name=wrap_cycle_marker,json=wrapCycleMarker" json:"wrap_cycle_marker,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cmerror) Reset()         { *m = Cmerror{} }
func (m *Cmerror) String() string { return proto.CompactTextString(m) }
func (*Cmerror) ProtoMessage()    {}
func (*Cmerror) Descriptor() ([]byte, []int) {
	return fileDescriptor_747f0735808ade43, []int{1}
}

func (m *Cmerror) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cmerror.Unmarshal(m, b)
}
func (m *Cmerror) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cmerror.Marshal(b, m, deterministic)
}
func (m *Cmerror) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cmerror.Merge(m, src)
}
func (m *Cmerror) XXX_Size() int {
	return xxx_messageInfo_Cmerror.Size(m)
}
func (m *Cmerror) XXX_DiscardUnknown() {
	xxx_messageInfo_Cmerror.DiscardUnknown(m)
}

var xxx_messageInfo_Cmerror proto.InternalMessageInfo

func (m *Cmerror) GetErrorItem() []*Error {
	if m != nil {
		return m.ErrorItem
	}
	return nil
}

func (m *Cmerror) GetLastConfigurationChange() uint64 {
	if m != nil && m.LastConfigurationChange != nil {
		return *m.LastConfigurationChange
	}
	return 0
}

func (m *Cmerror) GetWrapCycleMarker() bool {
	if m != nil && m.WrapCycleMarker != nil {
		return *m.WrapCycleMarker
	}
	return false
}

var E_JnprCmerrorExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*Cmerror)(nil),
	Field:         20,
	Name:          "jnpr_cmerror_ext",
	Tag:           "bytes,20,opt,name=jnpr_cmerror_ext",
	Filename:      "cmerror.proto",
}

func init() {
	proto.RegisterType((*Error)(nil), "Error")
	proto.RegisterType((*Cmerror)(nil), "Cmerror")
	proto.RegisterExtension(E_JnprCmerrorExt)
}

func init() { proto.RegisterFile("cmerror.proto", fileDescriptor_747f0735808ade43) }

var fileDescriptor_747f0735808ade43 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x93, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xe5, 0xb6, 0x49, 0x9c, 0x49, 0xd3, 0x24, 0xdb, 0x40, 0xb6, 0x15, 0x87, 0x50, 0x09,
	0x29, 0x80, 0x88, 0xa0, 0x07, 0x0e, 0x3d, 0x01, 0x51, 0x25, 0xca, 0xbf, 0x83, 0xdb, 0xbb, 0x65,
	0xd9, 0x93, 0x64, 0xa9, 0xbd, 0x6b, 0xad, 0x37, 0xb4, 0xbe, 0x72, 0xe0, 0x33, 0xf0, 0x71, 0xd1,
	0xce, 0xda, 0xb1, 0xb9, 0xed, 0xbc, 0xdf, 0x7b, 0x3b, 0xde, 0xd1, 0x18, 0x86, 0x71, 0x86, 0x5a,
	0x2b, 0xbd, 0xcc, 0xb5, 0x32, 0xea, 0xfc, 0xd4, 0x60, 0x8a, 0x19, 0x1a, 0x5d, 0x86, 0x46, 0xe5,
	0x4e, 0xbc, 0xf8, 0xd3, 0x81, 0xce, 0xb5, 0x35, 0xb1, 0x17, 0x00, 0x22, 0x41, 0x69, 0xc4, 0x5a,
	0xa0, 0xe6, 0xde, 0xdc, 0x5b, 0xf4, 0x3f, 0x75, 0x7e, 0x7f, 0x38, 0xf0, 0xbd, 0xa0, 0x05, 0x18,
	0x83, 0x23, 0x19, 0x65, 0xc8, 0x0f, 0xac, 0x21, 0xa0, 0x33, 0x7b, 0x0e, 0xc7, 0xb1, 0xca, 0x72,
	0x25, 0x51, 0x9a, 0x50, 0x24, 0xfc, 0x70, 0xee, 0x2d, 0x86, 0xc1, 0x60, 0xaf, 0xdd, 0x24, 0xec,
	0x0c, 0xfc, 0xb5, 0xde, 0x85, 0xa6, 0xcc, 0x91, 0x1f, 0x51, 0xb4, 0xb7, 0xd6, 0xbb, 0xbb, 0x32,
	0xc7, 0x1a, 0x15, 0xa9, 0x32, 0xbc, 0x43, 0x49, 0x8b, 0x6e, 0x53, 0x65, 0xd8, 0x14, 0x3a, 0x45,
	0xac, 0x72, 0xe4, 0x5d, 0x8a, 0xb8, 0x82, 0x9d, 0x83, 0x1f, 0x47, 0x06, 0x37, 0x4a, 0x97, 0xbc,
	0x47, 0x60, 0x5f, 0x5b, 0x56, 0xe0, 0x2f, 0xd4, 0xc2, 0x94, 0xdc, 0x77, 0xac, 0xae, 0xd9, 0x33,
	0xe8, 0x9b, 0xad, 0xc6, 0x62, 0xab, 0xd2, 0x84, 0xf7, 0xa9, 0x53, 0x23, 0xd8, 0x5e, 0xa9, 0xc8,
	0x84, 0xe1, 0x40, 0xc4, 0x15, 0xec, 0x35, 0x4c, 0x74, 0x24, 0x0a, 0x21, 0x37, 0x61, 0x93, 0x1d,
	0x90, 0x63, 0x5c, 0x81, 0xbb, 0xfd, 0x15, 0x6f, 0x80, 0xc5, 0x29, 0x46, 0xfa, 0x7f, 0xf7, 0x31,
	0xb9, 0x27, 0x35, 0x69, 0xec, 0x4f, 0xa1, 0x1b, 0xc5, 0x46, 0x28, 0xc9, 0x87, 0x64, 0xa9, 0x2a,
	0xf6, 0x16, 0xa6, 0xee, 0x14, 0x6e, 0x23, 0x99, 0xa4, 0x74, 0x9b, 0x9d, 0xdb, 0x09, 0xb9, 0x98,
	0x63, 0x9f, 0x2b, 0x44, 0x23, 0x7c, 0x07, 0xd3, 0x58, 0xc9, 0xb5, 0xd8, 0xec, 0x34, 0x26, 0xad,
	0xd6, 0x23, 0x4a, 0x9c, 0x36, 0xac, 0x69, 0xfe, 0x12, 0xc6, 0xad, 0x88, 0x7b, 0xf9, 0x98, 0xec,
	0xa3, 0x46, 0xff, 0x56, 0xcf, 0xa0, 0x65, 0xad, 0x3e, 0x79, 0xe2, 0x66, 0xd0, 0x80, 0x8f, 0xee,
	0xe3, 0xdf, 0xc3, 0xac, 0x65, 0xa6, 0x47, 0xd7, 0x11, 0x46, 0x91, 0x27, 0x0d, 0x5e, 0x59, 0xea,
	0x72, 0x17, 0x7f, 0x3d, 0xe8, 0xad, 0xdc, 0xbe, 0xda, 0x55, 0xa4, 0x43, 0x28, 0x0c, 0x66, 0xdc,
	0x9b, 0x1f, 0x2e, 0x06, 0x97, 0xdd, 0x25, 0xad, 0x69, 0xd0, 0x27, 0x72, 0x63, 0x30, 0x63, 0x57,
	0x70, 0x96, 0x46, 0x85, 0x09, 0xeb, 0x0b, 0x23, 0x9a, 0x59, 0xbc, 0x8d, 0xe4, 0xc6, 0xed, 0xe7,
	0x51, 0x30, 0xb3, 0x86, 0x55, 0x9b, 0xaf, 0x08, 0xb3, 0x57, 0x30, 0x79, 0xd0, 0x51, 0x1e, 0xc6,
	0x65, 0x9c, 0x62, 0x98, 0x45, 0xfa, 0x1e, 0x35, 0xed, 0xad, 0x1f, 0x8c, 0x2c, 0x58, 0x59, 0xfd,
	0x3b, 0xc9, 0x57, 0x5f, 0x61, 0xfc, 0x53, 0xe6, 0x3a, 0xac, 0x7e, 0xa7, 0x10, 0x1f, 0x0d, 0x9b,
	0x2d, 0xbf, 0xec, 0xa4, 0xc8, 0x51, 0xff, 0x40, 0xf3, 0xa0, 0xf4, 0x7d, 0x71, 0x8b, 0xb2, 0x50,
	0xba, 0xe0, 0xd3, 0xb9, 0xb7, 0x18, 0x5c, 0xfa, 0xcb, 0xea, 0x2d, 0xc1, 0x89, 0x8d, 0x56, 0xc5,
	0xf5, 0xa3, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x6f, 0xcb, 0xef, 0x95, 0x03, 0x00, 0x00,
}
