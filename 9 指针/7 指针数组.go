package main

import "fmt"

func main() {
	// 概念
	/*
	数组元素是指针类型
	指针数组指的是一个数组中存储的都是指针（也就是内存地址）， 也就是一个存储了地址的数组。
	var 数组名[下标数量]*类型
	 */

	// 1. 定义
	var p[2] * int

	var i int = 10
	var j int = 20

	p[0] = &i
	p[1] = &j
	fmt.Println(p)  // 执行结果；[0xc0000a4008 0xc0000a4010]
	fmt.Println(*p[0]) // 执行结果: 10


}
