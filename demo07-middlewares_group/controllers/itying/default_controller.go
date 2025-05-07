package itying

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (d DefaultController) Index(c *gin.Context) {
	c.String(http.StatusOK, "HOME") // 访问 / 时，输出 HOME
}

func (d DefaultController) News(c *gin.Context) {
	c.String(http.StatusOK, "NEWS Page") // 访问 /news 时，输出 NEWS Page
}
