// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lane.proto

package traffic_manage // import "github.com/polarismesh/specification/source/go/api/v1/traffic_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import model "github.com/polarismesh/specification/source/go/api/v1/model"
import anypb "google.golang.org/protobuf/types/known/anypb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 多个 SourceMatch 之间的判断关系
type TrafficMatchRule_TrafficMatchMode int32

const (
	// 与模式
	TrafficMatchRule_AND TrafficMatchRule_TrafficMatchMode = 0
	// 或模式
	TrafficMatchRule_OR TrafficMatchRule_TrafficMatchMode = 1
)

var TrafficMatchRule_TrafficMatchMode_name = map[int32]string{
	0: "AND",
	1: "OR",
}
var TrafficMatchRule_TrafficMatchMode_value = map[string]int32{
	"AND": 0,
	"OR":  1,
}

func (x TrafficMatchRule_TrafficMatchMode) String() string {
	return proto.EnumName(TrafficMatchRule_TrafficMatchMode_name, int32(x))
}
func (TrafficMatchRule_TrafficMatchMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{4, 0}
}

type LaneRule_LaneMatchMode int32

const (
	// 严格匹配模式
	LaneRule_STRICT LaneRule_LaneMatchMode = 0
	// 宽松匹配模式
	LaneRule_PERMISSIVE LaneRule_LaneMatchMode = 1
)

var LaneRule_LaneMatchMode_name = map[int32]string{
	0: "STRICT",
	1: "PERMISSIVE",
}
var LaneRule_LaneMatchMode_value = map[string]int32{
	"STRICT":     0,
	"PERMISSIVE": 1,
}

func (x LaneRule_LaneMatchMode) String() string {
	return proto.EnumName(LaneRule_LaneMatchMode_name, int32(x))
}
func (LaneRule_LaneMatchMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{5, 0}
}

// 流量入口
type TrafficEntry struct {
	// 标记流量入口类型
	// type == "polarismesh.cn/gateway/spring-cloud-gateway", 则 selector 为
	// ServiceGatewaySelector type == "polarismesh.cn/service, 则 selector 为
	// ServiceSelector
	Type                 string     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Selector             *anypb.Any `protobuf:"bytes,2,opt,name=selector,proto3" json:"selector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TrafficEntry) Reset()         { *m = TrafficEntry{} }
func (m *TrafficEntry) String() string { return proto.CompactTextString(m) }
func (*TrafficEntry) ProtoMessage()    {}
func (*TrafficEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{0}
}
func (m *TrafficEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficEntry.Unmarshal(m, b)
}
func (m *TrafficEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficEntry.Marshal(b, m, deterministic)
}
func (dst *TrafficEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficEntry.Merge(dst, src)
}
func (m *TrafficEntry) XXX_Size() int {
	return xxx_messageInfo_TrafficEntry.Size(m)
}
func (m *TrafficEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficEntry proto.InternalMessageInfo

func (m *TrafficEntry) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TrafficEntry) GetSelector() *anypb.Any {
	if m != nil {
		return m.Selector
	}
	return nil
}

// 微服务网关入口定义
type ServiceGatewaySelector struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Service   string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// 决定要不要部份
	Labels               map[string]*model.MatchString `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ServiceGatewaySelector) Reset()         { *m = ServiceGatewaySelector{} }
func (m *ServiceGatewaySelector) String() string { return proto.CompactTextString(m) }
func (*ServiceGatewaySelector) ProtoMessage()    {}
func (*ServiceGatewaySelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{1}
}
func (m *ServiceGatewaySelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceGatewaySelector.Unmarshal(m, b)
}
func (m *ServiceGatewaySelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceGatewaySelector.Marshal(b, m, deterministic)
}
func (dst *ServiceGatewaySelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceGatewaySelector.Merge(dst, src)
}
func (m *ServiceGatewaySelector) XXX_Size() int {
	return xxx_messageInfo_ServiceGatewaySelector.Size(m)
}
func (m *ServiceGatewaySelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceGatewaySelector.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceGatewaySelector proto.InternalMessageInfo

func (m *ServiceGatewaySelector) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceGatewaySelector) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServiceGatewaySelector) GetLabels() map[string]*model.MatchString {
	if m != nil {
		return m.Labels
	}
	return nil
}

// 普通服务入口定义
type ServiceSelector struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Service   string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// 决定要不要部份
	Labels               map[string]*model.MatchString `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ServiceSelector) Reset()         { *m = ServiceSelector{} }
func (m *ServiceSelector) String() string { return proto.CompactTextString(m) }
func (*ServiceSelector) ProtoMessage()    {}
func (*ServiceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{2}
}
func (m *ServiceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSelector.Unmarshal(m, b)
}
func (m *ServiceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSelector.Marshal(b, m, deterministic)
}
func (dst *ServiceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSelector.Merge(dst, src)
}
func (m *ServiceSelector) XXX_Size() int {
	return xxx_messageInfo_ServiceSelector.Size(m)
}
func (m *ServiceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSelector proto.InternalMessageInfo

func (m *ServiceSelector) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceSelector) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServiceSelector) GetLabels() map[string]*model.MatchString {
	if m != nil {
		return m.Labels
	}
	return nil
}

// 泳道组实体定义
type LaneGroup struct {
	// 泳道组 ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 泳道组名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 泳道组内的流量入口信息
	Entries []*TrafficEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	// 在泳道组内的服务列表信息
	Destinations []*DestinationGroup `protobuf:"bytes,4,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// 泳道组描述信息
	Revision string `protobuf:"bytes,7,opt,name=revision,proto3" json:"revision,omitempty"`
	// 泳道组描述信息
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// 泳道组的创建时间
	Ctime string `protobuf:"bytes,9,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 泳道组的更新时间
	Mtime string `protobuf:"bytes,10,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// 泳道组内的流量入口信息
	Rules []*LaneRule `protobuf:"bytes,11,rep,name=rules,proto3" json:"rules,omitempty"`
	// 泳道组标签信息
	Metadata map[string]string `protobuf:"bytes,12,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 操作标志位
	Editable             bool     `protobuf:"varint,20,opt,name=editable,proto3" json:"editable,omitempty"`
	Deleteable           bool     `protobuf:"varint,21,opt,name=deleteable,proto3" json:"deleteable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaneGroup) Reset()         { *m = LaneGroup{} }
func (m *LaneGroup) String() string { return proto.CompactTextString(m) }
func (*LaneGroup) ProtoMessage()    {}
func (*LaneGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{3}
}
func (m *LaneGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaneGroup.Unmarshal(m, b)
}
func (m *LaneGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaneGroup.Marshal(b, m, deterministic)
}
func (dst *LaneGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaneGroup.Merge(dst, src)
}
func (m *LaneGroup) XXX_Size() int {
	return xxx_messageInfo_LaneGroup.Size(m)
}
func (m *LaneGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_LaneGroup.DiscardUnknown(m)
}

var xxx_messageInfo_LaneGroup proto.InternalMessageInfo

func (m *LaneGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LaneGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LaneGroup) GetEntries() []*TrafficEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *LaneGroup) GetDestinations() []*DestinationGroup {
	if m != nil {
		return m.Destinations
	}
	return nil
}

func (m *LaneGroup) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *LaneGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LaneGroup) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *LaneGroup) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *LaneGroup) GetRules() []*LaneRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *LaneGroup) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LaneGroup) GetEditable() bool {
	if m != nil {
		return m.Editable
	}
	return false
}

func (m *LaneGroup) GetDeleteable() bool {
	if m != nil {
		return m.Deleteable
	}
	return false
}

// TrafficMatchRule 流量匹配规则
type TrafficMatchRule struct {
	// 流量匹配规则，判断哪些流量需要进入泳道
	Arguments            []*SourceMatch                    `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	MatchMode            TrafficMatchRule_TrafficMatchMode `protobuf:"varint,14,opt,name=matchMode,proto3,enum=v1.TrafficMatchRule_TrafficMatchMode" json:"matchMode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TrafficMatchRule) Reset()         { *m = TrafficMatchRule{} }
func (m *TrafficMatchRule) String() string { return proto.CompactTextString(m) }
func (*TrafficMatchRule) ProtoMessage()    {}
func (*TrafficMatchRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{4}
}
func (m *TrafficMatchRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficMatchRule.Unmarshal(m, b)
}
func (m *TrafficMatchRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficMatchRule.Marshal(b, m, deterministic)
}
func (dst *TrafficMatchRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficMatchRule.Merge(dst, src)
}
func (m *TrafficMatchRule) XXX_Size() int {
	return xxx_messageInfo_TrafficMatchRule.Size(m)
}
func (m *TrafficMatchRule) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficMatchRule.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficMatchRule proto.InternalMessageInfo

func (m *TrafficMatchRule) GetArguments() []*SourceMatch {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *TrafficMatchRule) GetMatchMode() TrafficMatchRule_TrafficMatchMode {
	if m != nil {
		return m.MatchMode
	}
	return TrafficMatchRule_AND
}

// 泳道规则
type LaneRule struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 所属泳道组的名称
	GroupName string `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// 流量匹配规则
	TrafficMatchRule *TrafficMatchRule `protobuf:"bytes,4,opt,name=traffic_match_rule,json=trafficMatchRule,proto3" json:"traffic_match_rule,omitempty"`
	// 保存这个泳道的默认实例标签
	DefaultLabelValue string `protobuf:"bytes,5,opt,name=default_label_value,json=defaultLabelValue,proto3" json:"default_label_value,omitempty"`
	// 泳道规则是否启用
	Enable    bool                   `protobuf:"varint,6,opt,name=enable,proto3" json:"enable,omitempty"`
	MatchMode LaneRule_LaneMatchMode `protobuf:"varint,7,opt,name=match_mode,json=matchMode,proto3,enum=v1.LaneRule_LaneMatchMode" json:"match_mode,omitempty"`
	// revision routing version
	Revision string `protobuf:"bytes,8,opt,name=revision,proto3" json:"revision,omitempty"`
	// ctime create time of the rules
	Ctime string `protobuf:"bytes,9,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// mtime modify time of the rules
	Mtime string `protobuf:"bytes,10,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// etime enable time of the rules
	Etime string `protobuf:"bytes,11,opt,name=etime,proto3" json:"etime,omitempty"`
	// priority rules priority
	Priority uint32 `protobuf:"varint,12,opt,name=priority,proto3" json:"priority,omitempty"`
	// description simple description rules
	Description string `protobuf:"bytes,13,opt,name=description,proto3" json:"description,omitempty"`
	// 属于该泳道的实例标签KEY，不填默认为lane
	LabelKey             string   `protobuf:"bytes,14,opt,name=label_key,json=labelKey,proto3" json:"label_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaneRule) Reset()         { *m = LaneRule{} }
func (m *LaneRule) String() string { return proto.CompactTextString(m) }
func (*LaneRule) ProtoMessage()    {}
func (*LaneRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_lane_031fa4e75520adba, []int{5}
}
func (m *LaneRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaneRule.Unmarshal(m, b)
}
func (m *LaneRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaneRule.Marshal(b, m, deterministic)
}
func (dst *LaneRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaneRule.Merge(dst, src)
}
func (m *LaneRule) XXX_Size() int {
	return xxx_messageInfo_LaneRule.Size(m)
}
func (m *LaneRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LaneRule.DiscardUnknown(m)
}

var xxx_messageInfo_LaneRule proto.InternalMessageInfo

func (m *LaneRule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LaneRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LaneRule) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *LaneRule) GetTrafficMatchRule() *TrafficMatchRule {
	if m != nil {
		return m.TrafficMatchRule
	}
	return nil
}

func (m *LaneRule) GetDefaultLabelValue() string {
	if m != nil {
		return m.DefaultLabelValue
	}
	return ""
}

func (m *LaneRule) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *LaneRule) GetMatchMode() LaneRule_LaneMatchMode {
	if m != nil {
		return m.MatchMode
	}
	return LaneRule_STRICT
}

func (m *LaneRule) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *LaneRule) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *LaneRule) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *LaneRule) GetEtime() string {
	if m != nil {
		return m.Etime
	}
	return ""
}

func (m *LaneRule) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LaneRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LaneRule) GetLabelKey() string {
	if m != nil {
		return m.LabelKey
	}
	return ""
}

func init() {
	proto.RegisterType((*TrafficEntry)(nil), "v1.TrafficEntry")
	proto.RegisterType((*ServiceGatewaySelector)(nil), "v1.ServiceGatewaySelector")
	proto.RegisterMapType((map[string]*model.MatchString)(nil), "v1.ServiceGatewaySelector.LabelsEntry")
	proto.RegisterType((*ServiceSelector)(nil), "v1.ServiceSelector")
	proto.RegisterMapType((map[string]*model.MatchString)(nil), "v1.ServiceSelector.LabelsEntry")
	proto.RegisterType((*LaneGroup)(nil), "v1.LaneGroup")
	proto.RegisterMapType((map[string]string)(nil), "v1.LaneGroup.MetadataEntry")
	proto.RegisterType((*TrafficMatchRule)(nil), "v1.TrafficMatchRule")
	proto.RegisterType((*LaneRule)(nil), "v1.LaneRule")
	proto.RegisterEnum("v1.TrafficMatchRule_TrafficMatchMode", TrafficMatchRule_TrafficMatchMode_name, TrafficMatchRule_TrafficMatchMode_value)
	proto.RegisterEnum("v1.LaneRule_LaneMatchMode", LaneRule_LaneMatchMode_name, LaneRule_LaneMatchMode_value)
}

func init() { proto.RegisterFile("lane.proto", fileDescriptor_lane_031fa4e75520adba) }

var fileDescriptor_lane_031fa4e75520adba = []byte{
	// 853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0xce, 0xd8, 0x6b, 0xaf, 0xa7, 0xfc, 0x13, 0xd3, 0x98, 0x68, 0x70, 0xf8, 0xb1, 0x8c, 0x82,
	0x2c, 0x10, 0x33, 0xec, 0x72, 0xd8, 0x00, 0x12, 0x52, 0x36, 0xbb, 0x8a, 0x16, 0xb2, 0x21, 0x1a,
	0xaf, 0x72, 0xe0, 0x62, 0xb5, 0x67, 0xca, 0xb3, 0x2d, 0xe6, 0x4f, 0x3d, 0x3d, 0x46, 0x7e, 0x05,
	0x5e, 0x85, 0x77, 0xe1, 0xc6, 0x8d, 0x27, 0xe1, 0x84, 0xba, 0x7a, 0x6c, 0x8f, 0x97, 0x3d, 0x80,
	0x90, 0x72, 0xeb, 0xfa, 0xbe, 0xea, 0xfa, 0xaf, 0x6e, 0x80, 0x98, 0xa7, 0xe8, 0xe6, 0x32, 0x53,
	0x19, 0x6b, 0xac, 0x4f, 0xc6, 0xef, 0x47, 0x59, 0x16, 0xc5, 0xe8, 0x11, 0xb2, 0x2c, 0x57, 0x1e,
	0x4f, 0x37, 0x86, 0x1e, 0x77, 0x93, 0x2c, 0xc4, 0xb8, 0x12, 0xfa, 0x32, 0x2b, 0x95, 0x48, 0x23,
	0x23, 0x4e, 0x6f, 0xa0, 0x77, 0x23, 0xf9, 0x6a, 0x25, 0x82, 0xcb, 0x54, 0xc9, 0x0d, 0x63, 0x70,
	0xa4, 0x36, 0x39, 0x3a, 0xd6, 0xc4, 0x9a, 0xd9, 0x3e, 0x9d, 0xd9, 0x97, 0xd0, 0x29, 0x30, 0xc6,
	0x40, 0x65, 0xd2, 0x69, 0x4c, 0xac, 0x59, 0xf7, 0x74, 0xe4, 0x1a, 0x6f, 0xee, 0xd6, 0x9b, 0xfb,
	0x2c, 0xdd, 0xf8, 0x3b, 0xad, 0xe9, 0x9f, 0x16, 0x3c, 0x9a, 0xa3, 0x5c, 0x8b, 0x00, 0x5f, 0x70,
	0x85, 0xbf, 0xf0, 0xcd, 0xbc, 0xa2, 0xd8, 0x07, 0x60, 0xa7, 0x3c, 0xc1, 0x22, 0xe7, 0xc1, 0xd6,
	0xcb, 0x1e, 0x60, 0x0e, 0x1c, 0x17, 0xe6, 0x1e, 0x79, 0xb2, 0xfd, 0xad, 0xc8, 0xbe, 0x83, 0x76,
	0xcc, 0x97, 0x18, 0x17, 0x4e, 0x73, 0xd2, 0x9c, 0x75, 0x4f, 0x3f, 0x75, 0xd7, 0x27, 0xee, 0xfd,
	0x3e, 0xdc, 0x97, 0xa4, 0x48, 0x09, 0xf9, 0xd5, 0xad, 0xf1, 0xf7, 0xd0, 0xad, 0xc1, 0x6c, 0x08,
	0xcd, 0x9f, 0x71, 0x53, 0x05, 0xa0, 0x8f, 0xec, 0x09, 0xb4, 0xd6, 0x3c, 0x2e, 0xb1, 0x4a, 0xf1,
	0xa1, 0xb6, 0x7f, 0xcd, 0x55, 0x70, 0x3b, 0x57, 0x52, 0xa4, 0x91, 0x6f, 0xd8, 0x6f, 0x1a, 0x4f,
	0xad, 0xe9, 0xef, 0x16, 0x3c, 0xac, 0x5c, 0xff, 0xef, 0xbc, 0xce, 0xee, 0xe4, 0xf5, 0x71, 0x2d,
	0xaf, 0xb7, 0x96, 0xd0, 0x1f, 0x4d, 0xb0, 0x5f, 0xf2, 0x14, 0x5f, 0xc8, 0xac, 0xcc, 0xd9, 0x00,
	0x1a, 0x22, 0xac, 0x2c, 0x35, 0x44, 0xa8, 0x67, 0x42, 0x67, 0x52, 0x45, 0x4e, 0x67, 0xf6, 0x19,
	0x1c, 0x63, 0xaa, 0xa4, 0xc0, 0x6d, 0xdc, 0x43, 0x6d, 0xbe, 0x3e, 0x4a, 0xfe, 0x56, 0x81, 0x3d,
	0x85, 0x5e, 0x88, 0x85, 0x12, 0x29, 0x57, 0x22, 0x4b, 0x0b, 0xe7, 0x88, 0x2e, 0x8c, 0xf4, 0x85,
	0x8b, 0x3d, 0x4e, 0xbe, 0xfd, 0x03, 0x4d, 0x36, 0x86, 0x8e, 0xc4, 0xb5, 0x28, 0x44, 0x96, 0x3a,
	0xc7, 0xe4, 0x7d, 0x27, 0xb3, 0x09, 0x74, 0x43, 0x2c, 0x02, 0x29, 0x72, 0xad, 0xeb, 0x74, 0x88,
	0xae, 0x43, 0x6c, 0x04, 0xad, 0x40, 0x89, 0x04, 0x1d, 0x9b, 0x38, 0x23, 0x68, 0x34, 0x21, 0x14,
	0x0c, 0x4a, 0x02, 0x9b, 0x42, 0x4b, 0x96, 0x31, 0x16, 0x4e, 0x97, 0x82, 0xeb, 0xe9, 0xe0, 0x74,
	0x45, 0xfc, 0x32, 0x46, 0xdf, 0x50, 0xec, 0x0c, 0x3a, 0x09, 0x2a, 0x1e, 0x72, 0xc5, 0x9d, 0x1e,
	0xa9, 0x3d, 0xde, 0xaa, 0x51, 0xf0, 0xee, 0x75, 0xc5, 0x9a, 0xfc, 0x77, 0xca, 0x3a, 0x0d, 0x0c,
	0x85, 0xe2, 0xcb, 0x18, 0x9d, 0xd1, 0xc4, 0x9a, 0x75, 0xfc, 0x9d, 0xcc, 0x3e, 0x02, 0x08, 0x31,
	0x46, 0x85, 0xc4, 0xbe, 0x47, 0x6c, 0x0d, 0x19, 0x7f, 0x0b, 0xfd, 0x03, 0xb3, 0xf7, 0x34, 0x7a,
	0x54, 0x6f, 0xb4, 0x5d, 0xef, 0xeb, 0x6f, 0x16, 0x0c, 0xab, 0x9e, 0x50, 0xe7, 0x75, 0x36, 0xec,
	0x0b, 0xb0, 0xb9, 0x8c, 0xca, 0x04, 0x53, 0xb5, 0xed, 0x05, 0xcd, 0xc6, 0x3c, 0x2b, 0x65, 0x80,
	0x46, 0x6f, 0xaf, 0xc1, 0x9e, 0x83, 0x9d, 0x68, 0xec, 0x3a, 0x0b, 0xd1, 0x19, 0x4c, 0xac, 0xd9,
	0xe0, 0xf4, 0x49, 0xad, 0xd7, 0x3b, 0xbb, 0x07, 0x80, 0x56, 0xf6, 0xf7, 0xf7, 0xa6, 0x9f, 0x1c,
	0xc6, 0xa1, 0x31, 0x76, 0x0c, 0xcd, 0x67, 0xaf, 0x2e, 0x86, 0x0f, 0x58, 0x1b, 0x1a, 0x3f, 0xfa,
	0x43, 0x6b, 0xfa, 0x57, 0x13, 0x3a, 0xdb, 0x9a, 0xff, 0xab, 0x21, 0xfc, 0x10, 0x20, 0xd2, 0x85,
	0x5f, 0x10, 0xd3, 0x34, 0x4b, 0x47, 0xc8, 0x2b, 0x4d, 0x9f, 0x03, 0x53, 0xc6, 0xe9, 0x82, 0x22,
	0x59, 0xe8, 0x36, 0x3a, 0x47, 0xd5, 0x0b, 0x76, 0x4f, 0x0a, 0xfe, 0x50, 0xdd, 0x2d, 0x96, 0x0b,
	0xef, 0x86, 0xb8, 0xe2, 0x65, 0xac, 0x16, 0xb4, 0x77, 0x0b, 0x53, 0xe9, 0x16, 0xf9, 0x7a, 0xa7,
	0xa2, 0x68, 0x0f, 0xdf, 0x68, 0x82, 0x3d, 0x82, 0x36, 0xa6, 0xd4, 0xca, 0x36, 0xb5, 0xb2, 0x92,
	0xd8, 0xd7, 0x00, 0x26, 0x06, 0xfd, 0x16, 0xd3, 0x2c, 0x0f, 0x4e, 0xc7, 0xf5, 0x21, 0xa3, 0xc3,
	0x7d, 0xb5, 0x3b, 0x58, 0x82, 0xce, 0x9d, 0x25, 0xf8, 0x2f, 0x23, 0x3e, 0x82, 0x16, 0x12, 0xda,
	0x35, 0x28, 0x09, 0xda, 0x7a, 0x2e, 0x45, 0x26, 0x85, 0xda, 0x38, 0xbd, 0x89, 0x35, 0xeb, 0xfb,
	0x3b, 0xf9, 0xee, 0x8a, 0xf5, 0xff, 0xb9, 0x62, 0x8f, 0xc1, 0x36, 0x65, 0xd1, 0x23, 0x39, 0x30,
	0xc1, 0x11, 0xf0, 0x03, 0x6e, 0xa6, 0x9f, 0x43, 0xff, 0x20, 0x29, 0x06, 0xd0, 0x9e, 0xdf, 0xf8,
	0x57, 0xcf, 0x6f, 0x86, 0x0f, 0xd8, 0x00, 0xe0, 0xf5, 0xa5, 0x7f, 0x7d, 0x35, 0x9f, 0x5f, 0xbd,
	0xb9, 0x1c, 0x5a, 0xe7, 0xbf, 0x5a, 0x70, 0x16, 0x64, 0x89, 0xab, 0x30, 0x0d, 0x30, 0x55, 0x6e,
	0x9e, 0xc5, 0x5c, 0x8a, 0xc2, 0x2d, 0x72, 0x0c, 0xc4, 0x4a, 0x04, 0xf4, 0x22, 0xb8, 0x3c, 0x17,
	0xba, 0x68, 0x55, 0x9b, 0xdc, 0x84, 0xa7, 0x3c, 0xc2, 0x73, 0x7a, 0xbb, 0x5e, 0xeb, 0xaf, 0xe8,
	0xa7, 0x8b, 0x48, 0xa8, 0xdb, 0x72, 0xe9, 0x06, 0x59, 0xe2, 0x55, 0x26, 0x12, 0x2c, 0x6e, 0xbd,
	0x03, 0x33, 0x5e, 0x41, 0xa3, 0xee, 0x45, 0x99, 0xc7, 0x73, 0xe1, 0xad, 0x4f, 0xbc, 0xfd, 0x9c,
	0x68, 0x83, 0xcb, 0x36, 0xfd, 0x6b, 0x5f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x35, 0xe8, 0x18,
	0x25, 0x65, 0x07, 0x00, 0x00,
}
