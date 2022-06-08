package main

import (
	"fmt"

	"github.com/singcl/gin-taro-api/configs"
	"github.com/singcl/gin-taro-api/routers"
)

func main() {
	r := routers.InitRouter()
	err := r.Run(configs.ProjectPort)
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
