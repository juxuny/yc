// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cos

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

// CosClient is the client API for Cos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CosClient interface {
	// @group: devops
	// @ignore-auth
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	// @group: auth
	// @ignore-auth
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// @group: user
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	// @group: user
	UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...grpc.CallOption) (*UpdateInfoResponse, error)
	// @group: user
	ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*ModifyPasswordResponse, error)
	// @group: user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// @group: namespace
	SaveNamespace(ctx context.Context, in *SaveNamespaceRequest, opts ...grpc.CallOption) (*SaveNamespaceResponse, error)
	// @group: namespace
	ListNamespace(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error)
	// @group: namespace
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
}

type cosClient struct {
	cc grpc.ClientConnInterface
}

func NewCosClient(cc grpc.ClientConnInterface) CosClient {
	return &cosClient{cc}
}

func (c *cosClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...grpc.CallOption) (*UpdateInfoResponse, error) {
	out := new(UpdateInfoResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/UpdateInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) ModifyPassword(ctx context.Context, in *ModifyPasswordRequest, opts ...grpc.CallOption) (*ModifyPasswordResponse, error) {
	out := new(ModifyPasswordResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/ModifyPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) SaveNamespace(ctx context.Context, in *SaveNamespaceRequest, opts ...grpc.CallOption) (*SaveNamespaceResponse, error) {
	out := new(SaveNamespaceResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/SaveNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) ListNamespace(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error) {
	out := new(ListNamespaceResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/ListNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/DeleteNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CosServer is the server API for Cos service.
// All implementations must embed UnimplementedCosServer
// for forward compatibility
type CosServer interface {
	// @group: devops
	// @ignore-auth
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	// @group: auth
	// @ignore-auth
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// @group: user
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	// @group: user
	UpdateInfo(context.Context, *UpdateInfoRequest) (*UpdateInfoResponse, error)
	// @group: user
	ModifyPassword(context.Context, *ModifyPasswordRequest) (*ModifyPasswordResponse, error)
	// @group: user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// @group: namespace
	SaveNamespace(context.Context, *SaveNamespaceRequest) (*SaveNamespaceResponse, error)
	// @group: namespace
	ListNamespace(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error)
	// @group: namespace
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	mustEmbedUnimplementedCosServer()
}

// UnimplementedCosServer must be embedded to have forward compatible implementations.
type UnimplementedCosServer struct {
}

func (UnimplementedCosServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedCosServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCosServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedCosServer) UpdateInfo(context.Context, *UpdateInfoRequest) (*UpdateInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInfo not implemented")
}
func (UnimplementedCosServer) ModifyPassword(context.Context, *ModifyPasswordRequest) (*ModifyPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyPassword not implemented")
}
func (UnimplementedCosServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCosServer) SaveNamespace(context.Context, *SaveNamespaceRequest) (*SaveNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveNamespace not implemented")
}
func (UnimplementedCosServer) ListNamespace(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespace not implemented")
}
func (UnimplementedCosServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedCosServer) mustEmbedUnimplementedCosServer() {}

// UnsafeCosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CosServer will
// result in compilation errors.
type UnsafeCosServer interface {
	mustEmbedUnimplementedCosServer()
}

func RegisterCosServer(s grpc.ServiceRegistrar, srv CosServer) {
	s.RegisterService(&Cos_ServiceDesc, srv)
}

func _Cos_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_UpdateInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).UpdateInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/UpdateInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).UpdateInfo(ctx, req.(*UpdateInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_ModifyPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).ModifyPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/ModifyPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).ModifyPassword(ctx, req.(*ModifyPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_SaveNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).SaveNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/SaveNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).SaveNamespace(ctx, req.(*SaveNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_ListNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).ListNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/ListNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).ListNamespace(ctx, req.(*ListNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/DeleteNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cos_ServiceDesc is the grpc.ServiceDesc for Cos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cos.Cos",
	HandlerType: (*CosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Cos_Health_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Cos_Login_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Cos_UserInfo_Handler,
		},
		{
			MethodName: "UpdateInfo",
			Handler:    _Cos_UpdateInfo_Handler,
		},
		{
			MethodName: "ModifyPassword",
			Handler:    _Cos_ModifyPassword_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Cos_CreateUser_Handler,
		},
		{
			MethodName: "SaveNamespace",
			Handler:    _Cos_SaveNamespace_Handler,
		},
		{
			MethodName: "ListNamespace",
			Handler:    _Cos_ListNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _Cos_DeleteNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cos.proto",
}