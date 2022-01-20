// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc_impl

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IConfigCenterClient is the client API for IConfigCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IConfigCenterClient interface {
	// for other services
	GetConfig(ctx context.Context, in *Config_GetConfigReq, opts ...grpc.CallOption) (*Config_GetConfigRes, error)
}

type iConfigCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewIConfigCenterClient(cc grpc.ClientConnInterface) IConfigCenterClient {
	return &iConfigCenterClient{cc}
}

func (c *iConfigCenterClient) GetConfig(ctx context.Context, in *Config_GetConfigReq, opts ...grpc.CallOption) (*Config_GetConfigRes, error) {
	out := new(Config_GetConfigRes)
	err := c.cc.Invoke(ctx, "/configCenter.IConfigCenter/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IConfigCenterServer is the server API for IConfigCenter service.
// All implementations must embed UnimplementedIConfigCenterServer
// for forward compatibility
type IConfigCenterServer interface {
	// for other services
	GetConfig(context.Context, *Config_GetConfigReq) (*Config_GetConfigRes, error)
	mustEmbedUnimplementedIConfigCenterServer()
}

// UnimplementedIConfigCenterServer must be embedded to have forward compatible implementations.
type UnimplementedIConfigCenterServer struct {
}

func (UnimplementedIConfigCenterServer) GetConfig(context.Context, *Config_GetConfigReq) (*Config_GetConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedIConfigCenterServer) mustEmbedUnimplementedIConfigCenterServer() {}

// UnsafeIConfigCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IConfigCenterServer will
// result in compilation errors.
type UnsafeIConfigCenterServer interface {
	mustEmbedUnimplementedIConfigCenterServer()
}

func RegisterIConfigCenterServer(s grpc.ServiceRegistrar, srv IConfigCenterServer) {
	s.RegisterService(&IConfigCenter_ServiceDesc, srv)
}

func _IConfigCenter_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config_GetConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IConfigCenterServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configCenter.IConfigCenter/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IConfigCenterServer).GetConfig(ctx, req.(*Config_GetConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IConfigCenter_ServiceDesc is the grpc.ServiceDesc for IConfigCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IConfigCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "configCenter.IConfigCenter",
	HandlerType: (*IConfigCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _IConfigCenter_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config_center.proto",
}
