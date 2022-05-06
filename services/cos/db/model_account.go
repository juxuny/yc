package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/orm"
	cos "github.com/juxuny/yc/services/cos"
)

var TableAccount = tableAccount{
	Id:          orm.FieldName("id"),
	Identifier:  orm.FieldName("identifier"),
	Credential:  orm.FieldName("credential"),
	AccountType: orm.FieldName("account_type"),
	CreateTime:  orm.FieldName("create_time"),
	UpdateTime:  orm.FieldName("update_time"),
	DeletedAt:   orm.FieldName("deleted_at"),
	IsDisabled:  orm.FieldName("is_disabled"),
	CreatorId:   orm.FieldName("creator_id"),
}

type ModelAccount struct {
	Id          dt.ID           `json:"id" orm:"id"`
	Identifier  string          `json:"identifier" orm:"identifier"`
	Credential  string          `json:"credential" orm:"credential"`
	AccountType cos.AccountType `json:"accountType" orm:"account_type"`
	CreateTime  int64           `json:"createTime" orm:"create_time"`
	UpdateTime  int64           `json:"updateTime" orm:"update_time"`
	DeletedAt   int64           `json:"deletedAt" orm:"deleted_at"`
	IsDisabled  bool            `json:"isDisabled" orm:"is_disabled"`
	CreatorId   dt.ID           `json:"creatorId" orm:"creator_id"`
}

func (ModelAccount) TableName() string {
	return cos.Name + "_" + "account"
}

type tableAccount struct {
	Id          orm.FieldName
	Identifier  orm.FieldName
	Credential  orm.FieldName
	AccountType orm.FieldName
	CreateTime  orm.FieldName
	UpdateTime  orm.FieldName
	DeletedAt   orm.FieldName
	IsDisabled  orm.FieldName
	CreatorId   orm.FieldName
}

func (tableAccount) TableName() string {
	return cos.Name + "_" + "account"
}

func (tableAccount) FindOneById(ctx context.Context, id dt.ID) (data ModelAccount, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccount.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableAccount) FindOneByIdentifier(ctx context.Context, identifier string) (data ModelAccount, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccount.Identifier, identifier)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableAccount) FindOneByAccountType(ctx context.Context, accountType cos.AccountType) (data ModelAccount, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccount.AccountType, accountType)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableAccount) FindOneByCreatorId(ctx context.Context, creatorId dt.ID) (data ModelAccount, found bool, err error) {
	w := orm.NewQueryWrapper(data).Limit(1)
	w.Eq(TableAccount.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &data)
	if err != nil {
		if e, ok := err.(errors.Error); ok && e.Code == errors.SystemError.DatabaseNoData.Code {
			return data, false, nil
		}
		return data, false, err
	}
	return data, true, nil
}

func (tableAccount) UpdateById(ctx context.Context, id dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.Eq(TableAccount.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) UpdateByIdentifier(ctx context.Context, identifier string, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.Eq(TableAccount.Identifier, identifier)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) UpdateByAccountType(ctx context.Context, accountType cos.AccountType, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.Eq(TableAccount.AccountType, accountType)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) UpdateByCreatorId(ctx context.Context, creatorId dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.Eq(TableAccount.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.SetWhere(where).Updates(update)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) DeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccount{})
	w.Eq(TableAccount.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) DeleteByIdentifier(ctx context.Context, identifier string) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccount{})
	w.Eq(TableAccount.Identifier, identifier)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) DeleteByAccountType(ctx context.Context, accountType cos.AccountType) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccount{})
	w.Eq(TableAccount.AccountType, accountType)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) DeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelAccount{})
	w.Eq(TableAccount.CreatorId, creatorId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) SoftDeleteById(ctx context.Context, id dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.SetValue(TableAccount.DeletedAt, orm.Now())
	w.Eq(TableAccount.Id, id)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) SoftDeleteByIdentifier(ctx context.Context, identifier string) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.SetValue(TableAccount.DeletedAt, orm.Now())
	w.Eq(TableAccount.Identifier, identifier)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) SoftDeleteByAccountType(ctx context.Context, accountType cos.AccountType) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.SetValue(TableAccount.DeletedAt, orm.Now())
	w.Eq(TableAccount.AccountType, accountType)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) SoftDeleteByCreatorId(ctx context.Context, creatorId dt.ID) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelAccount{})
	w.SetValue(TableAccount.DeletedAt, orm.Now())
	w.Eq(TableAccount.CreatorId, creatorId)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (tableAccount) Find(ctx context.Context, where orm.WhereWrapper) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) FindById(ctx context.Context, id dt.ID) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) FindByIdentifier(ctx context.Context, identifier string) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.Identifier, identifier)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) FindByAccountType(ctx context.Context, accountType cos.AccountType) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.AccountType, accountType)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) FindByCreatorId(ctx context.Context, creatorId dt.ID) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) Page(ctx context.Context, pageNum, pageSize int, where orm.WhereWrapper) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) PageById(ctx context.Context, pageNum, pageSize int, id dt.ID) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) PageByIdentifier(ctx context.Context, pageNum, pageSize int, identifier string) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.Identifier, identifier)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) PageByAccountType(ctx context.Context, pageNum, pageSize int, accountType cos.AccountType) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.AccountType, accountType)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) PageByCreatorId(ctx context.Context, pageNum, pageSize int, creatorId dt.ID) (list []ModelAccount, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableAccount) Count(ctx context.Context, where orm.WhereWrapper) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.SetWhere(where)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableAccount) CountById(ctx context.Context, id dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.Id, id)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableAccount) CountByIdentifier(ctx context.Context, identifier string) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.Identifier, identifier)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableAccount) CountByAccountType(ctx context.Context, accountType cos.AccountType) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.AccountType, accountType)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableAccount) CountByCreatorId(ctx context.Context, creatorId dt.ID) (count int, err error) {
	w := orm.NewQueryWrapper(ModelAccount{})
	w.Eq(TableAccount.CreatorId, creatorId)
	w.Nested(orm.NewOrWhereWrapper().Eq(TableAccount.DeletedAt, 0).IsNull(TableAccount.DeletedAt))
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	return count, err
}

func (tableAccount) Create(ctx context.Context, data ...ModelAccount) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelAccount{})
	for _, item := range data {
		w.Add(item)
	}
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
