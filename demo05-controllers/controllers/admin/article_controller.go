package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (ac ArticleController) Index(c *gin.Context) {
	c.String(http.StatusOK, "ADMIN/Article Page") // 访问 /admin/article 时触发
}

func (ac ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "ADMIN/Article/ADD Page") // 访问 /admin/article/add 时，输出 ADMIN/Article/ADD Page
}

func (ac ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "ADMIN/Article/EDIT Page") // 访问 /admin/article/edit 时，输出 ADMIN/Article/EDIT Page
}
