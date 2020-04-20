package main

// 定义: 有两个下标的数组，被称为二维数组。
// var 数组名[x][y] 类型

//func main() {
//	var arr[2][3] int
//	arr[0][1] = 123
//	arr[1][2] = 456
//	fmt.Println(arr)
//}

// 全部初始化
//func main() {
//	var arr[2][3] int = [2][3] int {{0, 123, 234},{1, 321, 345}}
//	fmt.Println(arr)
//}

// 部分初始化
//func main() {
//	var arr[2][3] int = [2][3] int {{0, 123 },{1 }}
//	fmt.Println(arr)
//}

//指定元素初始化
//func main() {
//	//var arr[2][3] int = [2][3] int {1:{5,6}} //给第二行的第一列和第二列赋值为5,6 ,结果为:[[0 0 0] [5 6 0]]
//	var arr [2][3] int = [2][3] int{1: {1: 6}} // 给第二行的第二列赋值为6， 结果为:[[0 0 0] [0 6 0]]
//	fmt.Println(arr)
//}

// 通过初始化确定二维数组行数
//func main() {
//	arr := [...][3] int {{1,2,3},{4,5}} // 行的下标可以用"..."代替，但是列的下标不能用"..."来代替。
//	fmt.Println(arr)
//}

// 下标获取方式及两种循环方式
//func main() {
//
//	arry := [2][4] int{{1, 2, 3, 4}, {5, 6, 7, 8}}
//	fmt.Println(len(arry)) // 获取arry的行数
//	fmt.Println(len(arry[0])) // 获取arry第一行的列数
//
//	// 循环二维数组，方式一
//	for _,v := range arry {
//		for _, data := range v {
//			fmt.Println(data)
//		}
//	}
//
//	// 循环二维数组， 方式二
//	for i:=0; i<len(arry); i++ {
//		for j:=0; j<len(arry[0]); j++ {
//			fmt.Println(arry[i][j])
//		}
//	}
//}

