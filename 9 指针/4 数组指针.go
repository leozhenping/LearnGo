package main

import "fmt"

func main() {
	// 1. 基本概念
	/*
	定义一个指针变量，指向数组
	var 数组指针变量 *[下标] 类型
	*/

	nums := [10]int{1,2,3,4,5,6,7,8,9,10}
	var p *[10]int = &nums
	fmt.Println(*p) // *p获取整个数组中的内容
	fmt.Println((*p)[0]) //*p加括号，是为了提高其优先级， 优先把*p的值获取到。然后再执行[0]。

	fmt.Println(p[0]) //等价((*p)[0]

	// 循环
	for i:=0; i<len(p);i++  {
		fmt.Println(p[i])

	}

	// 2. 函数中修改数组，利用数组指针完成。
	UpdateDate(p)
	fmt.Println(*p)
}

func UpdateDate(p *[10]int)  { //参处类型前必须加*, 表示一个指针类型
	p[0] = 1000
}

