package main

import "fmt"

// 定义父类和子类

// 抽象父类
type Object struct {
	numberA int
	numberB int
}

// 加法子类
type Add struct {
	Object
}

// 减法子类
type Sub struct {
	Object
}

// 实现加法match方法
func (n *Add) match() int {
	return n.numberA + n.numberB
}

// 实现减法match方法
func (n *Sub) match() int {
	return n.numberA - n.numberB
}

// 实现接口

type matcher interface {
	match() int  // 公共方法match, 返回int类型
}

// 实现多态
func getResult(n matcher) int {
	return n.match() // 子类的方法中明确返回字段时，需要同时返回，类型为子方法返回的类型
}

/*
// 当需要在加法或减法之间切换时， 需要调整好几行的代码。 使用以下封装的思想，可以减少代码量
func main() {
	add := Add{Object{3,1}}
	number := getResult(&add)
	fmt.Println(number)

	sub := Sub{Object{10,5}}
	number1 := getResult(&sub)
	fmt.Println(number1)
}
*/

// 定义一个操作类

type Oprator struct {

}

// 定义一个类方法，可以接受两个数值和"+", "-"操作

func (o *Oprator) createOprator(op string, numberA int, numberB int) int {
	switch op {
	case "+":
		add := Add{Object{numberA:numberA, numberB:numberB}}
		number := getResult(&add)
		return number
	case "-":
		sub := Sub{Object{numberA:numberA, numberB:numberB}}
		number := getResult(&sub)
		return number
	default:
		return 0
	}
}

func main() {
	var o Oprator
	number := o.createOprator("-", 20, 10)
	fmt.Println(number)
}
