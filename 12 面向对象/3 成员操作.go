package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person // 只有类型， 没有成员的名字，即称为匿名字段， Person为匿名字段
	score  float64
}

func main() {
	var stu Student = Student{Person{101, "zhenping", 19}, 88.8}

	// 获取成员的值
	fmt.Println(stu.score)     // 执行结果: 88.8
	fmt.Println(stu.Person.id) // 执行结果: 101
	fmt.Println(stu.id)        // 执行结果: 101 , 上一方法的简写

	// 修改成员的值

	stu.score = 100
	fmt.Println(stu.score) // 执行结果: 100

}
