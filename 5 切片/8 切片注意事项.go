package main
/*
   切片的使用注意事项:
   	1. 对原切片进行切片时，新切片与原底层共享同一数组；
	2. 对原切片进行切片时，新切片的容量继承旧切片的容量时，，当新切片append新元素时，旧的切片对应的元素也会被修改；
	3. 对原切片进行切片时，为了不引起数据混乱，需将新切片的容量指定为新切片的长度，对新切片再次进行append时，会将新切片的元素复制出来，并添加上新的元素，并组成新的底层数组。
*/

import (
	"fmt"
)

func main() {

	var Slice []int = []int{10, 20, 30, 40, 50}
	fmt.Println(Slice, len(Slice), cap(Slice))


	s1 := Slice[1:2]
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1: %v, len : %d, cap : %d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 35)
	fmt.Println(Slice)
	fmt.Println(s1)

	fmt.Printf("Slice 下标2的内存地址: %v\n", &Slice[2])
	fmt.Printf("s1 下标1的内存地址: %v\n", &s1[1])

	var Slice1 []int = []int{10, 20, 30, 40, 50}
	fmt.Println(Slice1, len(Slice1), cap(Slice1))

	s2 := Slice1[1:2:2]
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s2: %v, len : %d, cap : %d\n", s2, len(s2), cap(s2))
	s2 = append(s2, 35)
	fmt.Println(Slice1)
	fmt.Println(s2)

	fmt.Printf("Slice1 下标2的内存地址: %v\n", &Slice1[2])
	fmt.Printf("s2 下标1的内存地址: %v\n", &s2[1])

}
