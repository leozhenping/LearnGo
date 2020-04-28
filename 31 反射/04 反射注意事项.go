package main

/*
	1. reflect.Value.Kind() 获取变量的类别，返回的是一个常量
	2. Type是类型， Kind是类别， Type 和kind可能是相同的， 也可能是不同的
		var num int = 0  type: int, Kind: int
		var stu student type: 包名.student, Kind: struct
	3. 通过反射可以让变量在interface{}, reflect.Value之间相互转换
	4. 使用反射的方式来获取变量的值（并返回对应的类型）， 要求数据类型匹配，比如x是Int, 那么就应该使用reflect.value(x).Int(), 而不是使用其它的。 否则panic.
	5. 通过反射来修改变量， 注意当使用SetXXX方法来设置，需要通过对应的指针类型来完成，这样才能改变传入变量的值，同时需要使用到reflect.Value.Elem()方法
 */
func main() {

}
