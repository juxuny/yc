package handler

import (
	"context"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/trace"
	"runtime/debug"
)

func (t *wrapper) SaveValue(ctx context.Context, req *cos.SaveValueRequest) (resp *cos.SaveValueResponse, err error) {
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
	return t.handler.SaveValue(ctx, req)
}

func (t *wrapper) DeleteValue(ctx context.Context, req *cos.DeleteValueRequest) (resp *cos.DeleteValueRequest, err error) {
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
	return t.handler.DeleteValue(ctx, req)
}

func (t *wrapper) ListValue(ctx context.Context, req *cos.ListValueRequest) (resp *cos.ListValueResponse, err error) {
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
	return t.handler.ListValue(ctx, req)
}
func (t *wrapper) DisableValue(ctx context.Context, req *cos.DisableValueRequest) (resp *cos.DisableValueResponse, err error) {
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
	return t.handler.DisableValue(ctx, req)
}

func (t *wrapper) ListAllValue(ctx context.Context, req *cos.ListAllValueRequest) (resp *cos.ListAllValueResponse, err error) {
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
	return t.handler.ListAllValue(ctx, req)
}
