package handler

import (
	"context"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/trace"
	"runtime/debug"
)

func (t *wrapper) SaveConfig(ctx context.Context, req *cos.SaveConfigRequest) (resp *cos.SaveConfigResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
			debug.PrintStack()
			handleRecover(ctx, recoverError)
			return
		}
	}()
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
	return t.handler.SaveConfig(ctx, req)
}

func (t *wrapper) DeleteConfig(ctx context.Context, req *cos.DeleteConfigRequest) (resp *cos.DeleteConfigResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
			debug.PrintStack()
			handleRecover(ctx, recoverError)
			return
		}
	}()
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
	return t.handler.DeleteConfig(ctx, req)
}

func (t *wrapper) ListConfig(ctx context.Context, req *cos.ListConfigRequest) (resp *cos.ListConfigResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
			debug.PrintStack()
			handleRecover(ctx, recoverError)
			return
		}
	}()
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
	return t.handler.ListConfig(ctx, req)
}
func (t *wrapper) CloneConfig(ctx context.Context, req *cos.CloneConfigRequest) (resp *cos.CloneConfigResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
			debug.PrintStack()
			handleRecover(ctx, recoverError)
			return
		}
	}()
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
	return t.handler.CloneConfig(ctx, req)
}
func (t *wrapper) UpdateStatusConfig(ctx context.Context, req *cos.UpdateStatusConfigRequest) (resp *cos.UpdateStatusConfigResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
			debug.PrintStack()
			handleRecover(ctx, recoverError)
			return
		}
	}()
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
	return t.handler.UpdateStatusConfig(ctx, req)
}
