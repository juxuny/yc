package validator

import (
	"testing"
)

func TestValidator(t *testing.T) {
	type Level int
	type TestStruct struct {
		Level   Level
		Content string
		List    []int
	}
	v := &TestStruct{
		Level:   3,
		Content: "100",
		List:    []int{1},
	}
	action := Action{
		ValidatorFormulas: []string{
			"min",
		},
		ErrorTemplate: "minimum of level is 5, current value is: {{.Level}}",
	}
	if err := RegisterTemplate(action.ErrorTemplate); err != nil {
		t.Fatal(err)
	}
	if err := Run(v.Level, "5", action, v); err != nil {
		t.Fatal(err)
	}
}
