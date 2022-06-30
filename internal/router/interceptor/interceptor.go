package interceptor

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/proposal"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/internal/services/admin"
	"github.com/singcl/gin-taro-api/internal/services/authorized"
	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckSignature 验证签名是否合法，对用签名算法 pkg/signature
	CheckSignature() core.HandlerFunc

	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	// CheckWeixinLogin 验证是否微信登录
	CheckWeixinLogin(ctx core.Context) (info proposal.WeixinSessionUserInfo, err core.BusinessError)

	// CheckRBAC 验证 RBAC 权限是否合法
	CheckRBAC() core.HandlerFunc

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger            *zap.Logger
	cache             redis.Repo
	db                mysql.Repo
	authorizedService authorized.Service
	adminService      admin.Service
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo) Interceptor {
	return &interceptor{
		logger:            logger,
		cache:             cache,
		db:                db,
		authorizedService: authorized.New(db, cache),
		adminService:      admin.New(db, cache),
	}
}

func (i *interceptor) i() {}
