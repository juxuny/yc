// Code generated by yc@v0.0.1. DO NOT EDIT.
package db

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/orm"

	cos "github.com/juxuny/yc/services/cos"
)

var TableConfigRecord = tableConfigRecord{
	Id:         orm.FieldName("id"),
	ConfigId:   orm.FieldName("config_id"),
	CreateTime: orm.FieldName("create_time"),
	SeqNo:      orm.FieldName("seq_no"),
	RecordType: orm.FieldName("type"),
}

type ModelConfigRecord struct {
	Id         *dt.ID               `json:"id" orm:"id"`
	ConfigId   *dt.ID               `json:"configId" orm:"config_id"`
	CreateTime int64                `json:"createTime" orm:"create_time"`
	SeqNo      uint64               `json:"seqNo" orm:"seq_no"`
	RecordType cos.ConfigRecordType `json:"recordType" orm:"type"`
}

func (ModelConfigRecord) TableName() string {
	return cos.Name + "_" + "config_record"
}

type ModelConfigRecordList []ModelConfigRecord

func (t ModelConfigRecordList) Filter(f func(index int, item ModelConfigRecord) bool) ModelConfigRecordList {
	ret := make(ModelConfigRecordList, 0)
	for i, item := range t {
		if f(i, item) {
			ret = append(ret, item)
		}
	}
	return ret
}

func (t ModelConfigRecordList) MergeSort(list ModelConfigRecordList, less func(a, b ModelConfigRecord) bool) ModelConfigRecordList {
	ret := make(ModelConfigRecordList, 0)
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

type tableConfigRecord struct {
	suffix          []string
	checkCloneTable bool
	Id              orm.FieldName
	ConfigId        orm.FieldName
	CreateTime      orm.FieldName
	SeqNo           orm.FieldName
	RecordType      orm.FieldName
}

func (t tableConfigRecord) EnableHashTableNameAndCheckClone(suffix ...string) tableConfigRecord {
	return tableConfigRecord{
		suffix:          suffix,
		checkCloneTable: true,
		Id:              t.Id,
		ConfigId:        t.ConfigId,
		CreateTime:      t.CreateTime,
		SeqNo:           t.SeqNo,
		RecordType:      t.RecordType,
	}
}

func (t tableConfigRecord) EnableHash(suffix ...string) tableConfigRecord {
	return tableConfigRecord{
		suffix:          suffix,
		checkCloneTable: false,
		Id:              t.Id,
		ConfigId:        t.ConfigId,
		CreateTime:      t.CreateTime,
		SeqNo:           t.SeqNo,
		RecordType:      t.RecordType,
	}
}

func (t tableConfigRecord) BaseTableName() orm.TableName {
	return cos.Name + "_" + "config_record"
}

func (t tableConfigRecord) TableName() orm.TableName {
	ret := orm.TableName("config_record").Prefix(cos.Name)
	for _, s := range t.suffix {
		ret = ret.Suffix(s)
	}
	return ret
}

func (t tableConfigRecord) checkAndCloneTable(ctx context.Context) error {
	if t.checkCloneTable {
		tableNameList, err := orm.ShowTables(ctx, cos.Name)
		if err != nil {
			log.Error(err)
			return err
		}
		if !tableNameList.Contain(t.BaseTableName()) {
			return errors.SystemError.DatabaseCloneErrorNotFoundTemplate.WithField("tableName", t.BaseTableName().String())
		}
		if !tableNameList.Contain(t.TableName()) {
			w := orm.NewCloneWrapper(t.BaseTableName(), t.TableName())
			_, err := orm.Clone(ctx, cos.Name, w)
			return err
		}
	}
	return nil
}

func (t tableConfigRecord) FindOneById(ctx context.Context, id *dt.ID, orderBy ...orm.Order) (data ModelConfigRecord, found bool, err error) {
	err = t.checkAndCloneTable(ctx)
	if err != nil {
		return
	}
	w := orm.NewQueryWrapper(data).Limit(1).TableName(t.TableName())
	w.Eq(TableConfigRecord.Id, id)
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

func (t tableConfigRecord) FindOneByConfigId(ctx context.Context, configId *dt.ID, orderBy ...orm.Order) (data ModelConfigRecord, found bool, err error) {
	err = t.checkAndCloneTable(ctx)
	if err != nil {
		return
	}
	w := orm.NewQueryWrapper(data).Limit(1).TableName(t.TableName())
	w.Eq(TableConfigRecord.ConfigId, configId)
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

func (t tableConfigRecord) FindOneByRecordType(ctx context.Context, recordType cos.ConfigRecordType, orderBy ...orm.Order) (data ModelConfigRecord, found bool, err error) {
	err = t.checkAndCloneTable(ctx)
	if err != nil {
		return
	}
	w := orm.NewQueryWrapper(data).Limit(1).TableName(t.TableName())
	w.Eq(TableConfigRecord.RecordType, recordType)
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

func (tableConfigRecord) UpdateById(ctx context.Context, id *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.Id, id)
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) UpdateByConfigId(ctx context.Context, configId *dt.ID, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.ConfigId, configId)
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) UpdateByRecordType(ctx context.Context, recordType cos.ConfigRecordType, update orm.H) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.RecordType, recordType)
	w.Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) Update(ctx context.Context, update orm.H, where orm.WhereWrapper) (rowsAffected int64, err error) {
	w := orm.NewUpdateWrapper(ModelConfigRecord{})
	w.SetWhere(where).Updates(update)
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) DeleteById(ctx context.Context, id *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.Id, id)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) DeleteByConfigId(ctx context.Context, configId *dt.ID) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.ConfigId, configId)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) DeleteByRecordType(ctx context.Context, recordType cos.ConfigRecordType) (rowsAffected int64, err error) {
	w := orm.NewDeleteWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.RecordType, recordType)
	result, err := orm.Delete(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) Find(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Nested(where)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) FindOne(ctx context.Context, where orm.WhereWrapper, orderBy ...orm.Order) (ret ModelConfigRecord, found bool, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
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

func (tableConfigRecord) FindById(ctx context.Context, id *dt.ID, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.Id, id)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) FindByConfigId(ctx context.Context, configId *dt.ID, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.ConfigId, configId)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) FindByRecordType(ctx context.Context, recordType cos.ConfigRecordType, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.RecordType, recordType)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) Page(ctx context.Context, pageNum, pageSize int64, where orm.WhereWrapper, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.SetWhere(where).Offset((pageNum - 1) * pageSize).Limit(pageSize)
	w.Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		return nil, err
	}
	return
}

func (tableConfigRecord) PageById(ctx context.Context, pageNum, pageSize int64, id *dt.ID, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.Id, id)
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) PageByConfigId(ctx context.Context, pageNum, pageSize int64, configId *dt.ID, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.ConfigId, configId)
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) PageByRecordType(ctx context.Context, pageNum, pageSize int64, recordType cos.ConfigRecordType, orderBy ...orm.Order) (list ModelConfigRecordList, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.RecordType, recordType)
	w.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order(orderBy...)
	err = orm.Select(ctx, cos.Name, w, &list)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return
}

func (tableConfigRecord) Count(ctx context.Context, where orm.WhereWrapper) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.SetWhere(where)
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfigRecord) CountById(ctx context.Context, id *dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.Id, id)
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfigRecord) CountByConfigId(ctx context.Context, configId *dt.ID) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.ConfigId, configId)
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfigRecord) CountByRecordType(ctx context.Context, recordType cos.ConfigRecordType) (count int64, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Eq(TableConfigRecord.RecordType, recordType)
	w.Select("COUNT(*)")
	err = orm.Select(ctx, cos.Name, w, &count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (tableConfigRecord) Create(ctx context.Context, data ...ModelConfigRecord) (rowsAffected int64, err error) {
	w := orm.NewInsertWrapper(ModelConfigRecord{})
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

func (tableConfigRecord) CreateWithLastId(ctx context.Context, data ModelConfigRecord) (lastInsertId *dt.ID, err error) {
	w := orm.NewInsertWrapper(ModelConfigRecord{})
	w.Add(data)
	result, err := orm.Insert(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return dt.InvalidIDPointer(), err
	}
	if id, err := result.LastInsertId(); err != nil {
		return dt.InvalidIDPointer(), err
	} else {
		return dt.NewIDPointer(uint64(id)), nil
	}
}

func (tableConfigRecord) UpdateAdvance(ctx context.Context, update orm.UpdateWrapper) (rowsAffected int64, err error) {
	w := update
	result, err := orm.Update(ctx, cos.Name, w)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

func (tableConfigRecord) SumInt64(ctx context.Context, field orm.FieldName, where orm.WhereWrapper) (sum int64, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Select("SUM(" + field.Wrap() + ")")
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &sum)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return sum, err
}

func (tableConfigRecord) SumFloat64(ctx context.Context, field orm.FieldName, where orm.WhereWrapper) (sum float64, err error) {
	w := orm.NewQueryWrapper(ModelConfigRecord{})
	w.Select("SUM(" + field.Wrap() + ")")
	w.SetWhere(where)
	err = orm.Select(ctx, cos.Name, w, &sum)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	return sum, err
}
