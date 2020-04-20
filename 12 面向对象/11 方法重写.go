package main

import "fmt"

/*
方法重写:
	就是子类（结构体)中的方法，将父类中的相同名称的方法的功能重新给改写了
	注意: 在调用时，默认调用的是子类中的方法
 */

type Person struct {
	id int
	name string
	age int
}

func (p *Person) printInfo()  {
	fmt.Println("这是父类中的方法.")
}
type Student struct {
	Person
	score float64
}

func (s *Student) printInfo()  {
	fmt.Printf("这是子类中的方法")
}

func main() {
	var stu Student
	stu.printInfo() // 执行结果:这是子类中的方法 . 此处使用子类的方法覆盖父类的方法，即是方法重写。
	stu.Person.printInfo() // 执行结果: 这是父类中的方法. 此方法是调用父类中的重名方法。可以通过父类名称找到。


}
