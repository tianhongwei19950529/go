package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 忘记导入这个，在链接时候会报错
	"github.com/jmoiron/sqlx"
	"time"
)

// rows, err := db.Query()	// 查询单条，不返回错误
// row, err := db.QueryRow()  // 查询多条，不返回错误
// rows.Next()
// rows.Scan()
// err = rowsx, err := db.Queryx()	// 查询单条，不返回错误
// err = rowx := db.QueryRowx()  // 查询多条，不返回错误
// err = rowsx.Next()
// err = rowsx.StructScan()
// err = rowx.StructScan()

func main() {
	db := sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	rows, err := db.Query("select id, country, city, create_time, update_time, telcode from place")
	if err != nil {
		fmt.Println("query err...")
		panic(err)
	}
	for rows.Next() { // 对内存 友好
		var id int
		var country string
		// var city string		// Scan error on column index 2, name "city": converting NULL to string is unsupported
		var city sql.NullString
		var createTime time.Time
		var updateTime time.Time
		var telcode int
		err = rows.Scan(&id, &country, &city, &createTime, &updateTime, &telcode)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("country", country)
		fmt.Println("city", city)
		fmt.Println("telcode", telcode)
		fmt.Println("createTime", createTime)
		fmt.Println("updateTime", updateTime)
	}
	// check the error from rows
	err = rows.Err()
	fmt.Println(err)

	row := db.QueryRow("select id from place where telcode=?", 852)
	var telcode2 int
	err = row.Scan(&telcode2) // 查询不到报err，参数长度不够报err
	if err != nil {
		fmt.Println("row scan telcode2 error...")
		panic(err)
	}
	fmt.Println("telcode2", telcode2)

	type Place struct {
		Id            int64
		Country       string
		City          sql.NullString
		TelephoneCode int       `db:"telcode"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
	rowsx, err := db.Queryx("select * from place") // 只有带x才能用StructScan到结构体
	if err != nil {
		panic(err)
	}
	for rowsx.Next() {
		var p Place
		err = rowsx.StructScan(&p)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println("p", p)
	}
	var p Place
	rowx := db.QueryRowx("select city, telcode from place limit 1")
	err = rowx.StructScan(&p) // 可写成链式结构
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("rowx", p) // 只有 city，telcode字段，其它字段均为初始值

}
