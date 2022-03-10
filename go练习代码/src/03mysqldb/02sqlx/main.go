package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int
	Age  int
	Name string
}

//func (u User) Value() (driver.Value, error) {
//	return []interface{}{u.Name, u.Age}, nil
//}

var db *sqlx.DB

func InitSql() {
	var err error
	db, err = sqlx.Connect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
}

//查询单挑记录
func queryOneRow() {
	var user User
	sql := "select * from user where id = ?"
	err := db.Get(&user, sql, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

//查询多条记录
func queryMore() {
	var use []User
	sql := "select * from user where id > ?"
	err := db.Select(&use, sql, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(use)
}

func insertDemo() {
	sql := "insert into user(name,age) value (?,?)"
	row, err := db.Exec(sql, "额娘", 19)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row.LastInsertId())
}

func updateDemo() {
	sql := "update user set name = ?  where id = ?"
	row, err := db.Exec(sql, "皇额娘", 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row.RowsAffected())
}

func deleteDemo() {
	sql := "delete from  user  where id = ?"
	row, err := db.Exec(sql, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row.RowsAffected())
}
//
func BatchInsertUsers3(user []*User) error {
	//sql := "insert into user(name,age) value (:name,:age)"
	sql := "INSERT INTO user (name, age) VALUES (:name, :age)"
	rows, err := db.NamedExec(sql, user)
	fmt.Println(rows.RowsAffected())
	return err
}

func QueryByIDs(ids []int) (user []User,err error) {
	sql := "select id,name,age from user where id in (?)"
	query,args ,err := sqlx.In(sql,ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&user,query,args...)
	return
}
func main() {
	InitSql()
	defer db.Close()
	//queryOneRow()
	//queryMore()
	//insertDemo()
	//updateDemo()
	//deleteDemo()
	//u1 := User{Name: "xx", Age: 18}
	//u2 := User{Name: "xxx", Age: 28}
	//u3 := User{Name: "xxxx", Age: 38}
	//var u []User
	//u = []*User{&u1, &u2, &u3}
	//err := BatchInsertUsers3(u)
	//if err != nil {
	//	fmt.Println(err)
	//}
	u,err := QueryByIDs([]int{1,3,8,10,11})
	if err != nil {
		fmt.Println(err)
	}
	for _, user := range u {
		fmt.Println(user)
	}

}
