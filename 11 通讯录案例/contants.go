package main

import (
	"fmt"
)

type Person struct {
	name         string
	addressPhone map[string]string
}

var PersonList = make([]Person, 0)

func addPerson() {
	var name string
	var phoneType string
	var phoneNum string
	var addressPhone = make(map[string]string)
	var exitCommand string

	fmt.Println("请输入姓名: ")
	fmt.Scan(&name)

	for {
		fmt.Println("请输入电话类型: ")
		fmt.Scan(&phoneType)
		fmt.Println("请输入电话号码: ")
		fmt.Scan(&phoneNum)
		addressPhone[phoneType] = phoneNum

		fmt.Println("退出请按Q, 任何键继续。")
		fmt.Scan(&exitCommand)
		if exitCommand == "Q" {
			break
		}
	}
	PersonList = append(PersonList, Person{name: name, addressPhone: addressPhone})

	showPerson()
}

func showPerson() {
	for _, person := range PersonList {
		fmt.Printf("姓名:%s\n", person.name)
		for phoneType, number := range person.addressPhone {
			fmt.Printf("%s电话: %s\n", phoneType, number)
		}
		fmt.Println("----------")
	}

}

func deletePerson() {
	var name string
	var index int = -1

	fmt.Println("请输入需要删除联系人的姓名: ")
	fmt.Scan(&name)

	for i := 0; i < len(PersonList); i++ {
		if PersonList[i].name == name {
			index = i
			break
		}
	}
	PersonList = append(PersonList[:index], PersonList[index+1:]...) // append函数第二个参数如果是切片后面要给三个点。
	showPerson()
}

func queryPerson() *Person { // 返回一个Person指针
	var name string
	var index int = -1

	fmt.Println("请输入需要查找联系人的名称: ")
	fmt.Scan(&name)

	for k, person := range PersonList {
		if person.name == name {
			fmt.Printf("姓名: %s\n", person.name)
			for phoneType, phone := range person.addressPhone {
				fmt.Printf("%s号码: %s\n", phoneType, phone)
			}
			index = k
			break
		}
	}

	if index != -1 {
		return &PersonList[index] // 返回一个PersonList[index]的内存地址
	} else {
		fmt.Println("未找到联系人。")
		return nil
	}

}

func editPerson() {
	// 1. 查找联系人
	var name string
	var chooiceNumber int
	var p *Person
	var number int = 0
	var phoneTypeMap = make(map[int]string, 5)
	var choicePhoneTypeNumber int
	var newPhoneNumber string

	p = queryPerson() // 将返回的person内存地址，赋值给P指针，即获取到了person对象.
	if p != nil {
		fmt.Println("修改名称，请输入5, 修改电话，请输入6: ")
		fmt.Scan(&chooiceNumber)
		switch chooiceNumber {
		case 5:
			fmt.Println("请输入新的姓名: ")
			fmt.Scan(&name)
			p.name = name
		case 6:
			for phoneType, _ := range p.addressPhone {
				fmt.Printf("修改%s电话号码，请按%d\n", phoneType, number)
				phoneTypeMap[number] = phoneType // 获取联系人的电话类型，并赋值给phoneTypeMap,phoneTypeMap[0] = "home"
				number++                         // 每次number + 1， 与map的下标相匹配
			}
			fmt.Scan(&choicePhoneTypeNumber)

			fmt.Println("请输入新的电话号码: ")
			fmt.Scan(&newPhoneNumber)
			p.addressPhone[phoneTypeMap[choicePhoneTypeNumber]] = newPhoneNumber // p.addressPhone["home"] = newPhoneNumber
		}

	}
	showPerson()
}

func scanNum() {
	var num int
	fmt.Println("添加联系人信息,请按: 1")
	fmt.Println("删除联系人信息,请按: 2")
	fmt.Println("查询联系人信息,请按: 3")
	fmt.Println("编辑联系人信息,请按: 4")
	fmt.Scan(&num)

	switch num {
	case 1:
		// 添加联系人信息
		addPerson()
	case 2:
		deletePerson()
		// 删除联系人信息
	case 3:
		// 查询联系人信息
		queryPerson()
	case 4:
		// 编辑联系人信息
		editPerson()
	}

}
func main() {
	for {
		scanNum()
	}

}
