package weixin

import (
	"net/http"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/pkg/password"
	"github.com/singcl/gin-taro-api/internal/proposal"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/weixin"
)

type loginRequest struct {
	Code string `form:"code" binding:"required"` // 微信小程序临时登录凭证code
}

type loginResponse struct {
	Token string `json:"token"` // 微信小程序登录凭证token
}

// Sign 微信登录
// @Summary 微信登录
// @Description 微信登录
// @Tags Weixin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param code formData string true "微信小程序临时登录凭证code"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /weixin/login [get]
func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		// 微信Code2Session
		searchData := new(weixin.SearchCode2SessionData)
		searchData.JsCode = req.Code
		wxLoginData, err := h.weixinService.Login(c, searchData)

		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinLoginError,
				code.Text(code.WeixinLoginError)).WithError(err),
			)
			return
		}

		// 数据库查询该用户
		searchOneData := new(weixin.SearchOneData)
		searchOneData.Openid = wxLoginData.OpenID
		searchOneData.IsUsed = 1

		info, err := h.weixinService.Detail(c, searchOneData)

		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinLoginError,
				code.Text(code.WeixinLoginError)).WithError(err),
			)
			return
		}

		// TODO 找不到用户则插入，找到则更新
		if info == nil {
			createWeixinUserData := new(weixin.CreateWeixinUserData)
			createWeixinUserData.Openid = wxLoginData.OpenID
			createWeixinUserData.Unionid = wxLoginData.UnionID
			createWeixinUserData.SessionKey = wxLoginData.SessionKey

			_, err := h.weixinService.Create(c, createWeixinUserData)
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.WeixinLoginError,
					code.Text(code.WeixinLoginError)).WithError(err),
				)
				return
			}
		}

		token := password.GenerateWeixinLoginToken(wxLoginData.OpenID)

		// 用户信息
		sessionWeixinUserInfo := &proposal.WeixinSessionUserInfo{
			Openid:     wxLoginData.OpenID,
			SessionKey: wxLoginData.SessionKey,
			AvatarUrl:  info.AvatarUrl,
			Nickname:   info.Nickname,
			Mobile:     info.Mobile,
		}

		// 将用户信息记录到 Redis 中
		err = h.cache.SetR(
			configs.RedisKeyPrefixWeixinLoginUser+token,
			string(sessionWeixinUserInfo.Marshal()),
			configs.LoginSessionTTL,
			redis.WithTrace(c.Trace()))

		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinLoginError,
				code.Text(code.WeixinLoginError)).WithError(err),
			)
			return
		}

		res.Token = token
		c.Payload(res)
	}
}
