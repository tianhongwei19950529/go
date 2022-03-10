package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() (err error) {
	db, err = sql.Open("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	id   int
	age  int
	name string
}

//查询单行
func queryOneRow() {
	var use user
	err := db.QueryRow("select id,name,age from user where id = ?", 1).Scan(&use.id, &use.name, &use.age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", use.id, use.name, use.age)
}

//查询多行 不能用in???
func queryRows() {
	userSlice := make([]user, 0, 10)
	rows, err := db.Query("select * from user where id > ?", 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var use user
		err = rows.Scan(&use.id, &use.name, &use.age)
		if err != nil {
			fmt.Println(err.Error())
		}
		userSlice = append(userSlice, use)
	}
	fmt.Println(userSlice)
}

//更新数据
func insertSql() {
	row, err := db.Exec("insert into user (name,age) value (?,?)", "王子", 19)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, _ := row.LastInsertId()
	fmt.Println(id)
}

func PrePare() {
	item, err := db.Prepare("select id,name,age from user where id > ?")
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, _ := item.Query(0)
	defer rows.Close()
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(u)
	}
}
func main() {
	err := InitMysql()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	//queryOneRow()
	//queryRows()
	//insertSql()
	PrePare()

}
