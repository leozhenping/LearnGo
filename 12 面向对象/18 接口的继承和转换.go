package main

import "fmt"

type Humaner interface {
	sayHello()
}

type Person interface {
	Humaner
	hello()
}

type Student struct {

}

func (s *Student) hello()  {
	fmt.Println("hello")
}

func (s *Student) sayHello()  {
	fmt.Println("sayHello")

}

func main() {
	var s Student
	var p Person
	p = &s
	p.hello() // 返回hello
	p.sayHello() // 返回sayhello

	var h Humaner
	h = p
	h.sayHello()  // 返回sayHello

}