package main

import "fmt"

/*
空接口: 不包含任何的方法， 正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。
 */

func main() {
	var i interface{} // 空接口
	fmt.Println(i)

	// 空接口的切片。
	var s []interface{}
	s = append(s, 123, "abc", 88.8)
	for j:=0; j<len(s) ;j++  {
		fmt.Println(s[j])

	}

}
