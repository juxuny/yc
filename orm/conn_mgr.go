package orm

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juxuny/yc/errors"
	"runtime/debug"
)

type IManager interface {
	Add(c Config) error
	Exec(ctx context.Context, configName string, statement string, values ...interface{}) (result sql.Result, err error)
	Query(ctx context.Context, configName string, statement string, values ...interface{}) (result *sql.Rows, err error)
	scanRows(result *sql.Rows) (DataSet, error)
	QueryRows(ctx context.Context, configName string, statement string, values ...interface{}) (DataSet, error)
}

var connectManagerInstance = NewConnManager()

type connMgr struct {
	// config name => index of conns
	ns    map[string]int
	conns []*sql.DB
}

func NewConnManager() IManager {
	return &connMgr{
		ns:    make(map[string]int),
		conns: make([]*sql.DB, 0),
	}
}

func (t *connMgr) Add(c Config) error {
	_, b := t.ns[c.Name]
	if b {
		return errors.SystemError.DuplicatedConfigName.WithField("name", c.Name)
	}
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", c.User, c.Pass, c.Host, c.Port, c.Schema))
	if err != nil {
		return errors.SystemError.DatabaseConnectError.Wrap(err)
	}
	t.conns = append(t.conns, conn)
	idx := len(t.conns)
	t.ns[c.Name] = idx - 1
	return nil
}

func (t *connMgr) Exec(ctx context.Context, configName string, statement string, values ...interface{}) (result sql.Result, err error) {
	idx, b := t.ns[configName]
	if !b {
		return nil, errors.SystemError.DatabaseConfigNotFound.WithField("name", configName)
	}
	if idx >= len(t.conns) {
		return nil, errors.SystemError.DatabaseConnectionIndexError.WithField("idx", idx)
	}
	conn := t.conns[idx]
	return conn.ExecContext(ctx, statement, values...)
}

func (t *connMgr) Query(ctx context.Context, configName string, statement string, values ...interface{}) (result *sql.Rows, err error) {
	idx, b := t.ns[configName]
	if !b {
		return nil, errors.SystemError.DatabaseConfigNotFound.WithField("name", configName)
	}
	if idx >= len(t.conns) {
		return nil, errors.SystemError.DatabaseConnectionIndexError.WithField("idx", idx)
	}
	conn := t.conns[idx]
	result, err = conn.QueryContext(ctx, statement, values...)
	if err != nil {
		return nil, errors.SystemError.DatabaseQueryError.Wrap(err)
	}
	return result, err
}

func (t connMgr) scanRows(result *sql.Rows) (ret DataSet, err error) {
	defer func() {
		_ = result.Close()
	}()
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			return
		}
	}()
	columnTypes, err := result.ColumnTypes()
	ret = NewDataSet()
	if err != nil {
		return nil, errors.SystemError.DatabaseColumnTypeError.Wrap(err)
	}
	for result.Next() {
		values, holder := generateSlotFromColumnTypes(columnTypes)
		if err := result.Scan(holder...); err != nil {
			return ret, errors.SystemError.DatabaseScanError.Wrap(err)
		}
		row := make(Row, len(values))
		for i, v := range values {
			row[i] = Column{
				Value: v,
				Name:  columnTypes[i].Name(),
			}
		}
		ret.AppendRow(row)
	}
	return ret, nil
}

func (t *connMgr) QueryRows(ctx context.Context, configName string, statement string, values ...interface{}) (rows DataSet, err error) {
	result, err := t.Query(ctx, configName, statement, values...)
	if err != nil {
		return nil, err
	}
	return t.scanRows(result)
}
