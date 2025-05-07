package routers

import (
	"gin-demo/demo07-middlewares_group/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(router *gin.Engine) {
	apiRouter := router.Group("/api")
	apiRouter.GET("/", api.ApiController{}.Index)            // 访问 /api 时触发
	apiRouter.GET("/userList", api.ApiController{}.UserList) // 访问 /api/userList 时触发
}
