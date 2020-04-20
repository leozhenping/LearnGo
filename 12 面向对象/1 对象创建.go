package main

// student 类
type Student struct {
	id   int // 属性
	name string
	age  int
}

func main() {
	var stu Student = Student{101, "zhenping", 19} // stu是一个对象
	var stu1 Student = Student{102, "weizi", 18}
}
