package utils

import "testing"

func TestStructHelper_Patch(t *testing.T) {
	type data struct {
		UserId int64
		Name   string
	}
	var d = data{}
	userId := int64(100)
	if err := StructHelper.Patch(&d, map[string]interface{}{
		"UserId": userId,
	}); err != nil {
		t.Fatal(err)
	}
	t.Log(d.UserId)
}
