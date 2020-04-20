package main

import "fmt"

func main() {

	var level string
	var salary int = 5000
	var on bool = true

	fmt.Println("请输入评级等级: ")
	fmt.Scan(&level)

	//if level == "A" {
	//	salary += 500
	//} else if  level == "B" {
	//	salary += 200
	//} else if level == "C" {
	//
	//} else if level == "D" {
	//	salary -= 200
	//} else if level == "E" {
	//	salary -= 500
	//} else {
	//	on = false
	//	fmt.Println("输入的信息有误")
	//}

	switch level {
	case "A":
		salary += 500
	case "B":
		salary += 200
	case "C":
	case "D":
		salary -= 200
	case "E":
		salary -= 500
	default:
		on = false
		fmt.Println("输入的信息有误!")

	}
	if on {
		//fmt.Println("薪资是: ", salary)
		fmt.Printf("新资: %d", salary)
	}
}
