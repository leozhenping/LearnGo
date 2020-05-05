package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name   string `json:"name"`  // 返回给前端字段时，取tag中的字段。
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	r := gin.Default()

	// 1. 使用gin.H返回Map类型
	r.GET("/gin_h", func(c *gin.Context) {
		data := gin.H{
			"name":   "zhenping",
			"age":    18,
			"gender": "male",
		}
		c.JSON(http.StatusOK, data)
	})

	// 2. 使用struct结构体返回
	r.GET("/gin_struct", func(c *gin.Context) {
		data := Student{
			Name:   "zhenping",
			Age:    19,
			Gender: "male",
		}
		c.JSON(http.StatusOK, data)
	})

	// 3. 使用自定义map返回数据。
	r.GET("/gin_map", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":   "zhenping",
			"age":    19,
			"gender": "male",
		}
		c.JSON(http.StatusOK, data)

	})

	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("start gin server failed, %v", err)
	}
}
