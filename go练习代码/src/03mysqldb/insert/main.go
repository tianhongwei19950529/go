package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 忘记导入这个，在链接时候会报错
	"github.com/jmoiron/sqlx"
	"time"
)

// result, err := db.Exec()
// result := db.MustExec()
// result.LastInsertId()
// result.RowsAffected()
func main() {
	db := sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	schema := `create table if not exists place (
		id int primary key not null auto_increment, 
		country varchar(128),
		city varchar(128) null,
		create_time timestamp not null default current_timestamp,
		update_time timestamp not null default current_timestamp on update current_timestamp,
		telcode int(32) unsigned);`
	result, err := db.Exec(schema)
	if err != nil {
		fmt.Println("create table err...")
		panic(err)
	}
	truncateAll := "truncate place"
	result, err = db.Exec(truncateAll)
	if err != nil {
		fmt.Println("create table err...")
		panic(err)
	}
	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())

	cityState := `insert into place(country, telcode) values(?, ?)`
	countryCity := `insert into place(country, city, telcode) values(?, ?, ?)`
	db.MustExec(cityState, "Hong Kong", 852)
	db.MustExec(cityState, "Singapore", 65)
	db.MustExec(countryCity, "South Africa", "Johnnesburg", 27)
	countryTimeCity := `insert into place(country, city, telcode, create_time) values(?, ?, ?, ?)`
	db.MustExec(countryTimeCity, "China", "TaiWan", 1234, time.Now().Add(time.Hour*24))
	result = db.MustExec(countryTimeCity, "China", "Aomen", 1234, time.Now().AddDate(1, 2, 0))
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	// 插入 Null
	// 为了保证你所插入的值能如你所期望是NULL值，一定记得要将sql.Null***中Valid值置为false
}
