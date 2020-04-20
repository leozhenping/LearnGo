package main

import "fmt"

func main() {
	var s int = 10

	var p *int = &s // 一级指针

	var pp **int = &p  // 二级指针

	**pp = 200  // 修改s的值

	fmt.Println(s)
}