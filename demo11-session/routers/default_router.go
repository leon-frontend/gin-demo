package routers

import (
	"gin-demo/demo11-session/controllers/ccnu"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(router *gin.Engine) {
	// 路由分组
	defaultRouter := router.Group("/")
	defaultRouter.GET("/", ccnu.DefaultController{}.Index)
	defaultRouter.GET("/news", ccnu.DefaultController{}.News)
	defaultRouter.GET("/shop", ccnu.DefaultController{}.Shop)
}
