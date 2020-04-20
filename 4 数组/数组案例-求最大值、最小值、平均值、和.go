package main

import "fmt"

func main() {
	var Number [5] int = [5] int{3, 5, 9, 11, 13}
	var max int = Number[0]
	var min int = Number[0]
	var sum int

	for i:=1; i<len(Number) ; i++ {
		if Number[i] > max {
			max = Number[i]
		}

		if Number[i] < min {
			min = Number[i]
		}

		sum += Number[i]
	}

	fmt.Println("最大值为: ", max)
	fmt.Println("最小值为: ", min)
	fmt.Println("和的值为: ", sum)
	fmt.Println("平均值为： ", sum/len(Number))
}
