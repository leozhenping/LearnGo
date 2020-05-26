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
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func quickSort(s []int) []int {
	fmt.Println(s)
	if len(s) <= 1 {
		return s
	} else {
		splitData := s[0]
		low := make([]int, 0)
		mid := make([]int, 0)
		high := make([]int, 0)
		mid = append(mid, splitData)
		for i:=1; i< len(s); i++ {
			if s[i] <splitData {
				low = append(low, s[i])
			} else if s[i] > splitData {
				high = append(high, s[i])
			} else {
				mid = append(mid, s[i])
			}
		}
		low , high = quickSort(low), quickSort(high)
		s = append(append(low, mid...), high...)
		return s
	}
}

func main() {
	myArr := makeArr(10)
	fmt.Println(myArr)  // [12 81 97 38 46 16 56 78 84 17]
	fmt.Println(quickSort(myArr)) // [12 16 17 38 46 56 78 81 84 97]
}
