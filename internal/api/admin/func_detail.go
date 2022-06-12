package admin

import (
	"encoding/json"
	"net/http"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/pkg/password"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/admin"
)

type detailResponse struct {
	Username string                 `json:"username"` // 用户名
	Nickname string                 `json:"nickname"` // 昵称
	Mobile   string                 `json:"mobile"`   // 手机号
	Menu     []admin.ListMyMenuData `json:"menu"`     // 菜单栏
}

func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		res := new(detailResponse)

		searchOneData := new(admin.SearchOneData)
		searchOneData.Id = ctx.SessionUserInfo().UserID
		searchOneData.IsUsed = 1

		info, err := h.adminService.Detail(ctx, searchOneData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminDetailError,
				code.Text(code.AdminDetailError)).WithError(err),
			)
			return
		}

		menuCacheData, err := h.cache.Get(configs.RedisKeyPrefixLoginUser+password.GenerateLoginToken(searchOneData.Id)+":menu",
			redis.WithTrace(ctx.Trace()))
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminDetailError,
				code.Text(code.AdminDetailError)).WithError(err),
			)
			return
		}

		var menuData []admin.ListMyMenuData
		err = json.Unmarshal([]byte(menuCacheData), &menuData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminDetailError,
				code.Text(code.AdminDetailError)).WithError(err),
			)
			return
		}

		res.Username = info.Username
		res.Nickname = info.Nickname
		res.Mobile = info.Mobile
		res.Menu = menuData
		ctx.Payload(res)
	}
}
