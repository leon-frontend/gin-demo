package main

import (
	"gin-demo/demo04-router_group/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 路由分组 1: 默认路由组
	routers.DefaultRouterInit(router) // 初始化默认路由组

	// 路由分组 2: api 路由组
	routers.ApiRouterInit(router) // 初始化 api 路由组

	// 路由分组 3: admin 路由组
	routers.AdminRouterInit(router) // 初始化 admin 路由组

	router.Run(":8080")
}
