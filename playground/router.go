package playground

import (
	"github.com/gin-gonic/gin"
	"github.com/singcl/gin-taro-api/playground/api"
)

// 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// api路由
	api.InitApi(r)
	return r
}
