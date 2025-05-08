package ccnu

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (dc DefaultController) Index(c *gin.Context) {
	/*
		访问 HOME 页面时，使用 c.SetCookie() 函数设置 Cookie。该函数的参数解释如下：
			第一个参数：Cookie 名称；第二个参数：Cookie 值；
			第三个参数：Cookie 的过期时间，单位为秒；
			第四个参数：可以访问 cookie 的路径；第五个参数：可以访问 cookie 的域名；
			第六个参数：是否只能 HTTPS 协议访问，HTTP 协议无法访问；
			第七个参数：是否允许 JavaScript 访问 Cookie，设置为 false 时，可以防止 XSS 攻击。
	*/
	c.SetCookie("username", "ccnu", 3600, "/", "localhost", false, true)
	c.String(http.StatusOK, "HOME Page")
}

func (dc DefaultController) News(c *gin.Context) {
	// 在 /news 页面中获取 Cookie
	username, _ := c.Cookie("username")
	c.String(http.StatusOK, "NEWS Page's Cookie Value: %s", username)
}

func (dc DefaultController) Shop(c *gin.Context) {
	username, _ := c.Cookie("username")
	c.String(http.StatusOK, "SHOP Page's Cookie Value: %s", username)
}

func (dc DefaultController) DeleteCookie(c *gin.Context) {
	// 当 Cookie 的有效时间设置为小于 0 的值时，表示删除 Cookie
	c.SetCookie("username", "ccnu", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "DELETE Cookie Success!!!!!")
}
