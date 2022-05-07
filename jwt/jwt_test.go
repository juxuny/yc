package jwt

import (
	"github.com/juxuny/yc/dt"
	"testing"
)

func TestMain(m *testing.M) {
	SetSecret("123456")
	m.Run()
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(dt.NewID(100), "admin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
	claim, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(claim.UserId, claim.UserName)
}
