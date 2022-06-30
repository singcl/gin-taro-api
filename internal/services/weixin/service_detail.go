package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/weixin"
)

type SearchOneData struct {
	Id         int32  // 用户ID
	Openid     string // 微信Openid
	Unionid    string // 微信Unionid
	SessionKey string // 微信SessionKey
	Username   string // 用户名
	Nickname   string // 昵称
	AvatarUrl  string // 头像
	Mobile     string // 手机号
	IsUsed     int32  // 是否启用 1:是  -1:否
}

func (s *service) Detail(ctx core.Context, searchOneData *SearchOneData) (info *weixin.Weixin, err error) {
	qb := weixin.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	if searchOneData.Openid != "" {
		qb.WhereOpenid(mysql.EqualPredicate, searchOneData.Openid)
	}
	if searchOneData.Unionid != "" {
		qb.WhereUnionid(mysql.EqualPredicate, searchOneData.Unionid)
	}
	if searchOneData.Username != "" {
		qb.WhereUsername(mysql.EqualPredicate, searchOneData.Username)
	}

	if searchOneData.Nickname != "" {
		qb.WhereNickname(mysql.EqualPredicate, searchOneData.Nickname)
	}
	if searchOneData.AvatarUrl != "" {
		qb.WhereAvatarUrl(mysql.EqualPredicate, searchOneData.AvatarUrl)
	}

	if searchOneData.Mobile != "" {
		qb.WhereMobile(mysql.EqualPredicate, searchOneData.Mobile)
	}

	if searchOneData.IsUsed != 0 {
		qb.WhereIsUsed(mysql.EqualPredicate, searchOneData.IsUsed)
	}

	info, err = qb.QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
