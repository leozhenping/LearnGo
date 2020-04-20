package main

import "fmt"

// 使用面向对象方式 ，实现两个数加减的计算器

// 1. 抽取公共类
// 2. 实现各自的加、减方法

// 实现公共类，其都包含两个数字
type Object struct {
	num1 int
	num2 int
}

// 实现加法子类
type Add struct {
	Object
}

// 实现减法子类
type Sub struct {
	Object
}

// 实现加法方法
func (n *Add) match(num1 int, num2 int)  int {
	n.num1 = num1
	n.num2 = num2
	return n.num1 + n.num2
}

// 实现减法方法
func (n *Sub) match(num1 int, num2 int) int {
	n.num1 = num1
	n.num2 = num2
	return n.num1 - n.num2
}

func main() {
	var add Add  // 初始化加法对象
	// 完成加法运算
	n := add.match(20,10)
	fmt.Println(n)

	var sub Sub // 初始化减法对象
	// 完成减法运算
	n1 := sub.match(20,10)
	fmt.Println(n1)

}
