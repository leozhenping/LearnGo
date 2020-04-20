package main

// 浅拷贝: 仅仅拷贝的是变量的值，没有对指向的空间进行任何的拷贝
	// go语言的赋值、函数传参、函数返回值都是浅拷
// 深拷贝: 将原有的变量的空间全部拷贝一份


/*

type Student struct {
	name string
	age int
}

func main() {
	stu1 := Student{"zhenping", 18}
	stu2 := Student{"weizi", 19}
	stu3 := Student{"fukun", 7}

	s := []Student{stu1, stu2, stu3}
	for index,_ := range s {
		//value.age = value.age + 100
		s[index].age += 100

	}

	fmt.Println(s)
}

*/
