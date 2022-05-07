package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
)

var TableKeyValue = tableKeyValue{
	Id:          orm.FieldName("id"),
	CreateTime:  orm.FieldName("create_time"),
	UpdateTime:  orm.FieldName("update_time"),
	DeletedAt:   orm.FieldName("deleted_at"),
	IsDisabled:  orm.FieldName("is_disabled"),
	ConfigKey:   orm.FieldName("config_key"),
	ConfigValue: orm.FieldName("config_value"),
	ValueType:   orm.FieldName("value_type"),
	ConfigId:    orm.FieldName("config_id"),
	CreatorId:   orm.FieldName("creator_id"),
	IsHot:       orm.FieldName("is_hot"),
}

type ModelKeyValue struct {
	Id          *dt.ID        `json:"id" orm:"id"`
	CreateTime  int64         `json:"createTime" orm:"create_time"`
	UpdateTime  int64         `json:"updateTime" orm:"update_time"`
	DeletedAt   int64         `json:"deletedAt" orm:"deleted_at"`
	IsDisabled  bool          `json:"isDisabled" orm:"is_disabled"`
	ConfigKey   string        `json:"configKey" orm:"config_key"`
	ConfigValue string        `json:"configValue" orm:"config_value"`
	ValueType   cos.ValueType `json:"valueType" orm:"value_type"`
	ConfigId    string        `json:"configId" orm:"config_id"`
	CreatorId   *dt.ID        `json:"creatorId" orm:"creator_id"`
	IsHot       bool          `json:"isHot" orm:"is_hot"`
}

func (ModelKeyValue) TableName() string {
	return cos.Name + "_" + "key_value"
}

func (t ModelKeyValue) ToKeyValueResp() cos.KeyValueResp {
	return cos.KeyValueResp{
		Id:          t.Id,
		CreateTime:  t.CreateTime,
		UpdateTime:  t.UpdateTime,
		DeletedAt:   t.DeletedAt,
		IsDisabled:  t.IsDisabled,
		ConfigKey:   t.ConfigKey,
		ConfigValue: t.ConfigValue,
		ValueType:   t.ValueType,
		ConfigId:    t.ConfigId,
		CreatorId:   t.CreatorId,
		IsHot:       t.IsHot,
	}
}

func (t ModelKeyValue) ToKeyValueRespAsPointer() *cos.KeyValueResp {
	ret := t.ToKeyValueResp()
	return &ret
}

type tableKeyValue struct {
	Id          orm.FieldName
	CreateTime  orm.FieldName
	UpdateTime  orm.FieldName
	DeletedAt   orm.FieldName
	IsDisabled  orm.FieldName
	ConfigKey   orm.FieldName
	ConfigValue orm.FieldName
	ValueType   orm.FieldName
	ConfigId    orm.FieldName
	CreatorId   orm.FieldName
	IsHot       orm.FieldName
}

func (tableKeyValue) TableName() string {
	return cos.Name + "_" + "key_value"
}

func (tableKeyValue) FindOneById(ctx context.Context, id dt.ID) (data ModelKeyValue, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableKeyValue.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableKeyValue) FindOneByConfigKey(ctx context.Context, configKey string) (data ModelKeyValue, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableKeyValue.ConfigKey, configKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableKeyValue) FindOneByConfigId(ctx context.Context, configId string) (data ModelKeyValue, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableKeyValue.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableKeyValue) FindOneByCreatorId(ctx context.Context, creatorId dt.ID) (data ModelKeyValue, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableKeyValue.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableKeyValue) FindOneByIsHot(ctx context.Context, isHot bool) (data ModelKeyValue, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableKeyValue.IsHot, isHot)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableKeyValue) UpdateById(ctx context.Context, id dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) UpdateByConfigKey(ctx context.Context, configKey string, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigKey, configKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) UpdateByConfigId(ctx context.Context, configId string, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) UpdateByCreatorId(ctx context.Context, creatorId dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) UpdateByIsHot(ctx context.Context, isHot bool, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.IsHot, isHot)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.SetWhere(where).Updates(update)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) DeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) DeleteByConfigKey(ctx context.Context, configKey string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigKey, configKey)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) DeleteByConfigId(ctx context.Context, configId string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigId, configId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) DeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.CreatorId, creatorId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) DeleteByIsHot(ctx context.Context, isHot bool) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.IsHot, isHot)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) SoftDeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.SetValue(TableKeyValue.DeletedAt, orm.Now())
	w.Eq(TableKeyValue.Id, id)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) SoftDeleteByConfigKey(ctx context.Context, configKey string) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.SetValue(TableKeyValue.DeletedAt, orm.Now())
	w.Eq(TableKeyValue.ConfigKey, configKey)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) SoftDeleteByConfigId(ctx context.Context, configId string) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.SetValue(TableKeyValue.DeletedAt, orm.Now())
	w.Eq(TableKeyValue.ConfigId, configId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) SoftDeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.SetValue(TableKeyValue.DeletedAt, orm.Now())
	w.Eq(TableKeyValue.CreatorId, creatorId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) SoftDeleteByIsHot(ctx context.Context, isHot bool) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelKeyValue{})
	w.SetValue(TableKeyValue.DeletedAt, orm.Now())
	w.Eq(TableKeyValue.IsHot, isHot)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableKeyValue) Find(ctx context.Context, where orm.WhereWrapper) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) FindById(ctx context.Context, id dt.ID) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) FindByConfigKey(ctx context.Context, configKey string) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigKey, configKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) FindByConfigId(ctx context.Context, configId string) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) FindByCreatorId(ctx context.Context, creatorId dt.ID) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) FindByIsHot(ctx context.Context, isHot bool) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.IsHot, isHot)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) Page(ctx context.Context, pageNum, pageSize int, where orm.WhereWrapper) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) PageById(ctx context.Context, pageNum, pageSize int, id dt.ID) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) PageByConfigKey(ctx context.Context, pageNum, pageSize int, configKey string) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigKey, configKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) PageByConfigId(ctx context.Context, pageNum, pageSize int, configId string) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) PageByCreatorId(ctx context.Context, pageNum, pageSize int, creatorId dt.ID) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) PageByIsHot(ctx context.Context, pageNum, pageSize int, isHot bool) (list []ModelKeyValue, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.IsHot, isHot)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableKeyValue) Count(ctx context.Context, where orm.WhereWrapper) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableKeyValue) CountById(ctx context.Context, id dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableKeyValue) CountByConfigKey(ctx context.Context, configKey string) (count int, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigKey, configKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableKeyValue) CountByConfigId(ctx context.Context, configId string) (count int, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableKeyValue) CountByCreatorId(ctx context.Context, creatorId dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableKeyValue) CountByIsHot(ctx context.Context, isHot bool) (count int, err error) {
	w := orm.NewQueryWrapper(ModelKeyValue{})
	w.Eq(TableKeyValue.IsHot, isHot)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableKeyValue.DeletedAt, 0).IsNull(TableKeyValue.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableKeyValue) Create(ctx context.Context, data ...ModelKeyValue) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelKeyValue{})
	for _, item := range data {
		w.Add(item)
	}
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
