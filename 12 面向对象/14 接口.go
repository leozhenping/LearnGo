package main

import "fmt"

/*
接口就是一种规范与标准， 只是规定了要做哪些事情，具体怎么做，接口是不管的。 接口把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
接口最好的功能， 就是实现多态
 */

// 1. 定义语法

/*
type 接口名字 interface {
	方法声明
}
 */

type Personer interface {
	sayHello() // sayHello方法声名， 但不实现具体功能
}

type  Student struct {

}

func (s *Student) sayHello()  {
	fmt.Println("老师好")
}

type Teacher struct {

}

func (t *Teacher) sayHello()  {
	fmt.Println("同学们好.")
}
func main() {

	// 对象名.方法名
	var stu Student
	var t1 Teacher
	stu.sayHello()

	//通过接口变量来调用, 必须都实现接口中所声明的方法
	var person Personer // personer就是一个数据类型。
	person = &stu // 将stu的地址赋值给person变量
	person.sayHello() // 调用person的sayhello功能，却调用了stu.sayHello()方法

	person = &t1
	person.sayHello()

}
