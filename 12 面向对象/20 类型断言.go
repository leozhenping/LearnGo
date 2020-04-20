package main

import "fmt"

/*
通过类型断言，可以判断空接口中存储的数据类型
语法:  value, status := m.(T)

	m: 表空接口类型变量
	T: 是断言的类型
	value: 变量m中的值
	status: 布尔类型变量， 如断言成功为true, 否则为false.
 */

func main() {
	 //var i interface{}
	 //i = 123
	 //value, status := i.(int)
	 //if status {
	 //	fmt.Println(value)
	 //} else {
	 //	fmt.Println("断言失败")
	 //}

	 var i []interface{}
	 i = append(i,"abc")
	 value, status := i[0].(int)
	 if status {
	 	fmt.Println(value)
	 } else {
	 	fmt.Println("断言失败")
	 }

}
