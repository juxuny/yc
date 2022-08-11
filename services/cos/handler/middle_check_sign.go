package handler

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/router"
	"github.com/juxuny/yc/services/cos/db"
	"time"
)

func init() {
	router.Use(func(ctx *router.Context) {
		if !router.CheckIfOpenApi(ctx.Request.URL.Path) {
			ctx.Next()
			return
		}
		accessKey := ctx.GetAccessKey()
		if accessKey == "" {
			_, _ = ctx.ERROR(errors.SystemError.InvalidAccessKey)
			ctx.Abort()
			return
		}
		signTimestamp := ctx.GetTimestamp()
		if time.Now().UnixMilli()-signTimestamp > int64(yc.SignExpiresInSeconds*1000) {
			log.Error("sign expired")
			_, _ = ctx.ERROR(errors.SystemError.SignExpired)
			ctx.Abort()
			return
		}
		log.Debug("receive access key: " + accessKey)
		backgroundContext := context.Background()
		modelAccessKey, found, err := db.TableAccessKey.FindOneByAccessKey(backgroundContext, accessKey)
		if err != nil {
			log.Error(err)
			_, _ = ctx.ERROR(errors.SystemError.DatabaseQueryError.Wrap(err))
			ctx.Abort()
			return
		}
		if !found {
			log.Error("access key not found: ", accessKey)
			_, _ = ctx.ERROR(errors.SystemError.InvalidAccessKey.Wrap(fmt.Errorf("not found")))
			ctx.Abort()
			return
		}
		if modelAccessKey.IsDisabled {
			log.Error("access key is disabled")
			_, _ = ctx.ERROR(errors.SystemError.InvalidAccessKey.Wrap(fmt.Errorf("disabled")))
			ctx.Abort()
			return
		}
		body, err := ctx.GetCopyOfBodyBytes()
		if err != nil {
			log.Error(err)
			_, _ = ctx.ERROR(errors.SystemError.RpcCallErrorNetwork.Wrap(err))
			ctx.Abort()
			return
		}
		signHandler := yc.NewDefaultSignHandler(modelAccessKey.AccessKey, modelAccessKey.Secret)
		signMethod, signResult, err := signHandler.Sum(body)
		if err != nil {
			log.Error(err)
			_, _ = ctx.ERROR(errors.SystemError.SignError.Wrap(err))
			ctx.Abort()
			return
		}
		receivedSign := ctx.GetSign()
		receivedSignMethod := ctx.GetSignMethod()
		if signMethod != receivedSignMethod {
			log.Error("invalid sign method: ", receivedSignMethod)
			_, _ = ctx.ERROR(errors.SystemError.InvalidSignMethod.Wrap(fmt.Errorf("method: %v", receivedSignMethod)))
			ctx.Abort()
			return
		}
		if receivedSign != signResult {
			log.Error("invalid sign: received: %s, real: %s", receivedSign, signResult)
			return
		}
		log.Debug("sign verified")
		ctx.Next()
	})
}
