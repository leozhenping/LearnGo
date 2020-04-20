package main

import "fmt"

// 定义指针， 指向切片

func main() {
	// 定义

	s := []int{1, 2, 3, 4, 5, 6}

	var p *[]int

	p = &s

	fmt.Println(*p)
	fmt.Println((*p)[0]) // 执行结果: 打印切片的第一个值

	// 修改

	(*p)[0] = 200
	fmt.Println(s) // 执行结果: [200 2 3 4 5 6]

	// 循环
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}

	for _, value := range *p {
		fmt.Println(value)
	}

}
