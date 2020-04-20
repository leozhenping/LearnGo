package main

import "fmt"

func main() {
	var num int

	fmt.Println("请输入一个数字: ")
	_, _ = fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期天")
	default: //当所有条件不匹配时， 将执行default代码块
		fmt.Println("输入数值有误。。。")

	}
}
