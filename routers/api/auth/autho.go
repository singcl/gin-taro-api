package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuthRouter(g *gin.RouterGroup) {
	g.GET("/login", func(c *gin.Context) {
		c.JSON(
			http.StatusOK, gin.H{
				"message": "success",
				"data":    "登陆成功",
			},
		)
	})
}
