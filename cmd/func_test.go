package cmd

import "testing"

func TestCheckIfCommandExists(t *testing.T) {
	ok, err := CheckIfCommandExists("ls")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("ls is not found")
	}
	ok, err = CheckIfCommandExists("ls-abc")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("ls-abc is found")
	}
}
