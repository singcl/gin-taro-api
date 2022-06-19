package weixin

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	Login(ctx core.Context, searchCode2Session *SearchCode2SessionData) (info *Code2SessionData, err error)
}

type service struct {
	cache redis.Repo
}

func New(cache redis.Repo) Service {
	return &service{
		cache: cache,
	}
}

func (s *service) i() {}
