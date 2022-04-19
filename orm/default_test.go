package orm

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/utils"
	"testing"
	"time"
)

type testUserType int

const (
	testUserTypeDefault = testUserType(0)
	testUserTypeNormal  = testUserType(1)
)

type testUserInfo struct {
	Id          dt.ID
	Name        string
	CreateTime  *time.Time
	Price       float64
	Count       int64
	Remark      dt.NullString
	F1          dt.NullFloat64
	Type        testUserType
	Deleted     *int `orm:"is_deleted"`
	DeletedAt   *time.Time
	Description string
}

func TestQueryScan(t *testing.T) {
	var items []testUserInfo
	err := QueryScan(context.Background(), DefaultName, &items, "SELECT * FROM test_user_info LIMIT 2")
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Log(utils.ToJson(item))
	}
}

// goos: darwin
// goarch: amd64
// pkg: github.com/juxuny/yc/orm
// cpu: Intel(R) Core(TM) i7-8569U CPU @ 2.80GHz
// BenchmarkQueryScan
// BenchmarkQueryScan-8   	     604	   1739670 ns/op
func BenchmarkQueryScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var items []testUserInfo
		err := QueryScan(context.Background(), DefaultName, &items, "SELECT * FROM test_user_info")
		if err != nil {
			b.Fatal(err)
		}
	}
}
