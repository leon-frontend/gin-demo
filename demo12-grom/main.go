package main

import (
	"gin-demo/demo12-grom/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 注册路由分组
	routers.AdminRouterInit(router)

	router.Run()
}
