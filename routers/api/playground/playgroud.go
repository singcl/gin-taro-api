package playground

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type User struct {
// 	name   string `json:name`
// 	account int64  `json:account`
// }

func PlaygroundInitRouter(r *gin.RouterGroup) {
	playground := r.Group("/playground")
	playgroundUrlParams(playground)
	playgroundUrlQuery(playground)
	playgroundPost(playground)
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

//
func playgroundPost(r *gin.RouterGroup) {
	r.POST("/query", func(c *gin.Context) {
		/* // 获取表单参数
		message := c.PostForm("userName")              // 表单参数
		nick := c.DefaultPostForm("account", "123456") // 此方法可以设置默认值，和上面的get一样 */

		// 获取body中的参数方式一
		json := make(map[string]interface{}) //注意该结构接受的内容
		c.BindJSON(&json)
		log.Printf("%v", &json)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    json,
		})

		// // 获取body中的参数方式二
		// json := User{}
		// c.BindJSON(&json)
	})
}
