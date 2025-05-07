package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController // 使用嵌套结构体实现继承
}

func (uc UserController) Index(c *gin.Context) {
	uc.success(c)
	c.String(http.StatusOK, "ADMIN/USER Page") // 访问 /admin/user 时，输出 ADMIN/USER Page
}

func (uc UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "ADMIN/USER/ADD Page") // 访问 /admin/user/add 时，输出 ADMIN/USER Page
}

func (uc UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "ADMIN/USER/EDIT Page") // 访问 /admin/user/edit 时，输出 ADMIN/USER Page
}
