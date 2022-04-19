package orm

import (
	"context"
	"database/sql"
)

func Exec(ctx context.Context, configName string, statement string, values ...interface{}) (result sql.Result, err error) {
	return connectManagerInstance.Exec(ctx, configName, statement, values...)
}

func QueryRows(ctx context.Context, configName string, statement string, values ...interface{}) (DataSet, error) {
	return connectManagerInstance.QueryRows(ctx, configName, statement, values...)
}

func QueryScan(ctx context.Context, configName string, out interface{}, statement string, values ...interface{}) error {
	ds, err := QueryRows(ctx, configName, statement, values...)
	if err != nil {
		return err
	}
	return ds.Reform(out)
}
