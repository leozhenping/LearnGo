package main

import "fmt"

// 找出100 - 999间的水仙花数
// 水仙花: 是指一个三位数，它的每位数字的立方和等于其本身。

func main() {
	var h int // 百个
	var t int // 十位
	var b int // 个位

	for i := 100; i <= 999; i++ {
		h = i / 100
		t = i % 100 / 10
		b = i % 10

		if h*h*h+t*t*t+b*b*b == i {
			fmt.Println(i)
		}
	}
}
