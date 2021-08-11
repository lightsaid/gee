package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	// 注册路由，有访问进来执行回调
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h5> hello 你好</h5>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "你好 %s, 当前访问%s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9527")
}
