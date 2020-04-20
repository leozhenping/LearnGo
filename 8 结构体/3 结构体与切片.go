package main

import "fmt"

type Student struct {
	id int
	name string
	age int
	addr string
}

func main() {

	// 1. 定义
	var s []Student = []Student{
		{101,"zhenping",18, "shanghai"},
		{102, "weizi",19,"shanghai"},
	}

	fmt.Println(s[0])

	// 2. 修改
	// 与结构体数组操作一致

	// 3. 追加数据
	s = append(s,Student{103, "fukun",6, "shanghai"})
	fmt.Println(s) // 执行结果: 已经向s结构体中追加成功元素

}
