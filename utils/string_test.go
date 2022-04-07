package utils

import "testing"

func TestToUpperFirst(t *testing.T) {
	in := []struct {
		Input  string
		Output string
	}{
		{Input: "Aaaa", Output: "Aaaa"},
		{Input: "aa", Output: "Aa"},
		{Input: "", Output: ""},
	}
	for _, item := range in {
		ret := ToUpperFirst(item.Input)
		if ret != item.Output {
			t.Fatal("the real result: ", item.Output, " got: ", ret)
		}
	}
}
