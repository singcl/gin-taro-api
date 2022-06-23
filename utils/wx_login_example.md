### 后端代码核心逻辑示例

```go
// services/weixin/service.go
package weixin

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	Login(ctx core.Context, searchCode2Session *SearchCode2SessionData) (info *Code2SessionData, err error)
}

type service struct {
	cache       redis.Repo
	wc          *wechat.Wechat
	miniprogram *miniprogram.MiniProgram
}

func New(cache redis.Repo) Service {
	wc := wechat.NewWechat()
	wc.SetCache(cache)
	cfg := &miniConfig.Config{
		AppID:     configs.Get().Wechat.AppID,
		AppSecret: configs.Get().Wechat.Secret,
	}
	miniprogram := wc.GetMiniProgram(cfg)

	return &service{
		cache:       cache,
		wc:          wc,
		miniprogram: miniprogram,
	}
}

```

```go
// services/weixin/service_login.go
package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
)

type SearchCode2SessionData struct {
	JsCode string `json:"js_code"` // 登录时获取的 code
}

type Code2SessionData struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
	Errcode    int64  `json:"errcode"`     // 错误码
	Errmsg     string `json:"errmsg"`      // 错误信息
}

func (s *service) Login(ctx core.Context, searchCode2Session *SearchCode2SessionData) (info *Code2SessionData, err error) {
	auth := s.miniprogram.GetAuth()
	session, err := auth.Code2Session(searchCode2Session.JsCode)

	if err != nil {
		return nil, err
	}
	info = &Code2SessionData{
		session.OpenID,
		session.SessionKey,
		session.UnionID,
		session.ErrCode,
		session.ErrMsg,
	}
	return
}

```

```go
// api/weixin/handler.go
package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/weixin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	// login 微信登录
	Login() core.HandlerFunc
}

type handler struct {
	cache         redis.Repo
	weixinService weixin.Service
}

func New(cache redis.Repo) Handler {
	return &handler{
		cache:         cache,
		weixinService: weixin.New(cache),
	}
}
```

```go
// api/weixin/func_login.go
package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/services/weixin"
)

type loginRequest struct {
	Code string `form:"code" binding:"required"` // 微信小程序临时登录凭证code
}

type loginResponse struct {
	Token string `json:"token"` // 微信小程序登录凭证token
}

// Sign 微信登录
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)

		// 微信Code2Session
		searchData := new(weixin.SearchCode2SessionData)
		searchData.JsCode = req.Code
		wxLoginData, err := h.weixinService.Login(c, searchData)
	}
}

```

### 前端代码示例

```js
// 小程序需要调用 wx.login 获取到 code
App({
  globalData: {},
  onLaunch: function () {
    var token = wx.getStorageSync('token'); // 获取本地缓存
    if (!token) {
      wx.login({
        success: function (_a) {
          var code = _a.code; // 获取到登陆code
          if (code) {
            console.log(code);
            // 在这里请求API进行登陆
          }
        },
      });
      return;
    }
  },
});
```
