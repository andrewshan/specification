// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import fault_tolerance "github.com/polarismesh/specification/source/go/api/v1/fault_tolerance"
import model "github.com/polarismesh/specification/source/go/api/v1/model"
import security "github.com/polarismesh/specification/source/go/api/v1/security"
import traffic_manage "github.com/polarismesh/specification/source/go/api/v1/traffic_manage"
import anypb "google.golang.org/protobuf/types/known/anypb"
import wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DiscoverResponse_DiscoverResponseType int32

const (
	DiscoverResponse_UNKNOWN          DiscoverResponse_DiscoverResponseType = 0
	DiscoverResponse_INSTANCE         DiscoverResponse_DiscoverResponseType = 1
	DiscoverResponse_CLUSTER          DiscoverResponse_DiscoverResponseType = 2
	DiscoverResponse_ROUTING          DiscoverResponse_DiscoverResponseType = 3
	DiscoverResponse_RATE_LIMIT       DiscoverResponse_DiscoverResponseType = 4
	DiscoverResponse_CIRCUIT_BREAKER  DiscoverResponse_DiscoverResponseType = 5
	DiscoverResponse_SERVICES         DiscoverResponse_DiscoverResponseType = 6
	DiscoverResponse_NAMESPACES       DiscoverResponse_DiscoverResponseType = 12
	DiscoverResponse_FAULT_DETECTOR   DiscoverResponse_DiscoverResponseType = 13
	DiscoverResponse_SERVICE_CONTRACT DiscoverResponse_DiscoverResponseType = 14
)

var DiscoverResponse_DiscoverResponseType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "INSTANCE",
	2:  "CLUSTER",
	3:  "ROUTING",
	4:  "RATE_LIMIT",
	5:  "CIRCUIT_BREAKER",
	6:  "SERVICES",
	12: "NAMESPACES",
	13: "FAULT_DETECTOR",
	14: "SERVICE_CONTRACT",
}
var DiscoverResponse_DiscoverResponseType_value = map[string]int32{
	"UNKNOWN":          0,
	"INSTANCE":         1,
	"CLUSTER":          2,
	"ROUTING":          3,
	"RATE_LIMIT":       4,
	"CIRCUIT_BREAKER":  5,
	"SERVICES":         6,
	"NAMESPACES":       12,
	"FAULT_DETECTOR":   13,
	"SERVICE_CONTRACT": 14,
}

func (x DiscoverResponse_DiscoverResponseType) String() string {
	return proto.EnumName(DiscoverResponse_DiscoverResponseType_name, int32(x))
}
func (DiscoverResponse_DiscoverResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{3, 0}
}

type Response struct {
	Code                 *wrapperspb.UInt32Value         `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue         `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Client               *Client                         `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Namespace            *model.Namespace                `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Service              *Service                        `protobuf:"bytes,5,opt,name=service,proto3" json:"service,omitempty"`
	Instance             *Instance                       `protobuf:"bytes,6,opt,name=instance,proto3" json:"instance,omitempty"`
	Routing              *traffic_manage.Routing         `protobuf:"bytes,7,opt,name=routing,proto3" json:"routing,omitempty"`
	Alias                *ServiceAlias                   `protobuf:"bytes,8,opt,name=alias,proto3" json:"alias,omitempty"`
	RateLimit            *traffic_manage.Rule            `protobuf:"bytes,9,opt,name=rateLimit,proto3" json:"rateLimit,omitempty"`
	CircuitBreaker       *fault_tolerance.CircuitBreaker `protobuf:"bytes,10,opt,name=circuitBreaker,proto3" json:"circuitBreaker,omitempty"`
	ConfigRelease        *ConfigRelease                  `protobuf:"bytes,11,opt,name=configRelease,proto3" json:"configRelease,omitempty"`
	User                 *security.User                  `protobuf:"bytes,19,opt,name=user,proto3" json:"user,omitempty"`
	UserGroup            *security.UserGroup             `protobuf:"bytes,20,opt,name=userGroup,proto3" json:"userGroup,omitempty"`
	AuthStrategy         *security.AuthStrategy          `protobuf:"bytes,21,opt,name=authStrategy,proto3" json:"authStrategy,omitempty"`
	Relation             *security.UserGroupRelation     `protobuf:"bytes,22,opt,name=relation,proto3" json:"relation,omitempty"`
	LoginResponse        *security.LoginResponse         `protobuf:"bytes,23,opt,name=loginResponse,proto3" json:"loginResponse,omitempty"`
	ModifyAuthStrategy   *security.ModifyAuthStrategy    `protobuf:"bytes,24,opt,name=modifyAuthStrategy,proto3" json:"modifyAuthStrategy,omitempty"`
	ModifyUserGroup      *security.ModifyUserGroup       `protobuf:"bytes,25,opt,name=modifyUserGroup,proto3" json:"modifyUserGroup,omitempty"`
	Resources            *security.StrategyResources     `protobuf:"bytes,26,opt,name=resources,proto3" json:"resources,omitempty"`
	OptionSwitch         *OptionSwitch                   `protobuf:"bytes,27,opt,name=optionSwitch,proto3" json:"optionSwitch,omitempty"`
	InstanceLabels       *InstanceLabels                 `protobuf:"bytes,28,opt,name=instanceLabels,proto3" json:"instanceLabels,omitempty"`
	Data                 *anypb.Any                      `protobuf:"bytes,29,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{0}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Response) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Response) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Response) GetNamespace() *model.Namespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Response) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Response) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *Response) GetRouting() *traffic_manage.Routing {
	if m != nil {
		return m.Routing
	}
	return nil
}

func (m *Response) GetAlias() *ServiceAlias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Response) GetRateLimit() *traffic_manage.Rule {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

func (m *Response) GetCircuitBreaker() *fault_tolerance.CircuitBreaker {
	if m != nil {
		return m.CircuitBreaker
	}
	return nil
}

func (m *Response) GetConfigRelease() *ConfigRelease {
	if m != nil {
		return m.ConfigRelease
	}
	return nil
}

func (m *Response) GetUser() *security.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUserGroup() *security.UserGroup {
	if m != nil {
		return m.UserGroup
	}
	return nil
}

func (m *Response) GetAuthStrategy() *security.AuthStrategy {
	if m != nil {
		return m.AuthStrategy
	}
	return nil
}

func (m *Response) GetRelation() *security.UserGroupRelation {
	if m != nil {
		return m.Relation
	}
	return nil
}

func (m *Response) GetLoginResponse() *security.LoginResponse {
	if m != nil {
		return m.LoginResponse
	}
	return nil
}

func (m *Response) GetModifyAuthStrategy() *security.ModifyAuthStrategy {
	if m != nil {
		return m.ModifyAuthStrategy
	}
	return nil
}

func (m *Response) GetModifyUserGroup() *security.ModifyUserGroup {
	if m != nil {
		return m.ModifyUserGroup
	}
	return nil
}

func (m *Response) GetResources() *security.StrategyResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Response) GetOptionSwitch() *OptionSwitch {
	if m != nil {
		return m.OptionSwitch
	}
	return nil
}

func (m *Response) GetInstanceLabels() *InstanceLabels {
	if m != nil {
		return m.InstanceLabels
	}
	return nil
}

func (m *Response) GetData() *anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type BatchWriteResponse struct {
	Code                 *wrapperspb.UInt32Value `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Size                 *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
	Responses            []*Response             `protobuf:"bytes,4,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BatchWriteResponse) Reset()         { *m = BatchWriteResponse{} }
func (m *BatchWriteResponse) String() string { return proto.CompactTextString(m) }
func (*BatchWriteResponse) ProtoMessage()    {}
func (*BatchWriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{1}
}
func (m *BatchWriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchWriteResponse.Unmarshal(m, b)
}
func (m *BatchWriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchWriteResponse.Marshal(b, m, deterministic)
}
func (dst *BatchWriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchWriteResponse.Merge(dst, src)
}
func (m *BatchWriteResponse) XXX_Size() int {
	return xxx_messageInfo_BatchWriteResponse.Size(m)
}
func (m *BatchWriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchWriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchWriteResponse proto.InternalMessageInfo

func (m *BatchWriteResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *BatchWriteResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *BatchWriteResponse) GetSize() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *BatchWriteResponse) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type BatchQueryResponse struct {
	Code                 *wrapperspb.UInt32Value   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Amount               *wrapperspb.UInt32Value   `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Size                 *wrapperspb.UInt32Value   `protobuf:"bytes,4,opt,name=size,proto3" json:"size,omitempty"`
	Namespaces           []*model.Namespace        `protobuf:"bytes,5,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Services             []*Service                `protobuf:"bytes,6,rep,name=services,proto3" json:"services,omitempty"`
	Instances            []*Instance               `protobuf:"bytes,7,rep,name=instances,proto3" json:"instances,omitempty"`
	Routings             []*traffic_manage.Routing `protobuf:"bytes,8,rep,name=routings,proto3" json:"routings,omitempty"`
	Aliases              []*ServiceAlias           `protobuf:"bytes,9,rep,name=aliases,proto3" json:"aliases,omitempty"`
	RateLimits           []*traffic_manage.Rule    `protobuf:"bytes,10,rep,name=rateLimits,proto3" json:"rateLimits,omitempty"`
	ConfigWithServices   []*ConfigWithService      `protobuf:"bytes,11,rep,name=configWithServices,proto3" json:"configWithServices,omitempty"`
	Users                []*security.User          `protobuf:"bytes,18,rep,name=users,proto3" json:"users,omitempty"`
	UserGroups           []*security.UserGroup     `protobuf:"bytes,19,rep,name=userGroups,proto3" json:"userGroups,omitempty"`
	AuthStrategies       []*security.AuthStrategy  `protobuf:"bytes,20,rep,name=authStrategies,proto3" json:"authStrategies,omitempty"`
	Clients              []*Client                 `protobuf:"bytes,21,rep,name=clients,proto3" json:"clients,omitempty"`
	Data                 []*anypb.Any              `protobuf:"bytes,22,rep,name=data,proto3" json:"data,omitempty"`
	Summary              *model.Summary            `protobuf:"bytes,23,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BatchQueryResponse) Reset()         { *m = BatchQueryResponse{} }
func (m *BatchQueryResponse) String() string { return proto.CompactTextString(m) }
func (*BatchQueryResponse) ProtoMessage()    {}
func (*BatchQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{2}
}
func (m *BatchQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchQueryResponse.Unmarshal(m, b)
}
func (m *BatchQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchQueryResponse.Marshal(b, m, deterministic)
}
func (dst *BatchQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchQueryResponse.Merge(dst, src)
}
func (m *BatchQueryResponse) XXX_Size() int {
	return xxx_messageInfo_BatchQueryResponse.Size(m)
}
func (m *BatchQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchQueryResponse proto.InternalMessageInfo

func (m *BatchQueryResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *BatchQueryResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *BatchQueryResponse) GetAmount() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *BatchQueryResponse) GetSize() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Size
	}
	return nil
}

func (m *BatchQueryResponse) GetNamespaces() []*model.Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *BatchQueryResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *BatchQueryResponse) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *BatchQueryResponse) GetRoutings() []*traffic_manage.Routing {
	if m != nil {
		return m.Routings
	}
	return nil
}

func (m *BatchQueryResponse) GetAliases() []*ServiceAlias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *BatchQueryResponse) GetRateLimits() []*traffic_manage.Rule {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *BatchQueryResponse) GetConfigWithServices() []*ConfigWithService {
	if m != nil {
		return m.ConfigWithServices
	}
	return nil
}

func (m *BatchQueryResponse) GetUsers() []*security.User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BatchQueryResponse) GetUserGroups() []*security.UserGroup {
	if m != nil {
		return m.UserGroups
	}
	return nil
}

func (m *BatchQueryResponse) GetAuthStrategies() []*security.AuthStrategy {
	if m != nil {
		return m.AuthStrategies
	}
	return nil
}

func (m *BatchQueryResponse) GetClients() []*Client {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *BatchQueryResponse) GetData() []*anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BatchQueryResponse) GetSummary() *model.Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

type DiscoverResponse struct {
	Code                 *wrapperspb.UInt32Value               `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 *wrapperspb.StringValue               `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Type                 DiscoverResponse_DiscoverResponseType `protobuf:"varint,3,opt,name=type,proto3,enum=v1.DiscoverResponse_DiscoverResponseType" json:"type,omitempty"`
	Service              *Service                              `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	Instances            []*Instance                           `protobuf:"bytes,5,rep,name=instances,proto3" json:"instances,omitempty"`
	Routing              *traffic_manage.Routing               `protobuf:"bytes,6,opt,name=routing,proto3" json:"routing,omitempty"`
	RateLimit            *traffic_manage.RateLimit             `protobuf:"bytes,7,opt,name=rateLimit,proto3" json:"rateLimit,omitempty"`
	CircuitBreaker       *fault_tolerance.CircuitBreaker       `protobuf:"bytes,8,opt,name=circuitBreaker,proto3" json:"circuitBreaker,omitempty"`
	Services             []*Service                            `protobuf:"bytes,9,rep,name=services,proto3" json:"services,omitempty"`
	Namespaces           []*model.Namespace                    `protobuf:"bytes,10,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	FaultDetector        *fault_tolerance.FaultDetector        `protobuf:"bytes,11,opt,name=faultDetector,proto3" json:"faultDetector,omitempty"`
	AliasFor             *Service                              `protobuf:"bytes,21,opt,name=aliasFor,proto3" json:"aliasFor,omitempty"`
	ServiceContract      *ServiceContract                      `protobuf:"bytes,22,opt,name=serviceContract,proto3" json:"serviceContract,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DiscoverResponse) Reset()         { *m = DiscoverResponse{} }
func (m *DiscoverResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoverResponse) ProtoMessage()    {}
func (*DiscoverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{3}
}
func (m *DiscoverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverResponse.Unmarshal(m, b)
}
func (m *DiscoverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverResponse.Marshal(b, m, deterministic)
}
func (dst *DiscoverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverResponse.Merge(dst, src)
}
func (m *DiscoverResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoverResponse.Size(m)
}
func (m *DiscoverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverResponse proto.InternalMessageInfo

func (m *DiscoverResponse) GetCode() *wrapperspb.UInt32Value {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *DiscoverResponse) GetInfo() *wrapperspb.StringValue {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *DiscoverResponse) GetType() DiscoverResponse_DiscoverResponseType {
	if m != nil {
		return m.Type
	}
	return DiscoverResponse_UNKNOWN
}

func (m *DiscoverResponse) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *DiscoverResponse) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *DiscoverResponse) GetRouting() *traffic_manage.Routing {
	if m != nil {
		return m.Routing
	}
	return nil
}

func (m *DiscoverResponse) GetRateLimit() *traffic_manage.RateLimit {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

func (m *DiscoverResponse) GetCircuitBreaker() *fault_tolerance.CircuitBreaker {
	if m != nil {
		return m.CircuitBreaker
	}
	return nil
}

func (m *DiscoverResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *DiscoverResponse) GetNamespaces() []*model.Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *DiscoverResponse) GetFaultDetector() *fault_tolerance.FaultDetector {
	if m != nil {
		return m.FaultDetector
	}
	return nil
}

func (m *DiscoverResponse) GetAliasFor() *Service {
	if m != nil {
		return m.AliasFor
	}
	return nil
}

func (m *DiscoverResponse) GetServiceContract() *ServiceContract {
	if m != nil {
		return m.ServiceContract
	}
	return nil
}

type OptionSwitch struct {
	Options              map[string]string `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OptionSwitch) Reset()         { *m = OptionSwitch{} }
func (m *OptionSwitch) String() string { return proto.CompactTextString(m) }
func (*OptionSwitch) ProtoMessage()    {}
func (*OptionSwitch) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{4}
}
func (m *OptionSwitch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OptionSwitch.Unmarshal(m, b)
}
func (m *OptionSwitch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OptionSwitch.Marshal(b, m, deterministic)
}
func (dst *OptionSwitch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OptionSwitch.Merge(dst, src)
}
func (m *OptionSwitch) XXX_Size() int {
	return xxx_messageInfo_OptionSwitch.Size(m)
}
func (m *OptionSwitch) XXX_DiscardUnknown() {
	xxx_messageInfo_OptionSwitch.DiscardUnknown(m)
}

var xxx_messageInfo_OptionSwitch proto.InternalMessageInfo

func (m *OptionSwitch) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

type InstanceLabels struct {
	Labels               map[string]*model.StringList `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *InstanceLabels) Reset()         { *m = InstanceLabels{} }
func (m *InstanceLabels) String() string { return proto.CompactTextString(m) }
func (*InstanceLabels) ProtoMessage()    {}
func (*InstanceLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_fc83a74d7fb91365, []int{5}
}
func (m *InstanceLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceLabels.Unmarshal(m, b)
}
func (m *InstanceLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceLabels.Marshal(b, m, deterministic)
}
func (dst *InstanceLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceLabels.Merge(dst, src)
}
func (m *InstanceLabels) XXX_Size() int {
	return xxx_messageInfo_InstanceLabels.Size(m)
}
func (m *InstanceLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceLabels.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceLabels proto.InternalMessageInfo

func (m *InstanceLabels) GetLabels() map[string]*model.StringList {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "v1.Response")
	proto.RegisterType((*BatchWriteResponse)(nil), "v1.BatchWriteResponse")
	proto.RegisterType((*BatchQueryResponse)(nil), "v1.BatchQueryResponse")
	proto.RegisterType((*DiscoverResponse)(nil), "v1.DiscoverResponse")
	proto.RegisterType((*OptionSwitch)(nil), "v1.OptionSwitch")
	proto.RegisterMapType((map[string]string)(nil), "v1.OptionSwitch.OptionsEntry")
	proto.RegisterType((*InstanceLabels)(nil), "v1.InstanceLabels")
	proto.RegisterMapType((map[string]*model.StringList)(nil), "v1.InstanceLabels.LabelsEntry")
	proto.RegisterEnum("v1.DiscoverResponse_DiscoverResponseType", DiscoverResponse_DiscoverResponseType_name, DiscoverResponse_DiscoverResponseType_value)
}

func init() { proto.RegisterFile("response.proto", fileDescriptor_response_fc83a74d7fb91365) }

var fileDescriptor_response_fc83a74d7fb91365 = []byte{
	// 1409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0xfe, 0xc9, 0xa6, 0xf5, 0x67, 0x24, 0x2b, 0xfc, 0xad, 0xe4, 0x74, 0xa3, 0x26, 0x41, 0x20,
	0xa4, 0xa9, 0xeb, 0xa0, 0x52, 0xec, 0x04, 0x75, 0x10, 0x20, 0x07, 0x59, 0xa6, 0x52, 0x35, 0xb2,
	0x9c, 0xae, 0xa4, 0xa4, 0xe8, 0x45, 0xa0, 0xe9, 0xb5, 0xbc, 0x08, 0x45, 0x0a, 0x5c, 0xd2, 0x81,
	0x7a, 0xed, 0xbd, 0x87, 0xf6, 0x45, 0xfa, 0x16, 0x7d, 0x83, 0x3e, 0x4b, 0x8f, 0xc5, 0x2e, 0x97,
	0xff, 0x64, 0x57, 0xc9, 0x29, 0x17, 0x9b, 0x3b, 0xdf, 0x37, 0xc3, 0x9d, 0xd9, 0xd9, 0x8f, 0x23,
	0xa8, 0x7a, 0x94, 0x2f, 0x5c, 0x87, 0xd3, 0xd6, 0xc2, 0x73, 0x7d, 0x17, 0x6d, 0x5c, 0xed, 0x37,
	0xee, 0xcf, 0x5c, 0x77, 0x66, 0xd3, 0xb6, 0xb4, 0x9c, 0x05, 0x17, 0xed, 0x0f, 0x9e, 0xb9, 0x58,
	0x50, 0x8f, 0x87, 0x9c, 0xc6, 0x9d, 0x55, 0xdc, 0x74, 0x96, 0x0a, 0xba, 0xe5, 0x98, 0x73, 0xca,
	0x17, 0xa6, 0xa5, 0xe2, 0x35, 0xb6, 0x39, 0xf5, 0xae, 0x58, 0xb2, 0xf4, 0xdc, 0xc0, 0x67, 0xce,
	0x2c, 0xa2, 0x7b, 0xa6, 0x4f, 0x6d, 0x36, 0x67, 0xbe, 0x32, 0xd4, 0x2d, 0xe6, 0x59, 0x01, 0xf3,
	0xcf, 0x3c, 0x6a, 0xbe, 0xa7, 0x9e, 0xb2, 0x96, 0xe7, 0xee, 0x39, 0xb5, 0xd5, 0xa2, 0x62, 0xd9,
	0x8c, 0x3a, 0x91, 0x43, 0xcd, 0x72, 0x9d, 0x0b, 0x36, 0xf3, 0xa8, 0x4d, 0xcd, 0x28, 0x89, 0x46,
	0xfd, 0xc2, 0x0c, 0x6c, 0x7f, 0x7a, 0x4e, 0x7d, 0x6a, 0xf9, 0x6e, 0x14, 0x05, 0xcc, 0xc0, 0xbf,
	0x54, 0xcf, 0x55, 0xcb, 0x75, 0x7c, 0xcf, 0xb4, 0x54, 0x98, 0xe6, 0x5f, 0x45, 0x28, 0x12, 0x55,
	0x09, 0xf4, 0x04, 0x34, 0xcb, 0x3d, 0xa7, 0x38, 0xf7, 0x20, 0xb7, 0x5b, 0x3e, 0xb8, 0xdb, 0x0a,
	0xd3, 0x6d, 0x45, 0xe9, 0xb6, 0x26, 0x7d, 0xc7, 0x7f, 0x7a, 0xf0, 0xd6, 0xb4, 0x03, 0x4a, 0x24,
	0x53, 0x78, 0x30, 0xe7, 0xc2, 0xc5, 0x1b, 0xff, 0xe1, 0x31, 0xf2, 0x3d, 0xe6, 0xcc, 0x94, 0x87,
	0x60, 0xa2, 0x26, 0xe4, 0xc3, 0x3c, 0xf0, 0xa6, 0xf4, 0x81, 0xd6, 0xd5, 0x7e, 0xab, 0x2b, 0x2d,
	0x44, 0x21, 0xe8, 0x31, 0x94, 0xe2, 0x72, 0x62, 0x4d, 0xd2, 0xb6, 0x05, 0x6d, 0x18, 0x19, 0x49,
	0x82, 0xa3, 0xaf, 0xa0, 0xa0, 0x4a, 0x8d, 0xb7, 0x24, 0xb5, 0x2c, 0xa8, 0xa3, 0xd0, 0x44, 0x22,
	0x0c, 0xed, 0x42, 0x91, 0x39, 0xdc, 0x37, 0x1d, 0x8b, 0xe2, 0xbc, 0xe4, 0x55, 0x04, 0xaf, 0xaf,
	0x6c, 0x24, 0x46, 0x45, 0x40, 0x75, 0x58, 0xb8, 0x90, 0x04, 0x24, 0xa1, 0x89, 0x44, 0x18, 0x7a,
	0x04, 0x5b, 0xa6, 0xcd, 0x4c, 0x8e, 0x8b, 0x92, 0xa4, 0xa7, 0xde, 0xda, 0x11, 0x76, 0x12, 0xc2,
	0xe8, 0x11, 0x94, 0xc4, 0x61, 0x0f, 0xc4, 0x61, 0xe3, 0x92, 0xe4, 0x16, 0x65, 0xc0, 0xc0, 0xa6,
	0x24, 0x81, 0xd0, 0x0b, 0xa8, 0xaa, 0x1e, 0x38, 0x0a, 0x7b, 0x00, 0x83, 0x24, 0x23, 0x59, 0xa0,
	0x0c, 0x42, 0x56, 0x98, 0xe8, 0x10, 0xb6, 0xc3, 0x76, 0x20, 0x61, 0x3b, 0xe0, 0xb2, 0x74, 0xfd,
	0xbf, 0x74, 0x4d, 0x03, 0x24, 0xcb, 0x43, 0x77, 0x41, 0x0b, 0x38, 0xf5, 0x70, 0x2d, 0xd9, 0xd7,
	0x84, 0x53, 0x8f, 0x48, 0xab, 0x38, 0x07, 0xf1, 0xff, 0x95, 0xe7, 0x06, 0x0b, 0x5c, 0x4f, 0xce,
	0x61, 0x12, 0x19, 0x49, 0x82, 0xa3, 0x67, 0x50, 0x11, 0x7d, 0x36, 0xf2, 0x45, 0x4a, 0xb3, 0x25,
	0xde, 0x49, 0xca, 0xd2, 0x49, 0xd9, 0x49, 0x86, 0x85, 0xf6, 0xa1, 0xe8, 0x51, 0xdb, 0xf4, 0x99,
	0xeb, 0xe0, 0xdb, 0xd2, 0x63, 0x27, 0xfb, 0x06, 0x05, 0x92, 0x98, 0x26, 0x92, 0xb5, 0xdd, 0x19,
	0x73, 0xa2, 0xb6, 0xc5, 0x5f, 0x24, 0xc9, 0x0e, 0xd2, 0x00, 0xc9, 0xf2, 0x50, 0x0f, 0xd0, 0xdc,
	0x3d, 0x67, 0x17, 0xcb, 0xf4, 0x7e, 0x30, 0x96, 0xde, 0xb7, 0x85, 0xf7, 0xc9, 0x35, 0x94, 0xdc,
	0xe0, 0x81, 0x5e, 0xc2, 0xad, 0xd0, 0x1a, 0xef, 0x12, 0xdf, 0x91, 0x41, 0x6a, 0x49, 0x90, 0x24,
	0x81, 0x55, 0x2e, 0x7a, 0x0a, 0x25, 0x8f, 0x72, 0x37, 0xf0, 0x2c, 0xca, 0x71, 0x23, 0xc9, 0x39,
	0x7e, 0x67, 0x04, 0x92, 0x84, 0x27, 0xaa, 0xeb, 0x2e, 0x44, 0xfa, 0xa3, 0x0f, 0xcc, 0xb7, 0x2e,
	0xf1, 0x97, 0x49, 0x75, 0x4f, 0x53, 0x76, 0x92, 0x61, 0x89, 0x9e, 0x8a, 0xda, 0x7a, 0x60, 0x9e,
	0x51, 0x9b, 0xe3, 0xbb, 0x49, 0x4f, 0xf5, 0x33, 0x08, 0x59, 0x61, 0xa2, 0x5d, 0xd0, 0xce, 0x4d,
	0xdf, 0xc4, 0xf7, 0xa4, 0x47, 0xfd, 0xda, 0xd5, 0xee, 0x38, 0x4b, 0x22, 0x19, 0x3f, 0x68, 0xc5,
	0x8a, 0x5e, 0x6b, 0xfe, 0x9d, 0x03, 0x74, 0x64, 0xfa, 0xd6, 0xe5, 0x3b, 0x8f, 0xf9, 0xf4, 0xb3,
	0x6a, 0xca, 0x13, 0xd0, 0x38, 0xfb, 0x85, 0x2a, 0x45, 0xf9, 0xc8, 0x3b, 0x04, 0x13, 0xed, 0xc9,
	0x33, 0x90, 0x3b, 0xe4, 0x58, 0x7b, 0xb0, 0x19, 0xc9, 0x41, 0xdc, 0x3a, 0x09, 0xdc, 0xfc, 0x33,
	0xaf, 0x12, 0xfb, 0x31, 0xa0, 0xde, 0xf2, 0xb3, 0x26, 0xf6, 0x0c, 0xf2, 0xe6, 0xdc, 0x0d, 0x62,
	0xb1, 0x5c, 0xff, 0x16, 0xc5, 0x8d, 0xcb, 0xa1, 0x7d, 0x72, 0x39, 0xbe, 0x05, 0x88, 0x05, 0x95,
	0xe3, 0x2d, 0x59, 0x8f, 0x15, 0xc5, 0x4d, 0x11, 0xd0, 0xd7, 0x50, 0x54, 0xb2, 0xca, 0x71, 0x5e,
	0x92, 0x33, 0x9a, 0x1b, 0x83, 0xa2, 0xcc, 0x51, 0x57, 0x71, 0x5c, 0x48, 0xca, 0x1c, 0xab, 0x6e,
	0x02, 0x8b, 0xa0, 0x4a, 0x5a, 0x85, 0xa4, 0x6e, 0xae, 0xea, 0x6e, 0x0c, 0xa2, 0x3d, 0x28, 0x48,
	0x65, 0xa5, 0x1c, 0x97, 0x24, 0xef, 0xba, 0xf4, 0x46, 0x04, 0xb4, 0x0b, 0x10, 0x2b, 0x2c, 0xc7,
	0x20, 0xe9, 0x89, 0xfa, 0xa6, 0x30, 0x64, 0x00, 0x0a, 0xa5, 0xf1, 0x1d, 0xf3, 0x2f, 0x47, 0x51,
	0x76, 0x65, 0xe9, 0xb1, 0x93, 0xe8, 0x68, 0x0a, 0x25, 0x37, 0x38, 0xa0, 0xfb, 0xb0, 0x25, 0x24,
	0x91, 0x63, 0x94, 0xbc, 0x4b, 0x2a, 0x6a, 0x68, 0x16, 0x95, 0x8e, 0x25, 0x93, 0xe3, 0x5a, 0x52,
	0xe9, 0x44, 0x30, 0x52, 0x04, 0xf4, 0x1c, 0xaa, 0x29, 0xb9, 0x64, 0x94, 0xe3, 0x7a, 0x92, 0x72,
	0x46, 0xa8, 0x56, 0x78, 0xe8, 0x21, 0x14, 0xc2, 0xaf, 0x29, 0xc7, 0x3b, 0xd2, 0x25, 0xfd, 0xa1,
	0x8d, 0xa0, 0xf8, 0x92, 0xdf, 0x96, 0x94, 0x35, 0x97, 0x5c, 0x7e, 0x66, 0x83, 0xf9, 0xdc, 0xf4,
	0x96, 0x4a, 0x6f, 0xc3, 0x23, 0x0f, 0x4d, 0x24, 0xc2, 0xa4, 0x16, 0xa0, 0xe6, 0xef, 0x05, 0xd0,
	0x8f, 0x19, 0xb7, 0xdc, 0x2b, 0xea, 0x7d, 0xd6, 0x0b, 0xf3, 0x12, 0x34, 0x7f, 0xb9, 0x08, 0x95,
	0xa0, 0x7a, 0xf0, 0x8d, 0xd8, 0xe2, 0xea, 0x3e, 0xae, 0x19, 0xc6, 0xcb, 0x05, 0x25, 0xd2, 0x2d,
	0x3d, 0x4b, 0x68, 0x6b, 0x66, 0x89, 0x4c, 0x5b, 0x6f, 0xad, 0x6f, 0xeb, 0xd4, 0x34, 0x91, 0x5f,
	0x33, 0x4d, 0x3c, 0x4e, 0x4f, 0x09, 0x85, 0xe4, 0x53, 0x4b, 0x22, 0xe3, 0xfa, 0x51, 0xa1, 0xf8,
	0xc9, 0xa3, 0x42, 0xfa, 0xee, 0x96, 0xd6, 0xdd, 0xdd, 0xac, 0x26, 0xc0, 0xc7, 0x34, 0xe1, 0x10,
	0xb6, 0xe5, 0xf0, 0x79, 0xac, 0x66, 0xcf, 0xf4, 0x08, 0xd2, 0x4b, 0x03, 0x24, 0xcb, 0x13, 0x1b,
	0x92, 0xb7, 0xb5, 0xe7, 0x7a, 0x6a, 0x66, 0xc8, 0x6e, 0x28, 0x02, 0xc5, 0x67, 0x57, 0x6d, 0xae,
	0xab, 0x66, 0x58, 0x35, 0x31, 0xd4, 0x52, 0xfc, 0x08, 0x22, 0xab, 0xdc, 0xe6, 0x3f, 0x39, 0xa8,
	0xdf, 0x74, 0xf4, 0xa8, 0x0c, 0x85, 0xc9, 0xf0, 0xf5, 0xf0, 0xf4, 0xdd, 0x50, 0xff, 0x1f, 0xaa,
	0x40, 0xb1, 0x3f, 0x1c, 0x8d, 0x3b, 0xc3, 0xae, 0xa1, 0xe7, 0x04, 0xd4, 0x1d, 0x4c, 0x46, 0x63,
	0x83, 0xe8, 0x1b, 0x62, 0x41, 0x4e, 0x27, 0xe3, 0xfe, 0xf0, 0x95, 0xbe, 0x89, 0xaa, 0x00, 0xa4,
	0x33, 0x36, 0xa6, 0x83, 0xfe, 0x49, 0x7f, 0xac, 0x6b, 0xa8, 0x06, 0xb7, 0xba, 0x7d, 0xd2, 0x9d,
	0xf4, 0xc7, 0xd3, 0x23, 0x62, 0x74, 0x5e, 0x1b, 0x44, 0xdf, 0x12, 0xc1, 0x46, 0x06, 0x79, 0xdb,
	0xef, 0x1a, 0x23, 0x3d, 0x2f, 0x5c, 0x86, 0x9d, 0x13, 0x63, 0xf4, 0xa6, 0x23, 0xd6, 0x15, 0x84,
	0xa0, 0xda, 0xeb, 0x4c, 0x06, 0xe3, 0xe9, 0xb1, 0x31, 0x36, 0xba, 0xe3, 0x53, 0xa2, 0x6f, 0xa3,
	0x3a, 0xe8, 0xca, 0x63, 0xda, 0x3d, 0x1d, 0x8e, 0x49, 0xa7, 0x3b, 0xd6, 0xab, 0x4d, 0xad, 0x58,
	0xd0, 0xcb, 0x7b, 0xda, 0x89, 0x31, 0xfa, 0x7e, 0xaf, 0x2c, 0xfe, 0x0a, 0xb0, 0xd7, 0x7f, 0xb5,
	0x57, 0xed, 0x0d, 0x26, 0x3f, 0x4d, 0x8f, 0x8f, 0x88, 0xd1, 0x23, 0x02, 0x2c, 0xca, 0xf5, 0xe8,
	0xf8, 0xf5, 0x5e, 0x39, 0x7c, 0x32, 0xc8, 0x5b, 0x83, 0xc8, 0x4b, 0xb9, 0xd3, 0xfc, 0x35, 0x07,
	0x95, 0xf4, 0xac, 0x80, 0x0e, 0xa1, 0x10, 0x4e, 0x0b, 0x1c, 0xe7, 0xe4, 0xf1, 0xde, 0x5b, 0x1d,
	0x27, 0xd4, 0x82, 0x1b, 0x8e, 0x2f, 0x2e, 0xb9, 0x62, 0x37, 0x5e, 0x44, 0x81, 0x42, 0x00, 0xe9,
	0xb0, 0xf9, 0x9e, 0x2e, 0xe5, 0xc5, 0x2e, 0x11, 0xf1, 0x88, 0xea, 0xb0, 0x75, 0x25, 0xae, 0xa5,
	0xbc, 0xba, 0x25, 0x12, 0x2e, 0x5e, 0x6c, 0x3c, 0xcf, 0x35, 0xff, 0xc8, 0x41, 0x35, 0x3b, 0x79,
	0xa0, 0xef, 0x20, 0x6f, 0x87, 0xd3, 0x49, 0xb8, 0x8d, 0xfb, 0xd7, 0xa7, 0x93, 0x56, 0xf8, 0x2f,
	0xdc, 0x87, 0x62, 0x37, 0xfa, 0x50, 0x4e, 0x99, 0x6f, 0xd8, 0xc5, 0xc3, 0xf4, 0x2e, 0xca, 0x07,
	0x55, 0x35, 0x65, 0x31, 0x67, 0x36, 0x60, 0xdc, 0x4f, 0xed, 0xea, 0xe8, 0xb7, 0x1c, 0x1c, 0x5a,
	0xee, 0xbc, 0xe5, 0x53, 0xc7, 0x92, 0xbf, 0xb1, 0x5c, 0xdb, 0xf4, 0x18, 0x6f, 0xf1, 0x05, 0xb5,
	0xd8, 0x05, 0xb3, 0xe4, 0xe4, 0xd9, 0x32, 0x17, 0x4c, 0x84, 0x88, 0x7e, 0xd9, 0xcd, 0x4d, 0xc7,
	0x9c, 0xd1, 0xa3, 0xed, 0xa8, 0x9b, 0xde, 0x08, 0x5d, 0xfa, 0xf9, 0x78, 0xc6, 0xfc, 0xcb, 0xe0,
	0xac, 0x65, 0xb9, 0xf3, 0xb6, 0x0a, 0x33, 0xa7, 0xfc, 0xb2, 0x9d, 0x09, 0xd5, 0x0e, 0xe7, 0xba,
	0xf6, 0xcc, 0x6d, 0x9b, 0x0b, 0xd6, 0xbe, 0xda, 0x6f, 0xab, 0xa0, 0xd3, 0x30, 0xe8, 0x59, 0x5e,
	0x8a, 0xdc, 0xd3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x34, 0xfd, 0xda, 0xa7, 0x0e, 0x00,
	0x00,
}
