package handler

import (
	"context"
	"github.com/juxuny/yc/log"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/trace"
)

func (t *wrapper) SaveNamespace(ctx context.Context, req *cos.SaveNamespaceRequest) (resp *cos.SaveNamespaceResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	isEnd, err = t.authHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	isEnd, err = t.beforeHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	defer func() {
		_, err := t.afterHandler.Run(ctx)
		if err != nil {
			log.Error(err)
		}
	}()
	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()
	return t.handler.SaveNamespace(ctx, req)
}

func (t *wrapper) ListNamespace(ctx context.Context, req *cos.ListNamespaceRequest) (resp *cos.ListNamespaceResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	isEnd, err = t.authHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	isEnd, err = t.beforeHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	defer func() {
		_, err := t.afterHandler.Run(ctx)
		if err != nil {
			log.Error(err)
		}
	}()
	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()
	return t.handler.ListNamespace(ctx, req)
}

func (t *wrapper) DeleteNamespace(ctx context.Context, req *cos.DeleteNamespaceRequest) (resp *cos.DeleteNamespaceResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	isEnd, err = t.authHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	isEnd, err = t.beforeHandler.Run(ctx)
	if err != nil {
		return nil, err
	}
	if isEnd {
		return nil, nil
	}
	defer func() {
		_, err := t.afterHandler.Run(ctx)
		if err != nil {
			log.Error(err)
		}
	}()
	if err := req.Validate(); err != nil {
		log.Error(err)
		return nil, err
	}
	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()
	return t.handler.DeleteNamespace(ctx, req)
}
