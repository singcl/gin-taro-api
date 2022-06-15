package dashboard

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/repository/mysql"
	"github.com/singcl/gin-taro-api/internal/repository/redis"
	"go.uber.org/zap"
)

type handler struct {
	logger *zap.Logger
	cache  redis.Repo
	db     mysql.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) *handler {
	return &handler{
		logger: logger,
		db:     db,
		cache:  cache,
	}
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (h *handler) View() core.HandlerFunc {
	type mysqlVersion struct {
		Ver string
	}

	mysqlVer := new(mysqlVersion)
	if h.db != nil {
		h.db.GetDbR().Raw("SELECT version() as ver").Scan(mysqlVer)
	}

	redisVer := ""
	if h.cache != nil {
		redisVer = h.cache.Version()
	}

	type viewResponse struct {
		MemTotal       string
		MemUsed        string
		MemUsedPercent float64

		DiskTotal       string
		DiskUsed        string
		DiskUsedPercent float64

		HostOS   string
		HostName string

		CpuName        string
		CpuCores       int32
		CpuUsedPercent float64

		GoPath      string
		GoVersion   string
		Goroutine   int
		ProjectPath string
		Env         string
		Host        string
		GoOS        string
		GoArch      string

		ProjectVersion string
		MySQLVersion   string
		RedisVersion   string
	}

	return func(ctx core.Context) {
		obj := new(viewResponse)

		obj.MySQLVersion = mysqlVer.Ver
		obj.RedisVersion = redisVer

		ctx.HTML("dashboard", obj)
	}
}
