package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)                     // rVal是一个指针类别
	fmt.Printf("rVal kind is : %v\n", rVal.Kind()) // rVal kind is : ptr
	//rVal.SetInt(200) // 由于rVal的类别是ptr
	rVal.Elem().SetInt(200) // 类似 *rVal.SetInt()
}

func main() {

	var num int = 100
	reflect01(&num)
	fmt.Println(num) // 最终输出修改后的值 : 200

}
