package main

import "fmt"

// 根据用户输入年份和月份， 判断月份中的天数。
//知识点:  fallthrough.  case后面跟着的代码执行完毕后， 不会再执行后面的case, 面是跳出整个swtich结构。 但如果我们想执行完成某个case后，强势执行后面的case， 可以使用fallthrough.


func main() {
	var year int
	var month int
	var day int

	fmt.Println("请输入年份: ")
	_, _ = fmt.Scan(&year)

	fmt.Println("请输入月份: ")
	_, _ = fmt.Scan(&month)

	if month >= 1 && month <= 12 {
		switch month {
		case 1:
			fallthrough
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 10:
			fallthrough
		case 12:
			day = 31
		case 2:
			//判断闰年
			// 闰年判断方法: 年份能被400整队或者年份能被4整除并且不能被100整除。
			if year % 400 == 0 || year % 4 == 0 && year % 100 != 0 {
				day = 29
			} else {
				day = 28
			}
		default:
			day = 30
		}

		fmt.Println(day)
	}
}

