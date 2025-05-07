package main

import (
	"gin-demo/demo07-middlewares_group/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 注册路由分组中的路由
	routers.AdminRouterInit(router)
	routers.ApiRouterInit(router)
	routers.DefaultRouterInit(router)

	router.Run()
}
