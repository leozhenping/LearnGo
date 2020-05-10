package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id     int `gorm:"primary key"`
	Name   string
	Age    int
	Gender string `gorm:"default:'male'"` // 设定字段默认值
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/db2?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic(fmt.Sprintf("db connect error, %v", err))
	}
	db.AutoMigrate(&User{})
	// Gender写入字符串的零值，但又指定了默认值，此时会忽略零值，直接写入默认值。
	// 当字段未指定默认值时， 插入类型的零值，是可以直接插入的。
	user := User{Name: "", Age: 0, Hobby: "haha", Gender: ""}

	/*
		所有字段的零值，比如: 0, "", false或者其他零值，都不会保存到数据库内，但会使用他们的默认值。如果想避免类似情况，可以考虑使用指针或实现
		Scanner/Valuer接口

		如果指定了默认值，但又想保存类型的零值时，可以使用指针或实现scanner/valuer接口。如下:

		1. 指针
		type User struct {
			name *string `grom:"default:'zhenping'"`
		}

		user := User{Name: new(string)} // 因为name已经是指针类型，必须使用new方法返回一个指针，而不是直接写{name: ""}

		2. scanner/valuer接口
		type User struct {
			Name sql.NullString
		}

		user := User{Name:sql.NullString{String:"", Valid: true}} // 当Valid为true时， 告诉gorm必须使用此处定义的String:""
	*/
	if db.NewRecord(&user) { // db.NewRecord(&user), 主键为空返回true, 否则返回false
		db.Debug().Create(&user) //Debug显示执行的SQL, 结果为: INSERT INTO `users` (`name`,`age`,`hobby`) VALUES ('zhenping',18,'haha')
		fmt.Println(user.Id)
	} else {
		fmt.Println("user exist.")
	}
}
