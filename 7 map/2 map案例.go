// 统计每个字符出现的次数。


package main

import "fmt"

func main() {
	var str string = "helloworldhellogolang"
	m := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		ch := str[i]
		m[ch] = m[ch] + 1
	}

	for k,v := range m {
		fmt.Printf("%c: %d\n", k,v)
	}

}
