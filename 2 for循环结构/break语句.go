package main

import "fmt"

func main() {

	var userName string
	var passWd string
	var count int

	for { // 此处为死循环
		fmt.Println("请输入用户名: ")
		fmt.Scan(&userName)
		fmt.Println("请输入密码: ")
		fmt.Scan(&passWd)

		if userName == "admin" && passWd == "8888" {
			fmt.Println("用户名和密码正确.")
			break
		} else {
			count++
			if count >= 3 {
				fmt.Println("最多可以输错3次！！！")
				break
			}
			fmt.Println("用户名或密码错误.!")
		}
	}
}
