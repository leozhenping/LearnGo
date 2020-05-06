package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "index get",
			})
		})
		videoGroup.GET("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "detail get",
			})
		})
	}

	// 定制404路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
	})

	r.Run(":9090")
}
