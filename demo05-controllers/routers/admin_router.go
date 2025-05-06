package routers

import (
	"gin-demo/demo05-controllers/controllers/admin"

	"github.com/gin-gonic/gin"
)

// 形参 routet 来自于 router := gin.Default()
func AdminRouterInit(router *gin.Engine) {
	adminRouter := router.Group("/admin") // 路由分组

	adminRouter.GET("/", admin.IndexController{}.Index) // 访问 /admin 时触发

	adminRouter.GET("/user", admin.UserController{}.Index)     // 访问 /admin/user 时触发
	adminRouter.GET("/user/add", admin.UserController{}.Add)   // 访问 /admin/user/add 时触发
	adminRouter.GET("/user/edit", admin.UserController{}.Edit) // 访问 /admin/user/edit 时触发

	adminRouter.GET("/article", admin.ArticleController{}.Index)     // 访问 /admin/article 时触发
	adminRouter.GET("/article/add", admin.ArticleController{}.Add)   // 访问 /admin/article/add 时触发
	adminRouter.GET("/article/edit", admin.ArticleController{}.Edit) // 访问 /admin/article/edit 时触发
}
