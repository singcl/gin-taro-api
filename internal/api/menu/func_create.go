package menu

import (
	"net/http"

	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
)

type createRequest struct {
	Id    string `form:"id"`    // ID
	Pid   int32  `form:"pid"`   // 父类ID
	Name  string `form:"name"`  // 菜单名称
	Link  string `form:"link"`  // 链接地址
	Icon  string `form:"icon"`  // 图标
	Level int32  `form:"level"` // 菜单类型 1:一级菜单 2:二级菜单
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 创建/编辑菜单
// @Summary 创建/编辑菜单
// @Description 创建/编辑菜单
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu [post]
// @Security LoginToken
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		// res := new(createResponse)

		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
	}
}
