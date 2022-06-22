package main

import (
	"context"
	"github.com/juxuny/yc/redis"
	"github.com/spf13/cobra"
	"log"
)

type watchCommand struct {
	Host       string
	AccessKey  string
	Secret     string
	ConfigId   string
	RedisHost  string
	RedisPass  string
	RedisPort  int
	RedisIndex int
	RedisUser  string
}

func (t *watchCommand) Prepare(cmd *cobra.Command) {
}

func (t *watchCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.Host, "host", "p", "http://127.0.0.1:20080", "cos host")
	cmd.PersistentFlags().StringVarP(&t.AccessKey, "access-key", "a", "", "access key")
	cmd.PersistentFlags().StringVarP(&t.Secret, "secret", "s", "", "secret")
	cmd.PersistentFlags().StringVar(&t.ConfigId, "config-id", "", "config ID")

	cmd.PersistentFlags().IntVar(&t.RedisIndex, "redis-index", 0, "database index")
	cmd.PersistentFlags().IntVar(&t.RedisPort, "redis-port", 6379, "redis port")
	cmd.PersistentFlags().StringVar(&t.RedisHost, "redis-host", "127.0.0.1", "redis host")
	cmd.PersistentFlags().StringVar(&t.RedisUser, "redis-user", "", "redis auth user, redis v5 or earlier has no user name")
	cmd.PersistentFlags().StringVar(&t.RedisPass, "redis-pass", "", "redis auth password")
}

func (t *watchCommand) BeforeRun(cmd *cobra.Command) {
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
	if t.RedisHost == "" {
		log.Fatal("missing arguments: --redis-host")
	}
	if t.RedisPort == 0 {
		log.Fatal("missing arguments: --redis-port")
	}
	redis.InitConfig(redis.Config{
		UsePass: true,
		Host:    t.RedisHost,
		User:    t.RedisUser,
		Pass:    t.RedisPass,
		Port:    t.RedisPort,
		Index:   t.RedisIndex,
	})
}

func (t *watchCommand) Run() {
	ctx := context.Background()
	redis.Client().Subscription(ctx, KeyValue.NotifyChannel.Suffix(t.ConfigId), func(msg string) {
		log.Println(msg)
	})
}
