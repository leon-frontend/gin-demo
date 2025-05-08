package routers

import (
	"gin-demo/demo10-cookie/controllers/ccnu"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(router *gin.Engine) {
	// 路由分组
	defaultRouter := router.Group("/")
	defaultRouter.GET("/", ccnu.DefaultController{}.Index)
	defaultRouter.GET("/news", ccnu.DefaultController{}.News)
	defaultRouter.GET("/shop", ccnu.DefaultController{}.Shop)
	defaultRouter.GET("/deleteCookie", ccnu.DefaultController{}.DeleteCookie)
}
