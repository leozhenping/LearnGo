package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func getUpload(c *gin.Context)  {
	c.HTML(http.StatusOK, "file.html", nil)
}

func postFile(c *gin.Context)  {
	f, err := c.FormFile("f1") // f1名称为前端定义的Name
	if err != nil {
		fmt.Printf("get file failed. %v", err)
	}

	dst := path.Join("./", f.Filename)
	err = c.SaveUploadedFile(f,dst)
	if err != nil {
		fmt.Printf("save file failed. %v", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status" : "ok",
		})
	}
}
func main() {
	r := gin.Default()

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 前端的上传文件需要定义在form中， 并且类型必须为: multipart/form-data

	r.LoadHTMLFiles("file.html")
	r.GET("/file", getUpload)
	r.POST("/upload", postFile)
	r.Run(":9090")
}
