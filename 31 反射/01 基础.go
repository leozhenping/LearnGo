package main

import (
	"fmt"
	"reflect"
)

/*

反射场景1：
	Json的反序列号的tag
反射场景2:
	编写函数的适配器, 如下
	test 1 = func(v1 int, v2 int) {
		...
	}
	test 2 = func(v1 int, v2 int, s string) {
		...
	}

	briage := func(call interface{}, args...interface{}) {
		...
	}

	实现调用test1对应的函数:
		bridge(test1, 1, 2)
	实现调用test2对应的函数
		bridge(test2, 1, 2, "test2")



反射的作用:
	1. 可以在运行时动态获取变量的各种信息，比如变量的类型和类别
	2. 如果是结构体变量， 可以获取到结构体本身的信息，如字段，和方法
	3. 通过反射，可以修改变量的值，可以调用关联的方法
	4. 使用反射，需要import reflect


反射重要的函数及概念:
	1. reflect.TypeOf(变量名) 获取变量的类型， 返回reflect.Type类型
	2. reflect.ValueOf(变量名）， 获取变量的值 ，返回reflect.Value类型， reflect.Value是一个结构体。 通过reflect.Value可以获取到关于变量的很多信息。
	3. 变量、interface{} 和reflect.value是可以相互转换的。这点在实际开发中，会经常使用到。

 */

type student struct {

}
func main() {
	var stu student

	// 1. 将stu转换成reflect.value类型

	rValue := reflect.ValueOf(stu)

	// 2. 将reflect.value转换成interface{}类型
	iVal := rValue.Interface()

	// 3. 将interface{}类型转换成student类型

	v := iVal.(student)
	fmt.Print(v)
}
