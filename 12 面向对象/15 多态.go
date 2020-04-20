package main

import "fmt"

// 1. 多态定义
/*
多态的定义: 指的是多种表现形式
多态就是同一个接口，使用不同的实例而执行不同操作

例：
	两台打印机，一台黑色打印机，一台彩色打印机。他们都有同一接口打印功能。但一个打印的是黑色，一个打印的是彩色。
 */


// 2. 多态的实现

/*

func 函数名(参数名 接口类型) {

}
 */

// 定义接口
type Personer interface {
	sayHello()
}

// 定义学生类
type Student struct {

}

// 定义老师类
type Teacher struct {

}

// 定义学生sayHello方法
func (s *Student) sayHello()  { // 方法的实现要和接口中方法的声明保持一致, 接口中有参数，此处必须带参数。返回值也是一样的。
	fmt.Println("老师好")
}

// 定义老师sayHello方法

func (t *Teacher) sayHello()  {
	fmt.Println("同学好")
}

// 定义多态函数
func whoSayHello(p Personer)  {
	p.sayHello()
}

func main() {
	var s Student
	var t Teacher

	whoSayHello(&s) // 此处相当于 var person Personer, person = &s, 由于s的sayHello()的方法接收变量是一个Student类指针，所以需要在s前加&取地址符。
	whoSayHello(&t)
}
