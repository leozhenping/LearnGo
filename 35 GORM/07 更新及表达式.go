package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id     int
	Name   string
	Age    int
	Gender string
	Hobby  string
	Active bool `gorm:"default:true"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/db2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("db connect error, %v", err))
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	//var user1 = User{Name: "zhenping", Age: 18, Gender: "male", Hobby: "haha", Active: true}
	//db.Create(&user1)
	//var user2 = User{Name: "weizi", Age: 19, Gender: "female", Hobby: "haha", Active: true}
	//db.Create(&user2)

	// 更新所有字段
	// 更新操作时， User结构体必须有ID字段，或嵌套gorm.model字段。不然update就会成为Insert。
	// save()默认会更新该对象的所有字段，即使你没有赋值。

	/*

		var user User
		db.First(&user)

		user.Age = 20
		user.Hobby = "爬山"
		db.Debug().Save(&user)
		//  UPDATE `users` SET `name` = 'zhenping', `age` = 20, `gender` = 'male', `hobby` = '爬山'  WHERE `users`.`id` = 1

	*/

	// 更新指定字段，可以使用Update或updates
	/*

		var user User
		db.First(&user)
		db.Model(&user).Update("hobby", "hehe")

		db.Mode(&user) // 如果user是一个结构体，那么update是对user表的所有操作，如果user是一个查询后的对象，那么update是针对单个对象。

		db.Debug().Model(&user).Where("active = ?", true).Update("age", 30)
		//  UPDATE `users` SET `age` = 30  WHERE (active = true)

		// 使用map类型，更新多个字段
		db.Debug().Model(&user).Updates(map[string]interface{}{"name":"zhenping1", "age":32})
		// UPDATE `users` SET `age` = 32, `name` = 'zhenping1'  WHERE `users`.`id` = 1

		// 只会更新有变化的值，零值也可以被更新。
		db.Debug().Model(&user).Updates(map[string]interface{}{"name":"zhenping2", "age": 0})
		// UPDATE `users` SET `age` = 0, `name` = 'zhenping2'  WHERE `users`.`id` = 1

		// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
		db.Debug().Model(&user).Updates(User{Name: "zhenping", Age: 0})
		//  UPDATE `users` SET `name` = 'zhenping'  WHERE `users`.`id` = 1

		// 使用map更新类型零值, 零值正常被更新至表中
		db.Debug().Model(&user).Updates(map[string]interface{}{"name":"", "age":"", "hobby":"", "active":false})

		// 只更新选中的字段，如name
		db.Debug().Model(&user).Select("name").Updates(map[string]interface{}{"name":"zhenping", "age":18, "gender": "male", "hobby":"haha", "active":false})
		// UPDATE `users` SET `name` = 'zhenping'  WHERE `users`.`id` = 1

		// 排除指定字段，更新其他所有字段
		db.Debug().Model(&user).Omit("name").Updates(map[string]interface{}{"name": "zhenping", "age": 18, "hobby": "haha", "active": true, "gender": "male"})
		//  UPDATE `users` SET `active` = true, `age` = 18, `gender` = 'male', `hobby` = 'haha'  WHERE `users`.`id` = 1


		// 只更新单个字段，不执行beforeUpdate, AfterUpdate方法
		db.Debug().Model(&user).UpdateColumn("age", 22)
		// UPDATE `users` SET `age` = 22  WHERE `users`.`id` = 1

		// 同时更新多个字段，不执行beforeUpdate, AfterUpdate方法
		db.Debug().Model(&user).UpdateColumns(map[string]interface{}{"name":"zhenping11", "age": 18})
		//  UPDATE `users` SET `age` = 18, `name` = 'zhenping11'  WHERE `users`.`id` = 1

		// 批量更新
		db.Debug().Model(User{}).Where("active = ?", true).Update(map[string]interface{}{"age":100, "active": false}) // 此处model中的user是一个结构体，表示users表。
		db.Table("users").Where("age = ?", 100).Update("age" , 18)

	*/

	// gorm表达式

	/*
		db.Debug().Model(&user).Update("age", gorm.Expr("age + ?", 100))
		// UPDATE `users` SET `age` = age + 100  WHERE `users`.`id` = 1

		db.Debug().Model(User{}).Where("age < ?", gorm.Expr("100 - ?", 1)).Find(&users)
		// SELECT * FROM `users`  WHERE (age < 100 - 1)
	*/

}
