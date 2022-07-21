package admin

import (
	"net/http"

	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/pkg/password"
	"github.com/singcl/gin-taro-api/internal/services/admin"
)

type modifyPasswordRequest struct {
	OldPassword string `form:"old_password"` // 旧密码
	NewPassword string `form:"new_password"` // 新密码
}

type modifyPasswordResponse struct {
	Username string `json:"username"` // 用户账号
}

func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(modifyPasswordRequest)
		res := new(modifyPasswordResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchOneData := new(admin.SearchOneData)
		searchOneData.Id = ctx.SessionUserInfo().UserID
		searchOneData.Password = password.GeneratePassword(req.OldPassword)
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(ctx, searchOneData)
		if err != nil || info == nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithError(err),
			)
			return
		}

		if err := h.adminService.ModifyPassword(ctx, ctx.SessionUserInfo().UserID, req.NewPassword); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminModifyPasswordError,
				code.Text(code.AdminModifyPasswordError)).WithError(err),
			)
			return
		}

		res.Username = ctx.SessionUserInfo().UserName
		ctx.Payload(res)
	}
}
