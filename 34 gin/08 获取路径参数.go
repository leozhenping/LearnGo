package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	r := gin.Default()

	/*
		路径参数，对应前端访问URL格式 : /user/zhenping/19
	*/
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age": age,
		})

	})
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("start gin server failed. %v", err)
	}
}