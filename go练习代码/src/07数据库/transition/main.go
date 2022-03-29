package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 忘记导入这个，在链接时候会报错
	"github.com/jmoiron/sqlx"
)

// tx, err := db.Begin()
// tx.Exec()
// tx.Commit()
// tx.Rollback()

// tx := db.MustBegin()
// tx.MustExec()
// err := tx.Commit()

func main() {
	db := sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var hasAError bool
	cityState := `insert into place(country, telcode) values(?, ?)`
	for i := 3; i > -3; i-- {
		result, err := tx.Exec(cityState, "AAA", i)
		if err != nil {
			fmt.Println("exec error, will rollback", err)
			hasAError = true
			break
		}
		fmt.Println("last insert id")
		fmt.Println(result.LastInsertId())
	}
	if hasAError {
		fmt.Println("will rollback")
		err = tx.Rollback()
	} else {
		fmt.Println("will commit")
		err = tx.Commit()
	}
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	tx2 := db.MustBegin()
	for i := 3; i > -3; i-- {
		// todo How can i recover from MustExec panic
		result := tx2.MustExec(cityState, "AAA", i)
		fmt.Println("last insert id")
		fmt.Println(result.LastInsertId())
	}
	if err := tx2.Commit(); err != nil {
		fmt.Println("commit error...")
	} else {
		fmt.Println("commit success")
	}

	// 无法 recover...
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("recoverd in f", r)
	// 		if err := tx2.Rollback(); err != nil {
	// 			fmt.Println("tx2 rollback error...")
	// 		} else {
	// 			fmt.Println("tx2 rollback success...")
	// 		}
	// 	}
	// }()
}
