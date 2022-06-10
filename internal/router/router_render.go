package router

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/render/install"
)

func setRenderRouter(r *resource) {
	renderInstall := install.New(r.logger)

	// 无需记录日志，无需 RBAC 权限验证
	notRBAC := r.kiko.Group("", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 安装
		notRBAC.GET("install", renderInstall.View())
	}
}
