package handler

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/middle"
	"github.com/juxuny/yc/router"
	"github.com/juxuny/yc/services/cos/config"
	"github.com/juxuny/yc/services/cos/db"
	"github.com/juxuny/yc/trace"
	"runtime/debug"

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

func (t *wrapper) runMiddle(ctx context.Context, auth bool, requestValidator router.ValidatorHandler, apiHandler middle.Handler) (err error) {
	trace.WithContext(ctx)
	defer trace.Clean()
	defer func() {
		if recoverError := recover(); recoverError != nil {
			err = errors.SystemError.InternalError
			debug.PrintStack()
			handleRecover(ctx, recoverError)
		}
	}()
	groups := middle.NewGroup()
	if auth {
		groups.Add(t.authHandler)
	}
	groups.Add(t.beforeHandler, middle.NewValidatorHandler(requestValidator), apiHandler, t.afterHandler)
	_, _, err = groups.Run(ctx)
	return err
}

type levelValidator struct{}

func (t *levelValidator) Run(ctx context.Context) (next context.Context, isEnd bool, err error) {
	if config.Env.IgnoreCallLevel {
		return ctx, false, nil
	}
	callerLevel, err := yc.GetCallerLevelFromContext(ctx)
	if err != nil {
		return ctx, true, err
	}
	if callerLevel <= cos.Level {
		log.Error("caller service's level is smaller than current")
		return ctx, true, errors.SystemError.RpcCallErrorLevelNotAllow
	}
	return
}

type authValidator struct{}

func (t *authValidator) Run(ctx context.Context) (next context.Context, isEnd bool, err error) {
	var userId *dt.ID
	accessKey, found := yc.GetAccessKey(ctx)
	if found {
		modelAccessKey, found, err := db.TableAccessKey.FindOneByAccessKey(ctx, accessKey)
		if err != nil {
			return ctx, true, err
		}
		if !found {
			return ctx, true, cos.Error.AccessKeyNotFound
		}
		if modelAccessKey.IsDisabled {
			return ctx, true, cos.Error.AccessKeyDisabled
		}
		userId = modelAccessKey.UserId.Clone()
	} else {
		claims, err := yc.ParseJwt(ctx)
		if err != nil {
			log.Error(err)
			return ctx, true, err
		}
		userId = claims.UserId.Clone()
	}
	next = yc.SetIncomingUserId(ctx, userId)
	modelAccount, found, err := db.TableAccount.FindOneById(next, userId)
	if err != nil {
		log.Error(err)
		return next, true, err
	}
	if !found {
		log.Error("user not found: userId =", userId)
		return next, true, cos.Error.AccountNotFound
	}
	if modelAccount.IsDisabled {
		log.Errorf("account forbidden, userId:%v", userId)
		return next, true, cos.Error.AccountForbidden
	}
	return
}

func handleRecover(ctx context.Context, err interface{}) {
	log.Error(err)
}
