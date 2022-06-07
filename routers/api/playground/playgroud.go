package playground

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlQuery struct {
	name, email string
}

func PlaygroundInitRouter(r *gin.RouterGroup) {
	playground := r.Group("/playground")
	playgroundUrlParams(playground)
	playgroundUrlQuery(playground)
}

//  /playground/:param/ *action  形式的url
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

//  /playground?name=name&email=email  形式的url
func playgroundUrlQuery(r *gin.RouterGroup) {
	r.GET("query", func(ctx *gin.Context) {
		// 使用该方法获取query参数时需要设定一个默认值
		// 如果query参数有值则返回，无值则返回设置的默认值
		name := ctx.DefaultQuery("name", "singcl")
		// 是 c.Request.URL.Query().Get("email") 的简写
		email := ctx.Query("email")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    []string{name, email},
		})
	})
}
