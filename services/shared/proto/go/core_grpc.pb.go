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
	GetServiceConfig(ctx context.Context, in *ConfigCenter_GetServiceConfigReq, opts ...grpc.CallOption) (*ConfigCenter_GetServiceConfigRes, error)
}

type iConfigCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewIConfigCenterClient(cc grpc.ClientConnInterface) IConfigCenterClient {
	return &iConfigCenterClient{cc}
}

func (c *iConfigCenterClient) GetServiceConfig(ctx context.Context, in *ConfigCenter_GetServiceConfigReq, opts ...grpc.CallOption) (*ConfigCenter_GetServiceConfigRes, error) {
	out := new(ConfigCenter_GetServiceConfigRes)
	err := c.cc.Invoke(ctx, "/core.IConfigCenter/GetServiceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IConfigCenterServer is the server API for IConfigCenter service.
// All implementations must embed UnimplementedIConfigCenterServer
// for forward compatibility
type IConfigCenterServer interface {
	GetServiceConfig(context.Context, *ConfigCenter_GetServiceConfigReq) (*ConfigCenter_GetServiceConfigRes, error)
	mustEmbedUnimplementedIConfigCenterServer()
}

// UnimplementedIConfigCenterServer must be embedded to have forward compatible implementations.
type UnimplementedIConfigCenterServer struct {
}

func (UnimplementedIConfigCenterServer) GetServiceConfig(context.Context, *ConfigCenter_GetServiceConfigReq) (*ConfigCenter_GetServiceConfigRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceConfig not implemented")
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

func _IConfigCenter_GetServiceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigCenter_GetServiceConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IConfigCenterServer).GetServiceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.IConfigCenter/GetServiceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IConfigCenterServer).GetServiceConfig(ctx, req.(*ConfigCenter_GetServiceConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IConfigCenter_ServiceDesc is the grpc.ServiceDesc for IConfigCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IConfigCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.IConfigCenter",
	HandlerType: (*IConfigCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServiceConfig",
			Handler:    _IConfigCenter_GetServiceConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

// IRegistrationCenterCoreClient is the client API for IRegistrationCenterCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IRegistrationCenterCoreClient interface {
	Register(ctx context.Context, in *RegistrationCenterCore_RegisterReq, opts ...grpc.CallOption) (*RegistrationCenterCore_RegisterRes, error)
	ListServiceTarget(ctx context.Context, in *RegistrationCenterCore_ListServiceTargetReq, opts ...grpc.CallOption) (*RegistrationCenterCore_ListServiceTargetRes, error)
}

type iRegistrationCenterCoreClient struct {
	cc grpc.ClientConnInterface
}

func NewIRegistrationCenterCoreClient(cc grpc.ClientConnInterface) IRegistrationCenterCoreClient {
	return &iRegistrationCenterCoreClient{cc}
}

func (c *iRegistrationCenterCoreClient) Register(ctx context.Context, in *RegistrationCenterCore_RegisterReq, opts ...grpc.CallOption) (*RegistrationCenterCore_RegisterRes, error) {
	out := new(RegistrationCenterCore_RegisterRes)
	err := c.cc.Invoke(ctx, "/core.IRegistrationCenterCore/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iRegistrationCenterCoreClient) ListServiceTarget(ctx context.Context, in *RegistrationCenterCore_ListServiceTargetReq, opts ...grpc.CallOption) (*RegistrationCenterCore_ListServiceTargetRes, error) {
	out := new(RegistrationCenterCore_ListServiceTargetRes)
	err := c.cc.Invoke(ctx, "/core.IRegistrationCenterCore/ListServiceTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IRegistrationCenterCoreServer is the server API for IRegistrationCenterCore service.
// All implementations must embed UnimplementedIRegistrationCenterCoreServer
// for forward compatibility
type IRegistrationCenterCoreServer interface {
	Register(context.Context, *RegistrationCenterCore_RegisterReq) (*RegistrationCenterCore_RegisterRes, error)
	ListServiceTarget(context.Context, *RegistrationCenterCore_ListServiceTargetReq) (*RegistrationCenterCore_ListServiceTargetRes, error)
	mustEmbedUnimplementedIRegistrationCenterCoreServer()
}

// UnimplementedIRegistrationCenterCoreServer must be embedded to have forward compatible implementations.
type UnimplementedIRegistrationCenterCoreServer struct {
}

func (UnimplementedIRegistrationCenterCoreServer) Register(context.Context, *RegistrationCenterCore_RegisterReq) (*RegistrationCenterCore_RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIRegistrationCenterCoreServer) ListServiceTarget(context.Context, *RegistrationCenterCore_ListServiceTargetReq) (*RegistrationCenterCore_ListServiceTargetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServiceTarget not implemented")
}
func (UnimplementedIRegistrationCenterCoreServer) mustEmbedUnimplementedIRegistrationCenterCoreServer() {
}

// UnsafeIRegistrationCenterCoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IRegistrationCenterCoreServer will
// result in compilation errors.
type UnsafeIRegistrationCenterCoreServer interface {
	mustEmbedUnimplementedIRegistrationCenterCoreServer()
}

func RegisterIRegistrationCenterCoreServer(s grpc.ServiceRegistrar, srv IRegistrationCenterCoreServer) {
	s.RegisterService(&IRegistrationCenterCore_ServiceDesc, srv)
}

func _IRegistrationCenterCore_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationCenterCore_RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IRegistrationCenterCoreServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.IRegistrationCenterCore/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IRegistrationCenterCoreServer).Register(ctx, req.(*RegistrationCenterCore_RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IRegistrationCenterCore_ListServiceTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationCenterCore_ListServiceTargetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IRegistrationCenterCoreServer).ListServiceTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.IRegistrationCenterCore/ListServiceTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IRegistrationCenterCoreServer).ListServiceTarget(ctx, req.(*RegistrationCenterCore_ListServiceTargetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IRegistrationCenterCore_ServiceDesc is the grpc.ServiceDesc for IRegistrationCenterCore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IRegistrationCenterCore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.IRegistrationCenterCore",
	HandlerType: (*IRegistrationCenterCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _IRegistrationCenterCore_Register_Handler,
		},
		{
			MethodName: "ListServiceTarget",
			Handler:    _IRegistrationCenterCore_ListServiceTarget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

// IRegistrationCenterEmbeddedClient is the client API for IRegistrationCenterEmbedded service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IRegistrationCenterEmbeddedClient interface {
	CheckHealth(ctx context.Context, in *RegistrationCenterEmbedded_CheckHealthReq, opts ...grpc.CallOption) (*RegistrationCenterEmbedded_CheckHealthRes, error)
}

type iRegistrationCenterEmbeddedClient struct {
	cc grpc.ClientConnInterface
}

func NewIRegistrationCenterEmbeddedClient(cc grpc.ClientConnInterface) IRegistrationCenterEmbeddedClient {
	return &iRegistrationCenterEmbeddedClient{cc}
}

func (c *iRegistrationCenterEmbeddedClient) CheckHealth(ctx context.Context, in *RegistrationCenterEmbedded_CheckHealthReq, opts ...grpc.CallOption) (*RegistrationCenterEmbedded_CheckHealthRes, error) {
	out := new(RegistrationCenterEmbedded_CheckHealthRes)
	err := c.cc.Invoke(ctx, "/core.IRegistrationCenterEmbedded/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IRegistrationCenterEmbeddedServer is the server API for IRegistrationCenterEmbedded service.
// All implementations must embed UnimplementedIRegistrationCenterEmbeddedServer
// for forward compatibility
type IRegistrationCenterEmbeddedServer interface {
	CheckHealth(context.Context, *RegistrationCenterEmbedded_CheckHealthReq) (*RegistrationCenterEmbedded_CheckHealthRes, error)
	mustEmbedUnimplementedIRegistrationCenterEmbeddedServer()
}

// UnimplementedIRegistrationCenterEmbeddedServer must be embedded to have forward compatible implementations.
type UnimplementedIRegistrationCenterEmbeddedServer struct {
}

func (UnimplementedIRegistrationCenterEmbeddedServer) CheckHealth(context.Context, *RegistrationCenterEmbedded_CheckHealthReq) (*RegistrationCenterEmbedded_CheckHealthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedIRegistrationCenterEmbeddedServer) mustEmbedUnimplementedIRegistrationCenterEmbeddedServer() {
}

// UnsafeIRegistrationCenterEmbeddedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IRegistrationCenterEmbeddedServer will
// result in compilation errors.
type UnsafeIRegistrationCenterEmbeddedServer interface {
	mustEmbedUnimplementedIRegistrationCenterEmbeddedServer()
}

func RegisterIRegistrationCenterEmbeddedServer(s grpc.ServiceRegistrar, srv IRegistrationCenterEmbeddedServer) {
	s.RegisterService(&IRegistrationCenterEmbedded_ServiceDesc, srv)
}

func _IRegistrationCenterEmbedded_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationCenterEmbedded_CheckHealthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IRegistrationCenterEmbeddedServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.IRegistrationCenterEmbedded/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IRegistrationCenterEmbeddedServer).CheckHealth(ctx, req.(*RegistrationCenterEmbedded_CheckHealthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IRegistrationCenterEmbedded_ServiceDesc is the grpc.ServiceDesc for IRegistrationCenterEmbedded service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IRegistrationCenterEmbedded_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.IRegistrationCenterEmbedded",
	HandlerType: (*IRegistrationCenterEmbeddedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _IRegistrationCenterEmbedded_CheckHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
