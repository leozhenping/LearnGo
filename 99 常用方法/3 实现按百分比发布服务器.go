package main

import "fmt"

var arry = []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10"}

func splitArray(s []string, n int) [][]string {
	var newArray = make([][]string, 0)
	var indexArray = make([]int, 1)
	for i := 0; i < len(s); {
		if len(s)-i < n {
			indexArray = append(indexArray, len(s))
		} else {
			indexArray = append(indexArray, i+n)
		}
		i += n
	}
	j := 0
	for j <= len(indexArray)-2 {
		newArray = append(newArray, s[indexArray[j]:indexArray[j+1]])
		j++
	}
	return newArray
}

func main() {
	newArray := splitArray(arry, 5)
	fmt.Println(newArray)

}
