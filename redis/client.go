package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
	"time"
)

type client struct {
	redisClient *redis.Client
}

func (t *client) IncByFloat64(ctx context.Context, key Key, value float64) error {
	return t.redisClient.IncrByFloat(ctx, key.String(), value).Err()
}

func (t *client) HIncByFloat64(ctx context.Context, key Key, field Key, value float64) error {
	return t.redisClient.HIncrByFloat(ctx, key.String(), field.String(), value).Err()
}

func (t *client) HSetString(ctx context.Context, key Key, field Key, v string) error {
	return t.redisClient.HSet(ctx, key.String(), field.String(), v).Err()
}

func (t *client) HGetString(ctx context.Context, key Key, field Key) (string, error) {
	return t.redisClient.HGet(ctx, key.String(), field.String()).Result()
}

func (t *client) HSetInt64(ctx context.Context, key Key, field Key, v int64) error {
	return t.HSetString(ctx, key, field, fmt.Sprintf("%d", v))
}

func (t *client) HGetInt64(ctx context.Context, key Key, field Key) (int64, error) {
	v, err := t.HGetString(ctx, key, field)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(v, 10, 64)
}

func (t *client) HSetObject(ctx context.Context, key Key, field Key, v interface{}) error {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return t.HSetString(ctx, key, field, string(jsonData))
}

func (t *client) HGetObject(ctx context.Context, key Key, field Key, out interface{}) error {
	v, err := t.HGetString(ctx, key, field)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(v), out)
}

func (t *client) HSetBool(ctx context.Context, key Key, field Key, v bool) error {
	if v {
		return t.HSetString(ctx, key, field, "1")
	} else {
		return t.HSetString(ctx, key, field, "0")
	}
}

func (t *client) HGetBool(ctx context.Context, key Key, field Key) (bool, error) {
	v, err := t.HGetString(ctx, key, field)
	if err != nil {
		return false, err
	}
	return v == "1" || v == "true", nil
}

func (t *client) HDelete(ctx context.Context, key Key, field Key) error {
	return t.redisClient.HDel(ctx, key.String(), field.String()).Err()
}

func (t *client) HExists(ctx context.Context, key Key, field Key) (bool, error) {
	return t.redisClient.HExists(ctx, key.String(), field.String()).Result()
}

func (t *client) HIncBy(ctx context.Context, key Key, field Key, value int64) error {
	return t.redisClient.HIncrBy(ctx, key.String(), field.String(), value).Err()
}

func (t *client) Expire(ctx context.Context, key Key, expiration time.Duration) error {
	panic("implement me")
}

func (t *client) IncBy(ctx context.Context, key Key, value int64) error {
	return t.redisClient.IncrBy(ctx, key.String(), value).Err()
}

func (t *client) Delete(ctx context.Context, key ...Key) error {
	if len(key) == 0 {
		return fmt.Errorf("empty key list")
	}
	ks := convertKeysToStrings(key...)
	return t.redisClient.Del(ctx, ks...).Err()
}

func (t *client) Exists(ctx context.Context, key ...Key) (bool, error) {
	if len(key) == 0 {
		return false, fmt.Errorf("empty key list")
	}
	ks := convertKeysToStrings(key...)
	v, err := t.redisClient.Exists(ctx, ks...).Result()
	return v > 0, err
}

func (t *client) SetString(ctx context.Context, key Key, v string, expiration ...time.Duration) error {
	finalExpiration := defaultExpiration
	if len(expiration) > 0 {
		finalExpiration = expiration[0]
	}
	ret := t.redisClient.Set(ctx, key.String(), v, finalExpiration)
	return ret.Err()
}

func (t *client) GetString(ctx context.Context, key Key) (string, error) {
	v, err := t.redisClient.Get(ctx, key.String()).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (t *client) SetInt64(ctx context.Context, key Key, v int64, expiration ...time.Duration) error {
	return t.SetString(ctx, key, fmt.Sprintf("%d", v), expiration...)
}

func (t *client) GetInt64(ctx context.Context, key Key) (int64, error) {
	v, err := t.GetString(ctx, key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(v, 10, 64)
}

func (t *client) SetObject(ctx context.Context, key Key, v interface{}, expiration ...time.Duration) error {
	data, _ := json.Marshal(v)
	return t.SetString(ctx, key, string(data), expiration...)
}

func (t *client) GetObject(ctx context.Context, key Key, out interface{}) error {
	jsonData, err := t.GetString(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonData), out)
}

func (t *client) SetBool(ctx context.Context, key Key, v bool, expiration ...time.Duration) error {
	if v {
		return t.SetString(ctx, key, "1", expiration...)
	} else {
		return t.SetString(ctx, key, "0", expiration...)
	}
}

func (t *client) GetBool(ctx context.Context, key Key) (bool, error) {
	v, err := t.GetString(ctx, key)
	if err != nil {
		return false, err
	}
	if v == "0" || strings.ToLower(v) == "false" {
		return false, nil
	} else {
		return true, nil
	}
}

func NewClient(config Config) ClientInterface {
	c := &client{}
	var opt = &redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
		DB:   config.Index,
	}
	if config.UsePass {
		opt.Username = config.User
		opt.Password = config.Pass
	}
	c.redisClient = redis.NewClient(opt)
	return c
}
