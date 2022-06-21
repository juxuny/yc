package handler

import (
	"context"
	"github.com/juxuny/yc/middle"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *wrapper) SaveNamespace(ctx context.Context, req *cos.SaveNamespaceRequest) (resp *cos.SaveNamespaceResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.SaveNamespace(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) ListNamespace(ctx context.Context, req *cos.ListNamespaceRequest) (resp *cos.ListNamespaceResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.ListNamespace(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) DeleteNamespace(ctx context.Context, req *cos.DeleteNamespaceRequest) (resp *cos.DeleteNamespaceResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.DeleteNamespace(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) UpdateStatusNamespace(ctx context.Context, req *cos.UpdateStatusNamespaceRequest) (resp *cos.UpdateStatusNamespaceResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.UpdateStatusNamespace(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}

func (t *wrapper) SelectorNamespace(ctx context.Context, req *cos.SelectorRequest) (resp *cos.SelectorResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.SelectorNamespace(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}
