package main

import (
	"context"
	"fmt"
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/dt"
	"github.com/juxuny/yc/services/cos"
	"github.com/juxuny/yc/utils"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/metadata"
	"log"
)

type getCommand struct {
	Host      string
	AccessKey string
	Secret    string
	ConfigId  string
}

func (t *getCommand) Sum(data []byte) (method yc.SignMethod, signResult string, err error) {
	sign := utils.HashHelper.Sha256FromBytesToString(data, []byte(t.Secret))
	return yc.SignMethodSha256, sign, nil
}

func (t *getCommand) SelectOne() string {
	return t.Host
}

func (t *getCommand) Prepare(cmd *cobra.Command) {
}

func (t *getCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.Host, "host", "p", "http://127.0.0.1:20080", "cos host")
	cmd.PersistentFlags().StringVarP(&t.AccessKey, "access-key", "a", "", "access key")
	cmd.PersistentFlags().StringVarP(&t.Secret, "secret", "s", "", "secret")
	cmd.PersistentFlags().StringVar(&t.ConfigId, "config-id", "", "config ID")
}

func (t *getCommand) BeforeRun(cmd *cobra.Command) {
	if t.Host == "" {
		log.Fatal("missing arguments: --host")
	}
	if t.AccessKey == "" {
		log.Fatal("missing arguments: --access-key")
	}
	if t.Secret == "" {
		log.Fatal("missing arguments: --secret")
	}
	if t.ConfigId == "" {
		log.Fatal("missing arguments: --config-id")
	}
}

func (t *getCommand) Run() {
	cos.Config(t, t)
	ctx := context.Background()
	resp, err := cos.DefaultClient.ListAllValue(ctx, &cos.ListAllValueRequest{
		ConfigId:   dt.NewIDPointer(24),
		IsDisabled: &dt.NullBool{Valid: true, Bool: false},
	}, metadata.New(map[string]string{
		yc.MdContextAccessKey: t.AccessKey,
	}))
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range resp.List {
		log.Println(fmt.Sprintf("%s = %s", item.ConfigKey, item.ConfigValue))
	}
}
