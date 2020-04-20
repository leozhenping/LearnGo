package main

import "fmt"

func main() {

	// 1. 结构体的创建
	// 方法
	type student struct {
		// 成员名称前面不能添加var关键字
		id   int
		name string
		age  int
		addr string
	}

	//2. 初始化
	// A. 顺序初始化
	var s student = student{100, "zhenping",18, "shanghai"}
	fmt.Println(s) // 执行结果: {100 zhenping 18 shanghai}

	// B. 指定成员初始化
	var s1 student = student{name:"weizi", age:18}
	fmt.Println(s1) //执行结果: {0 weizi 18 }.

	// C. 通过结构体变量.名称的方式初始化
	var stu student
	stu.id = 101
	stu.name = "fukun"
	stu.addr = "shanghai"
	fmt.Println(stu) //执行结果: {101 fukun 0 shanghai}

	// 3. 比较
	// 可以使用 == 和 !=  , 其是比内部的成员。
}
