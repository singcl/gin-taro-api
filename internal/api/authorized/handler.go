package authorized

import (
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/authorized"
	"github.com/singcl/gin-taro-api/pkg/hash"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增调用方
	// @Tags API.authorized
	// @Router /api/authorized [post]
	Create() core.HandlerFunc

	// List 调用方列表
	// @Tags API.authorized
	// @Router /api/authorized [get]
	List() core.HandlerFunc

	// UpdateUsed 更新调用方为启用/禁用
	// @Tags API.authorized
	// @Router /api/authorized/used [patch]
	UpdateUsed() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	cache             redis.Repo
	hashids           hash.Hash
	authorizedService authorized.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		authorizedService: authorized.New(db, cache),
		hashids:           hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
