package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (a ApiController) Index(c *gin.Context) {
	c.String(http.StatusOK, "API Page") // 访问 /api 时，输出 API Page
}

func (a ApiController) UserList(c *gin.Context) {
	c.String(http.StatusOK, "API/UserList Page") // 访问 /api/userList 时，输出 API/UserList Page
}
