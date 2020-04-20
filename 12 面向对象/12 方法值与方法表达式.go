package main

import "fmt"

/*
调用方法的方式:
	1. 通过实例.方法名
	2. 方法值
	3. 方法表达式
 */

type Person struct {
	id int
	name string
	age int
}

func (p *Person) printInfo()  {
	fmt.Println(p)
}

func main() {
	var p = Person{101, "zhenping", 18}

	// 第一种调用方式: 对象.方法名
	p.printInfo()

	// 第二种调用方式: 方法值
	f := p.printInfo
	fmt.Printf("%T\n",f) // 执行结果: func().
	f() // 执行结果: &{101 zhenping 18}, 由于printInfo有一个接收都p, 在执行f()的时候，就默认将p传递了给printInfo的接收者.

	// 第三种调用方式: 方法表达式
	f1 := (*Person).printInfo // 直接使用方法名调用，*Person是与printInfo接收者的类型保持一致。另外: 点号的运行优先级要高于*, 所以需要加括号
	f1(&p) // 由于printInfo接收者是一个结构体指针，所以需要在传递对象时，把地址传递给接收者。
}
