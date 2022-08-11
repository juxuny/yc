package handler

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/jwt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/router"
	"github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
)

func init() {
	router.Use(func(ctx *router.Context) {
		if router.CheckIfIgnoreAuth(ctx.Request.URL.Path) {
			ctx.Next()
			return
		}
		var userId *dt.ID
		token := ctx.GetToken()
		if token == "" {
			log.Error("token is empty")
			_, err := ctx.ERROR(errors.SystemError.NotFoundRpcToken)
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			log.Error(err)
			_, err = ctx.ERROR(errors.SystemError.InvalidToken.Wrap(err))
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return
		}
		userId = claims.UserId.Clone()
		ctx.SetUserId(userId)
		backgroundContext := context.Background()
		modelAccount, found, err := db.TableAccount.FindOneById(backgroundContext, userId)
		if err != nil {
			log.Error(err)
			_, err = ctx.ERROR(errors.SystemError.DatabaseQueryError.Wrap(err))
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return
		}
		if !found {
			log.Error("user not found: userId =", userId)
			_, err = ctx.ERROR(cos.Error.AccountNotFound)
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return
		}
		if modelAccount.IsDisabled {
			log.Errorf("account forbidden, userId:%v", userId)
			_, err = ctx.ERROR(cos.Error.AccountForbidden)
			if err != nil {
				log.Error(err)
			}
			ctx.Abort()
			return
		}
		ctx.Next()
	})
}
