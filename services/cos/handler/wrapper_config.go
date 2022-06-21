package handler

import (
	"context"
	"github.com/juxuny/yc/middle"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *wrapper) SaveConfig(ctx context.Context, req *cos.SaveConfigRequest) (resp *cos.SaveConfigResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.SaveConfig(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) DeleteConfig(ctx context.Context, req *cos.DeleteConfigRequest) (resp *cos.DeleteConfigResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.DeleteConfig(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) ListConfig(ctx context.Context, req *cos.ListConfigRequest) (resp *cos.ListConfigResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.ListConfig(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) CloneConfig(ctx context.Context, req *cos.CloneConfigRequest) (resp *cos.CloneConfigResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.CloneConfig(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UpdateStatusConfig(ctx context.Context, req *cos.UpdateStatusConfigRequest) (resp *cos.UpdateStatusConfigResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UpdateStatusConfig(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}
