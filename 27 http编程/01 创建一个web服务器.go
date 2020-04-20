package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFile( w http.ResponseWriter, fileName string)  {
	filePath := "/Users/weizhenping/Downloads" + fileName
	fd, err := os.Open(filePath)
	if err != nil && err != io.EOF {
		fmt.Println("os.open err", err)
		w.Write([]byte("404 file not found."))
	}
	defer fd.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := fd.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}


}

func handler(w http.ResponseWriter, r *http.Request)  {
	// w: 写回给客户端（浏览器）的数据
	// r: 从客户端浏览器读到的数据。
	//w.Write([]byte("hello world"))
	//fmt.Println("r", r)
	//fmt.Println("method", r.Method)
	//fmt.Println("header", r.Header)

	fileName := r.URL.String()
	getFile(w, fileName)


}
func main() {

	// 1. 注册回调函数
	http.HandleFunc("/", handler)  // 参数1: 用户访问url, 参数2: 回调函数名, 要求函数必须是(w http.ResponseWriter, r *http.Request)作为参数。


	// 2. 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8000", nil) // nil: 表示会使用系统自已定义好的回调函数.
}
