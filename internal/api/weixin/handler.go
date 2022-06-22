package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/weixin"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// login 微信登录
	// @Tags weixin
	// @Router /weixin/login [get]
	Login() core.HandlerFunc
}

type handler struct {
	logger        *zap.Logger
	db            mysql.Repo
	cache         redis.Repo
	weixinService weixin.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:        logger,
		db:            db,
		cache:         cache,
		weixinService: weixin.New(db, cache),
	}
}

func (*handler) i() {}
