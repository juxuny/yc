package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"

	cos "github.com/juxuny/yc/services/cos"
)

var TableAccessKey = tableAccessKey{
	Id:             orm.FieldName("id"),
	CreateTime:     orm.FieldName("create_time"),
	UpdateTime:     orm.FieldName("update_time"),
	IsDisabled:     orm.FieldName("is_disabled"),
	DeletedAt:      orm.FieldName("deleted_at"),
	UserId:         orm.FieldName("user_id"),
	AccessKey:      orm.FieldName("access_key"),
	HasValidity:    orm.FieldName("has_validity"),
	ValidStartTime: orm.FieldName("valid_start_time"),
	ValidEndTime:   orm.FieldName("valid_end_time"),
	Remark:         orm.FieldName("remark"),
}

type ModelAccessKey struct {
	Id             *dt.ID `json:"id" orm:"id"`
	CreateTime     int64  `json:"createTime" orm:"create_time"`
	UpdateTime     int64  `json:"updateTime" orm:"update_time"`
	IsDisabled     bool   `json:"isDisabled" orm:"is_disabled"`
	DeletedAt      int64  `json:"deletedAt" orm:"deleted_at"`
	UserId         *dt.ID `json:"userId" orm:"user_id"`
	AccessKey      string `json:"accessKey" orm:"access_key"`
	HasValidity    bool   `json:"hasValidity" orm:"has_validity"`
	ValidStartTime int64  `json:"validStartTime" orm:"valid_start_time"`
	ValidEndTime   int64  `json:"validEndTime" orm:"valid_end_time"`
	Remark         string `json:"remark" orm:"remark"`
}

func (ModelAccessKey) TableName() string {
	return cos.Name + "_" + "access_key"
}

func (t ModelAccessKey) ToAccessKeyItem() cos.AccessKeyItem {
	return cos.AccessKeyItem{
		Id:             t.Id,
		CreateTime:     t.CreateTime,
		UpdateTime:     t.UpdateTime,
		IsDisabled:     t.IsDisabled,
		DeletedAt:      t.DeletedAt,
		UserId:         t.UserId,
		AccessKey:      t.AccessKey,
		HasValidity:    t.HasValidity,
		ValidStartTime: t.ValidStartTime,
		ValidEndTime:   t.ValidEndTime,
		Remark:         t.Remark,
	}
}

func (t ModelAccessKey) ToAccessKeyItemAsPointer() *cos.AccessKeyItem {
	ret := t.ToAccessKeyItem()
	return &ret
}

type ModelAccessKeyList []ModelAccessKey

func (t ModelAccessKeyList) Filter(f func(index int, item ModelAccessKey) bool) ModelAccessKeyList {
	ret := make(ModelAccessKeyList, 0)
	for i, item := range t {
		if f(i, item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func (t ModelAccessKeyList) MergeSort(list ModelAccessKeyList, less func(a, b ModelAccessKey) bool) ModelAccessKeyList {
	ret := make(ModelAccessKeyList, 0)
	i, j := 0, 0
	for i < len(t) || j < len(list) {
		if i < len(t) && j < len(list) {
			if less(t[i], list[j]) {
				ret = append(ret, t[i])
				i += 1
			} else {
				ret = append(ret, list[j])
				j += 1
			}
			continue
		} else if i < len(t) {
			ret = append(ret, t[i])
			i += 1
		} else if j < len(list) {
			ret = append(ret, list[j])
			j += 1
		}
	}
	return ret
}

func (t ModelAccessKeyList) MapToAccessKeyItemList() []*cos.AccessKeyItem {
	ret := make([]*cos.AccessKeyItem, 0)
	for _, item := range t {
		ret = append(ret, item.ToAccessKeyItemAsPointer())
	}
	return ret
}

type tableAccessKey struct {
	Id             orm.FieldName
	CreateTime     orm.FieldName
	UpdateTime     orm.FieldName
	IsDisabled     orm.FieldName
	DeletedAt      orm.FieldName
	UserId         orm.FieldName
	AccessKey      orm.FieldName
	HasValidity    orm.FieldName
	ValidStartTime orm.FieldName
	ValidEndTime   orm.FieldName
	Remark         orm.FieldName
}

func (tableAccessKey) TableName() string {
	return cos.Name + "_" + "access_key"
}

func (tableAccessKey) FindOneById(ctx context.Context, id *dt.ID, orderBy ...orm.Order) (data ModelAccessKey, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccessKey.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
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

func (tableAccessKey) FindOneByUserId(ctx context.Context, userId *dt.ID, orderBy ...orm.Order) (data ModelAccessKey, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccessKey.UserId, userId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
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

func (tableAccessKey) FindOneByAccessKey(ctx context.Context, accessKey string, orderBy ...orm.Order) (data ModelAccessKey, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccessKey.AccessKey, accessKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
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

func (tableAccessKey) UpdateById(ctx context.Context, id *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) UpdateByUserId(ctx context.Context, userId *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.UserId, userId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) UpdateByAccessKey(ctx context.Context, accessKey string, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.AccessKey, accessKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.SetWhere(where).Updates(update)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) DeleteById(ctx context.Context, id *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) DeleteByUserId(ctx context.Context, userId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.UserId, userId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) DeleteByAccessKey(ctx context.Context, accessKey string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.AccessKey, accessKey)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) SoftDeleteById(ctx context.Context, id *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.SetValue(TableAccessKey.DeletedAt, orm.Now())
	w.Eq(TableAccessKey.Id, id)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) SoftDeleteByUserId(ctx context.Context, userId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.SetValue(TableAccessKey.DeletedAt, orm.Now())
	w.Eq(TableAccessKey.UserId, userId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) SoftDeleteByAccessKey(ctx context.Context, accessKey string) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.SetValue(TableAccessKey.DeletedAt, orm.Now())
	w.Eq(TableAccessKey.AccessKey, accessKey)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) SoftDelete(ctx context.Context, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.SetValue(TableAccessKey.DeletedAt, orm.Now())
	w.SetWhere(where)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) Find(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Nested(where)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) FindOne(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (ret ModelAccessKey, found bool, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.SetWhere(where).Order(orderBy...)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
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

func (tableAccessKey) FindById(ctx context.Context, id *dt.ID, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) FindByUserId(ctx context.Context, userId *dt.ID, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.UserId, userId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) FindByAccessKey(ctx context.Context, accessKey string, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.AccessKey, accessKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) Page(ctx context.Context, pageNum, pageSize int64, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccessKey) PageById(ctx context.Context, pageNum, pageSize int64, id *dt.ID, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) PageByUserId(ctx context.Context, pageNum, pageSize int64, userId *dt.ID, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.UserId, userId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) PageByAccessKey(ctx context.Context, pageNum, pageSize int64, accessKey string, orderBy ...orm.Order) (list ModelAccessKeyList, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.AccessKey, accessKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableAccessKey) Count(ctx context.Context, where orm.WhereWrapper) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableAccessKey) CountById(ctx context.Context, id *dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableAccessKey) CountByUserId(ctx context.Context, userId *dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.UserId, userId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableAccessKey) CountByAccessKey(ctx context.Context, accessKey string) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Eq(TableAccessKey.AccessKey, accessKey)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableAccessKey) Create(ctx context.Context, data ...ModelAccessKey) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelAccessKey{})
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

func (tableAccessKey) CreateWithLastId(ctx context.Context, data ModelAccessKey) (lastInsertId dt.ID, err error) {
	w := orm.NewInsertWrapper(ModelAccessKey{})
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

func (tableAccessKey) ResetDeletedAt(ctx context.Context, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccessKey{})
	w.SetWhere(where)
	w.SetValue(TableConfig.DeletedAt, 0)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) UpdateAdvance(ctx context.Context, update orm.UpdateWrapper) (rowsAffected int64, err error) {
	w := update
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccessKey) SumInt64(ctx context.Context, field orm.FieldName, where orm.WhereWrapper) (sum int64, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Select("SUM(" + field.Wrap() + ")")
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &sum)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return sum, err
}

func (tableAccessKey) SumFloat64(ctx context.Context, field orm.FieldName, where orm.WhereWrapper) (sum float64, err error) {
	w := orm.NewQueryWrapper(ModelAccessKey{})
	w.Select("SUM(" + field.Wrap() + ")")
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccessKey.DeletedAt, 0).IsNull(TableAccessKey.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &sum)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return sum, err
}
