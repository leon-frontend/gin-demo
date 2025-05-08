package main

import (
	"gin-demo/demo11-session/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 配置 session 中间件
	/*
		创建一个基于 Cookie 存储方式的 session 存储引擎。
			[]byte("secret")：这是一个加密密钥，用于对 Cookie 中存储的 session 数据进行加密和验证签名，防止用户伪造 session 数据。
			cookie.NewStore(...)：用于创建一个 sessions.Store 类型的实例。
		工作机制：
			session 数据会被存储在客户端浏览器的 cookie 中。
			这个数据会被加密、签名，服务器通过这个 "secret" 验证数据是否被篡改。
			如果更改了这个 secret，原有的 cookie 数据就无法解密，用户将被视为未登录或 session 失效。
	*/
	store := cookie.NewStore([]byte("secret"))

	// 将 session 中间件全局注册到 Gin 路由中
	router.Use(sessions.Sessions("mysession", store))

	routers.DefaultRouterInit(router)

	router.Run()
}
