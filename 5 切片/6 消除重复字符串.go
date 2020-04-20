package main

import (
	"fmt"
)

func dedup(data []string) []string {
	out := data[:1]
	for _, word := range data {
		i := 0
		for ; i<len(out);i++  {
			if word == out[i] {
				break // 如果break被执行, i++不会被执行.
			}
		}

		if i == len(out) {  // 当i与len(out)相等时，说明out中没有word原素。
			out = append(out, word)
		}
	}

	return out
}

func main() {
	s := []string{"red", "red","black", "red", "pink", "blud", "pink", "blue"}
	s = dedup(s)
	fmt.Println(s)

	
}
