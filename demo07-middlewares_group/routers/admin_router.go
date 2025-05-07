package routers

import (
	"gin-demo/demo07-middlewares_group/controllers/admin"
	"gin-demo/demo07-middlewares_group/middlewares"

	"github.com/gin-gonic/gin"
)

// 形参 routet 来自于 router := gin.Default()
func AdminRouterInit(router *gin.Engine) {
	// 路由分组，并给路由配置中间件
	adminRouter := router.Group("/admin", middlewares.InitMiddleware)

	adminRouter.GET("/", admin.IndexController{}.Index) // 访问 /admin 时触发

	adminRouter.GET("/user", admin.UserController{}.Index)     // 访问 /admin/user 时触发
	adminRouter.GET("/user/add", admin.UserController{}.Add)   // 访问 /admin/user/add 时触发
	adminRouter.GET("/user/edit", admin.UserController{}.Edit) // 访问 /admin/user/edit 时触发

	adminRouter.GET("/article", admin.ArticleController{}.Index)     // 访问 /admin/article 时触发
	adminRouter.GET("/article/add", admin.ArticleController{}.Add)   // 访问 /admin/article/add 时触发
	adminRouter.GET("/article/edit", admin.ArticleController{}.Edit) // 访问 /admin/article/edit 时触发
}
