package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"time"
)

type Student struct {
	gorm.Model
	Name string
	Age  int
}

func (Student) TableName() string {
	return "student"
}

func buildStu(db *gorm.DB) {
	for i := 10; i < 20; i++ {
		name := fmt.Sprintf("stu%02d", i)
		age := rand.Intn(30)
		db.Create(&Student{Name: name, Age: age})
	}
}

func queryAge(db *gorm.DB) *gorm.DB {
	return db.Where("age = ?", 18)
}

func (Student)BeforeSave()  {

}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Student{})
	//buildStu(db)
	//定义查询要用的信息,返回结果一个的话就是一个结构体,返回结果多个的话就是一个切片
	var Stu Student
	var StuSlice []Student
	fmt.Println("最初的情况", StuSlice, Stu)
	//查询主键排序第一条的记录
	//db.First(&StuSlice)
	////fmt.Println(Stu.Age)
	//fmt.Println(StuSlice)
	//
	//db.Last(&StuSlice)
	//fmt.Println(StuSlice)
	//
	//db.Find(&StuSlice)
	//for _, student := range StuSlice {
	//	fmt.Println(student)
	//}
	//db.Where("name = ?", "stu01").First(&StuSlice)
	//fmt.Println(StuSlice)
	//db.Where(&Student{Name: "stu02"}).First(&StuSlice)
	//fmt.Println(StuSlice)
	//
	//db.Debug().Where("age in (?)", []int{20, 19}).Find(&StuSlice)
	//fmt.Println(StuSlice)
	//
	//db.Not("age", []int{20, 19}).Find(&StuSlice)
	//fmt.Println(StuSlice)
	//
	//db.Where("age = ? or name = ?", 20, "stu08").Find(&StuSlice)
	//fmt.Println(StuSlice)
	//
	//db.Where("age = ?", 20).Or("name = ?", "stu08").Find(&StuSlice)
	//fmt.Println(StuSlice)
	//db.First(&StuSlice, 23)
	//fmt.Println(StuSlice)
	//StuSlice = make([]Student, 0, 10)
	//Stu.Name = "stu99"
	//fmt.Println(Stu)
	//StuSlice = append(StuSlice, Stu)
	//db.Attrs(Student{Name: "stu99", Age: 19}).FirstOrCreate(&Stu, Student{Name: "stu99"})
	//fmt.Println(Stu)
	//
	//queruOne := db.Where("age = ?", 19)
	//lastQuery := queruOne.Or("name = ?", "stu01")
	//lastQuery.Find(&StuSlice)
	//fmt.Println(StuSlice)

	//db.Debug().Scopes(queryAge).Find(&StuSlice)
	//fmt.Println(StuSlice)
	//db.Debug().Where("age = ?",19).Delete(&Student{})

	//Stu.Age= 16
	//Stu.Name = "thw"
	//db.Debug().Model(&Stu).Save(&Stu)
	db.Table("student").Debug().Where("id in (?)", []int{2, 3, 4}).Update(map[string]interface{}{"name": "stu", "age": 19, "updated_at": time.Now()})
}
