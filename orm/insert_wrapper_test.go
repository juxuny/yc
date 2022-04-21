package orm

import (
	"context"
	"fmt"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/utils"
	"testing"
	"time"
)

func TestInsertWrapper_Build(t *testing.T) {
	w := NewInsertWrapper(&testUserInfo{})
	w.Add(&testUserInfo{
		Id:          dt.NewID(30),
		Name:        utils.RandHelper.RandString(4),
		CreateTime:  nil,
		Price:       10,
		Count:       40,
		Remark:      dt.NullString{Valid: true, String_: utils.RandHelper.RandString(10)},
		F1:          dt.NullFloat64{},
		Type:        1,
		Deleted:     nil,
		DeletedAt:   nil,
		Description: "description",
		Started:     1,
	}).OnDuplicatedUpdate("name", 100).OnDuplicatedUpdate("count").OnDuplicatedUpdate("deleted_at", time.Now())
	result, err := InsertWithWrapper(context.Background(), DefaultName, w)
	if err != nil {
		t.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	t.Log(fmt.Sprintf("rows affected: %v", rowsAffected))
}
