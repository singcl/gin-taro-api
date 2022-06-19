package weixin

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Login(ctx core.Context, searchCode2Session *SearchCode2SessionData) (info *Code2SessionData, err error)
}

type service struct {
	cache       redis.Repo
	wc          *wechat.Wechat
	miniprogram *miniprogram.MiniProgram
}

func New(cache redis.Repo) Service {
	wc := wechat.NewWechat()
	// TODO wechat cache interface
	wc.SetCache(cache)
	cfg := &miniConfig.Config{
		AppID:     configs.Get().Wechat.AppID,
		AppSecret: configs.Get().Wechat.Secret,
	}
	miniprogram := wc.GetMiniProgram(cfg)

	return &service{
		cache:       cache,
		wc:          wc,
		miniprogram: miniprogram,
	}
}

func (s *service) i() {}
