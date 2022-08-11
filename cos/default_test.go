package cos

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/env"
	"github.com/juxuny/yc/services/cos"
	"testing"
)

type TestEnv struct {
	AccessKey  string
	Secret     string
	Entrypoint string
	ConfigId   string
}

var testEnv TestEnv

func init() {
	env.Init(&testEnv, true, "COS")
}

func TestFetchValues(t *testing.T) {
	ctx := context.Background()
	cos.Config(yc.NewRandomEntrypointDispatcher([]string{testEnv.Entrypoint}), yc.NewDefaultSignHandler(testEnv.AccessKey, testEnv.Secret))
	resp, err := cos.ListAllValueByConfigId(ctx, &cos.ListAllValueByConfigIdRequest{
		ConfigId:   testEnv.ConfigId,
		IsDisabled: &dt.NullBool{Valid: true, Bool: false},
		IsHot:      nil,
		SearchKey:  "",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range resp.List {
		t.Log(item.ConfigKey, item.ConfigValue)
	}
}
