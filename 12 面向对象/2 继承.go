package main

import "fmt"

/*
继承是一种类间关系 ，描述一个类从另一个类获取成员信息的类间关系
继承必定发生在两个类之间， 参与继承关系的双方称为父类与子类
父类提供成员信息，子类获取成员信息
 */

type Person struct {
	id int
	name string
	age int
}

type Student struct {
	Person // 只有类型， 没有成员的名字，即称为匿名字段， Person为匿名字段
	score float64
}

type Teacher struct {
	Person
	salary float64
}
func main() {

	// 全部初始化
	var stu2  Student = Student{Person{102,"weizi", 18},99.9} // 执行结果；{{102 weizi 18} 99.9}
	fmt.Println(stu2)

	// 部分初始化
	var stu Student = Student{score:90}
	fmt.Println(stu) // 执行结果: {{0  0} 90}

	// 父类成员初始化
	var stu1 Student = Student{Person:Person{id:101}}
	fmt.Println(stu1) // 执行结果: {{101  0} 0}


}