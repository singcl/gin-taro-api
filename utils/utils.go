package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
)

func InitWechat(c *gin.Context) *wechat.Wechat {
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        "127.0.0.1:3306",
		Database:    0,  // redis db
		MaxActive:   10, // 连接池最大活跃连接数
		MaxIdle:     10, //连接池最大空闲连接数
		IdleTimeout: 60, //空闲连接超时时间，单位：second
	}
	redisCache := cache.NewRedis(c, redisOpts)
	wc.SetCache(redisCache)
	return wc
}
