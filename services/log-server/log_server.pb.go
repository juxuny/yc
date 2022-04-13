// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log_server.proto

package log_server // import "./"

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

type LogLevel int32

const (
	LogLevel_Unknown LogLevel = 0
	LogLevel_Debug   LogLevel = 1
	LogLevel_Info    LogLevel = 2
	LogLevel_Warn    LogLevel = 3
	LogLevel_Error   LogLevel = 4
)

var LogLevel_name = map[int32]string{
	0: "Unknown",
	1: "Debug",
	2: "Info",
	3: "Warn",
	4: "Error",
}
var LogLevel_value = map[string]int32{
	"Unknown": 0,
	"Debug":   1,
	"Info":    2,
	"Warn":    3,
	"Error":   4,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_log_server_620ea295793a1768, []int{0}
}

type HealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *HealthRequest) Reset()         { *m = HealthRequest{} }
func (m *HealthRequest) String() string { return proto.CompactTextString(m) }
func (*HealthRequest) ProtoMessage()    {}
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_server_620ea295793a1768, []int{0}
}
func (m *HealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthRequest.Unmarshal(m, b)
}
func (m *HealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthRequest.Marshal(b, m, deterministic)
}
func (dst *HealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthRequest.Merge(dst, src)
}
func (m *HealthRequest) XXX_Size() int {
	return xxx_messageInfo_HealthRequest.Size(m)
}
func (m *HealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthRequest proto.InternalMessageInfo

type HealthResponse struct {
	CurrentTime          string   `protobuf:"bytes,1,opt,name=currentTime" json:"currentTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *HealthResponse) Reset()         { *m = HealthResponse{} }
func (m *HealthResponse) String() string { return proto.CompactTextString(m) }
func (*HealthResponse) ProtoMessage()    {}
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_server_620ea295793a1768, []int{1}
}
func (m *HealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthResponse.Unmarshal(m, b)
}
func (m *HealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthResponse.Marshal(b, m, deterministic)
}
func (dst *HealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthResponse.Merge(dst, src)
}
func (m *HealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthResponse.Size(m)
}
func (m *HealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthResponse proto.InternalMessageInfo

func (m *HealthResponse) GetCurrentTime() string {
	if m != nil {
		return m.CurrentTime
	}
	return ""
}

type KeyPairs struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *KeyPairs) Reset()         { *m = KeyPairs{} }
func (m *KeyPairs) String() string { return proto.CompactTextString(m) }
func (*KeyPairs) ProtoMessage()    {}
func (*KeyPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_server_620ea295793a1768, []int{2}
}
func (m *KeyPairs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPairs.Unmarshal(m, b)
}
func (m *KeyPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPairs.Marshal(b, m, deterministic)
}
func (dst *KeyPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPairs.Merge(dst, src)
}
func (m *KeyPairs) XXX_Size() int {
	return xxx_messageInfo_KeyPairs.Size(m)
}
func (m *KeyPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPairs.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPairs proto.InternalMessageInfo

func (m *KeyPairs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyPairs) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PrintRequest struct {
	Level                LogLevel    `protobuf:"varint,1,opt,name=level,enum=log_server.LogLevel" json:"level,omitempty"`
	Content              string      `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	Extra                []*KeyPairs `protobuf:"bytes,3,rep,name=extra" json:"extra,omitempty"`
	FileLine             string      `protobuf:"bytes,4,opt,name=fileLine" json:"fileLine,omitempty"`
	ReqId                string      `protobuf:"bytes,5,opt,name=reqId" json:"reqId,omitempty"`
	DateTime             string      `protobuf:"bytes,6,opt,name=dateTime" json:"dateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-" bson:"-"`
	XXX_unrecognized     []byte      `json:"-" bson:"-"`
	XXX_sizecache        int32       `json:"-" bson:"-"`
}

func (m *PrintRequest) Reset()         { *m = PrintRequest{} }
func (m *PrintRequest) String() string { return proto.CompactTextString(m) }
func (*PrintRequest) ProtoMessage()    {}
func (*PrintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_server_620ea295793a1768, []int{3}
}
func (m *PrintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintRequest.Unmarshal(m, b)
}
func (m *PrintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintRequest.Marshal(b, m, deterministic)
}
func (dst *PrintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintRequest.Merge(dst, src)
}
func (m *PrintRequest) XXX_Size() int {
	return xxx_messageInfo_PrintRequest.Size(m)
}
func (m *PrintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrintRequest proto.InternalMessageInfo

func (m *PrintRequest) GetLevel() LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_Unknown
}

func (m *PrintRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PrintRequest) GetExtra() []*KeyPairs {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *PrintRequest) GetFileLine() string {
	if m != nil {
		return m.FileLine
	}
	return ""
}

func (m *PrintRequest) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

func (m *PrintRequest) GetDateTime() string {
	if m != nil {
		return m.DateTime
	}
	return ""
}

type PrintResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *PrintResponse) Reset()         { *m = PrintResponse{} }
func (m *PrintResponse) String() string { return proto.CompactTextString(m) }
func (*PrintResponse) ProtoMessage()    {}
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_server_620ea295793a1768, []int{4}
}
func (m *PrintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintResponse.Unmarshal(m, b)
}
func (m *PrintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintResponse.Marshal(b, m, deterministic)
}
func (dst *PrintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintResponse.Merge(dst, src)
}
func (m *PrintResponse) XXX_Size() int {
	return xxx_messageInfo_PrintResponse.Size(m)
}
func (m *PrintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrintResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HealthRequest)(nil), "log_server.HealthRequest")
	proto.RegisterType((*HealthResponse)(nil), "log_server.HealthResponse")
	proto.RegisterType((*KeyPairs)(nil), "log_server.KeyPairs")
	proto.RegisterType((*PrintRequest)(nil), "log_server.PrintRequest")
	proto.RegisterType((*PrintResponse)(nil), "log_server.PrintResponse")
	proto.RegisterEnum("log_server.LogLevel", LogLevel_name, LogLevel_value)
}

func init() { proto.RegisterFile("log_server.proto", fileDescriptor_log_server_620ea295793a1768) }

var fileDescriptor_log_server_620ea295793a1768 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xa5, 0xd0, 0x42, 0x19, 0x3e, 0xa0, 0xd9, 0x70, 0xe8, 0xc7, 0x89, 0xf4, 0x44, 0x38, 0x60,
	0x82, 0xde, 0x3c, 0x61, 0x34, 0x91, 0xd8, 0x18, 0x52, 0x35, 0x26, 0x5e, 0xcc, 0x02, 0x43, 0x6d,
	0xac, 0xbb, 0xb0, 0xdd, 0xa2, 0xfe, 0x49, 0x7f, 0x93, 0xd9, 0x5d, 0x0a, 0x25, 0xf1, 0xb6, 0xef,
	0xcd, 0x9b, 0xd7, 0x37, 0x33, 0x05, 0x2f, 0xe5, 0xf1, 0x6b, 0x86, 0x62, 0x87, 0x62, 0xbc, 0x11,
	0x5c, 0x72, 0x02, 0x47, 0x26, 0xe8, 0x42, 0xfb, 0x16, 0x69, 0x2a, 0xdf, 0x22, 0xdc, 0xe6, 0x98,
	0xc9, 0x60, 0x02, 0x9d, 0x82, 0xc8, 0x36, 0x9c, 0x65, 0x48, 0x06, 0xd0, 0x5a, 0xe6, 0x42, 0x20,
	0x93, 0x8f, 0xc9, 0x07, 0xfa, 0xd6, 0xc0, 0x1a, 0x36, 0xa3, 0x32, 0x15, 0x5c, 0x80, 0x7b, 0x87,
	0xdf, 0x73, 0x9a, 0x88, 0x8c, 0x10, 0xb0, 0x19, 0x3d, 0xc8, 0xf4, 0x9b, 0xf4, 0xc0, 0xd9, 0xd1,
	0x34, 0x47, 0xbf, 0xaa, 0x49, 0x03, 0x82, 0x1f, 0x0b, 0xfe, 0xcd, 0x45, 0xc2, 0xe4, 0xfe, 0xd3,
	0x64, 0x04, 0x4e, 0x8a, 0x3b, 0x4c, 0x75, 0x6f, 0x67, 0xd2, 0x1b, 0x97, 0x92, 0x87, 0x3c, 0x0e,
	0x55, 0x2d, 0x32, 0x12, 0xe2, 0x43, 0x63, 0xc9, 0x99, 0x44, 0x26, 0xf7, 0xa6, 0x05, 0x54, 0x2e,
	0xf8, 0x25, 0x05, 0xf5, 0x6b, 0x83, 0xda, 0xb0, 0x75, 0xea, 0x52, 0xa4, 0x8c, 0x8c, 0x84, 0xf4,
	0xc1, 0x5d, 0x27, 0x29, 0x86, 0x09, 0x43, 0xdf, 0xd6, 0x36, 0x07, 0xac, 0x42, 0x0b, 0xdc, 0xce,
	0x56, 0xbe, 0x63, 0x42, 0x6b, 0xa0, 0x3a, 0x56, 0x54, 0xa2, 0xde, 0x44, 0xdd, 0x74, 0x14, 0x58,
	0xed, 0x72, 0x3f, 0x8f, 0xd9, 0xdc, 0x68, 0x0a, 0x6e, 0x91, 0x9b, 0xb4, 0xa0, 0xf1, 0xc4, 0xde,
	0x19, 0xff, 0x64, 0x5e, 0x85, 0x34, 0xc1, 0xb9, 0xc6, 0x45, 0x1e, 0x7b, 0x16, 0x71, 0xc1, 0x9e,
	0xb1, 0x35, 0xf7, 0xaa, 0xea, 0xf5, 0x4c, 0x05, 0xf3, 0x6a, 0xaa, 0x7c, 0x23, 0x04, 0x17, 0x9e,
	0x3d, 0xb9, 0x87, 0x66, 0xc8, 0xe3, 0x07, 0x1d, 0x9f, 0x4c, 0xa1, 0x6e, 0x6e, 0x43, 0xfe, 0x97,
	0xa7, 0x3a, 0x39, 0x60, 0xbf, 0xff, 0x57, 0xc9, 0x04, 0x0a, 0x2a, 0x57, 0xdd, 0x97, 0xf6, 0xf8,
	0xec, 0xf2, 0xa8, 0x58, 0xd4, 0xf5, 0x3f, 0x71, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x08,
	0xd8, 0x8d, 0x27, 0x02, 0x00, 0x00,
}