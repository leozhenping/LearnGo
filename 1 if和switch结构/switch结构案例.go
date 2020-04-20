package main

import "fmt"

func main() {
	var num float64
	fmt.Println("请输入成绩: ")
	_, _ = fmt.Scan(&num)

	switch { // 此处switch后面没有跟expr表达式，需要在case中引用相应的表达式
	case num >= 90:
		fmt.Println("A")
	case num >= 80:
		fmt.Println("B")
	case num >= 70:
		fmt.Println("C")
	case num >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}
