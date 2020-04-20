package main

import "fmt"

// 语法

/*
func (对象 结构体类型) 方法名 (参数列表) (返回值) {
	代码体
}
 */

// 方法调用

/*
对象名。方法 // 完成调用
 */

type Student struct {
	id int
	name string
	age int
}

func (s Student) printShow()  { // 此处方法是为Student结构体添加的方法。 此处s也被称之为接收者。
	fmt.Println(s)
}

func (s *Student) editInfo()  { // 结构体类型添加指针，才可以修改s对象的值。否则修改只是s接收的值。而不是stu对象本身。
	s.age = 20
}

func main() {
	var stu Student = Student{101,"zhenping", 18}
	stu.printShow() // 方法调用, 执行结果: {101 zhenping 18}, 原理: 在方法调用时， 将stu中的值，传递给了方法的s .

	stu.editInfo() // 在方法调用时， 将stu的内存地址传递给了s的结构体指针了。
	stu.printShow() // 执行结果: // 20
}



