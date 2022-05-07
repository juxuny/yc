package handler

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/middle"

	{{.PackageAlias}} "{{.GoModuleName}}"
)

type wrapper struct {
	handler *handler
	authHandler middle.Group
	beforeHandler middle.Group
	afterHandler middle.Group
	{{.PackageAlias}}.Unimplemented{{.ServiceStruct}}Server
}

func NewWrapper() *wrapper {
	return &wrapper{
		authHandler: middle.NewGroup().Add(&authValidator{}),
		beforeHandler: middle.NewGroup().Add(&levelValidator{}),
		afterHandler: middle.NewGroup(),
		handler: &handler{},
	}
}

type levelValidator struct {}

func (t *levelValidator) Run(ctx context.Context) (isEnd bool, err error) {
	callerLevel, err := yc.GetCallerLevelFromContext(ctx)
	if err != nil {
		return true, err
	}
	if callerLevel {{.Le}} {{.PackageAlias}}.Level {
		log.Error("caller service's level is smaller than current")
		return true, errors.SystemError.RpcCallLevelNotAllow
	}
	return
}

type authValidator struct {}

func (t *authValidator) Run(ctx context.Context) (isEnd bool, err error) {
	return
}

func handleRecover(ctx context.Context, err interface{}) {
}
