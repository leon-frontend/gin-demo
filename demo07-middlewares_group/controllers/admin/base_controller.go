package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (bs BaseController) success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func (bs BaseController) error(c *gin.Context) {
	c.String(http.StatusOK, "error")
}
