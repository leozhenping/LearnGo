package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age int
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(localhost)/db2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("db connect err. %v", err))
	}

	db.AutoMigrate(&User{})



	//u1 := User{Name: "zhenping", Age: 18}
	//db.Create(&u1)
	//u2 := User{Name: "weizi", Age: 19}
	//db.Create(&u2)

	// 4. 查询
	/*


		var user User
		// 根据主键查询第一条记录
		db.Debug().First(&user)
		//  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL ORDER BY `users`.`id` ASC LIMIT 1
		fmt.Println(user)

		user = User{}
		db.Debug().First(&user, 2)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((id = 2)) ORDER BY `users`.`id` ASC LIMIT 1
		fmt.Println(user)

		var users []User
		db.Debug().Find(&users, "age > ?", 18)
		//  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL  AND ((age > 18))
		fmt.Println(users)

	*/


	// 5. where条件


	/*

		var users []User

		db.Debug().Where("name=?", "zhenping").First(&user)
		//  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name='zhenping')) ORDER BY `users`.`id` ASC LIMIT 1

		db.Debug().Where("name = ?", "zhenping").Find(&users)
		//  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name = 'zhenping'))

		db.Debug().Where("name <> ?", "zhenping").Find(&users)
	    //  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name <> 'zhenping'))

		db.Debug().Where("name in (?)", []string{"zhenping", "weizi"}).Find(&users)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name in ('zhenping','weizi')))

		db.Debug().Where("name like ?", "%wei%").Find(&users)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name like '%wei%'))

		db.Debug().Where("name = ? and age = ?", "zhenping", 18).Find(&users)
		//  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((name = 'zhenping' and age = 18))

		db.Debug().Where("created_at < ?", time.Now()).Find(&users)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((created_at < '2020-05-08 14:33:08'))


		//  Struct & map

		// struct
		db.Debug().Where(&User{Name: "zhenping", Age: 18}).Find(&users)
		//SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`name` = 'zhenping') AND (`users`.`age` = 18))

		// map
		db.Debug().Where(map[string]interface{}{"name": "zhenping", "age": 18}).Find(&users)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`name` = 'zhenping') AND (`users`.`age` = 18))

		// Slice
		db.Debug().Where([]int{1,2}).Find(&users)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` IN (1,2)))


		提示 当通过结构体进行查询时，GORM将会只通过非零值字段查询，这意味着如果你的字段值为0，''， false 或者其他 零值时，将不会被用于构建查询条件,
		可以使用指针或实现 Scanner/Valuer 接口来避免这个问题.
	*/

	// 6. Not条件

	/*
		db.Not("name", "weizi").Find(&users)
		// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`name` NOT IN ('weizi')))

		db.Not("name", []string{"zhenping"}).Find(&users)

		db.Debug().Not([]int{2}).Find(&users)
		//  SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND ((`users`.`id` NOT IN (2)))

		db.Not("name = ?", "zhenping").Find(&users)
		db.Not(&User{Name: "zhenping"}).Find(&users)
	*/

	// 7. Or条件
	/*

		// 普通字符
		db.Where("name = ?", "zhenping").Or("age > ?", 17).Find(&users)
		// 结构体
		db.Where("name = ? ", "zhenping").Or(&User{Age: 18}).Find(&users)
		// map
		db.Where("name = ? ", "zhenping").Or(map[string]interface{}{"age": 18}).Find(&users)

	*/


	// 8. 内联条件
	/*
		作用与 Where 类似

		// 通过主键获取
		db.First(&users, 2)
		// SELECT * FROM users WHERE id = 2;

		// 如果是一个非整数类型，则通过主键获取
		db.First(&user, "id = ?", "string_primary_key")
		// SELECT * FROM users WHERE id = 'string_primary_key';

		db.Find(&users, "name = ? AND age = ?", "zhenping", 18)

		db.Find(&users, &User{Name: "zhenping"})
	*/


	// 9. Extra Querying option 其它查询选项
	/*
		db.Set("gorm:query_option", "FOR UPDATE").First(&users, "name = ? and age > ?", "zhenping", 18)
		// SELECT * FROM `users`  WHERE name = 'zhenping' and age > 18 ORDER BY `users`.`id` ASC LIMIT 1 FOR UPDATE

	*/


	// 10. FirstOrInit配合Attrs和Assign
	/*

		var user User
		// FirstOrInit方法: 查询不到就初始化一个User对象, 并赋值给user, Attrs()是给User对象添加属性。此处添加一个Age=100, 此操作并不会写入数据库

		// Attrs() : 没有找到weizi对象时，才更新age字段
		db.Debug().Attrs(User{Age: 100}).FirstOrInit(&user, User{Name: "weizi"}) // 输出结果:  {weizi 19}

		// Assign() : 找到或未找到weizi对象时，始终会更新Age字段
		db.Debug().Assign(User{Age: 100}).FirstOrInit(&user, User{Name: "weizi"}) // 输出结果: {weizi 100}
	*/

	// 11. FirstOrCreate
	/*
		获取匹配的第一条记录，否则根据给定的条件创建一个新的记录(仅支持struct和map条件)

		// 查询是否有记录，如果有就返回给user, 没有就新建
		db.Debug().FirstOrCreate(&user, User{Name: "zhenping1", Age: 22}) // 此处必须是一个User对象，不能是切片
		//SELECT * FROM `users`  WHERE (`users`.`name` = 'zhenping1') AND (`users`.`age` = 22) LIMIT 1
		// INSERT INTO `users` (`name`,`age`) VALUES ('zhenping1',22)


		// 等同上面FirstOrCreate(&user, User{Name: "zhenping1", Age: 22})
		db.Where(User{"zhenping", 19}).FirstOrCreate(&user)

		// 查询Name="zhenping2", 如果存在就返回，并不会使用attrs中age=30的字段。 不存在就创建并使用attrs中age=30字段
		db.Where(User{Name: "zhenping2"}).Attrs(User{Age: 30}).FirstOrCreate(&user)

		// 查询Name="zhenping3"并且 assign中age=32 的记录是否存在，如果存在就返回， 不存在就创建
		db.Where(User{Name: "zhenping3"}).Assign(User{Age: 32}).FirstOrCreate(&userdb.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
		// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
		)
	*/

	// 12. 子查询
	/*

		db.Debug().Where("age > ?", db.Table("users").Select("age").Where("name = ?", "weizi").SubQuery()).Find(&users)
		// SELECT * FROM `users`  WHERE (age > (SELECT age FROM `users`  WHERE (name = 'weizi')))

	*/

	// 13. 选择字段

	/*
		db.Debug().Select("name, age").Find(&users)
		// SELECT name, age FROM `users`

		db.Select([]string{"name", "age"}).Find(&users)
		// SELECT name, age FROM `users`

	*/

	// 14. 排序
	/*

		db.Debug().Order("age desc, name").Find(&users)
		// SELECT * FROM `users`   ORDER BY age asc, name

		db.Debug().Order("age desc").Order("name asc").Find(&users)
		//  SELECT * FROM `users`   ORDER BY age desc,name asc

	*/

	// 15. limit

	/*

		db.Debug().Limit(3).Find(&users)
		// SELECT * FROM `users`   LIMIT 3

	*/

	// 16. Offset

	/*
		db.Offset(1).Find(&users)
		// SELECT * FROM users OFFSET 1;

	*/


	// 17. count

	/*

		db.Debug().Where("name = ?", "zhenping").Or("name = ?", "weizi").Find(&users).Count(&users)
		// SELECT * FROM `users`  WHERE (name = 'zhenping') OR (name = 'weizi')
		// SELECT count(*) FROM `users`  WHERE (name = 'zhenping') OR (name = 'weizi')
	*/

	var users []User
	fmt.Println(users)
}
