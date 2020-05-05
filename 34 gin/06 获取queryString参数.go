package main
/*
	queryString格式为:   /web?query=zhenping&age=18
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getWeb(c *gin.Context) {

	// 1. c.Query 获取指定字段， 如果字段为空，name的值为"".
	//name := c.Query("query")

	// 2. c.DefaultQuery 获取指定字段的值，如果值为空， 将返回指定的默认值。 如果前端输入query= , 后端获取到的name值为""， 即不为空。
	//name := c.DefaultQuery("query", "weizi")

	// 3. c.GetQuery 获取指定字段， 返回值和状态
	name, ok := c.GetQuery("query")
	if !ok {
		name = "weizi"
	}
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func main() {

	r := gin.Default()
	r.GET("/web", getWeb)
	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("start gin server failed. %v", err)
	}
}
