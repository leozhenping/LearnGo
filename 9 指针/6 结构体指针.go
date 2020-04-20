package main

import "fmt"

// 概念: 定义指针，指向结构体

type Student struct {
	id int
	name string
	age int
	addr string
}

func main() {

	s := Student{101,"zhenping",18,"shanghai"}
	var p *Student = &s


	fmt.Println(*p) // 执行结果: {101 zhenping 18 shanghai}
	fmt.Println((*p).name) // 执行结果: zhenping
	fmt.Println(p.name) // 执行结果: zhenping, 可以不使用*p的方式

	// 修改
	p.age = 20
	fmt.Println(s.age) // 执行结果: 20

	// 将结构体指针传入参数
	UpdateStruct(p)
	fmt.Println(s) // 执行结果: {101 zhenping 21 shanghai}

}

func UpdateStruct(s *Student)  { // 此处类型需要加一个*号，表示一个指针
	s.age = 21
}