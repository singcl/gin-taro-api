package api

import (
	"github.com/gin-gonic/gin"
	"github.com/singcl/gin-taro-api/routers/api/auth"
)

func InitApi(r *gin.Engine) *gin.RouterGroup {
	api := r.Group("/api")
	auth.UserInitRouter(api)
	return api
}
