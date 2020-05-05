package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func postsIndex(c *gin.Context)  {

	// gin.H是一个自定义的Map
	// name中的posts/index.tmpl是引用模板文件中,define字段的值
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "POSTS",
	})

}

func usersIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "<a href=\"https://www.zhenpin.me\">魏镇坪的博客</a>",
	})
}
func main() {
	r := gin.Default()

	//定义一个不转义相应内容的safe模板函数, 定义后， 模板文件中的safe关键字，才能被正常识别
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 定义静态文件的访问路径，当访问/xxx的路径，会将请求转发至当前statics目录中。html文件中的引用路径，必须是/xxx/index.css类似路径。
	r.Static("/xxx", "./statics")

	// 加载templates目录下的所有tmpl文件，** 表示目录， *表示所有文件
	//r.LoadHTMLFiles("templates/posts/index.tmpl"), 此文件为，只加载一个tmpl模板文件
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "首页")
	})

	r.GET("/posts/index", postsIndex)
	r.GET("/users/index", usersIndex)

	err := r.Run(":9090")
	if err != nil {
		fmt.Printf("gin server start failed. %v", err)
	}
}
