package admin

import (
	"net/http"

	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/services/admin"
)

type createAdminMenuRequest struct {
	Id      string `form:"id"`      // HashID
	Actions string `form:"actions"` // 功能权限ID,多个用,分割
}

type createAdminMenuResponse struct {
	Id int32 `json:"id"` // 主键ID
}

func (h *handler) CreateAdminMenu() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createAdminMenuRequest)
		res := new(createAdminMenuResponse)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		createData := new(admin.CreateMenuData)
		createData.AdminId = int32(ids[0])
		createData.Actions = req.Actions

		err = h.adminService.CreateMenu(c, createData)

		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminMenuCreateError,
				code.Text(code.AdminMenuCreateError)).WithError(err),
			)
			return
		}
		res.Id = int32(ids[0])
		c.Payload(res)
	}
}
