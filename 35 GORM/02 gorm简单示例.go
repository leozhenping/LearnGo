package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	gorm.Model
	Name   string
	Age    int
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/db2?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(fmt.Sprintf("db connection error, %v", err))
	}
	// 1. 创建表
	db.AutoMigrate(&UserInfo{})

	// 2. 创建 from user_infos where name='zhenping';行
	//db.Create(&UserInfo{Name: "zhenping", Age: 19, Gender: "male", Hobby: "haha"})

	// 3. 读取
	var user UserInfo
	//db.First(&user)
	//fmt.Println(user)

	//db.First(&user, "id = ?", 1)  //条件查询
	//fmt.Println(user)

	//db.First(&user, 1)	// 查询id为1的行
	//fmt.Println(user)

	db.First(&user)
	fmt.Println(user)

	// 4. 更新
	//db.Model(&user).Update("age", 20)
	//fmt.Println(user)

	// 5. 删除
	db.Delete(&user) // 如果嵌入了grom.model结构体，删除数据是逻辑删除，不会真正的删除数据。

}


