package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (ic IndexController) Index(c *gin.Context) {
	c.String(http.StatusOK, "ADMIN Page") // 访问 /admin 时，输出 ADMIN Page
}
