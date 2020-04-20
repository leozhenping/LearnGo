package main

import "fmt"

func main() {
	// 1. 空指针
	/*
	var p *int
	fmt.Println(p) // 执行结果: nil

	*p = 222 // 由于没有指向其他变量的内存地址， 不能向空指针赋值。
	fmt.Println(*p)
	*/


	// 3. new函数使用
	// 开辟数据类型对应的内存空间， 返回值为数据类型指针
	var p1 *int
	p1 = new(int)  //根据指定的类型，创建一个内存空间
	*p1 = 57 // 将57存放至新开辟出来的内存空间。
	fmt.Println(*p1)
}
