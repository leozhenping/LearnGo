package main

import "fmt"

type Person struct {
	id int
	name string
	age int
}

type Student struct {
	*Person // 匿名指针
	score float64
}

func main() {

	var stu Student = Student{&Person{101,"zhenping", 18}, 99.9} // 赋值时，必须使用&符号，将内存地址赋值给匿名指针
	fmt.Println(stu) // 执行结果: 0xc00000c060 99.9}, Person字段返回了内存地址。

	fmt.Println(stu.id) // 执行结果: 101, 可以直接取值
}