package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(router *gin.Engine) {
	defaultRouter := router.Group("/")

	defaultRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "HOME") // 访问 / 时，输出 HOME
	})

	defaultRouter.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "NEWS Page") // 访问 /news 时，输出 NEWS Page
	})
}
