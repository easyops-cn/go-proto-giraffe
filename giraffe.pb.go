// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/easyops-cn/go-proto-giraffe/giraffe.proto

package giraffeproto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Contract struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_45c7a1bac31e4345, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contract.Unmarshal(m, b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return xxx_messageInfo_Contract.Size(m)
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contract) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

var E_UrlPattern = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50001,
	Name:          "giraffeproto.url_pattern",
	Tag:           "bytes,50001,opt,name=url_pattern",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_ContractName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50002,
	Name:          "giraffeproto.contract_name",
	Tag:           "bytes,50002,opt,name=contract_name",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_ContractVersion = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50003,
	Name:          "giraffeproto.contract_version",
	Tag:           "bytes,50003,opt,name=contract_version",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_DataField = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50004,
	Name:          "giraffeproto.data_field",
	Tag:           "bytes,50004,opt,name=data_field",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_QueryField = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50005,
	Name:          "giraffeproto.query_field",
	Tag:           "bytes,50005,opt,name=query_field",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_ContentType = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50006,
	Name:          "giraffeproto.content_type",
	Tag:           "bytes,50006,opt,name=content_type",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_Contract = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Contract)(nil),
	Field:         72295000,
	Name:          "giraffeproto.contract",
	Tag:           "bytes,72295000,opt,name=contract",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         72295728,
	Name:          "giraffeproto.http",
	Tag:           "bytes,72295728,opt,name=http",
	Filename:      "github.com/easyops-cn/go-proto-giraffe/giraffe.proto",
}

func init() {
	proto.RegisterType((*Contract)(nil), "giraffeproto.Contract")
	proto.RegisterExtension(E_UrlPattern)
	proto.RegisterExtension(E_ContractName)
	proto.RegisterExtension(E_ContractVersion)
	proto.RegisterExtension(E_DataField)
	proto.RegisterExtension(E_QueryField)
	proto.RegisterExtension(E_ContentType)
	proto.RegisterExtension(E_Contract)
	proto.RegisterExtension(E_Http)
}

func init() {
	proto.RegisterFile("github.com/easyops-cn/go-proto-giraffe/giraffe.proto", fileDescriptor_45c7a1bac31e4345)
}

var fileDescriptor_45c7a1bac31e4345 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4a, 0xfb, 0x40,
	0x10, 0xc0, 0xe9, 0xff, 0x5f, 0xb4, 0xdd, 0x56, 0x94, 0x1c, 0x24, 0xf4, 0x20, 0xa5, 0xa7, 0x5e,
	0x9a, 0xa0, 0x55, 0x90, 0x78, 0x10, 0x2d, 0x8a, 0x20, 0x55, 0x09, 0xea, 0xc1, 0x4b, 0xd8, 0x26,
	0x93, 0x34, 0x90, 0x66, 0xd7, 0xcd, 0x44, 0xc8, 0x0b, 0xf8, 0x3c, 0x3e, 0x80, 0x47, 0x4f, 0xde,
	0xfc, 0xc4, 0xc7, 0x91, 0x6c, 0x36, 0xa5, 0x07, 0x21, 0x3d, 0x6d, 0x66, 0xc8, 0xef, 0x37, 0x1f,
	0x0c, 0xd9, 0x0d, 0x42, 0x9c, 0xa6, 0x13, 0xc3, 0x65, 0x33, 0x13, 0x68, 0x92, 0x31, 0x9e, 0x0c,
	0xdc, 0xd8, 0x0c, 0xd8, 0x80, 0x0b, 0x86, 0x6c, 0x10, 0x84, 0x82, 0xfa, 0x3e, 0x98, 0xea, 0x35,
	0x64, 0x56, 0x6b, 0xab, 0x50, 0x46, 0x9d, 0x6e, 0xc0, 0x58, 0x10, 0x81, 0x29, 0xa3, 0x49, 0xea,
	0x9b, 0x1e, 0x24, 0xae, 0x08, 0x39, 0x32, 0x51, 0xfc, 0xdf, 0xd9, 0x5e, 0xb2, 0xca, 0x14, 0x91,
	0x17, 0x48, 0x6f, 0x9f, 0x34, 0x46, 0x2c, 0x46, 0x41, 0x5d, 0xd4, 0x34, 0x52, 0x8f, 0xe9, 0x0c,
	0xf4, 0x5a, 0xb7, 0xd6, 0x6f, 0xda, 0xf2, 0x5b, 0xd3, 0xc9, 0xea, 0x03, 0x88, 0x24, 0x64, 0xb1,
	0xfe, 0x4f, 0xa6, 0xcb, 0xd0, 0x3a, 0x22, 0xad, 0x54, 0x44, 0x0e, 0xa7, 0x88, 0x20, 0x62, 0x6d,
	0xcb, 0x28, 0xda, 0x33, 0xca, 0xf6, 0x8c, 0x31, 0xe0, 0x94, 0x79, 0x97, 0x1c, 0x43, 0x16, 0x27,
	0xfa, 0xdb, 0xe3, 0x7f, 0xc9, 0x93, 0x54, 0x44, 0x57, 0x05, 0x63, 0x9d, 0x90, 0x35, 0x57, 0x15,
	0x77, 0x64, 0xb5, 0x2a, 0xc9, 0xbb, 0x92, 0xb4, 0x4b, 0xec, 0x82, 0xce, 0xc0, 0x3a, 0x27, 0x1b,
	0x73, 0x8d, 0xea, 0xae, 0xd2, 0xf4, 0xa1, 0x4c, 0xeb, 0x25, 0x79, 0xab, 0xc6, 0x3a, 0x24, 0xc4,
	0xa3, 0x48, 0x1d, 0x3f, 0x84, 0xc8, 0xab, 0xd4, 0x7c, 0x2a, 0x4d, 0x33, 0x67, 0x4e, 0x73, 0x24,
	0xdf, 0xcb, 0x7d, 0x0a, 0x22, 0x5b, 0xd2, 0xf0, 0x55, 0xee, 0x45, 0x42, 0x85, 0x62, 0x44, 0xe4,
	0x80, 0x10, 0xa3, 0x83, 0x19, 0xaf, 0x5e, 0xcb, 0xb7, 0x72, 0xb4, 0x14, 0x75, 0x9d, 0x71, 0xb0,
	0x6e, 0x48, 0xa3, 0x9c, 0xad, 0x52, 0xf0, 0xf3, 0xf2, 0xdc, 0xeb, 0xd6, 0xfa, 0xad, 0x9d, 0x4d,
	0x63, 0xf1, 0xe4, 0x8c, 0xf2, 0x34, 0xec, 0xb9, 0xca, 0x1a, 0x93, 0x7a, 0x7e, 0x3e, 0x95, 0xca,
	0xa7, 0xd7, 0xbf, 0x95, 0x67, 0x88, 0xdc, 0x4e, 0x23, 0xb0, 0xa5, 0xe6, 0x78, 0xef, 0x6e, 0xb8,
	0xdc, 0xd1, 0x1e, 0x2c, 0x5a, 0x26, 0x2b, 0xf2, 0x19, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x82, 0xe1, 0x5a, 0x58, 0x03, 0x00, 0x00,
}
