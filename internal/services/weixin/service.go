package weixin

import "github.com/singcl/gin-taro-api/internal/repository/redis"

var _ Service = (*service)(nil)

type Service interface {
	i()
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
