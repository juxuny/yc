package orm

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/juxuny/yc/env"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	var databaseEnv = struct {
		DbHost   string
		DbPort   int
		DbPass   string
		DbUser   string
		DbSchema string
	}{}
	env.Init(&databaseEnv, true)
	var config = Config{
		Name:   DefaultName,
		Host:   databaseEnv.DbHost,
		Port:   databaseEnv.DbPort,
		User:   databaseEnv.DbUser,
		Pass:   databaseEnv.DbPass,
		Schema: databaseEnv.DbSchema,
	}
	if err := InitConfig(config); err != nil {
		panic(err)
	}
	m.Run()
}

func TestQuery(t *testing.T) {
	//statement := fmt.Sprintf("INSERT INTO test_user_info (name, create_time, price, count, remark) VALUES (?, ?, ?, ?, ?)")
	//values := []interface{}{
	//	utils.StringHelper.RandString(5),
	//	time.Now(),
	//	utils.RandHelper.Float64(10),
	//	utils.RandHelper.Int63n(1000),
	//	utils.StringHelper.RandString(5),
	//}
	//if _, err := Exec(context.Background(), DefaultName, statement, values...); err != nil {
	//	t.Fatal(err)
	//}
	ds, err := QueryRows(context.Background(), DefaultName, "SELECT * FROM test_user_info")
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range ds {
		line := make([]string, 0)
		for _, col := range row {
			elem := col.Elem()
			switch elem.Kind() {
			case reflect.Uint64:
				line = append(line, fmt.Sprintf("%v", col.Elem().Interface()))
			default:
				t.Log(elem.Type().String())
				if elem.Type().String() == "sql.RawBytes" {
					line = append(line, fmt.Sprintf("%v", string(elem.Interface().(sql.RawBytes))))
				} else if elem.Type().String() == "sql.NullTime" {
					line = append(line, fmt.Sprintf("%v", elem.FieldByName("Time").Interface()))
				} else if elem.Type().String() == "sql.NullInt64" {
					line = append(line, fmt.Sprintf("%v", elem.FieldByName("Int64").Interface()))
				} else {
					line = append(line, fmt.Sprintf("%v", elem.FieldByName("Value").Interface()))
				}
			}
		}
		//fmt.Println(strings.Join(line, ","))
	}
}

// yc/orm: BenchmarkConnMgr_QueryRows-8   	     792	   1644782 ns/op
// gorm: BenchmarkQuery-8   	     320	   3977007 ns/op
func BenchmarkConnMgr_QueryRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := QueryRows(context.Background(), DefaultName, "SELECT * FROM test_user_info")
		if err != nil {
			b.Fatal(err)
		}
	}
}
