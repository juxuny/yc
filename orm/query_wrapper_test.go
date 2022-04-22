package orm

import (
	"context"
	"github.com/juxuny/yc/utils"
	"testing"
)

func TestQueryWrapper_Build(t *testing.T) {
	w := NewQueryWrapper(&testUserInfo{})
	w.Ge("price", 1).In("id", []int64{1, 2}).Nested(NewOrWhereWrapper().IsNull("is_deleted").IsNotNull("f1"))
	statement, values, err := w.Build()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(statement)
	t.Log(values...)

	var items []testUserInfo
	err = Select(context.Background(), DefaultName, w, &items)
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Log(utils.ToJson(item))
	}
}
