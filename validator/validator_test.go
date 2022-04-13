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
		Level:   5,
		Content: "100",
		List:    []int{1},
	}
	action := Action{
		ValidatorFormulas: "min",
		RefValue:          "5",
		ErrorTemplate:     "{{if lt .Level 5 }}minimum of level is 5, current value is: {{.Level}}{{else}}maximum of level is 10{{end}}",
	}
	if err := RegisterTemplate(action.ErrorTemplate); err != nil {
		t.Fatal(err)
	}
	if err := Run(v.Level, action, v); err != nil {
		t.Fatal(err)
	}
	v = &TestStruct{
		Level:   0,
		Content: "",
		List:    nil,
	}
	action = Action{
		ValidatorFormulas: "in",
		RefValue:          "1,2,3,4,5",
		ErrorTemplate:     "invalid level = {{.Level}}",
	}
	if err := RegisterTemplate(action.ErrorTemplate); err != nil {
		t.Fatal(err)
	}
	if err := Run(v.Level, action, v); err != nil {
		t.Fatal(err)
	}
}
