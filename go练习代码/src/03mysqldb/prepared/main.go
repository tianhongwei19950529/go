package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

// stmt, err := db.Prepare()
// stmt.Query()
// stmt.QueryRow()
// tx.Prepare()
// stmtex, err := db.Preparex()
// stmtex.Get(&x, params...)
func main() {
	db := sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	stmt, err := db.Prepare("select * from place where telcode=?")
	if err != nil {
		fmt.Println("prepare error...")
		panic(err)
	}
	row := stmt.QueryRow(65)
	fmt.Println(row)
	tx := db.MustBegin()
	txStmt, err := tx.Prepare("select id from place where telcode=?")
	rows, err := txStmt.Query(1234)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		fmt.Println(id)
	}
	type Place struct {
		Id            int64
		Country       string
		City          sql.NullString
		TelephoneCode int       `db:"telcode"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}
	// x means ???
	stmtex, err := db.Preparex("select * from place where telcode=?")
	var p Place
	err = stmtex.Get(&p, 852)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(p)
	}
}
