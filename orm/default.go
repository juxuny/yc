package orm

import (
	"context"
	"database/sql"
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

func Exec(ctx context.Context, configName string, statement string, values ...interface{}) (result sql.Result, err error) {
	return connectManagerInstance.Exec(ctx, configName, statement, values...)
}

func QueryRows(ctx context.Context, configName string, statement string, values ...interface{}) (DataSet, error) {
	return connectManagerInstance.QueryRows(ctx, configName, statement, values...)
}

func QueryScan(ctx context.Context, configName string, out interface{}, statement string, values ...interface{}) error {
	ds, err := connectManagerInstance.QueryRows(ctx, configName, statement, values...)
	if err != nil {
		return err
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
		return result, err
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
			return result, err
		} else {
			if rowsAffected, err := batchResult.RowsAffected(); err != nil {
				return totalResult, err
			} else {
				totalResult.rowsAffected += rowsAffected
			}
			if lastInsertId, err := batchResult.LastInsertId(); err != nil {
				return totalResult, err
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
	return connectManagerInstance.Exec(ctx, configName, statement, values...)
}

func Delete(ctx context.Context, configName string, w DeleteWrapper) (result sql.Result, err error) {
	statement, values, err := w.Build()
	if err != nil {
		return result, err
	}
	return connectManagerInstance.Exec(ctx, configName, statement, values...)
}