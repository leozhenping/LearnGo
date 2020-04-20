package main

import "fmt"

// 实现方式
// 成员操作

type Object struct {
	id int
}

type Person struct {
	Object
	name string
	age  int
}

type Student struct {
	Person
	score float64
}

func main() {
	var stu Student

	stu.age = 18
	fmt.Println(stu.Person.age) // 执行结果: 18

	stu.id = 101
	fmt.Println(stu.Person.Object.id) // 执行结果: 101
}
