package handler

import (
	"context"
	"github.com/juxuny/yc/middle"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *wrapper) Login(ctx context.Context, req *cos.LoginRequest) (resp *cos.LoginResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.Login(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}
