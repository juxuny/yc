package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
)

var TableNamespace = tableNamespace{
	Id:         orm.FieldName("id"),
	Namespace:  orm.FieldName("namespace"),
	CreateTime: orm.FieldName("create_time"),
	UpdateTime: orm.FieldName("update_time"),
	DeletedAt:  orm.FieldName("deleted_at"),
	IsDisabled: orm.FieldName("is_disabled"),
	CreatorId:  orm.FieldName("creator_id"),
}

type ModelNamespace struct {
	Id         dt.ID  `json:"id" orm:"id"`
	Namespace  string `json:"namespace" orm:"namespace"`
	CreateTime int64  `json:"createTime" orm:"create_time"`
	UpdateTime int64  `json:"updateTime" orm:"update_time"`
	DeletedAt  int64  `json:"deletedAt" orm:"deleted_at"`
	IsDisabled bool   `json:"isDisabled" orm:"is_disabled"`
	CreatorId  dt.ID  `json:"creatorId" orm:"creator_id"`
}

func (ModelNamespace) TableName() string {
	return cos.Name + "_" + "namespace"
}

type tableNamespace struct {
	Id         orm.FieldName
	Namespace  orm.FieldName
	CreateTime orm.FieldName
	UpdateTime orm.FieldName
	DeletedAt  orm.FieldName
	IsDisabled orm.FieldName
	CreatorId  orm.FieldName
}

func (tableNamespace) TableName() string {
	return cos.Name + "_" + "namespace"
}

func (tableNamespace) FindOneById(ctx context.Context, id dt.ID) (data ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableNamespace) FindOneByNamespace(ctx context.Context, namespace string) (data ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableNamespace) FindOneByCreatorId(ctx context.Context, creatorId dt.ID) (data ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableNamespace) UpdateById(ctx context.Context, id dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) UpdateByNamespace(ctx context.Context, namespace string, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) UpdateByCreatorId(ctx context.Context, creatorId dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.SetWhere(where).Updates(update)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) DeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) DeleteByNamespace(ctx context.Context, namespace string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) DeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) SoftDeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.SetValue(TableNamespace.DeletedAt, orm.Now())
	w.Eq(TableNamespace.Id, id)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) SoftDeleteByNamespace(ctx context.Context, namespace string) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.SetValue(TableNamespace.DeletedAt, orm.Now())
	w.Eq(TableNamespace.Namespace, namespace)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) SoftDeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.SetValue(TableNamespace.DeletedAt, orm.Now())
	w.Eq(TableNamespace.CreatorId, creatorId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) Find(ctx context.Context, where orm.WhereWrapper) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) FindById(ctx context.Context, id dt.ID) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) FindByNamespace(ctx context.Context, namespace string) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) FindByCreatorId(ctx context.Context, creatorId dt.ID) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) Page(ctx context.Context, pageNum, pageSize int, where orm.WhereWrapper) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) PageById(ctx context.Context, pageNum, pageSize int, id dt.ID) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) PageByNamespace(ctx context.Context, pageNum, pageSize int, namespace string) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) PageByCreatorId(ctx context.Context, pageNum, pageSize int, creatorId dt.ID) (list []ModelNamespace, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) Count(ctx context.Context, where orm.WhereWrapper) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableNamespace) CountById(ctx context.Context, id dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableNamespace) CountByNamespace(ctx context.Context, namespace string) (count int, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableNamespace) CountByCreatorId(ctx context.Context, creatorId dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableNamespace) Create(ctx context.Context, data ...ModelNamespace) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelNamespace{})
	for _, item := range data {
		w.Add(item)
	}
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}