package main

import (
	"fmt"
	"strings"
)

const (
	UPPER = 1 << iota // 大写字符串
	LOWER             // 小写字符串
	CAP               // 字符串单词首字母大写
	REV               // 反转字符串
)

func reversString(s string) string {
	res := []byte(s)
	num := len(res)
	for i := 0; i < num/2; i++ {
		res[i], res[num-1-i] = res[num-1-i], res[i]
	}
	return string(res)

}
func processString(s string, conf byte) string {
	fmt.Println(conf) // 14

	if (conf & UPPER) != 0 { // 1011 & 0001
		s = strings.ToUpper(s)
	}

	if (conf & LOWER) != 0 {
		s = strings.ToLower(s)
	}

	if (conf & CAP) != 0 {
		s = strings.Title(s)
	}

	if (conf & REV) != 0 {
		s = reversString(s)
	}
	return s
}

func main() {

	res := processString("HELLO WORLD!!!", LOWER|REV|CAP) // 计算: 0010 | 0100 | 1000 ,结果为: 1110
	fmt.Println(res)
}
