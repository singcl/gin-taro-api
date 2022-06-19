package weixin

import "github.com/singcl/gin-taro-api/internal/pkg/core"

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
		// TODO
	}
}
