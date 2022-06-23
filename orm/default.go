package orm

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/juxuny/yc/errors"
	"github.com/juxuny/yc/log"
	"reflect"
)

type execResult struct {
	rowsAffected int64
	lastInsertId int64
}

func (t execResult) RowsAffected() (int64, error) {
	return t.rowsAffected, nil
}

func (t execResult) LastInsertId() (int64, error) {
	return t.lastInsertId, nil
}

//func Exec(ctx context.Context, configName string, statement string, values ...interface{}) (result sql.Result, err error) {
//	return connectManagerInstance.Exec(ctx, configName, statement, values...)
//}

func QueryRows(ctx context.Context, configName string, statement string, values ...interface{}) (DataSet, error) {
	return connectManagerInstance.QueryRows(ctx, configName, statement, values...)
}

func QueryScan(ctx context.Context, configName string, out interface{}, statement string, values ...interface{}) error {
	ds, err := connectManagerInstance.QueryRows(ctx, configName, statement, values...)
	if err != nil {
		return err
	}
	if len(ds) > 0 && len(ds[0]) == 12 {
		log.Debug(ds[0][12].Elem().Convert(reflect.TypeOf("")).String())
	}
	return ds.Reform(out)
}

func Select(ctx context.Context, configName string, w QueryWrapper, out interface{}) error {
	statement, values, err := w.Build()
	if err != nil {
		return err
	}
	ds, err := connectManagerInstance.QueryRows(ctx, configName, statement, values...)
	if err != nil {
		return err
	}
	return ds.Reform(out)
}

func Insert(ctx context.Context, configName string, w InsertWrapper) (result sql.Result, err error) {
	statement, values, next, err := w.Build()
	if err != nil {
		return result, err
	}
	result, err = connectManagerInstance.Exec(ctx, configName, statement, values...)
	if err != nil {
		return result, errors.SystemError.DatabaseExecError.Wrap(err)
	}
	totalResult := execResult{}
	if affected, err := result.RowsAffected(); err != nil {
		return result, err
	} else {
		totalResult.rowsAffected = affected
	}
	if lastInsertId, err := result.LastInsertId(); err != nil {
		return result, err
	} else {
		totalResult.lastInsertId = lastInsertId
	}
	for next {
		statement, values, next, err = w.Build()
		if err != nil {
			return result, err
		}
		if batchResult, err := connectManagerInstance.Exec(ctx, configName, statement, values...); err != nil {
			return result, errors.SystemError.DatabaseExecError.Wrap(err)
		} else {
			if rowsAffected, err := batchResult.RowsAffected(); err != nil {
				return totalResult, errors.SystemError.DatabaseExecError.Wrap(err)
			} else {
				totalResult.rowsAffected += rowsAffected
			}
			if lastInsertId, err := batchResult.LastInsertId(); err != nil {
				return totalResult, errors.SystemError.DatabaseExecError.Wrap(err)
			} else {
				totalResult.lastInsertId = lastInsertId
			}
		}
	}
	return
}

func Update(ctx context.Context, configName string, w UpdateWrapper) (result sql.Result, err error) {
	statement, values, err := w.Build()
	if err != nil {
		return result, err
	}
	result, err = connectManagerInstance.Exec(ctx, configName, statement, values...)
	if err != nil {
		return result, errors.SystemError.DatabaseExecError.Wrap(err)
	}
	return
}

func Delete(ctx context.Context, configName string, w DeleteWrapper) (result sql.Result, err error) {
	statement, values, err := w.Build()
	if err != nil {
		return result, err
	}
	result, err = connectManagerInstance.Exec(ctx, configName, statement, values...)
	if err != nil {
		return result, errors.SystemError.DatabaseExecError.Wrap(err)
	}
	return
}

func Clone(ctx context.Context, configName string, w CloneWrapper) (result sql.Result, err error) {
	statement, values, err := w.Build()
	if err != nil {
		return result, err
	}
	result, err = connectManagerInstance.Exec(ctx, configName, statement, values...)
	if err != nil {
		return result, errors.SystemError.DatabaseExecError.Wrap(err)
	}
	return
}

func ShowTables(ctx context.Context, configName string) (tableNameList TableNameList, err error) {
	statement := fmt.Sprintf("SHOW TABLES")
	result, err := connectManagerInstance.Query(ctx, configName, statement)
	if err != nil {
		return nil, errors.SystemError.DatabaseQueryError.Wrap(err)
	}
	defer func() {
		_ = result.Close()
	}()
	for result.Next() {
		var tb string
		if err := result.Scan(&tb); err != nil {
			return nil, errors.SystemError.DatabaseScanError.Wrap(err)
		}
		tableNameList = append(tableNameList, TableName(tb))
	}
	return tableNameList, nil
}
