package redis

import (
	"context"
	"time"
)

var (
	// cache data for 28 days
	defaultExpiration = time.Hour * 24 * 28
)

type ClientInterface interface {
	SetString(ctx context.Context, key Key, v string, expiration ...time.Duration) error
	GetString(ctx context.Context, key Key) (string, error)
	SetInt64(ctx context.Context, key Key, v int64, expiration ...time.Duration) error
	GetInt64(ctx context.Context, key Key) (int64, error)
	SetObject(ctx context.Context, key Key, v interface{}, expiration ...time.Duration) error
	GetObject(ctx context.Context, key Key, out interface{}) error
	SetBool(ctx context.Context, key Key, v bool, expiration ...time.Duration) error
	GetBool(ctx context.Context, key Key) (bool, error)
	Delete(ctx context.Context, key ...Key) error
	Exists(ctx context.Context, key ...Key) (bool, error)
	IncBy(ctx context.Context, key Key, value int64) error
	IncByFloat64(ctx context.Context, key Key, value float64) error

	HSetString(ctx context.Context, key Key, field Key, v string) error
	HGetString(ctx context.Context, key Key, field Key) (string, error)
	HSetInt64(ctx context.Context, key Key, field Key, v int64) error
	HGetInt64(ctx context.Context, key Key, field Key) (int64, error)
	HSetObject(ctx context.Context, key Key, field Key, v interface{}) error
	HGetObject(ctx context.Context, key Key, field Key, out interface{}) error
	HSetBool(ctx context.Context, key Key, field Key, v bool) error
	HGetBool(ctx context.Context, key Key, field Key) (bool, error)
	HDelete(ctx context.Context, key Key, field Key) error
	HExists(ctx context.Context, key Key, field Key) (bool, error)
	HIncBy(ctx context.Context, key Key, field Key, value int64) error
	HIncByFloat64(ctx context.Context, key Key, field Key, value float64) error

	Expire(ctx context.Context, key Key, expiration time.Duration) error
}

var defaultInstance ClientInterface

func Client() ClientInterface {
	if defaultInstance == nil {
		panic("uninitialized redis client, please, call redis.InitConfig() first")
	}
	return defaultInstance
}

func InitConfig(c Config) {
	defaultInstance = NewClient(c)
}
