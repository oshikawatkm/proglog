// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/log.proto

package log_v1

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

type ProduceRequest struct {
	Record               *Record  `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{0}
}

func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRequest.Unmarshal(m, b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return xxx_messageInfo_ProduceRequest.Size(m)
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type ProduceResponse struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{1}
}

func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResponse.Unmarshal(m, b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
}
func (m *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(m, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return xxx_messageInfo_ProduceResponse.Size(m)
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

func (m *ProduceResponse) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{2}
}

func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRequest.Unmarshal(m, b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeRequest.Size(m)
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

func (m *ConsumeRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeResponse struct {
	Record               *Record  `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeResponse) Reset()         { *m = ConsumeResponse{} }
func (m *ConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeResponse) ProtoMessage()    {}
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{3}
}

func (m *ConsumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeResponse.Unmarshal(m, b)
}
func (m *ConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeResponse.Marshal(b, m, deterministic)
}
func (m *ConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeResponse.Merge(m, src)
}
func (m *ConsumeResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumeResponse.Size(m)
}
func (m *ConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeResponse proto.InternalMessageInfo

func (m *ConsumeResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Record struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{4}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*ProduceRequest)(nil), "log.v1.ProduceRequest")
	proto.RegisterType((*ProduceResponse)(nil), "log.v1.ProduceResponse")
	proto.RegisterType((*ConsumeRequest)(nil), "log.v1.ConsumeRequest")
	proto.RegisterType((*ConsumeResponse)(nil), "log.v1.ConsumeResponse")
	proto.RegisterType((*Record)(nil), "log.v1.Record")
}

func init() {
	proto.RegisterFile("api/v1/log.proto", fileDescriptor_19a5c3fde3f7ae80)
}

var fileDescriptor_19a5c3fde3f7ae80 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0x2e, 0xa8, 0x98, 0x8c, 0x95, 0x9a, 0x8d, 0xa9, 0xc6, 0x93, 0x62, 0x62, 0xf0, 0x02, 0xa5,
	0x26, 0x46, 0x13, 0x4f, 0x36, 0xf1, 0xe4, 0xc1, 0xd0, 0x9b, 0x17, 0xb3, 0xa5, 0x5b, 0x4a, 0x0a,
	0x1d, 0xdc, 0x1f, 0x7c, 0x0b, 0x9f, 0xd9, 0x00, 0x2b, 0x14, 0x8d, 0xc6, 0x78, 0x63, 0x3e, 0xe6,
	0xfb, 0xdb, 0x0c, 0x1c, 0xd0, 0x3c, 0xf1, 0x8b, 0xc0, 0x4f, 0x31, 0xf6, 0x72, 0x8e, 0x12, 0x89,
	0x55, 0x7e, 0x16, 0x81, 0x73, 0x03, 0xf6, 0x13, 0xc7, 0xb9, 0x8a, 0x58, 0xc8, 0x5e, 0x15, 0x13,
	0x92, 0x5c, 0x80, 0xc5, 0x59, 0x84, 0x7c, 0x7e, 0x6c, 0x9c, 0x1a, 0xee, 0xde, 0xd8, 0xf6, 0xea,
	0x55, 0x2f, 0xac, 0xd0, 0x50, 0xff, 0x75, 0x2e, 0x61, 0xd0, 0x30, 0x45, 0x8e, 0x6b, 0xc1, 0xc8,
	0x10, 0x2c, 0x5c, 0x2c, 0x04, 0x93, 0x15, 0x75, 0x3b, 0xd4, 0x93, 0xe3, 0x82, 0x3d, 0xc1, 0xb5,
	0x50, 0x59, 0x63, 0xf2, 0xd3, 0xe6, 0x2d, 0x0c, 0x9a, 0x4d, 0x2d, 0xda, 0xe6, 0x31, 0x7f, 0xcd,
	0x73, 0x0d, 0x56, 0x8d, 0x90, 0x43, 0xd8, 0x29, 0x68, 0xaa, 0x58, 0xa5, 0xdd, 0x0f, 0xeb, 0x61,
	0xc3, 0xd2, 0xdc, 0xb4, 0x1c, 0xbf, 0x9b, 0xb0, 0xf5, 0x88, 0x31, 0xb9, 0x83, 0x5d, 0xdd, 0x87,
	0x0c, 0x3f, 0x2d, 0xba, 0x4f, 0x73, 0x72, 0xf4, 0x0d, 0xaf, 0x33, 0x3a, 0xbd, 0x92, 0xad, 0x83,
	0xb7, 0xec, 0x6e, 0xe7, 0x96, 0xfd, 0xa5, 0xa1, 0xd3, 0x23, 0x13, 0xe8, 0x6b, 0x70, 0xca, 0x19,
	0xcd, 0xfe, 0x21, 0x31, 0x32, 0xc8, 0x03, 0xec, 0xeb, 0x5c, 0x53, 0xd9, 0x55, 0xf9, 0x73, 0x0d,
	0xd7, 0x18, 0x19, 0xf7, 0xe7, 0xcf, 0x67, 0x71, 0x22, 0x97, 0x6a, 0xe6, 0x45, 0x98, 0xf9, 0x28,
	0x96, 0xc9, 0x8a, 0xbe, 0x51, 0xb9, 0xca, 0xfc, 0xf2, 0x8a, 0x52, 0x8c, 0x5f, 0x8a, 0x60, 0x66,
	0x55, 0x67, 0x74, 0xf5, 0x11, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x61, 0x57, 0xcb, 0x5a, 0x02, 0x00,
	0x00,
}