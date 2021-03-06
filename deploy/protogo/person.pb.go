// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package person

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

type Person struct {
	Email                []*Person_Email `protobuf:"bytes,1,rep,name=email,proto3" json:"email,omitempty"`
	Person               *Person_Person  `protobuf:"bytes,2,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetEmail() []*Person_Email {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Person) GetPerson() *Person_Person {
	if m != nil {
		return m.Person
	}
	return nil
}

type Person_Email struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person_Email) Reset()         { *m = Person_Email{} }
func (m *Person_Email) String() string { return proto.CompactTextString(m) }
func (*Person_Email) ProtoMessage()    {}
func (*Person_Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 0}
}

func (m *Person_Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_Email.Unmarshal(m, b)
}
func (m *Person_Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_Email.Marshal(b, m, deterministic)
}
func (m *Person_Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_Email.Merge(m, src)
}
func (m *Person_Email) XXX_Size() int {
	return xxx_messageInfo_Person_Email.Size(m)
}
func (m *Person_Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Person_Email proto.InternalMessageInfo

func (m *Person_Email) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person_Email) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type Person_Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex                  string   `protobuf:"bytes,2,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person_Person) Reset()         { *m = Person_Person{} }
func (m *Person_Person) String() string { return proto.CompactTextString(m) }
func (*Person_Person) ProtoMessage()    {}
func (*Person_Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0, 1}
}

func (m *Person_Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_Person.Unmarshal(m, b)
}
func (m *Person_Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_Person.Marshal(b, m, deterministic)
}
func (m *Person_Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_Person.Merge(m, src)
}
func (m *Person_Person) XXX_Size() int {
	return xxx_messageInfo_Person_Person.Size(m)
}
func (m *Person_Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person_Person proto.InternalMessageInfo

func (m *Person_Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person_Person) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "person.Person")
	proto.RegisterType((*Person_Email)(nil), "person.Person.Email")
	proto.RegisterType((*Person_Person)(nil), "person.Person.Person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x0e, 0x30, 0x72,
	0xb1, 0x05, 0x80, 0x99, 0x42, 0x5a, 0x5c, 0xac, 0xa9, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x50, 0x0d, 0x10, 0x69, 0x3d, 0x57, 0x90, 0x5c, 0x10, 0x44,
	0x89, 0x90, 0x2e, 0x17, 0xd4, 0x00, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x51, 0x34, 0xc5,
	0x10, 0x2a, 0x08, 0xaa, 0x48, 0x4a, 0x9b, 0x8b, 0x15, 0xac, 0x5d, 0x88, 0x8f, 0x8b, 0x29, 0x33,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x3b,
	0x33, 0x2f, 0x05, 0x6c, 0x0a, 0x67, 0x10, 0x98, 0x2d, 0xa5, 0x07, 0x77, 0x91, 0x10, 0x17, 0x4b,
	0x5e, 0x62, 0x6e, 0x2a, 0x54, 0x3d, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x5c, 0x9c, 0x5a, 0x01, 0xd5,
	0x00, 0x62, 0x26, 0xb1, 0x81, 0x7d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xde, 0xef, 0x2b,
	0x1a, 0xe1, 0x00, 0x00, 0x00,
}
