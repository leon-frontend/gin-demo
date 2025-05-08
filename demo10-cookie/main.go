package main

import (
	"gin-demo/demo10-cookie/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 注册路由分组
	routers.DefaultRouterInit(router)

	router.Run()
}
