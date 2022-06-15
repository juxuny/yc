// Code generated by yc. DO NOT EDIT.
package cos

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/errors"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type Client interface {
	Login(ctx context.Context, req *LoginRequest, extensionMetadata ...metadata.MD) (resp *LoginResponse, err error)
	SaveConfig(ctx context.Context, req *SaveConfigRequest, extensionMetadata ...metadata.MD) (resp *SaveConfigResponse, err error)
	DeleteConfig(ctx context.Context, req *DeleteConfigRequest, extensionMetadata ...metadata.MD) (resp *DeleteConfigResponse, err error)
	ListConfig(ctx context.Context, req *ListConfigRequest, extensionMetadata ...metadata.MD) (resp *ListConfigResponse, err error)
	CloneConfig(ctx context.Context, req *CloneConfigRequest, extensionMetadata ...metadata.MD) (resp *CloneConfigResponse, err error)
	UpdateStatusConfig(ctx context.Context, req *UpdateStatusConfigRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusConfigResponse, err error)
	Health(ctx context.Context, req *HealthRequest, extensionMetadata ...metadata.MD) (resp *HealthResponse, err error)
	SaveValue(ctx context.Context, req *SaveValueRequest, extensionMetadata ...metadata.MD) (resp *SaveValueResponse, err error)
	DeleteValue(ctx context.Context, req *DeleteValueRequest, extensionMetadata ...metadata.MD) (resp *DeleteValueRequest, err error)
	ListValue(ctx context.Context, req *ListValueRequest, extensionMetadata ...metadata.MD) (resp *ListValueResponse, err error)
	DisableValue(ctx context.Context, req *DisableValueRequest, extensionMetadata ...metadata.MD) (resp *DisableValueResponse, err error)
	ListAllValue(ctx context.Context, req *ListAllValueRequest, extensionMetadata ...metadata.MD) (resp *ListAllValueResponse, err error)
	UpdateStatusValue(ctx context.Context, req *UpdateStatusValueRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusValueResponse, err error)
	SaveNamespace(ctx context.Context, req *SaveNamespaceRequest, extensionMetadata ...metadata.MD) (resp *SaveNamespaceResponse, err error)
	ListNamespace(ctx context.Context, req *ListNamespaceRequest, extensionMetadata ...metadata.MD) (resp *ListNamespaceResponse, err error)
	DeleteNamespace(ctx context.Context, req *DeleteNamespaceRequest, extensionMetadata ...metadata.MD) (resp *DeleteNamespaceResponse, err error)
	UpdateStatusNamespace(ctx context.Context, req *UpdateStatusNamespaceRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusNamespaceResponse, err error)
	SelectorNamespace(ctx context.Context, req *SelectorRequest, extensionMetadata ...metadata.MD) (resp *SelectorResponse, err error)
	UserInfo(ctx context.Context, req *UserInfoRequest, extensionMetadata ...metadata.MD) (resp *UserInfoResponse, err error)
	UpdateInfo(ctx context.Context, req *UpdateInfoRequest, extensionMetadata ...metadata.MD) (resp *UpdateInfoResponse, err error)
	ModifyPassword(ctx context.Context, req *ModifyPasswordRequest, extensionMetadata ...metadata.MD) (resp *ModifyPasswordResponse, err error)
	SaveOrCreateUser(ctx context.Context, req *SaveOrCreateUserRequest, extensionMetadata ...metadata.MD) (resp *SaveOrCreateUserResponse, err error)
	UserList(ctx context.Context, req *UserListRequest, extensionMetadata ...metadata.MD) (resp *UserListResponse, err error)
	UserUpdateStatus(ctx context.Context, req *UserUpdateStatusRequest, extensionMetadata ...metadata.MD) (resp *UserUpdateStatusResponse, err error)
	UserDelete(ctx context.Context, req *UserDeleteRequest, extensionMetadata ...metadata.MD) (resp *UserDeleteResponse, err error)
	AccessKeyList(ctx context.Context, req *AccessKeyListRequest, extensionMetadata ...metadata.MD) (resp *AccessKeyListResponse, err error)
	CreateAccessKey(ctx context.Context, req *CreateAccessKeyRequest, extensionMetadata ...metadata.MD) (resp *CreateAccessKeyResponse, err error)
	UpdateStatusAccessKey(ctx context.Context, req *UpdateStatusAccessKeyRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusAccessKeyResponse, err error)
	DeleteAccessKey(ctx context.Context, req *DeleteAccessKeyRequest, extensionMetadata ...metadata.MD) (resp *DeleteAccessKeyResponse, err error)
	SetRemarkAccessKey(ctx context.Context, req *SetAccessKeyRemarkRequest, extensionMetadata ...metadata.MD) (resp *SetAccessKeyRemarkResponse, err error)
}

type client struct {
	Service              string
	EntrypointDispatcher yc.EntrypointDispatcher
}

var DefaultClient Client

func NewClientWithDispatcher(entrypointDispatcher yc.EntrypointDispatcher) Client {
	return &client{
		Service:              Name,
		EntrypointDispatcher: entrypointDispatcher,
	}
}

func (t *client) Login(ctx context.Context, req *LoginRequest, extensionMetadata ...metadata.MD) (resp *LoginResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &LoginResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) SaveConfig(ctx context.Context, req *SaveConfigRequest, extensionMetadata ...metadata.MD) (resp *SaveConfigResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &SaveConfigResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) DeleteConfig(ctx context.Context, req *DeleteConfigRequest, extensionMetadata ...metadata.MD) (resp *DeleteConfigResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &DeleteConfigResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) ListConfig(ctx context.Context, req *ListConfigRequest, extensionMetadata ...metadata.MD) (resp *ListConfigResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &ListConfigResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) CloneConfig(ctx context.Context, req *CloneConfigRequest, extensionMetadata ...metadata.MD) (resp *CloneConfigResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &CloneConfigResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UpdateStatusConfig(ctx context.Context, req *UpdateStatusConfigRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusConfigResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UpdateStatusConfigResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) Health(ctx context.Context, req *HealthRequest, extensionMetadata ...metadata.MD) (resp *HealthResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &HealthResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) SaveValue(ctx context.Context, req *SaveValueRequest, extensionMetadata ...metadata.MD) (resp *SaveValueResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &SaveValueResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) DeleteValue(ctx context.Context, req *DeleteValueRequest, extensionMetadata ...metadata.MD) (resp *DeleteValueRequest, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &DeleteValueRequest{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) ListValue(ctx context.Context, req *ListValueRequest, extensionMetadata ...metadata.MD) (resp *ListValueResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &ListValueResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) DisableValue(ctx context.Context, req *DisableValueRequest, extensionMetadata ...metadata.MD) (resp *DisableValueResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &DisableValueResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) ListAllValue(ctx context.Context, req *ListAllValueRequest, extensionMetadata ...metadata.MD) (resp *ListAllValueResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &ListAllValueResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UpdateStatusValue(ctx context.Context, req *UpdateStatusValueRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusValueResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UpdateStatusValueResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) SaveNamespace(ctx context.Context, req *SaveNamespaceRequest, extensionMetadata ...metadata.MD) (resp *SaveNamespaceResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &SaveNamespaceResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) ListNamespace(ctx context.Context, req *ListNamespaceRequest, extensionMetadata ...metadata.MD) (resp *ListNamespaceResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &ListNamespaceResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) DeleteNamespace(ctx context.Context, req *DeleteNamespaceRequest, extensionMetadata ...metadata.MD) (resp *DeleteNamespaceResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &DeleteNamespaceResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UpdateStatusNamespace(ctx context.Context, req *UpdateStatusNamespaceRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusNamespaceResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UpdateStatusNamespaceResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) SelectorNamespace(ctx context.Context, req *SelectorRequest, extensionMetadata ...metadata.MD) (resp *SelectorResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &SelectorResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UserInfo(ctx context.Context, req *UserInfoRequest, extensionMetadata ...metadata.MD) (resp *UserInfoResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UserInfoResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UpdateInfo(ctx context.Context, req *UpdateInfoRequest, extensionMetadata ...metadata.MD) (resp *UpdateInfoResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UpdateInfoResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) ModifyPassword(ctx context.Context, req *ModifyPasswordRequest, extensionMetadata ...metadata.MD) (resp *ModifyPasswordResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &ModifyPasswordResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) SaveOrCreateUser(ctx context.Context, req *SaveOrCreateUserRequest, extensionMetadata ...metadata.MD) (resp *SaveOrCreateUserResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &SaveOrCreateUserResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UserList(ctx context.Context, req *UserListRequest, extensionMetadata ...metadata.MD) (resp *UserListResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UserListResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UserUpdateStatus(ctx context.Context, req *UserUpdateStatusRequest, extensionMetadata ...metadata.MD) (resp *UserUpdateStatusResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UserUpdateStatusResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UserDelete(ctx context.Context, req *UserDeleteRequest, extensionMetadata ...metadata.MD) (resp *UserDeleteResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UserDeleteResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) AccessKeyList(ctx context.Context, req *AccessKeyListRequest, extensionMetadata ...metadata.MD) (resp *AccessKeyListResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &AccessKeyListResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) CreateAccessKey(ctx context.Context, req *CreateAccessKeyRequest, extensionMetadata ...metadata.MD) (resp *CreateAccessKeyResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &CreateAccessKeyResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) UpdateStatusAccessKey(ctx context.Context, req *UpdateStatusAccessKeyRequest, extensionMetadata ...metadata.MD) (resp *UpdateStatusAccessKeyResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &UpdateStatusAccessKeyResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) DeleteAccessKey(ctx context.Context, req *DeleteAccessKeyRequest, extensionMetadata ...metadata.MD) (resp *DeleteAccessKeyResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &DeleteAccessKeyResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}

func (t *client) SetRemarkAccessKey(ctx context.Context, req *SetAccessKeyRemarkRequest, extensionMetadata ...metadata.MD) (resp *SetAccessKeyRemarkResponse, err error) {
	md := yc.GetHeader(ctx, extensionMetadata...)
	ctx = metadata.NewOutgoingContext(ctx, md)
	resp = &SetAccessKeyRemarkResponse{}
	var code int
	code, err = yc.RpcCall(ctx, t.EntrypointDispatcher.SelectOne(), "/api/"+t.Service+"/health", req, resp, md)
	if err != nil {
		return resp, err
	}
	if code != http.StatusOK {
		return resp, errors.SystemError.HttpError.Wrap(fmt.Errorf("http status = %v", code))
	}
	return
}
