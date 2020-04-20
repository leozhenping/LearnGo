package main

import "fmt"

type Person struct {
	id int
	name string
	age int
}

type Student struct {
	Person
	score float64

}

func (p *Person) showInfo()  {
	fmt.Println(p)
}

func main() {

	var stu Student = Student{Person{101, "zhenping", 18}, 99.9}
	stu.showInfo() // 继承了父类的方法, 执行结果: &{101 zhenping 18}

}
