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

func fixConfigLastSeqNo(ctx context.Context, configId dt.ID) (lastSeqNo uint64, recordType cos.ConfigRecordType, linkCount uint64, err error) {
	modelConfig, found, err := db.TableConfig.FindOneById(ctx, configId)
	if err != nil {
		log.Error(err)
		return lastSeqNo, recordType, 0, err
	}
	if !found {
		return lastSeqNo, recordType, 0, cos.Error.ConfigNotFound.WithField(db.TableConfig.Id.LowerFirstHump(), configId)
	}
	modelConfigRecord, found, err := db.TableConfigRecord.FindOneByConfigId(ctx, configId, orm.DESC(db.TableConfigRecord.SeqNo))
	if err != nil || !found {
		lastSeqNo = 1
		recordType = cos.ConfigRecordType_ConfigRecordTypeCreate
		_, _ = db.TableConfigRecord.Create(ctx, db.ModelConfigRecord{
			ConfigId:   &configId,
			CreateTime: orm.Now(),
			SeqNo:      lastSeqNo,
			RecordType: recordType,
		})
	} else {
		lastSeqNo = modelConfigRecord.SeqNo
		recordType = modelConfigRecord.RecordType
	}

	// calculate remain link count
	where := orm.NewAndWhereWrapper().Gt(db.TableConfigRecord.SeqNo, modelConfig.LastSeqNo).Le(db.TableConfigRecord.SeqNo, lastSeqNo).Eq(db.TableConfigRecord.ConfigId, configId)
	sumOfLink, err := db.TableConfigRecord.Count(ctx, where.Clone().Eq(db.TableConfigRecord.RecordType, cos.ConfigRecordType_ConfigRecordTypeLink))
	if err != nil {
		log.Error(err)
		return lastSeqNo, recordType, 0, err
	}
	sumOfUnlink, err := db.TableConfigRecord.Count(ctx, where.Clone().Eq(db.TableConfigRecord.RecordType, cos.ConfigRecordType_ConfigRecordTypeUnlink))
	if err != nil {
		log.Error(err)
		return lastSeqNo, recordType, 0, err
	}
	if modelConfig.LinkCount+uint64(sumOfLink) < uint64(sumOfUnlink) {
		linkCount = 0
	} else {
		linkCount = modelConfig.LinkCount + uint64(sumOfLink) - uint64(sumOfUnlink)
	}

	where = orm.NewAndWhereWrapper().Eq(db.TableConfig.Id, configId)
	if modelConfig.LastSeqNo > 0 {
		where.Eq(db.TableConfig.LastSeqNo, modelConfig.LastSeqNo)
	}
	update := orm.H{
		db.TableConfig.LastRecordType: recordType,
		db.TableConfig.LastSeqNo:      lastSeqNo,
		db.TableConfig.LinkCount:      linkCount,
	}
	if modelConfig.DeletedAt == 0 && recordType == cos.ConfigRecordType_ConfigRecordTypeDelete {
		update[db.TableConfig.DeletedAt] = orm.Now()
	}
	_, err = db.TableConfig.Update(ctx, update, where)
	return
}

func CreateConfig(ctx context.Context, modelConfig db.ModelConfig) (lastInsertId dt.ID, err error) {
	modelConfig.DeletedAt = orm.Now()
	modelConfig.LastSeqNo = 1
	modelConfig.LastRecordType = cos.ConfigRecordType_ConfigRecordTypeCreate
	lastInsertId, err = db.TableConfig.CreateWithLastId(ctx, modelConfig)
	if err != nil {
		log.Error(err)
		return
	}
	if modelConfig.BaseId != nil && modelConfig.BaseId.Valid {
		err = yc.Retry(func() (bool, error) {
			baseConfigLastSeqNo, baseConfigLastRecordType, _, err := fixConfigLastSeqNo(ctx, *modelConfig.BaseId)
			if err != nil {
				log.Error(err)
				return false, err
			}
			if baseConfigLastRecordType == cos.ConfigRecordType_ConfigRecordTypeDelete {
				return true, cos.Error.ConfigNotFound
			}
			records := []db.ModelConfigRecord{
				{
					ConfigId:   &lastInsertId,
					CreateTime: orm.Now(),
					SeqNo:      modelConfig.LastSeqNo,
					RecordType: modelConfig.LastRecordType,
				},
				{
					ConfigId:   modelConfig.BaseId,
					CreateTime: orm.Now(),
					SeqNo:      baseConfigLastSeqNo + 1,
					RecordType: cos.ConfigRecordType_ConfigRecordTypeLink,
				},
			}
			if _, err := db.TableConfigRecord.Create(ctx, records...); err != nil {
				log.Error(err)
				return false, err
			}
			return false, nil
		})
		if err != nil {
			log.Error(err)
			return lastInsertId, err
		}
		_, _, _, err = fixConfigLastSeqNo(ctx, *modelConfig.BaseId)
		if err != nil {
			log.Error(err)
		}
		_, err = db.TableConfig.ResetDeletedAt(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.Id, lastInsertId))
	} else {
		modelConfigRecord := db.ModelConfigRecord{
			ConfigId:   &lastInsertId,
			CreateTime: orm.Now(),
			SeqNo:      modelConfig.LastSeqNo,
			RecordType: modelConfig.LastRecordType,
		}
		_, err = db.TableConfigRecord.Create(ctx, modelConfigRecord)
		if err != nil {
			log.Error(err)
			return
		}
		_, err = db.TableConfig.ResetDeletedAt(ctx, orm.NewAndWhereWrapper().Eq(db.TableConfig.Id, lastInsertId))
	}
	return
}

func DeleteConfig(ctx context.Context, id dt.ID) error {
	err := yc.Retry(func() (isEnd bool, err error) {
		modelConfig, found, err := db.TableConfig.FindOneById(ctx, id)
		if err != nil {
			log.Error(err)
			return true, err
		}
		if !found {
			log.Error(err)
			return true, cos.Error.ConfigNotFound
		}
		configLastSeqNo, configLastRecordType, linkCount, err := fixConfigLastSeqNo(ctx, id)
		if err != nil {
			log.Error(err)
			return false, err
		}
		if linkCount > 0 {
			log.Error("delete not allowed, link count: ", linkCount)
			return true, cos.Error.DeleteNotAllowed.WithField(db.TableConfig.LinkCount.LowerFirstHump(), linkCount)
		}
		if configLastRecordType == cos.ConfigRecordType_ConfigRecordTypeDelete {
			return true, nil
		}
		records := []db.ModelConfigRecord{
			{
				ConfigId:   &id,
				CreateTime: orm.Now(),
				SeqNo:      configLastSeqNo + 1,
				RecordType: cos.ConfigRecordType_ConfigRecordTypeDelete,
			},
		}
		if modelConfig.BaseId != nil && modelConfig.BaseId.Valid {
			baseConfigLastSeqNo, _, _, err := fixConfigLastSeqNo(ctx, *modelConfig.BaseId)
			if err != nil {
				log.Error(err)
				return false, err
			}
			records = append(records, db.ModelConfigRecord{
				ConfigId:   modelConfig.BaseId,
				CreateTime: orm.Now(),
				SeqNo:      baseConfigLastSeqNo + 1,
				RecordType: cos.ConfigRecordType_ConfigRecordTypeUnlink,
			})
		}
		_, err = db.TableConfigRecord.Create(ctx, records...)
		if err != nil {
			log.Error(err)
			return false, err
		}
		_, _, _, err = fixConfigLastSeqNo(ctx, id)
		if err != nil {
			log.Error(err)
		}
		if modelConfig.BaseId != nil && modelConfig.BaseId.Valid {
			_, _, _, err = fixConfigLastSeqNo(ctx, *modelConfig.BaseId)
			if err != nil {
				log.Error(err)
			}
		}
		return false, nil
	})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
