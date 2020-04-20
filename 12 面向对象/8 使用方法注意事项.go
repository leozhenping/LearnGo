package main

import "fmt"

// 使用方法注意事项

/*
只要接收者类型不一样， 这个方法就算同名，也是不同方法
*/

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person
	score float64
}

type Teacher struct {
	Person
	salary float64
}

func (s *Student) Show() {
	fmt.Println(s)
}

func (t *Teacher) Show() {
	fmt.Println(t)
}

func main() {

	// 如果接收者类型不同， 即使方法名字是相同的，也是不同方法
	stu := Student{Person{101, "zhenping", 18}, 88.8}
	var t1 Teacher = Teacher{Person{101, "weizi", 28}, 5000}

	stu.Show() // 执行结果: &{{101 zhenping 18} 88.8}
	(&stu).Show() // 等同stu.show()

	t1.Show()  // 执行结果: &{{101 weizi 28} 5000}

}
