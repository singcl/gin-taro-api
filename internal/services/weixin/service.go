package weixin

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/mysql/weixin"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Login(ctx core.Context, searchCode2Session *SearchCode2SessionData) (info *Code2SessionData, err error)
	Detail(ctx core.Context, searchOneData *SearchOneData) (info *weixin.Weixin, err error)
	Create(ctx core.Context, weixinUserData *CreateWeixinUserData) (id int32, err error)
}

type service struct {
	db          mysql.Repo
	cache       redis.Repo
	wc          *wechat.Wechat
	miniprogram *miniprogram.MiniProgram
}

func New(db mysql.Repo, cache redis.Repo) Service {
	wc := wechat.NewWechat()
	wc.SetCache(cache)
	cfg := &miniConfig.Config{
		AppID:     configs.Get().Wechat.AppID,
		AppSecret: configs.Get().Wechat.Secret,
	}
	miniprogram := wc.GetMiniProgram(cfg)

	return &service{
		db:          db,
		cache:       cache,
		wc:          wc,
		miniprogram: miniprogram,
	}
}

func (s *service) i() {}
