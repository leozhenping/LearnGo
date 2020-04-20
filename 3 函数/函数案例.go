package main

import "fmt"

func Register() {
	var userNmae string = ""
	var passWd string = "8888"
	var email string = "wei@isesol.com"

	if checkInfo(userNmae, passWd, email) {
		fmt.Println("检验成功")
	} else {
		fmt.Println("检验失败")
	}

}

func checkInfo(name string, password string, email string) bool {
	if name != "" && password != "" && email != "" {
		return true
	} else {
		return false
	}
}

func main() {
	Register()

}
