package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数字转换成字符串
	s1 := strconv.Itoa(3)
	fmt.Println(s1)

	s2 := strconv.FormatBool(true)
	fmt.Println(s2)


	// 把字符串转换为其他类型
	// "true" -- bool
	// "123" -- int
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	num,_ := strconv.Atoi("123")
	fmt.Println(num)

}
