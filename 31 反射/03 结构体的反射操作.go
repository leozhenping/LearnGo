package main

import (
	"fmt"
	"reflect"
)

func reflectTest02(b interface{}) {
	// 1. 获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType= %v\n", rType) // rType= main.Student

	// 2. 获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal= %v\n", rVal) //rVal= {zhenping 18}

	// 3. 转换成interface{}
	iV := rVal.Interface()
	fmt.Printf("iV value: %v, iV Type: %T\n", iV, iV) // iV value: {zhenping 18}, iV Type: main.Student

	// 4. 转换成对应的类型
	stu, _:= iV.(Student)
	fmt.Println(stu.Name)
}

type Student struct {
	Name string
	Age  int
}

func main() {

	// 定义student实例
	stu := Student{"zhenping", 18}
	reflectTest02(stu)

}
