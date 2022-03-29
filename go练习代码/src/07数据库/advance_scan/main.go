package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 忘记导入这个，在链接时候会报错
	"github.com/jmoiron/sqlx"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	rows, err := db.Queryx("SELECT * FROM place")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rows.Next() {
		// cols is an []interface{} of all of the column results
		cols, err := rows.SliceScan()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v", cols)
	}
	fmt.Println("------------------------------------------")
	rows, err = db.Queryx("SELECT * FROM place")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Printf("%v", results)
	}
}
