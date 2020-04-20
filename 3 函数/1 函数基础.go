package main

// 1. 定义

/*

func 函数名() {
	函数体
}

*/

// 2. 调用

/*

func main() {
	PlayGame() //函数名
}

*/

// 3. 函数参数

/*

func 函数名(参数) {
	函数体
}

调用:
	函数名(参数)

//示例:

func SumAdd(num1 int, num2 int) {
	fmt.Println(num1 * num2)
}

func main() {
	SumAdd(2, 3)
}

*/

// 4.  函数不定参数列表

/*

func TestSum(args ...int) {

	//for 循环集合中的数据

	//for i :=0; i<len(args);i++ {
	//	fmt.Println(args[i])
	//}

	// for ranage循环集合
	for a, v := range args {
		fmt.Println(a, v) // a存储的是集合中的下标， v存的合集中的值
	}
}

func main() {
	TestSum(1, 2, 3, 4, 5)
}

*/

// 5. 函数返回值
/*

// 函数返回单个值
func AddResult(num1 int, num2 int) int { //表示指定函数返回的数据的类型为整形
	return num1 + num2
}

func AddResult1(num1 int, num2 int ) (sum int) { // 表明最终会返回整形变量sum中的值，在函数体中没有必要再创建sum变量
	sum = num1 + num2
	return sum
}
func AddResult3(num1 int, num2 int ) (sum int) {
	sum = num1 + num2
	return //如果已指定了返回的变量的名称， 那么return后面可以不用再加上变量的名称。
}

// 函数返回多个值
func GetResult() (int, int) {
	var num1 int = 10
	var num2 int = 20
	return num1, num2
}
// 调用
func main() {
	var s  int
	s = AddResult3(3,4)
	fmt.Println(s - 1)

	var s1 int
	var s2 int
	s1,s2 = GetResult()
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
}

*/
