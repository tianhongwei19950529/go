package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// db := sqlx.Open()
// db.Ping()
// sqlx.Connect()
// sqlx.MustConnect()

func main() {
	var db *sqlx.DB
	// 打开
	db, err := sqlx.Open("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open error...")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping error....")
	}

	db.Close()

	// sqlx.Connect: open and connect at the same time
	// 填错密码
	// db, err = sqlx.Connect("mysql", "root:1234567@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	// 正确密码测试
	db, err = sqlx.Connect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect error...")
	}
	db.Close()

	// sqlx.MustConnect: open and connect at the same time, panicing on error
	// 填错密码就把报 panic
	db = sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	// db = sqlx.MustConnect("mysql", "root:1234567@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	db.Close()
}
