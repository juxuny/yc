package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
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
	Id         *dt.ID `json:"id" orm:"id"`
	Namespace  string `json:"namespace" orm:"namespace"`
	CreateTime int64  `json:"createTime" orm:"create_time"`
	UpdateTime int64  `json:"updateTime" orm:"update_time"`
	DeletedAt  int64  `json:"deletedAt" orm:"deleted_at"`
	IsDisabled bool   `json:"isDisabled" orm:"is_disabled"`
	CreatorId  *dt.ID `json:"creatorId" orm:"creator_id"`
}

func (ModelNamespace) TableName() string {
	return cos.Name + "_" + "namespace"
}

func (t ModelNamespace) ToNamespaceResp() cos.NamespaceResp {
	return cos.NamespaceResp{
		Id:         t.Id,
		Namespace:  t.Namespace,
		CreateTime: t.CreateTime,
		UpdateTime: t.UpdateTime,
		IsDisabled: t.IsDisabled,
		CreatorId:  t.CreatorId,
	}
}
func (t ModelNamespace) ToListNamespaceItem() cos.ListNamespaceItem {
	return cos.ListNamespaceItem{
		Id:         t.Id,
		Namespace:  t.Namespace,
		CreateTime: t.CreateTime,
		UpdateTime: t.UpdateTime,
		IsDisabled: t.IsDisabled,
	}
}

func (t ModelNamespace) ToNamespaceRespAsPointer() *cos.NamespaceResp {
	ret := t.ToNamespaceResp()
	return &ret
}
func (t ModelNamespace) ToListNamespaceItemAsPointer() *cos.ListNamespaceItem {
	ret := t.ToListNamespaceItem()
	return &ret
}

type ModelNamespaceList []ModelNamespace

func (t ModelNamespaceList) MapToNamespaceRespList() []*cos.NamespaceResp {
	ret := make([]*cos.NamespaceResp, 0)
	for _, item := range t {
		ret = append(ret, item.ToNamespaceRespAsPointer())
	}
	return ret
}
func (t ModelNamespaceList) MapToListNamespaceItemList() []*cos.ListNamespaceItem {
	ret := make([]*cos.ListNamespaceItem, 0)
	for _, item := range t {
		ret = append(ret, item.ToListNamespaceItemAsPointer())
	}
	return ret
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

func (tableNamespace) FindOneById(ctx context.Context, id dt.ID, orderBy ...orm.Order) (data ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
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

func (tableNamespace) FindOneByNamespace(ctx context.Context, namespace string, orderBy ...orm.Order) (data ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
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

func (tableNamespace) FindOneByCreatorId(ctx context.Context, creatorId dt.ID, orderBy ...orm.Order) (data ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
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

func (tableNamespace) UpdateById(ctx context.Context, id dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
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
		log.Error(err)
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
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) DeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) DeleteByNamespace(ctx context.Context, namespace string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) DeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
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
		log.Error(err)
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
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) SoftDelete(ctx context.Context, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.SetValue(TableNamespace.DeletedAt, orm.Now())
	w.SetWhere(where)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) Find(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Nested(where)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) FindOne(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (ret ModelNamespace, found bool, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
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

func (tableNamespace) FindById(ctx context.Context, id dt.ID, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) FindByNamespace(ctx context.Context, namespace string, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) FindByCreatorId(ctx context.Context, creatorId dt.ID, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) Page(ctx context.Context, pageNum, pageSize int64, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableNamespace) PageById(ctx context.Context, pageNum, pageSize int64, id dt.ID, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) PageByNamespace(ctx context.Context, pageNum, pageSize int64, namespace string, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) PageByCreatorId(ctx context.Context, pageNum, pageSize int64, creatorId dt.ID, orderBy ...orm.Order) (list ModelNamespaceList, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableNamespace) Count(ctx context.Context, where orm.WhereWrapper) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableNamespace) CountById(ctx context.Context, id dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableNamespace) CountByNamespace(ctx context.Context, namespace string) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.Namespace, namespace)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableNamespace) CountByCreatorId(ctx context.Context, creatorId dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Eq(TableNamespace.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableNamespace) Create(ctx context.Context, data ...ModelNamespace) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelNamespace{})
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

func (tableNamespace) CreateWithLastId(ctx context.Context, data ModelNamespace) (lastInsertId dt.ID, err error) {
	w := orm.NewInsertWrapper(ModelNamespace{})
	w.Add(data)
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return dt.InvalidID(), err
	}
	if id, err := result.LastInsertId(); err != nil {
		return dt.InvalidID(), err
	} else {
		return dt.NewID(uint64(id)), nil
	}
}

func (tableNamespace) ResetDeletedAt(ctx context.Context, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelNamespace{})
	w.SetWhere(where)
	w.SetValue(TableConfig.DeletedAt, 0)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) UpdateAdvance(ctx context.Context, update orm.UpdateWrapper) (rowsAffected int64, err error) {
	w := update
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableNamespace) SumInt64(ctx context.Context, field orm.FieldName, where orm.WhereWrapper) (sum int64, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Select("SUM(" + field.Wrap() + ")")
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &sum)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return sum, err
}

func (tableNamespace) SumFloat64(ctx context.Context, field orm.FieldName, where orm.WhereWrapper) (sum float64, err error) {
	w := orm.NewQueryWrapper(ModelNamespace{})
	w.Select("SUM(" + field.Wrap() + ")")
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableNamespace.DeletedAt, 0).IsNull(TableNamespace.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &sum)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return sum, err
}
