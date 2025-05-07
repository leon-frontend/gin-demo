package main

import (
	"fmt"
	"gin-demo/demo06-middleware/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	router := gin.Default()

	// 注册全局中间件
	router.Use(middleware.InitMiddleware)

	// 在 GET 函数参数中，第一个参数和最后一个参数之间的都是中间件处理函数
	router.GET("/", func(c *gin.Context) {
		fmt.Println("MiddlewareInit ouput: 3")
		time.Sleep(time.Second)
		c.String(http.StatusOK, "hello HOME!!!")
	})

	router.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "hello NEWS!!!")
	})

	router.GET("/mul", middleware.MiddlewareOne, middleware.MiddlewareTwo, func(c *gin.Context) {
		fmt.Println("/mul router ouput: hello MULTIPLE MIDDLEWARES!!!")
		c.String(http.StatusOK, "hello MULTIPLE MIDDLEWARES!!!")
	})

	router.Run(":8080")
}
