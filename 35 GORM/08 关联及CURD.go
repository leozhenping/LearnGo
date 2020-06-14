package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID                uint       `gorm:"primary_key"`
	Birthday          string     `gorm:"column:birthday"`
	Age               int        `gorm:"column:age"`
	Name              string     `gorm:"column:name"`
	BillingAddress    Address    `gorm:"ForeignKey:BillingAddressId;"` // One-To-One (属于 - 本表的BillingAddressID作外键
	BillingAddressId  int        `gorm:"column:billing_address_id"`
	ShippingAddress   Address    `gorm:"ForeignKey:ShippingAddressId;"` // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressId int        `gorm:"column:shipping_address_id"`
	CreditCard        CreditCard `gorm:"ForeignKey:UserID;"`        // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails            []Email    `gorm:"ForeignKey:UserID;"`        // One-To-Many (拥有多个 - Email表的UserID作外键)
	Languages         []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID     int    `gorm:"primary_key"`
	UserID int    `gorm:"column:user_id"`            // 外键 (属于), tag `index`是为该列创建索引
	Email  string `gorm:"column:email;UNIQUE_INDEX"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
}

type Address struct {
	ID       int    `gorm:"primary_key"`
	Address1 string `gorm:"column:address1;UNIQUE_INDEX"` // 设置字段为非空并唯一
	Address2 string `gorm:"column:address2;UNIQUE_INDEX"`
	Post     string `gorm:"column:post"`
}

// Users是反向关联。方便后期反向查询
type Language struct {
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"column:name;UNIQUE_INDEX"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code  string `gorm:"column:code"`              // `unique_index` also works
	Users []User `gorm:"many2many:user_languages"`
}

type CreditCard struct {
	ID     int    `gorm:"primary_key"`
	UserID int    `gorm:"column:user_id;"`
	Number string `gorm:"column:number;UNIQUE_INDEX"`
}

func (u User) TableName() string {
	return "user"
}

func (u Email) TableName() string {
	return "email"
}

func (u Address) TableName() string {
	return "address"
}

func (u Language) TableName() string {
	return "language"
}

func (u CreditCard) TableName() string {
	return "creditcard"
}

func main() {
	url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "localhost", "go")
	db, err := gorm.Open("mysql", url)
	db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(60)
	db.AutoMigrate(&User{}, &Email{}, &Address{}, &Language{}, &CreditCard{})
	defer db.Close()
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	//t1 := time.Now()
	//u := User{ID: 1}
	//db.Model(&u).Find(&u)
	//db.Model(&u).Related(&u.CreditCard, "CreditCard")
	//db.Model(&u).Related(&u.Emails, "Emails")
	//db.Model(&u).Related(&u.Languages, "Languages")
	//db.Model(&u).Related(&u.BillingAddress, "BillingAddress")
	//db.Model(&u).Related(&u.ShippingAddress, "ShippingAddress")
	//fmt.Println(u)
	//t2 := time.Now()
	//fmt.Println("gorm耗时:", t2.Sub(t1))

	//var address Address
	//var language Language
	//db.Model(&address).Where("Address1 = ? and address2 = ?", "shanghai", "beijing").Scan(&address)
	//db.Model(&language).Where("name = ?", "CN").Scan(&language)
	//
	//user1 := User{
	//	Birthday:        time.Now().String(),
	//	Age:             18,
	//	Name:            "weibochen",
	//	//BillingAddress:  address,
	//	//ShippingAddress: address,
	//	CreditCard: CreditCard{
	//		Number: "111111",
	//	},
	//	Emails: []Email{
	//		{Email: "bochen@isesol.com"},
	//	},
	//	Languages: []Language{
	//		language,
	//	},
	//}
	//
	//db.Create(&user1)

	// 查询
	var user User
	db.Model(&user).Where("name = ?", "zhenping").Scan(&user)

	/*
		var user User
		var card CreditCard

		// 查询指定user的信用卡卡号
		creditCardAssociation := db.Model(&user).Association("CreditCard")
		creditCardAssociation.Find(&card)
		fmt.Println(card)
		fmt.Println(card.number)
	*/

	/*
		// 根据用户查询关联的地址
		var addr Address
		addressAssociation := db.Model(&user).Association("BillingAddress")
		addressAssociation.Find(&addr)
		fmt.Println(addr)
		fmt.Println(addr.Address2)
		fmt.Println(addr.Address1)
	*/

	/*
		// 查询指定用户的语言
		var language []Language
		languageAssociation := db.Model(&user).Association("Languages")
		languageAssociation.Find(&language)
		fmt.Println(language)
		fmt.Println(language.name)
	*/

	/*

		// 查询指定用户的邮箱地址
		var emails []Email
		emailAssociation := db.Model(&user).Association("Emails")
		emailAssociation.Find(&emails)
		for _, email := range emails {
			fmt.Println(email.Email)
		}
	*/

	/*

		// 反向关联，查找会CN语言的用户
		var language Language
		var  user []User
		db.Model(language).Where("name = ?", "CN").Scan(&language)
		languageAssociation := db.Model(&language).Association("Users")
		languageAssociation.Find(&user)
		for _, user := range user {
			fmt.Println(user.Name)
		}
	*/

	// 添加
	/*
		// 一对多添加
		// 给用户添加新的邮箱地址
		var emails []Email
		emailAssociation := db.Model(&user).Association("Emails")
		emailAssociation.Append(Email{UserID: int(user.ID), Email: "zhenping2.wei@isesol.com"})
		emailAssociation.Find(&emails)
		fmt.Println(emails)
	*/

	/*
		// 多对多添加
		// 给用户添加新的语言
		var languages []Language
		languageAssociation := db.Model(&user).Association("Languages")
		languageAssociation.Append(Language{Name: "EN", Code: "3"})
		languageAssociation.Find(&languages)
		fmt.Println(languages)
	*/

	/*
		// 一对一添加
		// 给用户添加地址
		var address Address
		db.Model(&address).Where("id = ?", 1).Scan(&address)
		addressAssociation := db.Model(&user).Association("BillingAddress")
		addressAssociation.Append(address) //  UPDATE `user` SET `billing_address_id` = 1  WHERE `user`.`id` = 3
	*/

	// 替换
	/*
		// 一对一关系替换
		var address Address
		db.Model(&Address{}).Where("id = ?", 2).Scan(&address)
		addressAssociation := db.Model(&user).Association("BillingAddress")
		addressAssociation.Replace(address) //  UPDATE `user` SET `billing_address_id` = 2  WHERE `user`.`id` = 1
		fmt.Println(user)
	*/

	/*
		// 一对多关系替换
		var emails = []Email{
			{Email: "w1@isesol.com"},
			{Email: "w2@isesol.com"},
		}
		emailAssociation := db.Model(&user).Association("Emails")
		emailAssociation.Replace(emails)
		//  INSERT INTO `email` (`user_id`,`email`) VALUES (1,'w1@isesol.com')
		//  UPDATE `email` SET `user_id` = 1, `email` = 'w1@isesol.com'  WHERE `email`.`id` = 4
		//  INSERT INTO `email` (`user_id`,`email`) VALUES (1,'w2@isesol.com')
		//  UPDATE `email` SET `user_id` = NULL  WHERE (`id` NOT IN (4,5)) AND (`user_id` = 1)
		fmt.Println(user)
	*/

	/*
		// 多对多关系替换
		CN1 := Language{
			Name: "CN4",
			Code: "4",
		}
		CN2 := Language{
			Name: "CN5",
			Code: "5",
		}
		CN3 := Language{
			Name: "CN6",
			Code: "6",
		}
		userAssociation := db.Model(&user).Association("Languages")
		userAssociation.Replace([]Language{CN1, CN2, CN3}) // 由于cn1, cn2, cn3在数据库中不存在， gorm会先创建并再创建关联。
		//  INSERT INTO `language` (`name`,`code`) VALUES ('CN6','6')
		//  INSERT INTO `user_languages` (`user_id`,`language_id`) SELECT 1,12 FROM DUAL WHERE NOT EXISTS (SELECT * FROM `user_languages` WHERE `user_id` = 1 AND `language_id` = 12)
	*/

	// 删除
	/*

		// 一对一关系删除
		var address Address
		db.Model(&address).Where("address1 = ? and address2 = ? ", "sichuan", "shanghai").Scan(&address)
		fmt.Println(address)
		addressAssociation := db.Model(&user).Association("BillingAddress")
		if err := addressAssociation.Delete(address).Error; err != nil {
			fmt.Println(err)
		}
	*/

	/*
		// 一对多关系删除
		var email Email
		db.Model(&email).Where("email = ?", "w1@isesol.com").Find(&email)
		emailAssociation := db.Model(&user).Association("Emails")
		emailAssociation.Delete(email) // UPDATE `email` SET `user_id` = NULL  WHERE (`user_id` IN (1)) AND (`id` IN (4))
	*/

	/*
		// 多对多删除
		var language Language
		db.Model(&language).Where("id = ?", 10).Find(&language)
		languageAssociation := db.Model(&user).Association("Languages")
		languageAssociation.Delete(language) // DELETE FROM `user_languages`  WHERE (`user_id` = 1) AND (`language_id` IN (10))
	*/

	// 清除
	/*
		// 一对一清除
		addressAssociation := db.Model(&user).Association("BillingAddress")
		addressAssociation.Clear()
	*/

	/*
		// 一对多清除
		emailAssociation := db.Model(&user).Association("Emails")
		emailAssociation.Clear() //  UPDATE `email` SET `user_id` = NULL  WHERE (`user_id` = 1)
	*/

	/*
		// 多对多清除
		languageAssociation := db.Model(&user).Association("Languages")
		languageAssociation.Clear() // DELETE FROM `user_languages`  WHERE (`user_id` IN (1))
	*/

	// 计数
	/*
		languageAssociation := db.Model(&user).Association("Languages")
		countNumber := languageAssociation.Count()
		fmt.Println(countNumber)
	*/
}
