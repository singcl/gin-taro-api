package weixin

import (
	"errors"
	"net/http"

	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/pkg/password"
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

		// TODO 找不到用户则插入，找到则更新
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinLoginError,
				code.Text(code.WeixinLoginError)).WithError(err),
			)
			return
		}

		if info == nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WeixinLoginError,
				code.Text(code.WeixinLoginError)).WithError(errors.New("Code2Session成功，但是微信用户表中未找到该用户")),
			)
			return
		}

		token := password.GenerateWeixinLoginToken(info.Openid)

		res.Token = token
		c.Payload(res)
	}
}
