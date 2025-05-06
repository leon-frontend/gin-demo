package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 形参 routet 来自于 router := gin.Default()
func AdminRouterInit(router *gin.Engine) {
	adminRouter := router.Group("/admin") // 路由分组

	adminRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ADMIN Page") // 访问 /admin 时，输出 ADMIN Page
	})

	adminRouter.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "ADMIN/USER Page") // 访问 /admin/user 时，输出 ADMIN/USER Page
	})

	adminRouter.GET("/article", func(c *gin.Context) {
		c.String(http.StatusOK, "ADMIN/Article Page") // 访问 /admin/article 时，输出 ADMIN/Article Page
	})
}
