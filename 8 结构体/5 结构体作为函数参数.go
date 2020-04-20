package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
	addr string
}

func PrintDemo(stu Student) {
	stu.age = 20 // 在函数中修改结构体中的值，是不会影响原来结构体的。
	fmt.Println(stu)
}
func main() {
	stu := Student{101, "zhenping", 18, "shanghai"}
	PrintDemo(stu)
	fmt.Println(stu)
}

