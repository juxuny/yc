package handler

import (
	"context"
	"github.com/juxuny/yc/middle"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *wrapper) UserInfo(ctx context.Context, req *cos.UserInfoRequest) (resp *cos.UserInfoResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UserInfo(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UpdateInfo(ctx context.Context, req *cos.UpdateInfoRequest) (resp *cos.UpdateInfoResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UpdateInfo(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) ModifyPassword(ctx context.Context, req *cos.ModifyPasswordRequest) (resp *cos.ModifyPasswordResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.ModifyPassword(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) SaveOrCreateUser(ctx context.Context, req *cos.SaveOrCreateUserRequest) (resp *cos.SaveOrCreateUserResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.SaveOrCreateUser(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UserList(ctx context.Context, req *cos.UserListRequest) (resp *cos.UserListResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UserList(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UserUpdateStatus(ctx context.Context, req *cos.UserUpdateStatusRequest) (resp *cos.UserUpdateStatusResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UserUpdateStatus(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UserDelete(ctx context.Context, req *cos.UserDeleteRequest) (resp *cos.UserDeleteResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UserDelete(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) AccessKeyList(ctx context.Context, req *cos.AccessKeyListRequest) (resp *cos.AccessKeyListResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.AccessKeyList(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) CreateAccessKey(ctx context.Context, req *cos.CreateAccessKeyRequest) (resp *cos.CreateAccessKeyResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.CreateAccessKey(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UpdateStatusAccessKey(ctx context.Context, req *cos.UpdateStatusAccessKeyRequest) (resp *cos.UpdateStatusAccessKeyResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UpdateStatusAccessKey(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) DeleteAccessKey(ctx context.Context, req *cos.DeleteAccessKeyRequest) (resp *cos.DeleteAccessKeyResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.DeleteAccessKey(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) SetRemarkAccessKey(ctx context.Context, req *cos.SetAccessKeyRemarkRequest) (resp *cos.SetAccessKeyRemarkResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.SetRemarkAccessKey(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}
