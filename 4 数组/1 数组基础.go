package main

// 数组: 是指一系列同一类型数据的集合

// 1. 数组定义
// var 数组名[元素数量] 类型
//var Numbers[5]  int

// 全部初始化
//func main() {
//	var Numbers[5] int = [5]int{1,2,3,4,5}  // 下标是从0开始计算的
//	fmt.Println(Numbers[3]) //获取下标3的元素
//}

// 部分初始化
//func main() {
//	Numbers := [5]int{1,2} // 只赋值了2个元素，其他元素默认值为0
//	fmt.Println(Numbers[0])
//}

// 定义元素初始化
//func main() {
//	Numbers := [5]int{2:5, 3:6} // 指定下标2的元素为5, 其他元素默认值为0
//	fmt.Println(Numbers[1])
//}

// 通过初始化确定数组长度
//func main() {
//	Numbers := [...] int {7,8,5}
//fmt.Println(len(Numbers))
//fmt.Println(Numbers[1])
//}

// 2. 数组遍历
// for...len()
//func main() {
//	Numbers := [...]int{1, 2, 3, 4, 5}
//	for i := 0; i < len(Numbers); i++ {
//		fmt.Println(Numbers[i])
//
//	}
//}

// for...range
//func main() {
//	Numbers := [...]int{1,2,3,4,5}
//	for _,v := range Numbers {
//		fmt.Println(v)
//	}
//}

// 3. 参数作为函数参数
// 函数中修改数组中的值，不会影响到原数组

//func main() {
//	Numbers := [5]int{1,2,3,4,5}
//	getNumber(Numbers)
//}
//
//func getNumber(n [5] int)  {
//	fmt.Println(n)
//}

// 4. 数组值比较
// 数组比较只支持 "=="， "!="。 数组的类型和长度要保持一致。
// 方法一: 循环对比
//func compare(n1 [5] int, n2 [5] int) bool {
//	var result = true
//	for i := 0; i < len(n1); i++ {
//		if n1[i] != n2[i] {
//			result = false
//			break
//		}
//	}
//	return result
//}
//
//func main() {
//	Number1 := [5]int{1, 2, 3, 4, 5}
//	Number2 := [5]int{1, 2, 3, 3, 5}
//	if compare(Number1, Number2) {
//		fmt.Println("数组一致")
//	} else {
//		fmt.Println("数组不一致")
//	}
//}

// 方法二: 直接使用Number1 == Number2比较即可。 返回true或false.
