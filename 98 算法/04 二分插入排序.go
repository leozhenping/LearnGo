package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeArray(n int) []int {
	var arr = make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}
func findMindIndex(s []int, start, end, cur int) int {
	//无限夹击模式，当s只有最后一个元素时，start和end就相等了，并对s中的元素与current进行比较。并返回索引
	if start == end {
		if s[start] > s[cur] {
			return start
		} else {
			return cur
		}
	}

	// 二分递归法， 查找符合条件元素的索引位置。
	mid := (start + end) / 2
	if s[mid] > s[cur] {
		return findMindIndex(s, start, mid, cur)
	} else {
		return findMindIndex(s, mid+1, end, cur)
	}
}
func main() {
	myArr := makeArray(100)
	fmt.Println(myArr)
	fmt.Println("开始排序")
	startTime := time.Now()
	for i:=1 ; i<len(myArr); i++ {
		p := findMindIndex(myArr, 0, i-1, i) // 0和i-1之间的切片是排序好的，i是current， 需要进行排序的数字。
		// 如果最合适的位置不是待排元素当前位置，那就一次把合适位置后的元素都向后移动一位
		if p != i { // 返回需要插入的索引
			for j:=i; j> p; j-- { // 从索引i的位置向前交换数字, 直到j小于或等于p时结束。
				myArr[j], myArr[j-1] = myArr[j-1], myArr[j]
			}
		}

	}
	elapsed := time.Since(startTime)
	fmt.Println(myArr)
	fmt.Println(elapsed)

}
