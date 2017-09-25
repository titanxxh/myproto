// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ProtoBufA
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoBufA struct {
	Name     string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	BirthDay int64   `protobuf:"varint,2,opt,name=birthDay" json:"birthDay,omitempty"`
	Phone    string  `protobuf:"bytes,3,opt,name=phone" json:"phone,omitempty"`
	Siblings int32   `protobuf:"varint,4,opt,name=siblings" json:"siblings,omitempty"`
	Spouse   bool    `protobuf:"varint,5,opt,name=spouse" json:"spouse,omitempty"`
	Money    float64 `protobuf:"fixed64,6,opt,name=money" json:"money,omitempty"`
}

func (m *ProtoBufA) Reset()                    { *m = ProtoBufA{} }
func (m *ProtoBufA) String() string            { return proto.CompactTextString(m) }
func (*ProtoBufA) ProtoMessage()               {}
func (*ProtoBufA) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoBufA) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtoBufA) GetBirthDay() int64 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

func (m *ProtoBufA) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *ProtoBufA) GetSiblings() int32 {
	if m != nil {
		return m.Siblings
	}
	return 0
}

func (m *ProtoBufA) GetSpouse() bool {
	if m != nil {
		return m.Spouse
	}
	return false
}

func (m *ProtoBufA) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoBufA)(nil), "api.ProtoBufA")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x9a, 0xcd, 0xc8, 0xc5, 0x19,
	0x00, 0xe2, 0x3a, 0x95, 0xa6, 0x39, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x52, 0x5c, 0x1c, 0x49, 0x99, 0x45, 0x25, 0x19,
	0x2e, 0x89, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x6b,
	0x41, 0x46, 0x7e, 0x5e, 0xaa, 0x04, 0x33, 0x58, 0x03, 0x84, 0x03, 0xd2, 0x51, 0x9c, 0x99, 0x94,
	0x93, 0x99, 0x97, 0x5e, 0x2c, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe7, 0x0b, 0x89, 0x71,
	0xb1, 0x15, 0x17, 0xe4, 0x97, 0x16, 0xa7, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x41, 0x79,
	0x20, 0x93, 0x72, 0xf3, 0xf3, 0x52, 0x2b, 0x25, 0xd8, 0x14, 0x18, 0x35, 0x18, 0x83, 0x20, 0x9c,
	0x24, 0x36, 0xb0, 0x4b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x69, 0xc2, 0x5b, 0xb6,
	0x00, 0x00, 0x00,
}
