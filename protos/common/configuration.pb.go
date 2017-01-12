// Code generated by protoc-gen-go.
// source: common/configuration.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConfigurationItem_ConfigurationType int32

const (
	ConfigurationItem_Policy  ConfigurationItem_ConfigurationType = 0
	ConfigurationItem_Chain   ConfigurationItem_ConfigurationType = 1
	ConfigurationItem_Orderer ConfigurationItem_ConfigurationType = 2
	ConfigurationItem_Fabric  ConfigurationItem_ConfigurationType = 3
)

var ConfigurationItem_ConfigurationType_name = map[int32]string{
	0: "Policy",
	1: "Chain",
	2: "Orderer",
	3: "Fabric",
}
var ConfigurationItem_ConfigurationType_value = map[string]int32{
	"Policy":  0,
	"Chain":   1,
	"Orderer": 2,
	"Fabric":  3,
}

func (x ConfigurationItem_ConfigurationType) String() string {
	return proto.EnumName(ConfigurationItem_ConfigurationType_name, int32(x))
}
func (ConfigurationItem_ConfigurationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{2, 0}
}

type Policy_PolicyType int32

const (
	Policy_UNKNOWN   Policy_PolicyType = 0
	Policy_SIGNATURE Policy_PolicyType = 1
	Policy_MSP       Policy_PolicyType = 2
)

var Policy_PolicyType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SIGNATURE",
	2: "MSP",
}
var Policy_PolicyType_value = map[string]int32{
	"UNKNOWN":   0,
	"SIGNATURE": 1,
	"MSP":       2,
}

func (x Policy_PolicyType) String() string {
	return proto.EnumName(Policy_PolicyType_name, int32(x))
}
func (Policy_PolicyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4, 0} }

// ConfigurationEnvelope is designed to contain _all_ configuration for a chain with no dependency
// on previous configuration transactions.
//
// It is generated with the following scheme:
//   1. Retrieve the existing configuration
//   2. Note the highest configuration sequence number, store it and increment it by one
//   3. Modify desired ConfigurationItems, setting each LastModified to the stored and incremented sequence number
//     a) Note that the ConfigurationItem has a ChainHeader header attached to it, who's type is set to CONFIGURATION_ITEM
//   4. Update SignedConfigurationItem with appropriate signatures over the modified ConfigurationItem
//     a) Each signature is of type ConfigurationSignature
//     b) The ConfigurationSignature signature is over the concatenation of signatureHeader and the ConfigurationItem bytes (which includes a ChainHeader)
//   5. Submit new Configuration for ordering in Envelope signed by submitter
//     a) The Envelope Payload has data set to the marshaled ConfigurationEnvelope
//     b) The Envelope Payload has a header of type Header.Type.CONFIGURATION_TRANSACTION
//
// The configuration manager will verify:
//   1. All configuration items and the envelope refer to the correct chain
//   2. Some configuration item has been added or modified
//   3. No existing configuration item has been ommitted
//   4. All configuration changes have a LastModification of one more than the last configuration's highest LastModification number
//   5. All configuration changes satisfy the corresponding modification policy
type ConfigurationEnvelope struct {
	Items []*SignedConfigurationItem `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
}

func (m *ConfigurationEnvelope) Reset()                    { *m = ConfigurationEnvelope{} }
func (m *ConfigurationEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ConfigurationEnvelope) ProtoMessage()               {}
func (*ConfigurationEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ConfigurationEnvelope) GetItems() []*SignedConfigurationItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// This message may change slightly depending on the finalization of signature schemes for transactions
type SignedConfigurationItem struct {
	ConfigurationItem []byte                    `protobuf:"bytes,1,opt,name=ConfigurationItem,proto3" json:"ConfigurationItem,omitempty"`
	Signatures        []*ConfigurationSignature `protobuf:"bytes,2,rep,name=Signatures" json:"Signatures,omitempty"`
}

func (m *SignedConfigurationItem) Reset()                    { *m = SignedConfigurationItem{} }
func (m *SignedConfigurationItem) String() string            { return proto.CompactTextString(m) }
func (*SignedConfigurationItem) ProtoMessage()               {}
func (*SignedConfigurationItem) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SignedConfigurationItem) GetSignatures() []*ConfigurationSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type ConfigurationItem struct {
	Header             *ChainHeader                        `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Type               ConfigurationItem_ConfigurationType `protobuf:"varint,2,opt,name=Type,enum=common.ConfigurationItem_ConfigurationType" json:"Type,omitempty"`
	LastModified       uint64                              `protobuf:"varint,3,opt,name=LastModified" json:"LastModified,omitempty"`
	ModificationPolicy string                              `protobuf:"bytes,4,opt,name=ModificationPolicy" json:"ModificationPolicy,omitempty"`
	Key                string                              `protobuf:"bytes,5,opt,name=Key" json:"Key,omitempty"`
	Value              []byte                              `protobuf:"bytes,6,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *ConfigurationItem) Reset()                    { *m = ConfigurationItem{} }
func (m *ConfigurationItem) String() string            { return proto.CompactTextString(m) }
func (*ConfigurationItem) ProtoMessage()               {}
func (*ConfigurationItem) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ConfigurationItem) GetHeader() *ChainHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ConfigurationSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signatureHeader,proto3" json:"signatureHeader,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ConfigurationSignature) Reset()                    { *m = ConfigurationSignature{} }
func (m *ConfigurationSignature) String() string            { return proto.CompactTextString(m) }
func (*ConfigurationSignature) ProtoMessage()               {}
func (*ConfigurationSignature) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// Policy expresses a policy which the orderer can evaluate, because there has been some desire expressed to support
// multiple policy engines, this is typed as a oneof for now
type Policy struct {
	Type   int32  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Policy []byte `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

// SignaturePolicyEnvelope wraps a SignaturePolicy and includes a version for future enhancements
type SignaturePolicyEnvelope struct {
	Version    int32            `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Policy     *SignaturePolicy `protobuf:"bytes,2,opt,name=Policy" json:"Policy,omitempty"`
	Identities []*MSPPrincipal  `protobuf:"bytes,3,rep,name=Identities" json:"Identities,omitempty"`
}

func (m *SignaturePolicyEnvelope) Reset()                    { *m = SignaturePolicyEnvelope{} }
func (m *SignaturePolicyEnvelope) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicyEnvelope) ProtoMessage()               {}
func (*SignaturePolicyEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *SignaturePolicyEnvelope) GetPolicy() *SignaturePolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *SignaturePolicyEnvelope) GetIdentities() []*MSPPrincipal {
	if m != nil {
		return m.Identities
	}
	return nil
}

// SignaturePolicy is a recursive message structure which defines a featherweight DSL for describing
// policies which are more complicated than 'exactly this signature'.  The NOutOf operator is sufficent
// to express AND as well as OR, as well as of course N out of the following M policies
// SignedBy implies that the signature is from a valid certificate which is signed by the trusted
// authority specified in the bytes.  This will be the certificate itself for a self-signed certificate
// and will be the CA for more traditional certificates
type SignaturePolicy struct {
	// Types that are valid to be assigned to Type:
	//	*SignaturePolicy_SignedBy
	//	*SignaturePolicy_From
	Type isSignaturePolicy_Type `protobuf_oneof:"Type"`
}

func (m *SignaturePolicy) Reset()                    { *m = SignaturePolicy{} }
func (m *SignaturePolicy) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy) ProtoMessage()               {}
func (*SignaturePolicy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type isSignaturePolicy_Type interface {
	isSignaturePolicy_Type()
}

type SignaturePolicy_SignedBy struct {
	SignedBy int32 `protobuf:"varint,1,opt,name=SignedBy,oneof"`
}
type SignaturePolicy_From struct {
	From *SignaturePolicy_NOutOf `protobuf:"bytes,2,opt,name=From,oneof"`
}

func (*SignaturePolicy_SignedBy) isSignaturePolicy_Type() {}
func (*SignaturePolicy_From) isSignaturePolicy_Type()     {}

func (m *SignaturePolicy) GetType() isSignaturePolicy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SignaturePolicy) GetSignedBy() int32 {
	if x, ok := m.GetType().(*SignaturePolicy_SignedBy); ok {
		return x.SignedBy
	}
	return 0
}

func (m *SignaturePolicy) GetFrom() *SignaturePolicy_NOutOf {
	if x, ok := m.GetType().(*SignaturePolicy_From); ok {
		return x.From
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SignaturePolicy) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SignaturePolicy_OneofMarshaler, _SignaturePolicy_OneofUnmarshaler, _SignaturePolicy_OneofSizer, []interface{}{
		(*SignaturePolicy_SignedBy)(nil),
		(*SignaturePolicy_From)(nil),
	}
}

func _SignaturePolicy_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_From:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.From); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SignaturePolicy.Type has unexpected type %T", x)
	}
	return nil
}

func _SignaturePolicy_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SignaturePolicy)
	switch tag {
	case 1: // Type.SignedBy
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &SignaturePolicy_SignedBy{int32(x)}
		return true, err
	case 2: // Type.From
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SignaturePolicy_NOutOf)
		err := b.DecodeMessage(msg)
		m.Type = &SignaturePolicy_From{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SignaturePolicy_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SignaturePolicy)
	// Type
	switch x := m.Type.(type) {
	case *SignaturePolicy_SignedBy:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SignedBy))
	case *SignaturePolicy_From:
		s := proto.Size(x.From)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SignaturePolicy_NOutOf struct {
	N        int32              `protobuf:"varint,1,opt,name=N" json:"N,omitempty"`
	Policies []*SignaturePolicy `protobuf:"bytes,2,rep,name=Policies" json:"Policies,omitempty"`
}

func (m *SignaturePolicy_NOutOf) Reset()                    { *m = SignaturePolicy_NOutOf{} }
func (m *SignaturePolicy_NOutOf) String() string            { return proto.CompactTextString(m) }
func (*SignaturePolicy_NOutOf) ProtoMessage()               {}
func (*SignaturePolicy_NOutOf) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6, 0} }

func (m *SignaturePolicy_NOutOf) GetPolicies() []*SignaturePolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigurationEnvelope)(nil), "common.ConfigurationEnvelope")
	proto.RegisterType((*SignedConfigurationItem)(nil), "common.SignedConfigurationItem")
	proto.RegisterType((*ConfigurationItem)(nil), "common.ConfigurationItem")
	proto.RegisterType((*ConfigurationSignature)(nil), "common.ConfigurationSignature")
	proto.RegisterType((*Policy)(nil), "common.Policy")
	proto.RegisterType((*SignaturePolicyEnvelope)(nil), "common.SignaturePolicyEnvelope")
	proto.RegisterType((*SignaturePolicy)(nil), "common.SignaturePolicy")
	proto.RegisterType((*SignaturePolicy_NOutOf)(nil), "common.SignaturePolicy.NOutOf")
	proto.RegisterEnum("common.ConfigurationItem_ConfigurationType", ConfigurationItem_ConfigurationType_name, ConfigurationItem_ConfigurationType_value)
	proto.RegisterEnum("common.Policy_PolicyType", Policy_PolicyType_name, Policy_PolicyType_value)
}

func init() { proto.RegisterFile("common/configuration.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x54, 0xc1, 0x4e, 0xdb, 0x40,
	0x10, 0xcd, 0xc6, 0x89, 0x43, 0x26, 0x69, 0x71, 0x07, 0x0a, 0x6e, 0x84, 0x68, 0xe4, 0x93, 0x25,
	0x4a, 0xa2, 0x06, 0x7a, 0x6d, 0x55, 0x10, 0x14, 0x44, 0x71, 0xa2, 0x0d, 0x50, 0xa9, 0xa7, 0x1a,
	0x7b, 0x13, 0x56, 0x4a, 0x6c, 0x6b, 0xed, 0x54, 0xca, 0x17, 0xf4, 0x1f, 0xfa, 0x29, 0x3d, 0xf6,
	0xcb, 0x2a, 0xef, 0xda, 0xae, 0x03, 0xe1, 0xe4, 0x9d, 0x99, 0xf7, 0xde, 0xbc, 0x19, 0xed, 0x1a,
	0x3a, 0x5e, 0x38, 0x9f, 0x87, 0x41, 0xdf, 0x0b, 0x83, 0x09, 0x9f, 0x2e, 0x84, 0x9b, 0xf0, 0x30,
	0xe8, 0x45, 0x22, 0x4c, 0x42, 0xd4, 0x55, 0xad, 0xb3, 0x55, 0x60, 0xd2, 0x8f, 0x2a, 0x76, 0xde,
	0xe4, 0xc9, 0x07, 0x97, 0x07, 0x87, 0x8a, 0xae, 0x4a, 0x96, 0x03, 0xaf, 0x4f, 0xcb, 0x72, 0x67,
	0xc1, 0x4f, 0x36, 0x0b, 0x23, 0x86, 0x1f, 0xa0, 0x7e, 0x99, 0xb0, 0x79, 0x6c, 0x92, 0xae, 0x66,
	0xb7, 0x06, 0x6f, 0x7b, 0x99, 0xe2, 0x98, 0x4f, 0x03, 0xe6, 0xaf, 0x70, 0x52, 0x1c, 0x55, 0x68,
	0xeb, 0x17, 0x81, 0xdd, 0x67, 0x20, 0xf8, 0x0e, 0x5e, 0x3d, 0x49, 0x9a, 0xa4, 0x4b, 0xec, 0x36,
	0x7d, 0x5a, 0xc0, 0x8f, 0x00, 0xa9, 0x90, 0x9b, 0x2c, 0x04, 0x8b, 0xcd, 0xaa, 0x74, 0xb1, 0x9f,
	0xbb, 0x58, 0x81, 0x17, 0x30, 0x5a, 0x62, 0x58, 0x7f, 0xab, 0x6b, 0xda, 0xe1, 0x01, 0xe8, 0x17,
	0xcc, 0xf5, 0x99, 0x90, 0x8d, 0x5b, 0x83, 0xad, 0x42, 0x31, 0xdd, 0x8d, 0x2a, 0xd1, 0x0c, 0x82,
	0x9f, 0xa0, 0x76, 0xb3, 0x8c, 0x98, 0x59, 0xed, 0x12, 0xfb, 0xe5, 0xe0, 0x60, 0x6d, 0xf3, 0x54,
	0x75, 0x35, 0x93, 0x52, 0xa8, 0x24, 0xa2, 0x05, 0xed, 0xaf, 0x6e, 0x9c, 0x5c, 0x87, 0x3e, 0x9f,
	0x70, 0xe6, 0x9b, 0x5a, 0x97, 0xd8, 0x35, 0xba, 0x92, 0xc3, 0x1e, 0xa0, 0x3a, 0x7b, 0x92, 0x3d,
	0x0a, 0x67, 0xdc, 0x5b, 0x9a, 0xb5, 0x2e, 0xb1, 0x9b, 0x74, 0x4d, 0x05, 0x0d, 0xd0, 0xae, 0xd8,
	0xd2, 0xac, 0x4b, 0x40, 0x7a, 0xc4, 0x6d, 0xa8, 0xdf, 0xb9, 0xb3, 0x05, 0x33, 0x75, 0xb9, 0x4b,
	0x15, 0x58, 0xa7, 0x8f, 0xc6, 0x97, 0x86, 0x00, 0x74, 0x25, 0x63, 0x54, 0xb0, 0x09, 0x75, 0x39,
	0xb4, 0x41, 0xb0, 0x05, 0x8d, 0xa1, 0xf0, 0x99, 0x60, 0xc2, 0xa8, 0xa6, 0x98, 0x73, 0xf7, 0x5e,
	0x70, 0xcf, 0xd0, 0xac, 0x1f, 0xb0, 0xb3, 0x7e, 0xd5, 0x68, 0xc3, 0x66, 0x9c, 0x07, 0xa5, 0x8d,
	0xb6, 0xe9, 0xe3, 0x34, 0xee, 0x41, 0xb3, 0x48, 0xc9, 0x55, 0xb6, 0xe9, 0xff, 0x84, 0x35, 0xcd,
	0x1d, 0x21, 0x42, 0x2d, 0x49, 0xb7, 0x9d, 0xca, 0xd4, 0xa9, 0x3c, 0xe3, 0x0e, 0xe8, 0x91, 0x5a,
	0x88, 0x22, 0x66, 0x91, 0xf5, 0x1e, 0x40, 0xb1, 0xe4, 0x54, 0x2d, 0x68, 0xdc, 0x3a, 0x57, 0xce,
	0xf0, 0x9b, 0x63, 0x54, 0xf0, 0x05, 0x34, 0xc7, 0x97, 0x5f, 0x9c, 0xcf, 0x37, 0xb7, 0xf4, 0xcc,
	0x20, 0xd8, 0x00, 0xed, 0x7a, 0x3c, 0x32, 0xaa, 0xd6, 0xef, 0xec, 0x66, 0xca, 0xb6, 0x8a, 0x5c,
	0x5c, 0x76, 0x13, 0x1a, 0x77, 0x4c, 0xc4, 0x3c, 0x0c, 0xb2, 0xee, 0x79, 0x88, 0xfd, 0xdc, 0x9e,
	0x34, 0xd0, 0x1a, 0xec, 0x96, 0xdf, 0x41, 0x49, 0x8a, 0xe6, 0x53, 0x1c, 0x03, 0x5c, 0xfa, 0x2c,
	0x48, 0x78, 0xc2, 0x59, 0x6c, 0x6a, 0xf2, 0xda, 0x6e, 0xe7, 0xa4, 0xeb, 0xf1, 0x68, 0x24, 0x78,
	0xe0, 0xf1, 0xc8, 0x9d, 0xd1, 0x12, 0xce, 0xfa, 0x43, 0x60, 0xf3, 0x91, 0x22, 0xee, 0xc1, 0x86,
	0x7a, 0x49, 0x27, 0x4b, 0xe5, 0xea, 0xa2, 0x42, 0x8b, 0x0c, 0x1e, 0x43, 0xed, 0x5c, 0x84, 0xf3,
	0xcc, 0xd6, 0xfe, 0x33, 0xb6, 0x7a, 0xce, 0x70, 0x91, 0x0c, 0x27, 0x17, 0x15, 0x2a, 0xd1, 0x9d,
	0x2b, 0xd0, 0x55, 0x06, 0xdb, 0x40, 0x9c, 0x6c, 0x58, 0xe2, 0xe0, 0x11, 0x6c, 0x48, 0x02, 0x2f,
	0x9e, 0xda, 0xb3, 0x83, 0x16, 0xc0, 0x13, 0x5d, 0x3d, 0x8f, 0x93, 0xc3, 0xef, 0x07, 0x53, 0x9e,
	0x3c, 0x2c, 0xee, 0x53, 0x4a, 0xff, 0x61, 0x19, 0x31, 0x31, 0x63, 0xfe, 0x94, 0x89, 0xfe, 0x44,
	0xde, 0xa3, 0xbe, 0xfc, 0xd5, 0xc4, 0xd9, 0x3f, 0xe9, 0x5e, 0x97, 0xe1, 0xd1, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x94, 0x93, 0x24, 0x92, 0xcf, 0x04, 0x00, 0x00,
}