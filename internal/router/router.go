package router

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"go.uber.org/zap"
)

type resource struct {
	kiko   core.Kiko
	logger *zap.Logger
	db     mysql.Repo
	cache  redis.Repo
}

type Server struct {
	Kiko  core.Kiko
	Db    mysql.Repo
	Cache redis.Repo
}

func NewHTTPServer() (*Server, error) {
	r := new(resource)

	kiko, err := core.New()

	if err != nil {
		panic(err)
	}

	r.kiko = kiko

	// 设置API路由
	setApiRouter(r)

	s := new(Server)
	s.Kiko = kiko

	return s, nil
}
