package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 1. 定义模板文件

func sayHello(w http.ResponseWriter, r *http.Request)  {

	// 3. 解析模板
	t, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("template file parse failed.")
	}

	// 4. 渲染html文件
	name := "zhenping"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("template file execute failed.")
	}




}

func main() {
	// 2. 启动http server

	http.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Println("http listen failed.")
	}


}


