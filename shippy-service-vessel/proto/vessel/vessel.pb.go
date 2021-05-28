// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() {
	proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716)
}

var fileDescriptor_04ef66862bb50716 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0xd3, 0x36, 0x4d, 0x46, 0x52, 0x64, 0x40, 0xd8, 0x16, 0x85, 0x90, 0x83, 0xe4, 0x54,
	0xa1, 0xde, 0xbc, 0x79, 0x11, 0xd4, 0x5b, 0x0a, 0x7a, 0x11, 0xca, 0x36, 0x19, 0x75, 0x20, 0x5f,
	0x24, 0x21, 0xad, 0xff, 0xc6, 0x9f, 0x2a, 0x4c, 0x92, 0x4a, 0x45, 0x3c, 0xe5, 0x7d, 0x64, 0x86,
	0x37, 0x6f, 0x61, 0x5e, 0x56, 0x45, 0x53, 0x5c, 0xb7, 0x54, 0xd7, 0x94, 0xf6, 0x9f, 0xa5, 0x68,
	0x68, 0x77, 0x2c, 0xf8, 0x52, 0x60, 0x3f, 0x0b, 0xc4, 0x19, 0x58, 0x9c, 0x68, 0xe5, 0xab, 0xd0,
	0x8d, 0x2c, 0x4e, 0x70, 0x01, 0x4e, 0x6c, 0x4a, 0x13, 0x73, 0xf3, 0xa9, 0x2d, 0x5f, 0x85, 0x93,
	0xe8, 0xc0, 0xf1, 0x12, 0x20, 0x33, 0xfb, 0xcd, 0x8e, 0xf8, 0xfd, 0xa3, 0xd1, 0x23, 0x71, 0xdd,
	0xcc, 0xec, 0x5f, 0x44, 0x40, 0x84, 0x71, 0x6e, 0x32, 0xd2, 0x63, 0x59, 0x26, 0x18, 0x2f, 0xc0,
	0x35, 0xad, 0xe1, 0xd4, 0x6c, 0x53, 0xd2, 0x13, 0x5f, 0x85, 0x4e, 0xf4, 0x23, 0xe0, 0x1c, 0x9c,
	0x62, 0x97, 0x53, 0xb5, 0xe1, 0x44, 0xdb, 0x32, 0x35, 0x15, 0xfe, 0x90, 0x04, 0x8f, 0xe0, 0xad,
	0x4b, 0x8a, 0xf9, 0x8d, 0x63, 0xd3, 0x70, 0x91, 0x1f, 0x05, 0x53, 0xff, 0x06, 0xb3, 0x7e, 0x05,
	0x0b, 0x5e, 0xc1, 0x89, 0xa8, 0x2e, 0x8b, 0xbc, 0x26, 0xbc, 0x82, 0xbe, 0x04, 0x59, 0x72, 0xba,
	0x9a, 0x2d, 0xfb, 0x86, 0xba, 0x3e, 0xa2, 0xde, 0xc5, 0x10, 0xa6, 0x1d, 0xaa, 0xb5, 0xe5, 0x8f,
	0xfe, 0xf8, 0x71, 0xb0, 0x57, 0x4f, 0xe0, 0x75, 0xd2, 0x9a, 0xaa, 0x96, 0x63, 0xc2, 0x5b, 0xf0,
	0xee, 0x39, 0x4f, 0xee, 0x0e, 0x67, 0x9e, 0x0f, 0xa3, 0x47, 0x17, 0x2d, 0xce, 0x06, 0x79, 0x08,
	0x17, 0x9c, 0x6c, 0x6d, 0x79, 0xa8, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x8c, 0xe2,
	0x74, 0xc5, 0x01, 0x00, 0x00,
}
