package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义路由中间件
func InitMiddleware(c *gin.Context) {
	start := time.Now().UnixNano() // 记录当前时间，单位是纳秒
	fmt.Println("MiddlewareInit ouput: 1")

	// c.Next() 执行后续中间件处理函数（包括最后一个函数），等待后续中间件处理函数执行完毕后，再执行当前中间件处理函数后面的代码
	c.Next()

	fmt.Println("MiddlewareInit ouput: 2")
	end := time.Now().UnixNano() // 记录当前时间，单位是纳秒

	fmt.Printf("请求处理前到请求处理后的执行时间: %v 纳秒\n", end-start) // 请求处理前到请求处理后的执行时间: 1000371100 纳秒
}

func MiddlewareOne(c *gin.Context) {
	fmt.Println("MiddlewareOne ouput: One-1")
	c.Next()
	fmt.Println("MiddlewareOne ouput: One-2")
}

func MiddlewareTwo(c *gin.Context) {
	fmt.Println("MiddlewareTwo ouput: Two-1")
	c.Next()
	fmt.Println("MiddlewareTwo ouput: Two-2")
}
