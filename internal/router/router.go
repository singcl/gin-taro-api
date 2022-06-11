package router

import (
	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/alert"
	"github.com/singcl/gin-taro-api/internal/metrics"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"github.com/singcl/gin-taro-api/pkg/errors"
	"github.com/singcl/gin-taro-api/pkg/file"
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

func NewHTTPServer(logger *zap.Logger, cronLogger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}
	r := new(resource)
	r.logger = logger

	openBrowserUri := configs.ProjectDomain + configs.ProjectPort

	_, ok := file.IsExists(configs.ProjectInstallMark)

	if !ok {
		// 未安装
		openBrowserUri += "/install"
	} else {
		// 已安装

		// 初始化 DB
		dbRepo, err := mysql.New()
		if err != nil {
			logger.Fatal("new db err", zap.Error(err))
		}
		r.db = dbRepo

		// 初始化 Cache
		cacheRepo, err := redis.New()
		if err != nil {
			logger.Fatal("new cache err", zap.Error(err))
		}
		r.cache = cacheRepo

		// // 初始化 CRON Server
		// cronServer, err := cron.New(cronLogger, dbRepo, cacheRepo)
		// if err != nil {
		// 	logger.Fatal("new cron err", zap.Error(err))
		// }
		// cronServer.Start()
		// r.cronServer = cronServer
	}

	kiko, err := core.New(logger,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithAlertNotify(alert.NotifyHandler(logger)),
		core.WithRecordMetrics(metrics.RecordHandler(logger)),
	)

	if err != nil {
		panic(err)
	}

	r.kiko = kiko

	// 设置 Render 路由
	setRenderRouter(r)

	// 设置API路由
	setApiRouter(r)

	s := new(Server)
	s.Kiko = kiko

	return s, nil
}
