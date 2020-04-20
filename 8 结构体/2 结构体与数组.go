package main

import "fmt"

type student struct {
	id int
	name string
	age int
	addr string
}

func main() {

	// 1. 定义
	var arr[3]student = [3]student{
		{101,"zhenping",18, "shanghai"},
		{102,"weizi",18, "shanghai"},
	}
	fmt.Println(arr) // 执行结果: [{101 zhenping 18 shanghai} {102 weizi 18 shanghai} {0  0 }]
	fmt.Println(arr[0]) //执行结果: {101 zhenping 18 shanghai}
	fmt.Println(arr[0].name) //执行结果: zhenping


	// 2. 修改结构体的值
	arr[0].age = 20
	fmt.Println(arr[0]) // 执行结果: {101 zhenping 20 shanghai}

	// 3. 通过循环输出结构体中的内容
	// A: 方法一
	for i:=0; i<len(arr) ;i++  {
		fmt.Println(arr[i])
		fmt.Println(arr[i].name) // 执行结果: weizi
	}

	// B: 方法二
	for _,value := range arr {
		fmt.Println(value)
		fmt.Println(value.name, value.age)
	}
}
