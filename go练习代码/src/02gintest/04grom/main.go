package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserInfo struct {
	Id     int
	Name   string
	Gender string
	Hobby  string
}

type YouzanErrorRecord struct {
	Id         int
	BizId      string
	HandleType int
	Reason     string
	UserId     string
	Points     string
	CreateTime time.Time
	UpdateTime time.Time
}

func (UserInfo)  TableName()  string  {
	return "user"
}


type Animal struct {
	AnimalId    int64     `gorm:"column:beast_id;size:150"`         // set column name to `beast_id`
	Birthday    time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
	Age         int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}


func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer db.Close()
	db.SingularTable(true)
	//将表名指定成user_info
	db.AutoMigrate(&UserInfo{})
	//db.AutoMigrate(&UserInfo{})
	//u1 := UserInfo{1, "小王子", "男", "篮球"}
	//u2 := UserInfo{2, "白龙马", "男", "跑步"}
	//db.Create(&u1)
	//db.Create(&u2)
	//var use UserInfo
	//db.First(&use)
	//fmt.Println(use)
	//db.Model(&use).Update("hobby", "跑步")
	//var recode1 YouzanErrorRecord
	//db.First(&recode1)
	//fmt.Println(recode1)
	//var user UserInfo
	//db.First(&user)
	//fmt.Println(user)
	db.AutoMigrate(&Animal{})
}
