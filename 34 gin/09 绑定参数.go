package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Age    int    `form:"age" json:"age" binding:"required"`
	Gender string `form:"gender" json:"gender" binding:"required"`
}

func getParams(c *gin.Context) {
	var user UserInfo
	if err := c.ShouldBind(&user); err == nil { // 前端按照userInfo结构体中tag字段传过来的字段名称，与user实例自动完成参数绑定。绑定完成后，可通过user实例中的字段来调用。
		c.JSON(http.StatusOK, gin.H{
			"name":   user.Name,
			"age":    user.Age,
			"gender": user.Gender,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func main() {

	r := gin.Default()
	r.GET("/query", getParams) // 前端传递query params参数时，名称必须按UserInfo中的json tag名称来。
	r.POST("/json", getParams)  // 前端传递raw json字段时，名称可以按json tag或userInfo结构体定义的字段来
	r.POST("/form", getParams)  // 前端传递form-data字段时，名称必须按form tag名称来
	r.Run(":9090")
}
