package redis

import (
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/pkg/errors"
	"github.com/singcl/gin-taro-api/pkg/timeutil"
	"github.com/singcl/gin-taro-api/pkg/trace"
)

type Option func(*option)
type Trace = trace.T

var _ Repo = (*cacheRepo)(nil)

type option struct {
	Trace *trace.Trace
	Redis *trace.Redis
}

func newOption() *option {
	return &option{}
}

type Repo interface {
	i()
	Exists(keys ...string) bool
	GetR(key string, options ...Option) (string, error)
	SetR(key, value string, ttl time.Duration, options ...Option) error
	Del(key string, options ...Option) bool
	Version() string
	//
	IsExist(key string) bool
	Delete(key string) error
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
}

// https://silenceper.com/wechat/officialaccount/configuration.html
// 自定义微信缓存数据接口
type WeixinCache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}

type cacheRepo struct {
	client *redis.Client
}

func New() (Repo, error) {
	client, err := redisConnect()
	if err != nil {
		return nil, err
	}

	return &cacheRepo{
		client: client,
	}, nil
}

func redisConnect() (*redis.Client, error) {
	cfg := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Pass,
		DB:           cfg.Db,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "ping redis err")
	}

	return client, nil
}

func (c *cacheRepo) i() {}

func (c *cacheRepo) Exists(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(keys...).Result()
	return value > 0
}

func (c *cacheRepo) IsExist(key string) bool {
	if key == "" {
		return true
	}
	value, _ := c.client.Exists(key).Result()
	return value > 0
}

// Set set some <key,value> into redis
func (c *cacheRepo) Set(key string, value interface{}, ttl time.Duration) error {
	if err := c.client.Set(key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}

	return nil
}

// Set set some <key,value> into redis
func (c *cacheRepo) SetR(key, value string, ttl time.Duration, options ...Option) error {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "set"
			opt.Redis.Key = key
			opt.Redis.Value = value
			opt.Redis.TTL = ttl.Minutes()
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	if err := c.client.Set(key, value, ttl).Err(); err != nil {
		return errors.Wrapf(err, "redis set key: %s err", key)
	}

	return nil
}

// WithTrace 设置trace信息
func WithTrace(t Trace) Option {
	return func(opt *option) {
		if t != nil {
			opt.Trace = t.(*trace.Trace)
			opt.Redis = new(trace.Redis)
		}
	}
}

// Get get some key from redis
func (c *cacheRepo) GetR(key string, options ...Option) (string, error) {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "get"
			opt.Redis.Key = key
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	value, err := c.client.Get(key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "redis get key: %s err", key)
	}

	return value, nil
}

// Get get some key from redis
func (c *cacheRepo) Get(key string) interface{} {
	value, err := c.client.Get(key).Result()
	if err != nil {
		return ""
	}

	return value
}

func (c *cacheRepo) Del(key string, options ...Option) bool {
	ts := time.Now()
	opt := newOption()
	defer func() {
		if opt.Trace != nil {
			opt.Redis.Timestamp = timeutil.CSTLayoutString()
			opt.Redis.Handle = "del"
			opt.Redis.Key = key
			opt.Redis.CostSeconds = time.Since(ts).Seconds()
			opt.Trace.AppendRedis(opt.Redis)
		}
	}()

	for _, f := range options {
		f(opt)
	}

	if key == "" {
		return true
	}

	value, _ := c.client.Del(key).Result()
	return value > 0
}

func (c *cacheRepo) Delete(key string) error {
	if key == "" {
		return nil
	}

	value, _ := c.client.Del(key).Result()

	if value > 0 {
		return nil
	}
	return errors.New("删除失败")
}

// Version redis server version
func (c *cacheRepo) Version() string {
	server := c.client.Info("server").Val()
	spl1 := strings.Split(server, "# Server")
	spl2 := strings.Split(spl1[1], "redis_version:")
	spl3 := strings.Split(spl2[1], "redis_git_sha1:")
	return spl3[0]
}
