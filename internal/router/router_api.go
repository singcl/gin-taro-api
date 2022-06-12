package router

import (
	"github.com/singcl/gin-taro-api/internal/api/admin"
	"github.com/singcl/gin-taro-api/internal/api/helper"
	"github.com/singcl/gin-taro-api/internal/pkg/core"
)

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helpers := r.kiko.Group("helper")
	{
		helpers.GET("/md5/:str", helperHandler.Md5())
		helpers.POST("/sign", helperHandler.Sign())
	}

	// admin
	adminHandler := admin.New(r.logger, r.db, r.cache)

	// 需要签名验证，无需登录验证，无需 RBAC 权限验证
	login := r.kiko.Group("/api", r.interceptors.CheckSignature())
	{
		login.POST("/login", adminHandler.Login())
	}

	// 需要签名验证、登录验证，无需 RBAC 权限验证
	notRBAC := r.kiko.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature())
	{
		notRBAC.POST("/admin/logout", adminHandler.Logout())
		notRBAC.GET("/admin/info", adminHandler.Detail())
	}
}
