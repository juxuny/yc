package errors

import "testing"

type TestErrorConst struct {
	Success Error `code:"0" msg:""`
	System  Error `code:"-1" msg:"system error"`
	Proto   Error `code:"-2" msg:"proto error"`
}

func TestInitErrorStruct(t *testing.T) {
	var obj TestErrorConst
	if err := InitErrorStruct(&obj); err != nil {
		t.Fatal(err)
	}
	if obj.Success.Code != 0 {
		t.Fatal("Success.Code is error: ", obj.Success.Code)
	}
	if obj.System.Code != -1 || obj.System.Msg != "system error" {
		t.Fatal(obj.System.Code, obj.System.Msg)
	}
	t.Log(obj.Success)
	t.Log(obj.System)
	t.Log(obj.System.WithField("reqId", "123"))
	t.Log(obj.System)
	t.Log(FromError(obj.System.Err()))
	t.Log(FromError(obj.Success.Err()))
}
