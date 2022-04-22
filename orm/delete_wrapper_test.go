package orm

import (
	"context"
	"testing"
)

func TestDeleteWrapper_Build(t *testing.T) {
	w := NewDeleteWrapper().TableName("test_user_info")
	w.Gt("id", 30)
	result, err := Delete(context.Background(), DefaultName, w)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.RowsAffected())
}
