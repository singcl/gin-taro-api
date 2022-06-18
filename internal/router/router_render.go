package router

import (
	"github.com/singcl/gin-taro-api/internal/pkg/core"
	"github.com/singcl/gin-taro-api/internal/render/admin"
	"github.com/singcl/gin-taro-api/internal/render/authorized"
	"github.com/singcl/gin-taro-api/internal/render/dashboard"
	"github.com/singcl/gin-taro-api/internal/render/index"
	"github.com/singcl/gin-taro-api/internal/render/install"
)

func setRenderRouter(r *resource) {
	renderInstall := install.New(r.logger)
	renderIndex := index.New(r.logger, r.db, r.cache)
	renderDashboard := dashboard.New(r.logger, r.db, r.cache)
	renderAdmin := admin.New(r.logger, r.db, r.cache)
	renderAuthorized := authorized.New(r.logger, r.db, r.cache)

	// 无需记录日志，无需 RBAC 权限验证
	notRBAC := r.kiko.Group("", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 首页
		notRBAC.GET("", renderIndex.Index())

		// 仪表盘
		notRBAC.GET("/dashboard", renderDashboard.View())

		// 安装
		notRBAC.GET("install", renderInstall.View())
		notRBAC.POST("/install/execute", renderInstall.Execute())

		// 管理员
		notRBAC.GET("/login", renderAdmin.Login())
	}

	// 无需记录日志，需要 RBAC 权限验证
	render := r.kiko.Group("", core.DisableTraceLog, core.DisableRecordMetrics)
	{
		// 调用方
		render.GET("/authorized/list", renderAuthorized.List())
		// render.GET("/authorized/add", renderAuthorized.Add()) // 新增不用单独页面，弹窗即可
	}
}
