package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体
type Article struct {
	Title   string `json:"title"` // json:"title" 表示在 JSON 数据中，该字段的名称为 title (小写)
	Content string
}

func main() {
	// 1. 创建一个默认的路由引擎
	router := gin.Default()

	// 2. 渲染 HTML 模板时，告诉 gin 去 templates 文件夹加载模板文件
	router.LoadHTMLGlob("templates/*")

	// 2.1 绑定路由规则和回调函数
	router.GET("/", func(ctx *gin.Context) {
		// String 接收三个参数：1. 状态码 2. 格式化模板 3. 模板参数
		ctx.String(200, "Value: %v", "hello GIN!!!!hhh")
	})

	// 2.2 返回 JSON 数据
	router.GET("/json", func(ctx *gin.Context) {
		// gin.H 是 map[string]interface{} 的简写
		ctx.JSON(200, gin.H{
			"name": "Gin",
			"age":  18,
		})
	})

	// 2.3 返回结构体数据
	router.GET("/struct", func(ctx *gin.Context) {
		// 实例化 Artical 结构体
		article := Article{Title: "Gin", Content: "Gin is a web framework written in Go"}
		ctx.JSON(200, article)
	})

	// 2.4 响应 JSONP 请求（主要用于解决跨域问题）
	// 若前端的 URL 为 http://example.com/jsonp?callback=cbFn，则会将 Map 数据作为参数传递给 cbFn，输出 cbFn({"name":"Gin"})
	router.GET("/jsonp", func(ctx *gin.Context) {
		// 返回 JSONP 数据
		ctx.JSONP(200, gin.H{
			"name": "Gin",
		})
	})

	// 2.5 响应 XML 数据
	router.GET("/xml", func(ctx *gin.Context) {
		// 返回 XML 数据
		ctx.XML(200, gin.H{
			"name": "Gin",
		})
	})

	// 2.6 渲染 HTML 模板
	router.GET("/news", func(ctx *gin.Context) {

		/*
			使用 ctx.HTML() 函数渲染模板
			第二个参数是模板名称，应该和 templates 文件夹下的模板文件名保持一致
			第三个参数中的数据会传给模板，模板中可以通过 {{.title}} 获取数据
		*/
		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后端的 Title 数据",
		})
	})

	// 3. 启动路由，并监听默认的 8080 端口
	router.Run(":8080")
}
