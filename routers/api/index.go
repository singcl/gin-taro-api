package api

import (
	"github.com/gin-gonic/gin"
	"github.com/singcl/gin-taro-api/routers/api/auth"
	"github.com/singcl/gin-taro-api/routers/api/playground"
)

func InitApi(r *gin.Engine) *gin.RouterGroup {
	api := r.Group("/api")
	auth.UserInitRouter(api)
	playground.PlaygroundInitRouter(api)
	return api
}
