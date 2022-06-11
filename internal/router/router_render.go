package router

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/render/index"
	"github.com/singcl/gin-taro-api/internal/render/install"
)

func setRenderRouter(r *resource) {
	renderInstall := install.New(r.logger)
	renderIndex := index.New(r.logger, r.db, r.cache)

	// 无需记录日志，无需 RBAC 权限验证
	notRBAC := r.kiko.Group("", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 首页
		notRBAC.GET("", renderIndex.Index())

		// 安装
		notRBAC.GET("install", renderInstall.View())
		notRBAC.POST("/install/execute", renderInstall.Execute())
	}
}
