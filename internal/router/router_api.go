package router

import (
	"github.com/singcl/gin-taro-api/internal/api/admin"
	"github.com/singcl/gin-taro-api/internal/api/authorized"
	"github.com/singcl/gin-taro-api/internal/api/helper"
	"github.com/singcl/gin-taro-api/internal/api/menu"
	"github.com/singcl/gin-taro-api/internal/api/weixin"
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

	//////////////////////////////////////////////////////////////////////////

	// 微信
	weixinHandler := weixin.New(r.logger, r.db, r.cache)
	weixin := r.kiko.Group("/weixin")
	{
		// 微信登录
		// 无需签名验证，无需登录验证，无需 RBAC 权限验证
		weixin.GET("/login", weixinHandler.Login())
	}

	// 无需签名验证、无需 RBAC 权限验证
	// 需要登录验证
	weixinCheckLogin := r.kiko.Group("/weixin", core.WrapWeixinAuthHandler(r.interceptors.CheckWeixinLogin))
	{
		// weixinCheckLogin.POST("/auth/logout", weixinHandler.Logout())
		weixinCheckLogin.GET("/auth/info", weixinHandler.Detail())
		weixinCheckLogin.POST("/auth/avatar", weixinHandler.Avatar())
	}

	/////////////////////////////////////////////////////////////////////////

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
		notRBAC.PATCH("/admin/modify_password", adminHandler.ModifyPassword())
		notRBAC.PATCH("/admin/modify_personal_info", adminHandler.ModifyPersonalInfo())
	}

	// 需要签名验证、登录验证、RBAC 权限验证
	api := r.kiko.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin), r.interceptors.CheckSignature(), r.interceptors.CheckRBAC())
	{
		// authorized
		authorizedHandler := authorized.New(r.logger, r.db, r.cache)
		api.POST("/authorized", authorizedHandler.Create())
		api.GET("/authorized", authorizedHandler.List())
		api.PATCH("/authorized/used", authorizedHandler.UpdateUsed())
		api.DELETE("/authorized/:id", core.AliasForRecordMetrics("/api/authorized/info"), authorizedHandler.Delete())

		api.GET("/authorized_api", authorizedHandler.ListAPI())
		api.POST("/authorized_api", authorizedHandler.CreateAPI())
		api.DELETE("/authorized_api/:id", core.AliasForRecordMetrics("/api/authorized_api/info"), authorizedHandler.DeleteAPI())

		// admin
		api.GET("/admin", adminHandler.List())
		api.POST("/admin", adminHandler.Create())
		api.DELETE("/admin/:id", core.AliasForRecordMetrics("/api/admin"), adminHandler.Delete())
		api.PATCH("/admin/used", adminHandler.UpdateUsed())
		api.PATCH("/admin/offline", adminHandler.Offline())
		api.PATCH("/admin/reset_password/:id", core.AliasForRecordMetrics("/api/admin/reset_password"), adminHandler.ResetPassword())

		api.POST("/admin/menu", adminHandler.CreateAdminMenu())
		api.GET("/admin/menu/:id", core.AliasForRecordMetrics("/api/admin/menu"), adminHandler.ListAdminMenu())

		// menu
		menuHandler := menu.New(r.logger, r.db, r.cache)
		api.POST("/menu", menuHandler.Create())
		api.GET("/menu", menuHandler.List())
	}
}
