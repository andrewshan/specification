// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

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

type DiscoverRequest_DiscoverRequestType int32

const (
	DiscoverRequest_UNKNOWN         DiscoverRequest_DiscoverRequestType = 0
	DiscoverRequest_INSTANCE        DiscoverRequest_DiscoverRequestType = 1
	DiscoverRequest_CLUSTER         DiscoverRequest_DiscoverRequestType = 2
	DiscoverRequest_ROUTING         DiscoverRequest_DiscoverRequestType = 3
	DiscoverRequest_RATE_LIMIT      DiscoverRequest_DiscoverRequestType = 4
	DiscoverRequest_CIRCUIT_BREAKER DiscoverRequest_DiscoverRequestType = 5
	DiscoverRequest_SERVICES        DiscoverRequest_DiscoverRequestType = 6
	DiscoverRequest_NAMESPACES      DiscoverRequest_DiscoverRequestType = 12
	DiscoverRequest_FAULT_DETECTOR  DiscoverRequest_DiscoverRequestType = 13
	DiscoverRequest_LANE            DiscoverRequest_DiscoverRequestType = 100
	// 自定义路由规则
	DiscoverRequest_CUSTOM_ROUTE_RULE DiscoverRequest_DiscoverRequestType = 101
	// 就近路由规则
	DiscoverRequest_NEARBY_ROUTE_RULE DiscoverRequest_DiscoverRequestType = 102
	// 无损上下线规则
	DiscoverRequest_LOSSLESS DiscoverRequest_DiscoverRequestType = 103
	// 服务黑白名单规则
	DiscoverRequest_BLOCK_ALLOW_RULE DiscoverRequest_DiscoverRequestType = 104
)

var DiscoverRequest_DiscoverRequestType_name = map[int32]string{
	0:   "UNKNOWN",
	1:   "INSTANCE",
	2:   "CLUSTER",
	3:   "ROUTING",
	4:   "RATE_LIMIT",
	5:   "CIRCUIT_BREAKER",
	6:   "SERVICES",
	12:  "NAMESPACES",
	13:  "FAULT_DETECTOR",
	100: "LANE",
	101: "CUSTOM_ROUTE_RULE",
	102: "NEARBY_ROUTE_RULE",
	103: "LOSSLESS",
	104: "BLOCK_ALLOW_RULE",
}
var DiscoverRequest_DiscoverRequestType_value = map[string]int32{
	"UNKNOWN":           0,
	"INSTANCE":          1,
	"CLUSTER":           2,
	"ROUTING":           3,
	"RATE_LIMIT":        4,
	"CIRCUIT_BREAKER":   5,
	"SERVICES":          6,
	"NAMESPACES":        12,
	"FAULT_DETECTOR":    13,
	"LANE":              100,
	"CUSTOM_ROUTE_RULE": 101,
	"NEARBY_ROUTE_RULE": 102,
	"LOSSLESS":          103,
	"BLOCK_ALLOW_RULE":  104,
}

func (x DiscoverRequest_DiscoverRequestType) String() string {
	return proto.EnumName(DiscoverRequest_DiscoverRequestType_name, int32(x))
}
func (DiscoverRequest_DiscoverRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_request_290a05d3b6f37256, []int{1, 0}
}

type DiscoverFilter struct {
	OnlyHealthyInstance  bool     `protobuf:"varint,1,opt,name=OnlyHealthyInstance,proto3" json:"OnlyHealthyInstance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoverFilter) Reset()         { *m = DiscoverFilter{} }
func (m *DiscoverFilter) String() string { return proto.CompactTextString(m) }
func (*DiscoverFilter) ProtoMessage()    {}
func (*DiscoverFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_290a05d3b6f37256, []int{0}
}
func (m *DiscoverFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverFilter.Unmarshal(m, b)
}
func (m *DiscoverFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverFilter.Marshal(b, m, deterministic)
}
func (dst *DiscoverFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverFilter.Merge(dst, src)
}
func (m *DiscoverFilter) XXX_Size() int {
	return xxx_messageInfo_DiscoverFilter.Size(m)
}
func (m *DiscoverFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverFilter.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverFilter proto.InternalMessageInfo

func (m *DiscoverFilter) GetOnlyHealthyInstance() bool {
	if m != nil {
		return m.OnlyHealthyInstance
	}
	return false
}

type DiscoverRequest struct {
	Type                 DiscoverRequest_DiscoverRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.DiscoverRequest_DiscoverRequestType" json:"type,omitempty"`
	Service              *Service                            `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Filter               *DiscoverFilter                     `protobuf:"bytes,30,opt,name=Filter,proto3" json:"Filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *DiscoverRequest) Reset()         { *m = DiscoverRequest{} }
func (m *DiscoverRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoverRequest) ProtoMessage()    {}
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_290a05d3b6f37256, []int{1}
}
func (m *DiscoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRequest.Unmarshal(m, b)
}
func (m *DiscoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRequest.Marshal(b, m, deterministic)
}
func (dst *DiscoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRequest.Merge(dst, src)
}
func (m *DiscoverRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoverRequest.Size(m)
}
func (m *DiscoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRequest proto.InternalMessageInfo

func (m *DiscoverRequest) GetType() DiscoverRequest_DiscoverRequestType {
	if m != nil {
		return m.Type
	}
	return DiscoverRequest_UNKNOWN
}

func (m *DiscoverRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *DiscoverRequest) GetFilter() *DiscoverFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoverFilter)(nil), "v1.DiscoverFilter")
	proto.RegisterType((*DiscoverRequest)(nil), "v1.DiscoverRequest")
	proto.RegisterEnum("v1.DiscoverRequest_DiscoverRequestType", DiscoverRequest_DiscoverRequestType_name, DiscoverRequest_DiscoverRequestType_value)
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_request_290a05d3b6f37256) }

var fileDescriptor_request_290a05d3b6f37256 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xa7, 0x6d, 0xd6, 0x46, 0xee, 0x9f, 0x19, 0x17, 0xa4, 0x8a, 0x03, 0x9a, 0x2a, 0x21, 0x26,
	0x0e, 0x09, 0x1d, 0x07, 0x0e, 0x9c, 0x92, 0xd4, 0x83, 0xd0, 0x34, 0x99, 0x6c, 0x87, 0x09, 0x2e,
	0x51, 0x66, 0xbc, 0xd6, 0x52, 0x9b, 0x84, 0xc4, 0xad, 0xd4, 0x8f, 0xc0, 0x81, 0x4f, 0xc3, 0x17,
	0x44, 0x6e, 0x52, 0x69, 0x4c, 0x3b, 0xbe, 0xdf, 0xdf, 0xa7, 0xa7, 0x07, 0x86, 0xa5, 0xf8, 0xb5,
	0x13, 0x95, 0xb2, 0x8a, 0x32, 0x57, 0x39, 0x6a, 0xef, 0x67, 0xaf, 0x86, 0x95, 0x28, 0xf7, 0x92,
	0x8b, 0x1a, 0x9a, 0xba, 0x60, 0x34, 0x97, 0x15, 0xcf, 0xf7, 0xa2, 0xbc, 0x96, 0x1b, 0x25, 0x4a,
	0xf4, 0x1e, 0x8c, 0xa3, 0x6c, 0x73, 0xf8, 0x22, 0xd2, 0x8d, 0x5a, 0x1f, 0xfc, 0xac, 0x52, 0x69,
	0xc6, 0xc5, 0xa4, 0x75, 0xd1, 0xba, 0x34, 0xc9, 0x53, 0xd4, 0xf4, 0x6f, 0x07, 0x9c, 0x9f, 0x42,
	0x48, 0x5d, 0x88, 0x3e, 0x01, 0x43, 0x1d, 0x8a, 0xda, 0x36, 0xba, 0x7a, 0x6b, 0xed, 0x67, 0xd6,
	0x23, 0xc9, 0xe3, 0x99, 0x1d, 0x0a, 0x41, 0x8e, 0x26, 0xf4, 0x06, 0xf4, 0x9a, 0x2d, 0x27, 0xed,
	0x8b, 0xd6, 0x65, 0xff, 0xaa, 0xaf, 0xfd, 0xb4, 0x86, 0xc8, 0x89, 0x43, 0xef, 0x40, 0xb7, 0xde,
	0x79, 0xf2, 0xfa, 0xa8, 0x42, 0x0f, 0x5b, 0x6a, 0x86, 0x34, 0x8a, 0xe9, 0xef, 0x36, 0x18, 0x3f,
	0x51, 0x88, 0xfa, 0xa0, 0x17, 0x87, 0x8b, 0x30, 0xba, 0x0d, 0xe1, 0x33, 0x34, 0x00, 0xa6, 0x1f,
	0x52, 0xe6, 0x84, 0x1e, 0x86, 0x2d, 0x4d, 0x79, 0x41, 0x4c, 0x19, 0x26, 0xb0, 0xad, 0x07, 0x12,
	0xc5, 0xcc, 0x0f, 0x3f, 0xc3, 0x0e, 0x1a, 0x01, 0x40, 0x1c, 0x86, 0x93, 0xc0, 0x5f, 0xfa, 0x0c,
	0x1a, 0x68, 0x0c, 0xce, 0x3d, 0x9f, 0x78, 0xb1, 0xcf, 0x12, 0x97, 0x60, 0x67, 0x81, 0x09, 0x3c,
	0xd3, 0x61, 0x14, 0x93, 0x6f, 0xbe, 0x87, 0x29, 0xec, 0x6a, 0x4b, 0xe8, 0x2c, 0x31, 0xbd, 0x71,
	0xf4, 0x3c, 0x40, 0x08, 0x8c, 0xae, 0x9d, 0x38, 0x60, 0xc9, 0x1c, 0x33, 0xec, 0xb1, 0x88, 0xc0,
	0x21, 0x32, 0x81, 0x11, 0x38, 0x21, 0x86, 0x3f, 0xd1, 0x4b, 0xf0, 0xdc, 0x8b, 0x29, 0x8b, 0x96,
	0x89, 0x2e, 0xc5, 0x09, 0x89, 0x03, 0x0c, 0x85, 0x86, 0x43, 0xec, 0x10, 0xf7, 0xfb, 0x43, 0xf8,
	0x5e, 0x37, 0x05, 0x11, 0xa5, 0x01, 0xa6, 0x14, 0xae, 0xd0, 0x0b, 0x00, 0xdd, 0x20, 0xf2, 0x16,
	0x89, 0x13, 0x04, 0xd1, 0x6d, 0xad, 0x59, 0x4f, 0x0d, 0xb3, 0x07, 0xfb, 0x53, 0xc3, 0x1c, 0x41,
	0xfe, 0xd5, 0x30, 0x3b, 0xf0, 0xcc, 0xfd, 0xd3, 0x02, 0x1f, 0x79, 0xbe, 0xb5, 0x94, 0xc8, 0xb8,
	0xc8, 0x94, 0x55, 0xe4, 0x9b, 0xb4, 0x94, 0x95, 0x55, 0x15, 0x82, 0xcb, 0x7b, 0xc9, 0x53, 0x25,
	0xf3, 0xcc, 0x4a, 0x0b, 0xa9, 0xaf, 0x7a, 0x7a, 0x9a, 0x6d, 0x9a, 0xa5, 0x2b, 0xe1, 0x0e, 0x9a,
	0x13, 0xde, 0xe8, 0x1f, 0xfa, 0x31, 0x5f, 0x49, 0xb5, 0xde, 0xdd, 0x59, 0x3c, 0xdf, 0xda, 0x4d,
	0xca, 0x56, 0x54, 0x6b, 0xfb, 0xbf, 0x24, 0xbb, 0xca, 0x77, 0x25, 0x17, 0xf6, 0x2a, 0xb7, 0xd3,
	0x42, 0xda, 0xfb, 0x99, 0xdd, 0x64, 0x26, 0x75, 0xe6, 0x5d, 0xf7, 0xf8, 0x90, 0x1f, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xae, 0x0d, 0x76, 0x20, 0xb4, 0x02, 0x00, 0x00,
}
