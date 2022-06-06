package impl

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"
	"github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/services/cos/db"
)

func getMaxSeqNoFromKeyValue(ctx context.Context, configId dt.ID) (maxSeqNo uint64) {
	modelKeyValue, found, err := db.TableKeyValue.FindOneByConfigId(ctx, configId, orm.DESC(db.TableKeyValue.SeqNo))
	if err != nil {
		log.Error(err)
		return 0
	}
	if !found {
		return 0
	}
	return modelKeyValue.SeqNo
}

func SaveValue(ctx context.Context, modelKeyValue db.ModelKeyValue) error {
	// update key-value record
	if modelKeyValue.Id != nil && modelKeyValue.Id.Valid {
		_, err := db.TableKeyValue.UpdateById(ctx, *modelKeyValue.Id, orm.H{
			db.TableKeyValue.ConfigValue: modelKeyValue.ConfigValue,
			db.TableKeyValue.IsHot:       modelKeyValue.IsHot,
			db.TableKeyValue.ValueType:   modelKeyValue.ValueType,
		})
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	}

	// create new key-value
	if modelKeyValue.ConfigId == nil || !modelKeyValue.ConfigId.Valid {
		return cos.Error.MissingConfigId
	}
	return yc.Retry(func() (isEnd bool, err error) {
		maxSeqNo := getMaxSeqNoFromKeyValue(ctx, *modelKeyValue.ConfigId)
		// check duplicated key
		count, err := db.TableKeyValue.Count(ctx, orm.NewAndWhereWrapper().Eq(db.TableKeyValue.ConfigId, modelKeyValue.ConfigId).Le(db.TableKeyValue.SeqNo, maxSeqNo).Eq(db.TableKeyValue.ConfigKey, modelKeyValue.ConfigKey))
		if err != nil {
			log.Error(err)
			return true, err
		}
		if count > 0 {
			return true, cos.Error.KeyDuplicated.WithField(db.TableKeyValue.ConfigKey.LowerFirstHump(), modelKeyValue.ConfigKey)
		}
		modelKeyValue.SeqNo = maxSeqNo + 1
		_, err = db.TableKeyValue.Create(ctx, modelKeyValue)
		if err != nil {
			log.Error(err)
			return false, err
		}
		return false, nil
	})
}
