package main

import (
	"fmt"
	"reflect"
)

/*
	演示对（基本数据类型， interface{} reflect.value)进行反射的基本操作
*/

func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量的type, kind, 值
	// 1. 先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType= %v\n", rType) // 返回int
	fmt.Printf("rType= %T\n", rType) // 返回 *reflect.rtype

	// 2. 获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal= %v, rVal type= %T\n", rVal, rVal) // 返回100,  reflect.Value

	// 3. 获取变量真正的类型和值
	rVal.Int()

	// 4. 获取变量对应的Kind, 两种方法
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("Kind1 = %v, kind2 = %v\n", kind1, kind2) // Kind1 = int, kind2 = int

	// 4. 将reflect.value转换成interface{}
	iV := rVal.Interface()
	fmt.Printf("iV value: %v, iV Type: %T\n", iV, iV)
	// 将interface{}通过断言转换成需要的类型
	num := iV.(int)
	fmt.Println(num)
}

func main() {

	// 1. 定义Int
	var num int = 100

	reflectTest01(num)

}
