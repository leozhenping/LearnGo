package main

/*
map: 字被称为字典结构, 由键和值构成。键是不允许重复的。
是一种无序的键值对的集合。
map最重要的一点是通过key来快速检索数据。key类似于索引， 指向数据的值。
*/

// 1. map的创建
/*

func main() {
	//方式一
	var m map[int]string = map[int]string{1:"张三",2:"李四",3:"王五"}
	fmt.Println(m)

	//方式二
	m1 := map[int]string{1:"张三",2:"李四",3:"王五"}
	fmt.Println(m1)

	//方式三
	m2 := make(map[string]int,10)
	// 赋值
	m2["张三"] = 12
	m2["李四"] = 15
	fmt.Println(m2)
	fmt.Println(len(m2)) // len()返回的是Map中已有的键值对个数
}

*/

// 2. map常用操作

/*

func main() {
	m := map[int]string{1:"王五", 2:"李四"}

	// 通过key获取值
	fmt.Println(m[2])

	// 通过key获取值时，判断是否存在
	value, ok := m[3] // 当键存在时， 返回值和true. 如果键不存在时， 返回空值和false
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("键不存在")

	}

	//循环map
	for key,v := range m {
		fmt.Println(key,v)
	}

	//删除键与值
	delete(m,2)
	fmt.Println(m)
}

*/

// map作为函数参数
// 注意: 在函数中修改map， 会修改原map.

/*
func main() {
	m := map[int]string{1: "张三", 2: "李四", 3: "王五"}
	PrintMap(m)
	deleteMap(m) // 函数删除了2的键， 原map也就被删除了
	fmt.Println(m)
}

func PrintMap(m map[int]string) {
	for key, value := range m {
		fmt.Println(key,value)
	}
}

func deleteMap(m map[int]string)  {
	delete(m, 2)
}
*/

