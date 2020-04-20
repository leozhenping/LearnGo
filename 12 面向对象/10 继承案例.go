package main

import "fmt"

/*
题目；
根据以下信息， 实现对应的继承关系
记者:  我叫张三，我的爱好是偷拍，我的年龄是34， 我是一个男狗仔
程序员: 我叫孙全，我的年龄是23， 我是男生，我的工作年限是3年

抽取父类:
	name : 姓名
    age : 年龄
    sex : 性别

抽取子类:
	记者子类:
		hobby: 爱好
	程序员子类:
		workYear: 工作年限
 */

// 定义父类
type Person struct {
	name string
	age int
	sex string
}

// 定义记者子类
type Rep struct {
	Person
	hobby string
}

// 定义程序员子类
type Pro struct {
	Person
	workYear int
}

// 定义初始化类公共方法
func (p *Person) setValue(name string, age int, sex string)  {
	p.name = name
	p.age = age
	p.sex = sex
}

// 定义记者子类打招呼的方法
func (r *Rep) sayHello(hobby string)  {
	r.hobby = hobby
	fmt.Printf("我叫%s，我的爱好是%s，我的年龄是%d， 我是一个%s狗仔\n", r.name,r.hobby, r.age, r.sex )
}

// 定义程序员打招呼的方法

func (p *Pro) sayHello(workYear int)  {
	p.workYear = workYear
	fmt.Printf("我叫%s，我的年龄是%d， 我是%s生，我的工作年限是%d年\n", p.name, p.age, p.sex, p.workYear)
}

func main() {

	// 初始化记者子类
	var rep Rep
	rep.setValue("zhenping", 18,"男")
	rep.sayHello("偷拍")

	var pro Pro
	pro.setValue("weizi", 18, "女")
	pro.sayHello(3)
}
