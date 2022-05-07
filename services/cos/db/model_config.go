package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
)

var TableConfig = tableConfig{
	Id:         orm.FieldName("id"),
	CreateTime: orm.FieldName("create_time"),
	UpdateTime: orm.FieldName("update_time"),
	DeletedAt:  orm.FieldName("deleted_at"),
	ConfigId:   orm.FieldName("config_id"),
	IsDisabled: orm.FieldName("is_disabled"),
	CreatorId:  orm.FieldName("creator_id"),
	BaseId:     orm.FieldName("base_id"),
}

type ModelConfig struct {
	Id         *dt.ID `json:"id" orm:"id"`
	CreateTime int64  `json:"createTime" orm:"create_time"`
	UpdateTime int64  `json:"updateTime" orm:"update_time"`
	DeletedAt  int64  `json:"deletedAt" orm:"deleted_at"`
	ConfigId   string `json:"configId" orm:"config_id"`
	IsDisabled bool   `json:"isDisabled" orm:"is_disabled"`
	CreatorId  *dt.ID `json:"creatorId" orm:"creator_id"`
	BaseId     *dt.ID `json:"baseId" orm:"base_id"`
}

func (ModelConfig) TableName() string {
	return cos.Name + "_" + "config"
}

type tableConfig struct {
	Id         orm.FieldName
	CreateTime orm.FieldName
	UpdateTime orm.FieldName
	DeletedAt  orm.FieldName
	ConfigId   orm.FieldName
	IsDisabled orm.FieldName
	CreatorId  orm.FieldName
	BaseId     orm.FieldName
}

func (tableConfig) TableName() string {
	return cos.Name + "_" + "config"
}

func (tableConfig) FindOneById(ctx context.Context, id *dt.ID) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByConfigId(ctx context.Context, configId string) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByCreatorId(ctx context.Context, creatorId *dt.ID) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByBaseId(ctx context.Context, baseId *dt.ID) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) UpdateById(ctx context.Context, id *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) UpdateByConfigId(ctx context.Context, configId string, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) UpdateByCreatorId(ctx context.Context, creatorId *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) UpdateByBaseId(ctx context.Context, baseId *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetWhere(where).Updates(update)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteById(ctx context.Context, id *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByConfigId(ctx context.Context, configId string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByCreatorId(ctx context.Context, creatorId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByBaseId(ctx context.Context, baseId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteById(ctx context.Context, id *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.Id, id)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteByConfigId(ctx context.Context, configId string) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.ConfigId, configId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteByCreatorId(ctx context.Context, creatorId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.CreatorId, creatorId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteByBaseId(ctx context.Context, baseId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.BaseId, baseId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) Find(ctx context.Context, where orm.WhereWrapper) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) FindById(ctx context.Context, id *dt.ID) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) FindByConfigId(ctx context.Context, configId string) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) FindByCreatorId(ctx context.Context, creatorId *dt.ID) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) FindByBaseId(ctx context.Context, baseId *dt.ID) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) Page(ctx context.Context, pageNum, pageSize int, where orm.WhereWrapper) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) PageById(ctx context.Context, pageNum, pageSize int, id *dt.ID) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) PageByConfigId(ctx context.Context, pageNum, pageSize int, configId string) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) PageByCreatorId(ctx context.Context, pageNum, pageSize int, creatorId *dt.ID) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) PageByBaseId(ctx context.Context, pageNum, pageSize int, baseId *dt.ID) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) Count(ctx context.Context, where orm.WhereWrapper) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableConfig) CountById(ctx context.Context, id *dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableConfig) CountByConfigId(ctx context.Context, configId string) (count int, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableConfig) CountByCreatorId(ctx context.Context, creatorId *dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableConfig) CountByBaseId(ctx context.Context, baseId *dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableConfig) Create(ctx context.Context, data ...ModelConfig) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelConfig{})
	for _, item := range data {
		w.Add(item)
	}
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
