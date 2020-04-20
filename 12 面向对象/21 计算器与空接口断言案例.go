package main

import "fmt"

// 定义父类
type Object struct {
	numberA int
	numberB int
}
// 增加父类的checkData方法, 所有子类都可继承方法.
func (o *Object) checkData(args...interface{}) bool {
	var b bool = true
	if len(args) != 2 {
		fmt.Println("请传入正常的数值个数.")
		b = false
	}

	value1, status1 := args[0].(int)
	if !status1 {
		fmt.Println("第一个值的类型不是整数类型.")
		b = false
	}

	value2, status2 := args[1].(int)
	if !status2 {
		fmt.Println("第二个值的类型不是整数类型.")
		b = false
	}
	o.numberA = value1
	o.numberB = value2

	return b

}

// 定义加法子类
type Add struct {
	Object
}

// 定义減法子类
type Sub struct {
	Object
}

// 定义加法计算方法
func (a *Add) match()  int {
	return a.numberA + a.numberB
}

// 定义减法计算方法
func (s *Sub) match() int {
	return s.numberA - s.numberB
}

// 定义接口
type matcher interface {
	match() int
	checkData(args...interface{}) bool // 增加接口，检验数据
}

// 定义多态
func getsult(m matcher) int {
	return m.match()
}


// 定义操作类
type Oprator struct {

}

// 定义操作子类创建方法

func (o *Oprator) opratorCreate(op string )  matcher { // 由于主体代码使用new创建了一个各类型的指针， 返回的也是一个对象。所以此处需要返回一个指定的接口名称。
	switch op {
	case "+":
		add := new(Add)
		return add
	case "-":
		sub := new(Sub)
		return sub
	default:
		return nil
	}
}

func main() {
	var o Oprator
	obj := o.opratorCreate("-")
	status :=  obj.checkData("abc","ddd")
	if status {
		number := obj.match()
		fmt.Println(number)
	}

	/*
	思路:
		1. 创建一个Oprator的空类
		2. 为oprator的空类添加一个opratorCreate的方法，方法根据传入的op， 判断并创建相应的对象，并返回matcher接口类型.
		3. 返回的matcher接口对象，可以直接调用各自的checkData()方法，检验数据
		4. 再调用对象的match()方法进行计算。
	 */

}


