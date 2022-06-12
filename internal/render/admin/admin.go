package admin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	Login() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	db     mysql.Repo
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_login", nil)
	}
}

func (h *handler) i() {}
