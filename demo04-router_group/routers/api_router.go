package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(router *gin.Engine) {
	apiRouter := router.Group("/api")

	apiRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API Page") // 访问 /api 时，输出 API Page
	})

	apiRouter.GET("/userList", func(c *gin.Context) {
		c.String(http.StatusOK, "API/UserList Page") // 访问 /api/userList 时，输出 API/UserList Page
	})
}
