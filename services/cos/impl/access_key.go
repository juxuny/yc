package impl

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
)

func CheckIfAllowToAccessAccessKey(ctx context.Context, id *dt.ID) error {
	userId, _ := yc.GetUserId(ctx)
	modelAccessKey, found, err := db.TableAccessKey.FindOneById(ctx, id)
	if err != nil {
		log.Error(err)
		return err
	}
	if !found {
		return cos.Error.AccessKeyNotFound
	}
	if modelAccessKey.UserId == nil || !modelAccessKey.UserId.Equal(userId) {
		log.Debug("no permission access access_key id =", id.NumberAsString(), " accessKey:", modelAccessKey.AccessKey)
		return cos.Error.NoPermissionToAccessAccessKey
	}
	return nil
}
