package ccnu

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (dc DefaultController) Index(c *gin.Context) {
	// 访问 / （首页）时，设置 session 的值，设置完成后必须保存 session 值
	session := sessions.Default(c)
	// 配置 session 的过期时间
	session.Options(sessions.Options{
		MaxAge: 60, // 设置 session 过期时间，单位为秒
	})
	session.Set("username", "ccnu")
	session.Save()

	c.String(http.StatusOK, "HOME Page")
}

func (dc DefaultController) News(c *gin.Context) {
	// 访问 /news 页面时，获取 session 的值
	session := sessions.Default(c)
	username := session.Get("username")

	c.String(http.StatusOK, "NEWS Page's Seesion Value: %s", username)
}

func (dc DefaultController) Shop(c *gin.Context) {
	c.String(http.StatusOK, "SHOP Page")
}
