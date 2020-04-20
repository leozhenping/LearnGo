package main

import (
	"fmt"
)

func delElemt(data []int, index int) []int {
	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}

func delElemt2(data []int, index int) []int {
	result := append(data[:index], data[index+1:]...)
	return result

}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	//data := delElemt(s1, 4)
	data := delElemt2(s1, 4)
	fmt.Println(data)
}
