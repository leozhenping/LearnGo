// 计算一组整数之和。 使用切片

package main

import (
	"fmt"
)

func InitData(number [] int) {
	for i:=0; i<len(number);i++  {
		fmt.Printf("请输入第%d个整数: ", i+1)
		fmt.Scan(&number[i])
	}
}

func SumCompare(number [] int) int {
	var sum int
	for i:=0; i<len(number);i++  {
		sum += number[i]
	}

	return sum
}

func main() {
	var count int
	fmt.Println("请输入需要计算的整数个数: ")
	fmt.Scan(&count)

	s := make([]int, count)
	InitData(s)

	sum := SumCompare(s)
	fmt.Printf("sum的值是: %d", sum)
}


