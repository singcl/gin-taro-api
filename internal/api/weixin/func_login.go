package weixin

import (
	"net/http"

	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/services/weixin"
)

type loginRequest struct {
	Code string `form:"code" binding:"required"` // 微信小程序临时登录凭证code
}

type loginResponse struct {
	Code string `json:"code"` // 微信小程序临时登录凭证code
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

		searchData := new(weixin.SearchCode2SessionData)
		searchData.JsCode = req.Code
		_, err := h.weixinService.Login(c, searchData)

		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminLoginError,
				code.Text(code.AdminLoginError)).WithError(err),
			)
			return
		}

		res.Code = req.Code
		c.Payload(res)
	}
}
