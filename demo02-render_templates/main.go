package main

import (
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

// UnixToTime 函数会将时间戳转换为年月日时分秒格式
func UnixToTime(timestamp int) string {
	// 将一个 Unix 时间戳（以秒为单位）转换为 Go 语言中的 time.Time 类型对象。
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	// 1. 创建一个默认的路由引擎
	router := gin.Default()

	// 自定义模板函数，需要将这个操作放在加载模板 LoadHTMLGlob 函数之前
	router.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime, // 注册自定义模板函数，然后在模板中直接使用即可
	})

	// 配置静态文件服务，第一个参数是网站访问的 URL；第二个参数是匹配的静态文件目录路径
	// 当网站的 URL 访问 /static 时，会去 static 文件夹中寻找匹配的文件
	router.Static("/static", "./static")

	// 2. 渲染 HTML 模板时，需要配置模板目录
	// 加载多级模板文件目录：/** 表示加载 templates 文件夹下的子目录
	router.LoadHTMLGlob("templates/**/*")

	// 3. 渲染 HTML 模板
	router.GET("/", func(ctx *gin.Context) {
		// 加载 templates 文件夹下的单个模板文件，会覆盖之前加载的模板集合
		router.LoadHTMLFiles("templates/index.html", "templates/public/header.html")
		ctx.HTML(200, "index.html", gin.H{
			"title":     "HOME",
			"score":     89,
			"hobby":     []string{"football", "basketball", "swimming"},
			"news":      &Article{Content: "this is a news", Title: "news"},
			"timestamp": 1746434101,
		})
	})

	// 渲染 templates/news.html 模板文件
	router.GET("/news", func(ctx *gin.Context) {
		// 加载 templates 文件夹下的单个模板文件
		router.LoadHTMLFiles("templates/news.html")
		ctx.HTML(200, "news.html", gin.H{
			"title":   "NEWS",
			"content": Article{Title: "我是标题", Content: "我是内容"},
		})
	})

	// 渲染 templates/admin/news.html 模板文件
	router.GET("/admin/news", func(ctx *gin.Context) {
		ctx.HTML(200, "admin/news.html", gin.H{
			"title": "ADMIN/NEWS",
		})
	})

	router.Run()
}
