package routers

import (
	"gin-demo/demo09-upload_file/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	adminRouter.GET("/uploadSingle", admin.AdminController{}.UploadSingleFile)
	adminRouter.POST("/uploadSingle/addFile", admin.AdminController{}.AddSingleFile) // 处理文件上传请求

	adminRouter.GET("/uploadMultiple", admin.AdminController{}.UploadMultipleFiles)
	adminRouter.POST("/uploadMultiple/addFiles", admin.AdminController{}.AddMultipleSameNameFiles)
}
