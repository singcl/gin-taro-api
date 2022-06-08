package router

import "github.com/singcl/gin-taro-api/internal/api/helper"

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helpers := r.kiko.Group("helper")
	{
		helpers.GET("/md5/:str", helperHandler.Md5())
	}
}
