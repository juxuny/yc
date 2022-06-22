package handler

import (
	"context"
	"github.com/juxuny/yc/middle"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *wrapper) SaveValue(ctx context.Context, req *cos.SaveValueRequest) (resp *cos.SaveValueResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.SaveValue(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) DeleteValue(ctx context.Context, req *cos.DeleteValueRequest) (resp *cos.DeleteValueRequest, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.DeleteValue(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) ListValue(ctx context.Context, req *cos.ListValueRequest) (resp *cos.ListValueResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.ListValue(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) DisableValue(ctx context.Context, req *cos.DisableValueRequest) (resp *cos.DisableValueResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.DisableValue(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) ListAllValue(ctx context.Context, req *cos.ListAllValueRequest) (resp *cos.ListAllValueResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.ListAllValue(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UpdateStatusValue(ctx context.Context, req *cos.UpdateStatusValueRequest) (resp *cos.UpdateStatusValueResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UpdateStatusValue(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) ListAllValueByConfigId(ctx context.Context, req *cos.ListAllValueByConfigIdRequest) (resp *cos.ListAllValueByConfigIdResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.ListAllValueByConfigId(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}
