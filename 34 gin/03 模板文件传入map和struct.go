package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

var s1 Student = Student{
	"zhenping",
	18,
	"male",
}

var m1  = map[string]interface{}{"name": "weizi", "age": 19, "gender": "female"}

func sayHello1(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模板
	t, err := template.ParseFiles("./test.tmpl")
	if err != nil {
		fmt.Printf("Parse tpml file failed. %v", err)
	}

	//3. 渲染模板
	data := map[string]interface{}{"s1": s1, "m1": m1}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Printf("template exceute failed. %v", err)
	}
}
func main() {

	// 1. 启动httpServer
	http.HandleFunc("/", sayHello1)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("start http server failed , %v\n", err)
	}
}
