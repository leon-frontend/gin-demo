package main

import (
	"gin-demo/demo08-model/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 调用 models.UnixToTime() 方法将时间戳转换为日期格式
	router.GET("/", func(c *gin.Context) {
		c.String(200, "将时间戳转换为日期：%s", models.UnixToTime(1746606824))
	})

	router.Run()
}
