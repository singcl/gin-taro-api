package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// https://gin-gonic.com/zh-cn/docs/examples/ascii-json/
	r.GET("/asciiJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<button/>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})
	r.Run() // // 监听并在 0.0.0.0:8080 上启动服务
}
