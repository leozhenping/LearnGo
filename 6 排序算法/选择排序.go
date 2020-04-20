package main

import "fmt"

func main() {
	s := []int{84, 61, 88, 19, 22, 33, 42}

	for j := 0; j < len(s)-1; j++ { // 只会循环元素总的个数-1次

		var min int = s[j]
		var minIndex int

		for i := j + 1; i < len(s); i++ { // j+1, 表示，每次都从排列好的第后一位开始循环。
			// 每次找出最小值，并获取最小值的下标。
			if s[i] < min {
				min = s[i]
				minIndex = i
			}
		}

		// 如果最小值与相当循环的第一个元素不相等时，才进入条件
		if min != s[j] {
			s[j], s[minIndex] = s[minIndex], s[j] // 交换元素， 把最小值与当时循环的第一个元素进行调换。
		}

	}
	fmt.Println(s) //执行结果: [19 22 33 42 61 84 88], 数值从小到大排列
}