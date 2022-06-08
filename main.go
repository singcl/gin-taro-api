package main

import (
	"net/http"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/internal/router"
)

func main() {
	// r := routers.InitRouter()
	// err := r.Run(configs.ProjectPort)
	// if err != nil {
	// 	fmt.Println("服务器启动失败！")
	// }

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer()
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Kiko,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// logger
		}
	}()
}
