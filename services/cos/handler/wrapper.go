package handler

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/middle"
	"github.com/juxuny/yc/services/cos/config"
	"github.com/juxuny/yc/services/cos/db"

	cos "github.com/juxuny/yc/services/cos"
)

type wrapper struct {
	handler       *handler
	authHandler   middle.Group
	beforeHandler middle.Group
	afterHandler  middle.Group
	cos.UnimplementedCosServer
}

func NewWrapper() *wrapper {
	return &wrapper{
		authHandler:   middle.NewGroup().Add(&authValidator{}),
		beforeHandler: middle.NewGroup().Add(&levelValidator{}),
		afterHandler:  middle.NewGroup(),
		handler:       &handler{},
	}
}

type levelValidator struct{}

func (t *levelValidator) Run(ctx context.Context) (isEnd bool, err error) {
	if config.Env.IgnoreCallLevel {
		return false, nil
	}
	callerLevel, err := yc.GetCallerLevelFromContext(ctx)
	if err != nil {
		return true, err
	}
	if callerLevel <= cos.Level {
		log.Error("caller service's level is smaller than current")
		return true, errors.SystemError.RpcCallErrorLevelNotAllow
	}
	return
}

type authValidator struct{}

func (t *authValidator) Run(ctx context.Context) (isEnd bool, err error) {
	claims, err := yc.ParseJwt(ctx)
	if err != nil {
		log.Error(err)
		return true, err
	}
	modelAccount, found, err := db.TableAccount.FindOneById(ctx, claims.UserId)
	if err != nil {
		log.Error(err)
		return true, err
	}
	if !found {
		log.Error("user not found: userId =", claims.UserId)
		return true, cos.Error.AccountNotFound
	}
	if modelAccount.IsDisabled {
		log.Errorf("account forbidden, userId:%v, userName:%v", claims.UserId, claims.UserName)
		return true, cos.Error.AccountForbidden
	}
	return
}

func handleRecover(ctx context.Context, err interface{}) {
	log.Error(err)
}
