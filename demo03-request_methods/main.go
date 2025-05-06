package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体
type USerInfo struct {
	Username string `json:"username" form:"username"` // Tag 标签用于给 ShouldBind 方法进行解析
	Password string `json:"password" form:"password"`
}

func main() {
	// 1. 创建一个默认的路由引擎
	router := gin.Default()

	// 加载 HTML 模板时，需要如下配置
	router.LoadHTMLGlob("templates/*")

	// 2. GET 请求传值
	// 2.1 获取 URL 传递的查询参数 QueryString
	// 若 URL 为 http://127.0.0.1:8080/?username=zhangsan，则 username 的值为 zhangsan
	router.GET("/", func(ctx *gin.Context) {
		username := ctx.Query("username")                                       // 获取 URL 传递的 QueryString 参数
		page := ctx.DefaultQuery("page", "1")                                   // 如果 QueryString 中没有名为 page 的参数，则 page 的默认值为 1
		ctx.String(http.StatusOK, "hello %s, Page Value is %s", username, page) // 输出 hello zhangsan, Page Value is 1
	})

	// 3. POST 请求传值
	// 3.1 渲染模板
	router.GET("/user", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "templates/user.html", gin.H{})
	})

	// 3.2 在 templates/user.html 模板中点击表单提交按钮时，会触发 /addUser 路由
	router.POST("/addUser", func(ctx *gin.Context) {
		username := ctx.PostForm("username")    // 获取表单中的 username 数据
		password := ctx.PostForm("password")    // 获取表单中的 password 数据
		age := ctx.DefaultPostForm("age", "18") // 获取表单中的 age 数据，如果表单中没有 age 数据，则使用默认值 18
		// 返回 JSON 类型的数据
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	// 4. 将 GET 和 POST 传递的值绑定到结构体中
	router.GET("/getUser", func(ctx *gin.Context) {
		user := &USerInfo{} // 因为要修改 user，所以要使用指针
		if err := ctx.ShouldBind(user); err == nil {
			// 若浏览器 URL 为 http://localhost:8080/getUser?username=11111&password=777
			fmt.Printf("user struct value: %#v", user) // 输出 user struct value: &main.USerInfo{Username:"11111", Password:"777"}
			ctx.JSON(http.StatusOK, user)              // 返回结构体数据
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	router.Run()
}
