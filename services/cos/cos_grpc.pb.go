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
	SaveOrCreateUser(ctx context.Context, in *SaveOrCreateUserRequest, opts ...grpc.CallOption) (*SaveOrCreateUserResponse, error)
	// @group: user
	UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error)
	// @group: user
	UserUpdateStatus(ctx context.Context, in *UserUpdateStatusRequest, opts ...grpc.CallOption) (*UserUpdateStatusResponse, error)
	// @group: namespace
	SaveNamespace(ctx context.Context, in *SaveNamespaceRequest, opts ...grpc.CallOption) (*SaveNamespaceResponse, error)
	// @group: namespace
	ListNamespace(ctx context.Context, in *ListNamespaceRequest, opts ...grpc.CallOption) (*ListNamespaceResponse, error)
	// @group: namespace
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	// @group: config
	SaveConfig(ctx context.Context, in *SaveConfigRequest, opts ...grpc.CallOption) (*SaveConfigResponse, error)
	// @group: config
	DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*DeleteConfigResponse, error)
	// @group: config
	ListConfig(ctx context.Context, in *ListConfigRequest, opts ...grpc.CallOption) (*ListConfigResponse, error)
	// @group: config
	CloneConfig(ctx context.Context, in *CloneConfigRequest, opts ...grpc.CallOption) (*CloneConfigResponse, error)
	// @group: key_value
	SaveValue(ctx context.Context, in *SaveValueRequest, opts ...grpc.CallOption) (*SaveValueResponse, error)
	// @group: key_value
	DeleteValue(ctx context.Context, in *DeleteValueRequest, opts ...grpc.CallOption) (*DeleteValueRequest, error)
	// @group: key_value
	ListValue(ctx context.Context, in *ListValueRequest, opts ...grpc.CallOption) (*ListValueResponse, error)
	// @group: key_value
	DisableValue(ctx context.Context, in *DisableValueRequest, opts ...grpc.CallOption) (*DisableValueResponse, error)
	// @group: key_value
	ListAllValue(ctx context.Context, in *ListAllValueRequest, opts ...grpc.CallOption) (*ListAllValueResponse, error)
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

func (c *cosClient) SaveOrCreateUser(ctx context.Context, in *SaveOrCreateUserRequest, opts ...grpc.CallOption) (*SaveOrCreateUserResponse, error) {
	out := new(SaveOrCreateUserResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/SaveOrCreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) UserList(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) UserUpdateStatus(ctx context.Context, in *UserUpdateStatusRequest, opts ...grpc.CallOption) (*UserUpdateStatusResponse, error) {
	out := new(UserUpdateStatusResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/UserUpdateStatus", in, out, opts...)
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

func (c *cosClient) SaveConfig(ctx context.Context, in *SaveConfigRequest, opts ...grpc.CallOption) (*SaveConfigResponse, error) {
	out := new(SaveConfigResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/SaveConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*DeleteConfigResponse, error) {
	out := new(DeleteConfigResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) ListConfig(ctx context.Context, in *ListConfigRequest, opts ...grpc.CallOption) (*ListConfigResponse, error) {
	out := new(ListConfigResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/ListConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) CloneConfig(ctx context.Context, in *CloneConfigRequest, opts ...grpc.CallOption) (*CloneConfigResponse, error) {
	out := new(CloneConfigResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/CloneConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) SaveValue(ctx context.Context, in *SaveValueRequest, opts ...grpc.CallOption) (*SaveValueResponse, error) {
	out := new(SaveValueResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/SaveValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) DeleteValue(ctx context.Context, in *DeleteValueRequest, opts ...grpc.CallOption) (*DeleteValueRequest, error) {
	out := new(DeleteValueRequest)
	err := c.cc.Invoke(ctx, "/cos.Cos/DeleteValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) ListValue(ctx context.Context, in *ListValueRequest, opts ...grpc.CallOption) (*ListValueResponse, error) {
	out := new(ListValueResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/ListValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) DisableValue(ctx context.Context, in *DisableValueRequest, opts ...grpc.CallOption) (*DisableValueResponse, error) {
	out := new(DisableValueResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/DisableValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cosClient) ListAllValue(ctx context.Context, in *ListAllValueRequest, opts ...grpc.CallOption) (*ListAllValueResponse, error) {
	out := new(ListAllValueResponse)
	err := c.cc.Invoke(ctx, "/cos.Cos/ListAllValue", in, out, opts...)
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
	SaveOrCreateUser(context.Context, *SaveOrCreateUserRequest) (*SaveOrCreateUserResponse, error)
	// @group: user
	UserList(context.Context, *UserListRequest) (*UserListResponse, error)
	// @group: user
	UserUpdateStatus(context.Context, *UserUpdateStatusRequest) (*UserUpdateStatusResponse, error)
	// @group: namespace
	SaveNamespace(context.Context, *SaveNamespaceRequest) (*SaveNamespaceResponse, error)
	// @group: namespace
	ListNamespace(context.Context, *ListNamespaceRequest) (*ListNamespaceResponse, error)
	// @group: namespace
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	// @group: config
	SaveConfig(context.Context, *SaveConfigRequest) (*SaveConfigResponse, error)
	// @group: config
	DeleteConfig(context.Context, *DeleteConfigRequest) (*DeleteConfigResponse, error)
	// @group: config
	ListConfig(context.Context, *ListConfigRequest) (*ListConfigResponse, error)
	// @group: config
	CloneConfig(context.Context, *CloneConfigRequest) (*CloneConfigResponse, error)
	// @group: key_value
	SaveValue(context.Context, *SaveValueRequest) (*SaveValueResponse, error)
	// @group: key_value
	DeleteValue(context.Context, *DeleteValueRequest) (*DeleteValueRequest, error)
	// @group: key_value
	ListValue(context.Context, *ListValueRequest) (*ListValueResponse, error)
	// @group: key_value
	DisableValue(context.Context, *DisableValueRequest) (*DisableValueResponse, error)
	// @group: key_value
	ListAllValue(context.Context, *ListAllValueRequest) (*ListAllValueResponse, error)
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
func (UnimplementedCosServer) SaveOrCreateUser(context.Context, *SaveOrCreateUserRequest) (*SaveOrCreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveOrCreateUser not implemented")
}
func (UnimplementedCosServer) UserList(context.Context, *UserListRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedCosServer) UserUpdateStatus(context.Context, *UserUpdateStatusRequest) (*UserUpdateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdateStatus not implemented")
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
func (UnimplementedCosServer) SaveConfig(context.Context, *SaveConfigRequest) (*SaveConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveConfig not implemented")
}
func (UnimplementedCosServer) DeleteConfig(context.Context, *DeleteConfigRequest) (*DeleteConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedCosServer) ListConfig(context.Context, *ListConfigRequest) (*ListConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConfig not implemented")
}
func (UnimplementedCosServer) CloneConfig(context.Context, *CloneConfigRequest) (*CloneConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneConfig not implemented")
}
func (UnimplementedCosServer) SaveValue(context.Context, *SaveValueRequest) (*SaveValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveValue not implemented")
}
func (UnimplementedCosServer) DeleteValue(context.Context, *DeleteValueRequest) (*DeleteValueRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteValue not implemented")
}
func (UnimplementedCosServer) ListValue(context.Context, *ListValueRequest) (*ListValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListValue not implemented")
}
func (UnimplementedCosServer) DisableValue(context.Context, *DisableValueRequest) (*DisableValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableValue not implemented")
}
func (UnimplementedCosServer) ListAllValue(context.Context, *ListAllValueRequest) (*ListAllValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllValue not implemented")
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

func _Cos_SaveOrCreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveOrCreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).SaveOrCreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/SaveOrCreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).SaveOrCreateUser(ctx, req.(*SaveOrCreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).UserList(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_UserUpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).UserUpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/UserUpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).UserUpdateStatus(ctx, req.(*UserUpdateStatusRequest))
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

func _Cos_SaveConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).SaveConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/SaveConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).SaveConfig(ctx, req.(*SaveConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).DeleteConfig(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_ListConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).ListConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/ListConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).ListConfig(ctx, req.(*ListConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_CloneConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).CloneConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/CloneConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).CloneConfig(ctx, req.(*CloneConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_SaveValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).SaveValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/SaveValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).SaveValue(ctx, req.(*SaveValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_DeleteValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).DeleteValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/DeleteValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).DeleteValue(ctx, req.(*DeleteValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_ListValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).ListValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/ListValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).ListValue(ctx, req.(*ListValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_DisableValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).DisableValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/DisableValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).DisableValue(ctx, req.(*DisableValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cos_ListAllValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CosServer).ListAllValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cos.Cos/ListAllValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CosServer).ListAllValue(ctx, req.(*ListAllValueRequest))
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
			MethodName: "SaveOrCreateUser",
			Handler:    _Cos_SaveOrCreateUser_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Cos_UserList_Handler,
		},
		{
			MethodName: "UserUpdateStatus",
			Handler:    _Cos_UserUpdateStatus_Handler,
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
		{
			MethodName: "SaveConfig",
			Handler:    _Cos_SaveConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _Cos_DeleteConfig_Handler,
		},
		{
			MethodName: "ListConfig",
			Handler:    _Cos_ListConfig_Handler,
		},
		{
			MethodName: "CloneConfig",
			Handler:    _Cos_CloneConfig_Handler,
		},
		{
			MethodName: "SaveValue",
			Handler:    _Cos_SaveValue_Handler,
		},
		{
			MethodName: "DeleteValue",
			Handler:    _Cos_DeleteValue_Handler,
		},
		{
			MethodName: "ListValue",
			Handler:    _Cos_ListValue_Handler,
		},
		{
			MethodName: "DisableValue",
			Handler:    _Cos_DisableValue_Handler,
		},
		{
			MethodName: "ListAllValue",
			Handler:    _Cos_ListAllValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cos.proto",
}
