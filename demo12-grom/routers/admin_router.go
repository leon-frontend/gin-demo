package routers

import (
	"gin-demo/demo12-grom/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(router *gin.Engine) {
	// 添加路由分组
	adminGroup := router.Group("/admin")

	// 添加路由
	adminGroup.GET("/user", admin.UserController{}.Index)
	adminGroup.GET("/user/add", admin.UserController{}.Add)
	adminGroup.GET("/user/update", admin.UserController{}.Update)
	adminGroup.GET("/user/delete", admin.UserController{}.Delete)

	adminGroup.GET("/article", admin.ArticleController{}.Index)

	adminGroup.GET("/student", admin.StudentController{}.Index)

	adminGroup.GET("/bank", admin.BankController{}.Index)
}
