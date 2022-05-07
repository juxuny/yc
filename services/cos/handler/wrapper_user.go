package handler

import (
	"context"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/trace"
	"runtime/debug"
)

func (t *wrapper) UserInfo(ctx context.Context, req *cos.UserInfoRequest) (resp *cos.UserInfoResponse, err error) {
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
	return t.handler.UserInfo(ctx, req)
}

func (t *wrapper) UpdateInfo(ctx context.Context, req *cos.UpdateInfoRequest) (resp *cos.UpdateInfoResponse, err error) {
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
	return t.handler.UpdateInfo(ctx, req)
}

func (t *wrapper) ModifyPassword(ctx context.Context, req *cos.ModifyPasswordRequest) (resp *cos.ModifyPasswordResponse, err error) {
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
	return t.handler.ModifyPassword(ctx, req)
}

func (t *wrapper) SaveOrCreateUser(ctx context.Context, req *cos.SaveOrCreateUserRequest) (resp *cos.SaveOrCreateUserResponse, err error) {
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
	return t.handler.SaveOrCreateUser(ctx, req)
}
