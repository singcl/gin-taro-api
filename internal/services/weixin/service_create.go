package weixin

import (
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/weixin"
)

type CreateWeixinUserData struct {
	Openid     string `json:"openid"`      // 微信Openid
	Unionid    string `json:"unionid"`     // 微信Unionid
	SessionKey string `json:"session_key"` // 微信SessionKey
}

func (s *service) Create(ctx core.Context, weixinUserData *CreateWeixinUserData) (id int32, err error) {
	model := weixin.NewModel()

	model.Openid = weixinUserData.Openid
	model.Unionid = weixinUserData.Unionid
	model.SessionKey = weixinUserData.SessionKey

	model.Nickname = configs.Get().Wechat.Nickname
	model.AvatarUrl = configs.Get().Wechat.AvatarUrl
	model.Mobile = configs.Get().Wechat.Mobile

	model.CreatedUser = ctx.SessionUserInfo().UserName
	model.IsUsed = 1
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
