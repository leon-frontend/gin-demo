package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println("InitMiddleware ouput: 11111111")

	// 在中间件中设置数据
	c.Set("username", "zhangsan")

	// 定义一个 goroutine 统计日志
	// Gin 的主线程不会立即结束请求，而是会等所有中间件和处理函数执行完毕。
	// go func() 启动的是一个新的 goroutine，它的生命周期不是依附在主 goroutine 上的，而是由整个 Go 程序的生命周期控制。
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		// 这里使用你创建的副本
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
