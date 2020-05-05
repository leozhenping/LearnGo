package  main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getLogin(c *gin.Context)  {
	c.HTML(http.StatusOK, "login.html", nil)
}

func postLogin(c *gin.Context)  {
	// 1. c.PostForm获取参数值
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	// 2. c.defaultPostForm获取参数值, 如果值不存在， 将赋值默认值
	//username := c.DefaultPostForm("username", "somebody")
	//password := c.DefaultPostForm("password", "***")

	// 3. c.GetPostForm获取参数值，返回值和状态, 值不存在， 即返回false
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")

	c.HTML(http.StatusOK, "index.html", gin.H{
		"username": username,
		"password": password,
	})

}


func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html", "index.html")

	r.GET("/login", getLogin)
	r.POST("/login", postLogin)


	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("start gin server failed. %v" , err)
	}
}