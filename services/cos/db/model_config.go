package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"

	cos "github.com/juxuny/yc/services/cos"
)

var TableConfig = tableConfig{
	Id:          orm.FieldName("id"),
	CreateTime:  orm.FieldName("create_time"),
	UpdateTime:  orm.FieldName("update_time"),
	DeletedAt:   orm.FieldName("deleted_at"),
	ConfigId:    orm.FieldName("config_id"),
	IsDisabled:  orm.FieldName("is_disabled"),
	CreatorId:   orm.FieldName("creator_id"),
	BaseId:      orm.FieldName("base_id"),
	NamespaceId: orm.FieldName("namespace_id"),
}

type ModelConfig struct {
	Id          *dt.ID `json:"id" orm:"id"`
	CreateTime  int64  `json:"createTime" orm:"create_time"`
	UpdateTime  int64  `json:"updateTime" orm:"update_time"`
	DeletedAt   int64  `json:"deletedAt" orm:"deleted_at"`
	ConfigId    string `json:"configId" orm:"config_id"`
	IsDisabled  bool   `json:"isDisabled" orm:"is_disabled"`
	CreatorId   *dt.ID `json:"creatorId" orm:"creator_id"`
	BaseId      *dt.ID `json:"baseId" orm:"base_id"`
	NamespaceId *dt.ID `json:"namespaceId" orm:"namespace_id"`
}

func (ModelConfig) TableName() string {
	return cos.Name + "_" + "config"
}

func (t ModelConfig) ToListConfigItem() cos.ListConfigItem {
	return cos.ListConfigItem{
		Id:          t.Id,
		CreateTime:  t.CreateTime,
		UpdateTime:  t.UpdateTime,
		BaseId:      t.BaseId,
		NamespaceId: t.NamespaceId,
		ConfigId:    t.ConfigId,
	}
}

func (t ModelConfig) ToListConfigItemAsPointer() *cos.ListConfigItem {
	ret := t.ToListConfigItem()
	return &ret
}

type ModelConfigList []ModelConfig

func (t ModelConfigList) MapToListConfigItemList() []*cos.ListConfigItem {
	ret := make([]*cos.ListConfigItem, 0)
	for _, item := range t {
		ret = append(ret, item.ToListConfigItemAsPointer())
	}
	return ret
}

type tableConfig struct {
	Id          orm.FieldName
	CreateTime  orm.FieldName
	UpdateTime  orm.FieldName
	DeletedAt   orm.FieldName
	ConfigId    orm.FieldName
	IsDisabled  orm.FieldName
	CreatorId   orm.FieldName
	BaseId      orm.FieldName
	NamespaceId orm.FieldName
}

func (tableConfig) TableName() string {
	return cos.Name + "_" + "config"
}

func (tableConfig) FindOneById(ctx context.Context, id dt.ID, orderBy ...orm.Order) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		log.Error(err)
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByConfigId(ctx context.Context, configId string, orderBy ...orm.Order) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		log.Error(err)
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByCreatorId(ctx context.Context, creatorId dt.ID, orderBy ...orm.Order) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		log.Error(err)
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByBaseId(ctx context.Context, baseId dt.ID, orderBy ...orm.Order) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		log.Error(err)
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) FindOneByNamespaceId(ctx context.Context, namespaceId dt.ID, orderBy ...orm.Order) (data ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableConfig.NamespaceId, namespaceId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		log.Error(err)
		return data, false, err
	}
	return data, true, nil
}

func (tableConfig) UpdateById(ctx context.Context, id dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) UpdateByCreatorId(ctx context.Context, creatorId dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) UpdateByBaseId(ctx context.Context, baseId dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) UpdateByNamespaceId(ctx context.Context, namespaceId dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.Eq(TableConfig.NamespaceId, namespaceId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByConfigId(ctx context.Context, configId string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByBaseId(ctx context.Context, baseId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) DeleteByNamespaceId(ctx context.Context, namespaceId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfig{})
	w.Eq(TableConfig.NamespaceId, namespaceId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.Id, id)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.CreatorId, creatorId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteByBaseId(ctx context.Context, baseId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.BaseId, baseId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDeleteByNamespaceId(ctx context.Context, namespaceId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.Eq(TableConfig.NamespaceId, namespaceId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) SoftDelete(ctx context.Context, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetValue(TableConfig.DeletedAt, orm.Now())
	w.SetWhere(where)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) Find(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.SetWhere(where).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) FindOne(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (ret ModelConfig, found bool, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.SetWhere(where).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &ret)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return ret, false, nil
		}
		log.Error(err)
		return ret, false, err
	}
	return ret, true, nil
}

func (tableConfig) FindById(ctx context.Context, id dt.ID, orderBy ...orm.Order) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) FindByConfigId(ctx context.Context, configId string, orderBy ...orm.Order) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) FindByCreatorId(ctx context.Context, creatorId dt.ID, orderBy ...orm.Order) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) FindByBaseId(ctx context.Context, baseId dt.ID, orderBy ...orm.Order) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) FindByNamespaceId(ctx context.Context, namespaceId dt.ID, orderBy ...orm.Order) (list []ModelConfig, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.NamespaceId, namespaceId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) Page(ctx context.Context, pageNum, pageSize int64, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelConfigList, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfig) PageById(ctx context.Context, pageNum, pageSize int64, id dt.ID, orderBy ...orm.Order) (list ModelConfigList, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) PageByConfigId(ctx context.Context, pageNum, pageSize int64, configId string, orderBy ...orm.Order) (list ModelConfigList, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) PageByCreatorId(ctx context.Context, pageNum, pageSize int64, creatorId dt.ID, orderBy ...orm.Order) (list ModelConfigList, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) PageByBaseId(ctx context.Context, pageNum, pageSize int64, baseId dt.ID, orderBy ...orm.Order) (list ModelConfigList, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) PageByNamespaceId(ctx context.Context, pageNum, pageSize int64, namespaceId dt.ID, orderBy ...orm.Order) (list ModelConfigList, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.NamespaceId, namespaceId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfig) Count(ctx context.Context, where orm.WhereWrapper) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfig) CountById(ctx context.Context, id dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfig) CountByConfigId(ctx context.Context, configId string) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.ConfigId, configId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfig) CountByCreatorId(ctx context.Context, creatorId dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfig) CountByBaseId(ctx context.Context, baseId dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.BaseId, baseId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfig) CountByNamespaceId(ctx context.Context, namespaceId dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfig{})
	w.Eq(TableConfig.NamespaceId, namespaceId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableConfig.DeletedAt, 0).IsNull(TableConfig.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfig) Create(ctx context.Context, data ...ModelConfig) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelConfig{})
	for _, item := range data {
		w.Add(item)
	}
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfig) ResetDeletedAt(ctx context.Context, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfig{})
	w.SetWhere(where)
	w.SetValue(TableConfig.DeletedAt, 0)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}
