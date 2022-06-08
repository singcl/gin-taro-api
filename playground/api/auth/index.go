package auth

import "github.com/gin-gonic/gin"

func UserInitRouter(r *gin.RouterGroup) {
	userAuth := r.Group("/auth")
	UserAuthRouter(userAuth)
}
