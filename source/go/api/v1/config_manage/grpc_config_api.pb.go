// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_config_api.proto

package config_manage // import "github.com/polarismesh/specification/source/go/api/v1/config_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolarisConfigGRPCClient is the client API for PolarisConfigGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisConfigGRPCClient interface {
	// 拉取配置
	GetConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 创建配置
	CreateConfigFile(ctx context.Context, in *ConfigFile, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 更新配置
	UpdateConfigFile(ctx context.Context, in *ConfigFile, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 发布配置
	PublishConfigFile(ctx context.Context, in *ConfigFileRelease, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 发布配置
	UpsertAndPublishConfigFile(ctx context.Context, in *ConfigFilePublishInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 订阅配置变更
	WatchConfigFiles(ctx context.Context, in *ClientWatchConfigFileRequest, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 拉取指定配置分组下的配置文件列表
	GetConfigFileMetadataList(ctx context.Context, in *ConfigFileGroupRequest, opts ...grpc.CallOption) (*ConfigClientListResponse, error)
	// 统一发现接口
	Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisConfigGRPC_DiscoverClient, error)
}

type polarisConfigGRPCClient struct {
	cc *grpc.ClientConn
}

func NewPolarisConfigGRPCClient(cc *grpc.ClientConn) PolarisConfigGRPCClient {
	return &polarisConfigGRPCClient{cc}
}

func (c *polarisConfigGRPCClient) GetConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/GetConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) CreateConfigFile(ctx context.Context, in *ConfigFile, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/CreateConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) UpdateConfigFile(ctx context.Context, in *ConfigFile, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/UpdateConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) PublishConfigFile(ctx context.Context, in *ConfigFileRelease, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/PublishConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) UpsertAndPublishConfigFile(ctx context.Context, in *ConfigFilePublishInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/UpsertAndPublishConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) WatchConfigFiles(ctx context.Context, in *ClientWatchConfigFileRequest, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/WatchConfigFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) GetConfigFileMetadataList(ctx context.Context, in *ConfigFileGroupRequest, opts ...grpc.CallOption) (*ConfigClientListResponse, error) {
	out := new(ConfigClientListResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/GetConfigFileMetadataList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisConfigGRPC_DiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PolarisConfigGRPC_serviceDesc.Streams[0], "/v1.PolarisConfigGRPC/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &polarisConfigGRPCDiscoverClient{stream}
	return x, nil
}

type PolarisConfigGRPC_DiscoverClient interface {
	Send(*ConfigDiscoverRequest) error
	Recv() (*ConfigDiscoverResponse, error)
	grpc.ClientStream
}

type polarisConfigGRPCDiscoverClient struct {
	grpc.ClientStream
}

func (x *polarisConfigGRPCDiscoverClient) Send(m *ConfigDiscoverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *polarisConfigGRPCDiscoverClient) Recv() (*ConfigDiscoverResponse, error) {
	m := new(ConfigDiscoverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PolarisConfigGRPCServer is the server API for PolarisConfigGRPC service.
type PolarisConfigGRPCServer interface {
	// 拉取配置
	GetConfigFile(context.Context, *ClientConfigFileInfo) (*ConfigClientResponse, error)
	// 创建配置
	CreateConfigFile(context.Context, *ConfigFile) (*ConfigClientResponse, error)
	// 更新配置
	UpdateConfigFile(context.Context, *ConfigFile) (*ConfigClientResponse, error)
	// 发布配置
	PublishConfigFile(context.Context, *ConfigFileRelease) (*ConfigClientResponse, error)
	// 发布配置
	UpsertAndPublishConfigFile(context.Context, *ConfigFilePublishInfo) (*ConfigClientResponse, error)
	// 订阅配置变更
	WatchConfigFiles(context.Context, *ClientWatchConfigFileRequest) (*ConfigClientResponse, error)
	// 拉取指定配置分组下的配置文件列表
	GetConfigFileMetadataList(context.Context, *ConfigFileGroupRequest) (*ConfigClientListResponse, error)
	// 统一发现接口
	Discover(PolarisConfigGRPC_DiscoverServer) error
}

func RegisterPolarisConfigGRPCServer(s *grpc.Server, srv PolarisConfigGRPCServer) {
	s.RegisterService(&_PolarisConfigGRPC_serviceDesc, srv)
}

func _PolarisConfigGRPC_GetConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).GetConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/GetConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).GetConfigFile(ctx, req.(*ClientConfigFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_CreateConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).CreateConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/CreateConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).CreateConfigFile(ctx, req.(*ConfigFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_UpdateConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).UpdateConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/UpdateConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).UpdateConfigFile(ctx, req.(*ConfigFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_PublishConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigFileRelease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).PublishConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/PublishConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).PublishConfigFile(ctx, req.(*ConfigFileRelease))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_UpsertAndPublishConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigFilePublishInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).UpsertAndPublishConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/UpsertAndPublishConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).UpsertAndPublishConfigFile(ctx, req.(*ConfigFilePublishInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_WatchConfigFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientWatchConfigFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).WatchConfigFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/WatchConfigFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).WatchConfigFiles(ctx, req.(*ClientWatchConfigFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_GetConfigFileMetadataList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigFileGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).GetConfigFileMetadataList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/GetConfigFileMetadataList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).GetConfigFileMetadataList(ctx, req.(*ConfigFileGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolarisConfigGRPCServer).Discover(&polarisConfigGRPCDiscoverServer{stream})
}

type PolarisConfigGRPC_DiscoverServer interface {
	Send(*ConfigDiscoverResponse) error
	Recv() (*ConfigDiscoverRequest, error)
	grpc.ServerStream
}

type polarisConfigGRPCDiscoverServer struct {
	grpc.ServerStream
}

func (x *polarisConfigGRPCDiscoverServer) Send(m *ConfigDiscoverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *polarisConfigGRPCDiscoverServer) Recv() (*ConfigDiscoverRequest, error) {
	m := new(ConfigDiscoverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PolarisConfigGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisConfigGRPC",
	HandlerType: (*PolarisConfigGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigFile",
			Handler:    _PolarisConfigGRPC_GetConfigFile_Handler,
		},
		{
			MethodName: "CreateConfigFile",
			Handler:    _PolarisConfigGRPC_CreateConfigFile_Handler,
		},
		{
			MethodName: "UpdateConfigFile",
			Handler:    _PolarisConfigGRPC_UpdateConfigFile_Handler,
		},
		{
			MethodName: "PublishConfigFile",
			Handler:    _PolarisConfigGRPC_PublishConfigFile_Handler,
		},
		{
			MethodName: "UpsertAndPublishConfigFile",
			Handler:    _PolarisConfigGRPC_UpsertAndPublishConfigFile_Handler,
		},
		{
			MethodName: "WatchConfigFiles",
			Handler:    _PolarisConfigGRPC_WatchConfigFiles_Handler,
		},
		{
			MethodName: "GetConfigFileMetadataList",
			Handler:    _PolarisConfigGRPC_GetConfigFileMetadataList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _PolarisConfigGRPC_Discover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_config_api.proto",
}

func init() {
	proto.RegisterFile("grpc_config_api.proto", fileDescriptor_grpc_config_api_75804dc4d93d6af6)
}

var fileDescriptor_grpc_config_api_75804dc4d93d6af6 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x86, 0x37, 0xcb, 0xb2, 0x2c, 0x82, 0x5d, 0x12, 0x43, 0x20, 0x31, 0x7b, 0x58, 0xf6, 0xd4,
	0x93, 0x5c, 0xb7, 0xd0, 0x63, 0xa1, 0x71, 0x1b, 0x53, 0x68, 0xc1, 0x38, 0x84, 0x40, 0x2f, 0x41,
	0x51, 0xc6, 0xce, 0x80, 0x2d, 0xa9, 0x92, 0xec, 0x57, 0xe9, 0x43, 0xf6, 0x25, 0x4a, 0x62, 0x9b,
	0xd8, 0x49, 0x49, 0xa0, 0x47, 0xff, 0xff, 0xcc, 0x67, 0x69, 0xf4, 0x0f, 0x19, 0xa6, 0x5a, 0xf1,
	0x25, 0x97, 0x22, 0xc1, 0x74, 0xc9, 0x14, 0x52, 0xa5, 0xa5, 0x95, 0xce, 0xf7, 0xd2, 0x77, 0x07,
	0xb5, 0x9a, 0x60, 0x06, 0x95, 0xec, 0xba, 0x2d, 0x69, 0xa9, 0xc1, 0x28, 0x29, 0x4c, 0xed, 0x5d,
	0xbd, 0xff, 0x20, 0x83, 0x48, 0x66, 0x4c, 0xa3, 0x09, 0x76, 0x55, 0x61, 0x1c, 0x05, 0xce, 0x03,
	0xf9, 0x1d, 0x82, 0xad, 0x84, 0x29, 0x66, 0xe0, 0x8c, 0x68, 0xe9, 0xd3, 0x20, 0x43, 0x10, 0x2d,
	0xf5, 0x51, 0x24, 0xd2, 0xad, 0x9c, 0x9d, 0x56, 0xf9, 0x71, 0xfd, 0x83, 0xff, 0xdf, 0x9c, 0x5b,
	0xd2, 0x0f, 0x34, 0x30, 0x0b, 0x2d, 0xd2, 0x9f, 0x7d, 0xfd, 0xf6, 0xfb, 0x5c, 0xff, 0x5c, 0xad,
	0xbf, 0xde, 0x3f, 0x25, 0x83, 0xa8, 0x58, 0x65, 0x68, 0x36, 0x2d, 0xc0, 0xb0, 0x0b, 0x88, 0x21,
	0x03, 0x66, 0x4e, 0x73, 0x66, 0xc4, 0x9d, 0x2b, 0x03, 0xda, 0xde, 0x89, 0xf5, 0x31, 0x70, 0xdc,
	0x05, 0xd6, 0x05, 0x67, 0x87, 0x13, 0x91, 0xfe, 0x82, 0x59, 0xde, 0x22, 0x19, 0xe7, 0xdf, 0x7e,
	0xcc, 0x07, 0x5e, 0x0c, 0xaf, 0x05, 0x18, 0x7b, 0x92, 0xb8, 0x20, 0xe3, 0xce, 0xab, 0x3d, 0x83,
	0x65, 0x6b, 0x66, 0xd9, 0x13, 0x1a, 0xeb, 0xb8, 0xdd, 0x53, 0x86, 0x5a, 0x16, 0xaa, 0x81, 0xfe,
	0x3d, 0x84, 0x6e, 0x3b, 0x5a, 0xe0, 0x90, 0xfc, 0xba, 0x47, 0xc3, 0x65, 0x09, 0xba, 0x7d, 0xdb,
	0x46, 0x6b, 0x30, 0xee, 0x67, 0x56, 0x03, 0xb9, 0xe8, 0x5d, 0xf6, 0x26, 0x6f, 0x3d, 0x72, 0xc3,
	0x65, 0x4e, 0x2d, 0x08, 0x0e, 0xc2, 0x52, 0x55, 0x25, 0x8f, 0x1a, 0x05, 0x1c, 0x13, 0xe4, 0xcc,
	0xa2, 0x14, 0x74, 0x1b, 0xe8, 0xd2, 0xa7, 0x55, 0x6a, 0x69, 0xce, 0x04, 0x4b, 0x61, 0x32, 0x3a,
	0x4a, 0xe9, 0x0c, 0x74, 0x89, 0x1c, 0x5e, 0x82, 0x14, 0xed, 0xa6, 0x58, 0x51, 0x2e, 0x73, 0xaf,
	0x06, 0xe6, 0x60, 0x36, 0x5e, 0x07, 0xea, 0x19, 0x59, 0x68, 0x0e, 0x5e, 0x2a, 0x3d, 0xa6, 0xd0,
	0x2b, 0x7d, 0xaf, 0x5e, 0x8a, 0x0a, 0xbf, 0xfa, 0xb9, 0x5b, 0x87, 0xeb, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0xe3, 0x11, 0x0b, 0x5a, 0x03, 0x00, 0x00,
}
