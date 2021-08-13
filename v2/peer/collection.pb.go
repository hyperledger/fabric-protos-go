// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/collection.proto

package peer

import (
	fmt "fmt"
	common "github.com/SmartBFT-Go/fabric-protos-go/v2/common"
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

// CollectionConfigPackage represents an array of CollectionConfig
// messages; the extra struct is required because repeated oneof is
// forbidden by the protobuf syntax
type CollectionConfigPackage struct {
	Config               []*CollectionConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CollectionConfigPackage) Reset()         { *m = CollectionConfigPackage{} }
func (m *CollectionConfigPackage) String() string { return proto.CompactTextString(m) }
func (*CollectionConfigPackage) ProtoMessage()    {}
func (*CollectionConfigPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{0}
}

func (m *CollectionConfigPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfigPackage.Unmarshal(m, b)
}
func (m *CollectionConfigPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfigPackage.Marshal(b, m, deterministic)
}
func (m *CollectionConfigPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfigPackage.Merge(m, src)
}
func (m *CollectionConfigPackage) XXX_Size() int {
	return xxx_messageInfo_CollectionConfigPackage.Size(m)
}
func (m *CollectionConfigPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfigPackage.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionConfigPackage proto.InternalMessageInfo

func (m *CollectionConfigPackage) GetConfig() []*CollectionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
type CollectionConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionConfig_StaticCollectionConfig
	Payload              isCollectionConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *CollectionConfig) Reset()         { *m = CollectionConfig{} }
func (m *CollectionConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionConfig) ProtoMessage()    {}
func (*CollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{1}
}

func (m *CollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionConfig.Unmarshal(m, b)
}
func (m *CollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionConfig.Marshal(b, m, deterministic)
}
func (m *CollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionConfig.Merge(m, src)
}
func (m *CollectionConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionConfig.Size(m)
}
func (m *CollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionConfig proto.InternalMessageInfo

type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}

type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,proto3,oneof"`
}

func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}

func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := m.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
}

// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
type StaticCollectionConfig struct {
	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy,proto3" json:"member_orgs_policy,omitempty"`
	// The minimum number of peers private data will be sent to upon
	// endorsement. The endorsement would fail if dissemination to at least
	// this number of peers is not achieved.
	RequiredPeerCount int32 `protobuf:"varint,3,opt,name=required_peer_count,json=requiredPeerCount,proto3" json:"required_peer_count,omitempty"`
	// The maximum number of peers that private data will be sent to
	// upon endorsement. This number has to be bigger than required_peer_count.
	MaximumPeerCount int32 `protobuf:"varint,4,opt,name=maximum_peer_count,json=maximumPeerCount,proto3" json:"maximum_peer_count,omitempty"`
	// The number of blocks after which the collection data expires.
	// For instance if the value is set to 10, a key last modified by block number 100
	// will be purged at block number 111. A zero value is treated same as MaxUint64
	BlockToLive uint64 `protobuf:"varint,5,opt,name=block_to_live,json=blockToLive,proto3" json:"block_to_live,omitempty"`
	// The member only read access denotes whether only collection member clients
	// can read the private data (if set to true), or even non members can
	// read the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyRead bool `protobuf:"varint,6,opt,name=member_only_read,json=memberOnlyRead,proto3" json:"member_only_read,omitempty"`
	// The member only write access denotes whether only collection member clients
	// can write the private data (if set to true), or even non members can
	// write the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyWrite bool `protobuf:"varint,7,opt,name=member_only_write,json=memberOnlyWrite,proto3" json:"member_only_write,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define the endorsement policy for this collection
	EndorsementPolicy    *ApplicationPolicy `protobuf:"bytes,8,opt,name=endorsement_policy,json=endorsementPolicy,proto3" json:"endorsement_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StaticCollectionConfig) Reset()         { *m = StaticCollectionConfig{} }
func (m *StaticCollectionConfig) String() string { return proto.CompactTextString(m) }
func (*StaticCollectionConfig) ProtoMessage()    {}
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{2}
}

func (m *StaticCollectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticCollectionConfig.Unmarshal(m, b)
}
func (m *StaticCollectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticCollectionConfig.Marshal(b, m, deterministic)
}
func (m *StaticCollectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticCollectionConfig.Merge(m, src)
}
func (m *StaticCollectionConfig) XXX_Size() int {
	return xxx_messageInfo_StaticCollectionConfig.Size(m)
}
func (m *StaticCollectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticCollectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StaticCollectionConfig proto.InternalMessageInfo

func (m *StaticCollectionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if m != nil {
		return m.MemberOrgsPolicy
	}
	return nil
}

func (m *StaticCollectionConfig) GetRequiredPeerCount() int32 {
	if m != nil {
		return m.RequiredPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetMaximumPeerCount() int32 {
	if m != nil {
		return m.MaximumPeerCount
	}
	return 0
}

func (m *StaticCollectionConfig) GetBlockToLive() uint64 {
	if m != nil {
		return m.BlockToLive
	}
	return 0
}

func (m *StaticCollectionConfig) GetMemberOnlyRead() bool {
	if m != nil {
		return m.MemberOnlyRead
	}
	return false
}

func (m *StaticCollectionConfig) GetMemberOnlyWrite() bool {
	if m != nil {
		return m.MemberOnlyWrite
	}
	return false
}

func (m *StaticCollectionConfig) GetEndorsementPolicy() *ApplicationPolicy {
	if m != nil {
		return m.EndorsementPolicy
	}
	return nil
}

// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
type CollectionPolicyConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload              isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CollectionPolicyConfig) Reset()         { *m = CollectionPolicyConfig{} }
func (m *CollectionPolicyConfig) String() string { return proto.CompactTextString(m) }
func (*CollectionPolicyConfig) ProtoMessage()    {}
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8182e05ac5917d8, []int{3}
}

func (m *CollectionPolicyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionPolicyConfig.Unmarshal(m, b)
}
func (m *CollectionPolicyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionPolicyConfig.Marshal(b, m, deterministic)
}
func (m *CollectionPolicyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionPolicyConfig.Merge(m, src)
}
func (m *CollectionPolicyConfig) XXX_Size() int {
	return xxx_messageInfo_CollectionPolicyConfig.Size(m)
}
func (m *CollectionPolicyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionPolicyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionPolicyConfig proto.InternalMessageInfo

type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}

type CollectionPolicyConfig_SignaturePolicy struct {
	SignaturePolicy *common.SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}

func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CollectionPolicyConfig) GetSignaturePolicy() *common.SignaturePolicyEnvelope {
	if x, ok := m.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CollectionPolicyConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
}

func init() {
	proto.RegisterType((*CollectionConfigPackage)(nil), "protos.CollectionConfigPackage")
	proto.RegisterType((*CollectionConfig)(nil), "protos.CollectionConfig")
	proto.RegisterType((*StaticCollectionConfig)(nil), "protos.StaticCollectionConfig")
	proto.RegisterType((*CollectionPolicyConfig)(nil), "protos.CollectionPolicyConfig")
}

func init() { proto.RegisterFile("peer/collection.proto", fileDescriptor_d8182e05ac5917d8) }

var fileDescriptor_d8182e05ac5917d8 = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xd1, 0x6a, 0xdb, 0x3e,
	0x18, 0xc5, 0x9b, 0x7f, 0xd2, 0xb4, 0x55, 0xf8, 0xaf, 0x89, 0x46, 0x33, 0xaf, 0x17, 0x5b, 0xf0,
	0x95, 0x19, 0xad, 0x33, 0xba, 0x27, 0x58, 0xc3, 0x20, 0xb0, 0xc0, 0x82, 0x3a, 0x18, 0xf4, 0xc6,
	0x28, 0xf2, 0x57, 0x57, 0x54, 0x96, 0x5c, 0x49, 0xce, 0xe6, 0x77, 0xd9, 0xc3, 0x0e, 0x4b, 0x71,
	0xed, 0x86, 0x5c, 0x25, 0x9c, 0xf3, 0x3b, 0x9f, 0x3e, 0x1d, 0x64, 0x74, 0x51, 0x00, 0xe8, 0x39,
	0x53, 0x42, 0x00, 0xb3, 0x5c, 0xc9, 0xb8, 0xd0, 0xca, 0x2a, 0x3c, 0x74, 0x3f, 0xe6, 0xf2, 0x82,
	0xa9, 0x3c, 0x57, 0x72, 0x5e, 0x28, 0xc1, 0x19, 0x07, 0xe3, 0xed, 0xcb, 0x89, 0x4b, 0x39, 0xb1,
	0xf2, 0x52, 0xf8, 0x1d, 0xbd, 0x5b, 0xbc, 0x4c, 0x59, 0x28, 0xf9, 0xc0, 0xb3, 0x35, 0x65, 0x4f,
	0x34, 0x03, 0xfc, 0x19, 0x0d, 0x99, 0x13, 0x82, 0xde, 0xac, 0x1f, 0x8d, 0x6e, 0x02, 0x1f, 0x31,
	0xf1, 0x7e, 0x80, 0xec, 0xb8, 0xb0, 0x42, 0xe3, 0x7d, 0x0f, 0xdf, 0xa3, 0xc0, 0x58, 0x6a, 0x39,
	0x4b, 0xda, 0x6d, 0x93, 0x97, 0xb9, 0xbd, 0x68, 0x74, 0xf3, 0xa1, 0x99, 0x7b, 0xe7, 0xb8, 0xfd,
	0x09, 0xcb, 0x23, 0x32, 0x35, 0x07, 0x9d, 0xdb, 0x33, 0x74, 0x52, 0xd0, 0x4a, 0x28, 0x9a, 0x86,
	0x7f, 0xfb, 0x68, 0x7a, 0x38, 0x8f, 0x31, 0x1a, 0x48, 0x9a, 0x83, 0x3b, 0xed, 0x8c, 0xb8, 0xff,
	0x78, 0x85, 0x70, 0x0e, 0xf9, 0x06, 0x74, 0xa2, 0x74, 0x66, 0x12, 0x5f, 0x49, 0xf0, 0xdf, 0xeb,
	0x7d, 0xda, 0x49, 0x6b, 0xe7, 0xef, 0x6e, 0x3b, 0xf6, 0xc9, 0x1f, 0x3a, 0x33, 0x5e, 0xc7, 0x31,
	0x7a, 0xab, 0xe1, 0xb9, 0xe4, 0x1a, 0xd2, 0xa4, 0xae, 0x38, 0x61, 0xaa, 0x94, 0x36, 0xe8, 0xcf,
	0x7a, 0xd1, 0x31, 0x99, 0x34, 0xd6, 0x1a, 0x40, 0x2f, 0x6a, 0x03, 0x5f, 0x21, 0x9c, 0xd3, 0x3f,
	0x3c, 0x2f, 0xf3, 0x2e, 0x3e, 0x70, 0xf8, 0x78, 0xe7, 0xb4, 0x74, 0x88, 0xfe, 0xdf, 0x08, 0xc5,
	0x9e, 0x12, 0xab, 0x12, 0xc1, 0xb7, 0x10, 0x1c, 0xcf, 0x7a, 0xd1, 0x80, 0x8c, 0x9c, 0xf8, 0x53,
	0xad, 0xf8, 0x16, 0x70, 0x84, 0xc6, 0xcd, 0x7d, 0xa4, 0xa8, 0x12, 0x0d, 0x34, 0x0d, 0x86, 0xb3,
	0x5e, 0x74, 0x4a, 0xde, 0xec, 0xb6, 0x95, 0xa2, 0x22, 0x40, 0x53, 0xfc, 0x09, 0x4d, 0xba, 0xe4,
	0x6f, 0xcd, 0x2d, 0x04, 0x27, 0x0e, 0x3d, 0x6f, 0xd1, 0x5f, 0xb5, 0x8c, 0x97, 0x08, 0x83, 0x4c,
	0x95, 0x36, 0x90, 0x83, 0xb4, 0x4d, 0x4b, 0xa7, 0xae, 0xa5, 0xf7, 0x4d, 0x4b, 0x5f, 0x8b, 0x42,
	0x70, 0x46, 0xdb, 0x9a, 0xc8, 0xa4, 0x13, 0xf2, 0x52, 0xf8, 0x8c, 0xa6, 0x87, 0xdb, 0xc4, 0x2b,
	0x34, 0x36, 0x3c, 0x93, 0xd4, 0x96, 0x1a, 0x9a, 0x13, 0xfc, 0xbb, 0xf8, 0x18, 0xfb, 0x57, 0x1c,
	0xdf, 0x35, 0xbe, 0x0f, 0x7e, 0x93, 0x5b, 0x10, 0xaa, 0x80, 0xe5, 0x11, 0x39, 0x37, 0xaf, 0xad,
	0xce, 0x8b, 0xb8, 0x25, 0x28, 0x54, 0x3a, 0x8b, 0x1f, 0xab, 0x02, 0xb4, 0x80, 0x34, 0x03, 0x1d,
	0x3f, 0xd0, 0x8d, 0xe6, 0xac, 0x59, 0xbc, 0xae, 0xfe, 0xfe, 0x2a, 0xe3, 0xf6, 0xb1, 0xdc, 0xd4,
	0x47, 0xcd, 0x3b, 0xe8, 0xdc, 0xa3, 0xd7, 0x1e, 0xbd, 0xce, 0xd4, 0xbc, 0xa6, 0x37, 0xfe, 0xfb,
	0xfa, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x67, 0x04, 0x1f, 0x7f, 0x03, 0x00, 0x00,
}
