package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` // sql.NullInt64表示: 零值, column定义列名
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号(member number) 唯一且不能为空
	Num          int     `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               //忽略本本字段，不写入数据库中
}

// 定义表的名称
func (u *User) TableName() string {
	return "user"
}

func main() {

	db, err := gorm.Open("mysql", "root:123456@(localhost)/db2?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	db.SingularTable(true) // 禁用表的复数
	if err != nil {
		panic(fmt.Sprintf("db connect error, %v", err))
	}

	db.AutoMigrate(&User{})

}
