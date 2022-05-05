package utils

import "testing"

func TestBcrypt(t *testing.T) {
	in := "XPr2yJk9T6WFUTJX"
	h := Bcrypt(in)
	t.Log(h)
}
