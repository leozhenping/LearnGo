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
	m := make(map[int]Student) // 值的类型指定为Student类型

	// 2. 初始化
	m[1] = Student{101, "zhenping", 18, "shanghai"}
	m[2] = Student{102, "weizi", 18, "shanghai"}

	fmt.Println(m) // 执行结果: map[1:{101 zhenping 18 shanghai} 2:{102 weizi 18 shanghai}]
	fmt.Println(m[1]) // 执行结果: {101 zhenping 18 shanghai}
	fmt.Println(m[1].name) // 执行结果； zhenping


	// 3. 循环
	// for...len
	// for...range


	// 4. 删除
	delete(m, 2)
	fmt.Println(m) // 执行结果: map[1:{101 zhenping 18 shanghai}], 键为2的结构体已经被删除


}
