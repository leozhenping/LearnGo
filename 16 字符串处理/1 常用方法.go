package main

import (
	"fmt"
	"strings"
)

/*
	studygolang.com/pkgdoc // 文档学习网址
常用方法:
	Contains() // 判断字符串是否包含子串, 返回bool类型
	Join() // 字符串连接
	Index() // 在一个字符串中查找某个字符串的位置
	Repeat() // 某个字符串重复多少次，返回的是重复后的字符串
	replace() // 在s字符串中，把old字符串替换为new字符串, n表示替换的次数，小于0表示全部替换
	split() // 把s字符串按照sep分割，返回slice（切片）

*/

func main() {

	// 字符串包含判断
	var str string = "hellogo"
	b := strings.Contains(str, "go") // 判断go字符串是否在str中存在，如果存在返回true,否则返回false
	fmt.Println(b)                   // 返回true

	// 字符串连接, Join方法
	s := []string{"acb", "hello", "world"}
	s1 := strings.Join(s, "|") // 以"|"对切片的内容进行分隔
	fmt.Println(s1)            // 执行结果:  acb|hello|world

	// 获取指定字符串的索引, Index方法

	s2 := "abchello"
	n := strings.Index(s2, "hello") // 判断hello在str中出现的位置。
	fmt.Println(n)                  // 执行结果: 3

	// 查询字符串重复的次数。Repeat方法
	s3 := strings.Repeat("go", 3) //go重复3次
	fmt.Println(s3)               // 执行结果: gogogo

	// 查找替换，Replace方法
	s4 := "hello world"
	s4 = strings.Replace(s4, "l", "w", 2) // 查找s4字符串中的l字符，替换成w, 只替换2次。 -1表示全部替换(小于0的数字都可以全部替换)
	fmt.Println(s4)                       // 执行结果: hewwo world

	// 字符串分隔, Split()方法
	s5 := strings.Split("xyz", "")
	s6 := strings.Split("hello,world", ",")
	fmt.Println(s5) // 执行结果: [x y z]
	fmt.Println(s6) // 执行结果: [hello world]

}
