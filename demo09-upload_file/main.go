package main

import (
	"gin-demo/demo09-upload_file/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 加载模版
	router.LoadHTMLGlob("templates/*")

	// 注册路由分组
	routers.AdminRouterInit(router)

	router.Run()
}
