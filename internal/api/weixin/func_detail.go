package weixin

import (
	"net/http"

	"github.com/singcl/gin-taro-api/internal/code"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/services/weixin"
)

type detailResponse struct {
	Username  string `json:"username"`   // 用户名
	Nickname  string `json:"nickname"`   // 昵称
	Mobile    string `json:"mobile"`     // 手机号
	AvatarUrl string `json:"avatar_url"` // 头像
}

func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		res := new(detailResponse)

		searchOneData := new(weixin.SearchOneData)
		searchOneData.Openid = ctx.SessionWeixinUserInfo().Openid
		searchOneData.IsUsed = 1

		// 这里从数据库查询，也可以直接缓存中拿，缓存中已经保存一些用户信息
		info, err := h.weixinService.Detail(ctx, searchOneData)
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
		res.AvatarUrl = info.AvatarUrl
		ctx.Payload(res)
	}
}
