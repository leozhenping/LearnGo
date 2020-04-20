package main

import "fmt"

// 基本语法: copy(切片1，切片2)

func main() {
	s1 := []int{1,2}
	s2 := []int{4,5,6,7,8}

	copy(s1,s2) //将s2拷贝到s1。
	fmt.Println(s1) //执行结果: [4 5], 其是替换相同下标的元素

	copy(s2,s1)
	fmt.Println(s2) //执行结果: [1 2 6 7 8]
}

