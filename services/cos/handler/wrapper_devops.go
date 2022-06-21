package handler

import (
	"context"
	"github.com/juxuny/yc/middle"
	cos "github.com/juxuny/yc/services/cos"
)

func (t *wrapper) Health(ctx context.Context, req *cos.HealthRequest) (resp *cos.HealthResponse, err error) {
	if err := t.runMiddle(ctx, true, req, middle.NewApiHandler(func(ctx context.Context) {
		resp, err = t.handler.Health(ctx, req)
	})); err != nil {
		return nil, err
	}
	return
}
