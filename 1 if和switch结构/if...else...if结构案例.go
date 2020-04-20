package main

import "fmt"

func main() {

	var userName string
	var passWd string

	fmt.Println("请输入用户名: ")
	_, _ = fmt.Scan(&userName)
	fmt.Println("请输入密码: ")
	_, _ = fmt.Scan(&passWd)

	if userName == "admin" && passWd == "8888" {
		fmt.Println("可以登陆系统")
	} else if userName == "admin" {
		fmt.Println("密码错误")
	} else if passWd == "8888" {
		fmt.Println("用户名错误")
	} else {
		fmt.Println("用户名和密码都错误")
	}
}
