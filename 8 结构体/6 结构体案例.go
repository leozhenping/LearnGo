// 比较年龄最大的学生，并打印出学生的详细信息

package main

import (
	"fmt"
)

type Student struct {
	id int
	name string
	age int
	addr string
}

func main() {

	stu := make([]Student, 3)

	for i:=0; i<len(stu);i++  {
		fmt.Println("请输入第%d个学生的信息: ", i+1)
		_, _ = fmt.Scan(&stu[i].id, &stu[i].name, &stu[i].age, &stu[i].addr)
	}

	InitData(stu)

}

func InitData(stu[]Student)  {
	var maxAge int = stu[0].age
	var maxIndex int
	for i:=0; i<len(stu);i++  {
		if stu[i].age > maxAge {
			maxAge = stu[i].age
			maxIndex = i
		}
	}
	fmt.Println(stu[maxIndex])
}
