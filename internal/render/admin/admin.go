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
	List() core.HandlerFunc
	ModifyPassword() core.HandlerFunc
	ModifyInfo() core.HandlerFunc
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

func (h *handler) i() {}

func (h *handler) Login() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_login", nil)
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_list", nil)
	}
}

func (h *handler) ModifyPassword() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_modify_password", nil)
	}
}

func (h *handler) ModifyInfo() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("admin_modify_info", nil)
	}
}
