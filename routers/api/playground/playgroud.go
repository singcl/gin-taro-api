package playground

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaygroundInitRouter(r *gin.RouterGroup) {
	playground := r.Group("/playground")
	playgroundUrlParams(playground)
}

//
func playgroundUrlParams(r *gin.RouterGroup) {
	r.GET("/:param/*action", func(ctx *gin.Context) {
		param := ctx.Param("param")
		action := ctx.Param("action")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    [2]string{param, action},
		})
	})
}
