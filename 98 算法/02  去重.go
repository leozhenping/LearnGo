package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeArr(n int) []int {
	var arr = make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10))
	}
	return arr
}

func quickSort(s []int) []int {
	if len(s) <= 1 {
		return s
	} else {
		splitData := s[0]
		low := make([]int, 0)
		mid := make([]int, 0)
		high := make([]int, 0)
		mid = append(mid, splitData)
		for i := 1; i < len(s); i++ {
			if s[i] < splitData {
				low = append(low, s[i])
			} else if s[i] > splitData {
				high = append(high, s[i])
			} else {
				mid = append(mid, s[i])
			}
		}
		low, high = quickSort(low), quickSort(high)
		s = append(append(low, mid...), high...)
		return s
	}
}

func duplicate(s []int) []int {
	arrLen := len(s)
	newArr := make([]int, 0)
	i := 0
	for i < arrLen {
		fieldNum := s[i]
		newArr = append(newArr, fieldNum)
		times := 1
		for i+1 < arrLen && s[i] == s[i+1] {
			i++
			times++
		}
		i++
		fmt.Println(fieldNum, times)
	}

	return newArr
}

func main() {
	myArr := makeArr(1000)
	// 去重需要先进行排序
	myArr = quickSort(myArr)
	fmt.Println(myArr)
	fmt.Println(duplicate(myArr))
}
