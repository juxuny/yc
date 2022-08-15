package cos

import (
	"context"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/log"
	"github.com/juxuny/yc/services/cos"
)

type client struct {
	Options Options
}

func NewClient(opt Options) *client {
	cos.Config(yc.NewRandomEntrypointDispatcher([]string{opt.CosEntrypoint}), yc.NewDefaultSignHandler(opt.AccessKey, opt.Secret))
	return &client{
		Options: opt,
	}
}

func (t *client) fetch() (data map[string]string, err error) {
	resp, err := cos.ListAllValueByConfigId(context.Background(), &cos.ListAllValueByConfigIdRequest{
		ConfigId:   t.Options.ConfigId,
		IsDisabled: &dt.NullBool{Valid: true, Bool: false},
		SearchKey:  "",
		IsHot:      nil,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	data = make(map[string]string)
	for _, item := range resp.List {
		data[item.ConfigKey] = item.ConfigValue
	}
	return data, nil
}

func (t *client) Parse(out interface{}) error {
	configMap, err := t.fetch()
	if err != nil {
		log.Error(err)
		return nil
	}
	return Parse(configMap, out)
}
