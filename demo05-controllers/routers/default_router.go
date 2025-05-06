package routers

import (
	"gin-demo/demo05-controllers/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(router *gin.Engine) {
	defaultRouter := router.Group("/")

	defaultRouter.GET("/", itying.DefaultController{}.Index)    // 访问 / 时，输出 HOME
	defaultRouter.GET("/news", itying.DefaultController{}.News) // 访问 /news 时，输出 NEWS Page
}
