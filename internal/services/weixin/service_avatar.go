package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/weixin"
	"gorm.io/gorm"
)

func (s *service) Avatar(ctx core.Context, avatarUrl string, openid string) (err error) {
	_, err = weixin.NewQueryBuilder().
		WhereIsDeleted(mysql.EqualPredicate, -1).
		WhereOpenid(mysql.EqualPredicate, openid).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		// 更新时候没有找到该条记录时，依然返回更新成功 ？
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	data := map[string]interface{}{
		"avatar_url": avatarUrl,
		// TODO: 登录成功redis存储Username
		// "updated_user": ctx.SessionWeixinUserInfo().Username
	}

	qb := weixin.NewQueryBuilder()
	qb.WhereOpenid(mysql.EqualPredicate, openid)

	err = qb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}
	return
}
