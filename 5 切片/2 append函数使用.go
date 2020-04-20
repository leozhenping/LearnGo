package main

import "fmt"

// append函数追加元素时，当容量不够使用时， 一般会自动扩容。扩容方式为上一次： 容量*2 ， 如果超过1024字节，每次扩容上一次的1/4
// append函数向切片末尾追回数据。
func main() {
	s := make([]int, 5, 8)
	s = append(s,1,2,3,4,5)
	fmt.Println("len: ", len(s))  // 结果为10
	fmt.Println("cap: ", cap(s)) //结果为16
}
