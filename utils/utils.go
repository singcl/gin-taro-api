package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
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

func MiniProgramLogin(c *gin.Context, code string) (map[string]interface{}, error) {
	wc := InitWechat(c)
	cfg := &miniConfig.Config{
		AppID:     "APPID",
		AppSecret: "APPSECRET",
	}
	miniprogram := wc.GetMiniProgram(cfg)
	auth := miniprogram.GetAuth()
	session, err := auth.Code2Session(code)

	result := map[string]interface{}{}
	if err != nil {
		return result, err
	}
	result["openid"] = session.OpenID
	result["session_key"] = session.SessionKey
	result["unionid"] = session.UnionID
	return result, nil
}

// utils.MiniProgramLogin(c, code)

// 小程序需要调用 wx.login 获取到 code
// App({
// 	globalData: {},
// 	onLaunch: function () {
// 		var token = wx.getStorageSync('token'); // 获取本地缓存
// 		if (!token) {
// 			wx.login({
// 				success: function (_a) {
// 					var code = _a.code; // 获取到登陆code
// 					if (code) {
// 						console.log(code);
// 						// 在这里请求API进行登陆
// 					}
// 				}
// 			});
// 			return;
// 		}
// 	}
//   });
