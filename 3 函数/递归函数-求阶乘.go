package main

import "fmt"

// 阶乘公式:  n! = 1 * 2 * 3 * .... * n

var s int = 1

func TestDeom(n int)  {
	if n == 1 {
		return //结束函数
	}
	s *= n
	TestDeom(n-1)
}

func main() {
	TestDeom(5)
	fmt.Println(s)
}

//总结：
// 1. 自己调用自己
// 2. 一定要有出口