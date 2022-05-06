package handler

import (
	"context"
	"github.com/juxuny/yc/log"
	cos "github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/trace"
)

func (t *wrapper) Login(ctx context.Context, req *cos.LoginRequest) (resp *cos.LoginResponse, err error) {
	var isEnd bool
	trace.WithContext(ctx)
	defer trace.Clean()
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
	return t.handler.Login(ctx, req)
}
