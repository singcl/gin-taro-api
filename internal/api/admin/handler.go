package admin

import (
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/admin"
	"github.com/singcl/gin-taro-api/pkg/hash"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Login 管理员登录
	// @Tags API.admin
	// @Router /api/login [post]
	Login() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	cache        redis.Repo
	hashids      hash.Hash
	adminService admin.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:       logger,
		cache:        cache,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		adminService: admin.New(db, cache),
	}
}

func (h *handler) i() {}