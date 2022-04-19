package orm

import (
	"database/sql"
	"reflect"
	"strings"
)

func generateSlotFromColumnTypes(columnTypes []*sql.ColumnType) (values []reflect.Value, holder []interface{}) {
	for _, col := range columnTypes {
		v := reflect.New(col.ScanType())
		values = append(values, v)
		holder = append(holder, v.Interface())
	}
	return
}

func getOrmColumnNameMap(model reflect.Type) map[string]string {
	ret := make(map[string]string)
	for i := 0; i < model.NumField(); i++ {
		ft := model.Field(i)
		tag := ft.Tag
		if v, b := tag.Lookup("orm"); b {
			ret[strings.TrimSpace(v)] = ft.Name
		}
	}
	return ret
}
