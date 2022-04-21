package orm

import (
	"context"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/utils"
	"testing"
)

func TestUpdateWrapper_Updates(t *testing.T) {
	item := testUserInfo{
		Id:   dt.NewID(uint64(utils.RandHelper.Int63())),
		Name: utils.RandHelper.RandString(40),
	}
	w := NewUpdateWrapper(&item)
	w.Updates(item).Eq("id", 37).Eq("name", "Lpr1")
	result, err := UpdateWithWrapper(context.Background(), DefaultName, w)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.RowsAffected())
}
