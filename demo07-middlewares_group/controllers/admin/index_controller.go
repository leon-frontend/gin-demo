package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (ic IndexController) Index(c *gin.Context) {
	// 获取中间件中的数据
	username, _ := c.Get("username") // 返回的是 any 类型
	fmt.Println("中间件中的数据 username:", username)

	usernameStr, _ := username.(string)
	c.String(http.StatusOK, "ADMIN Page -- "+usernameStr) // 访问 /admin 时，输出 ADMIN Page
}
