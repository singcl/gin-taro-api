package authorized

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"go.uber.org/zap"
)

type handler struct {
	db     mysql.Repo
	logger *zap.Logger
	cache  redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_list", nil)
	}
}

func (h *handler) Add() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("authorized_add", nil)
	}
}
