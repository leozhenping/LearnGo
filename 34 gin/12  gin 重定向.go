package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")  // 外部重定向。

	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"  // 修改url
		r.HandleContext(c) // 继续执行接下来的请求
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg" : "b",
		})

	})
	r.Run(":9090")
}
