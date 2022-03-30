package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

// fuck sql in keyword, please use sqlx.In
// 	rows, args, err := sqlx.In(sql, params)
// 	sql = db.Rebind(sql)
// 	db.Query(sql, args)
// what the fuck NamedQuery, NamedExec, PrepareNamed, Named
//

func main() {
	db := sqlx.MustConnect("mysql", "root:123456@(127.0.0.1)/sql_test?charset=utf8mb4&parseTime=True&loc=Local")
	var levels []int = []int{852, 65, 27, 1234, 0, 3, 2, 1}
	rows, err := db.Query("select id from place where telcode in (?)", levels)
	if err != nil {
		fmt.Println("query has err...")
		fmt.Println("fuck in query, use sqlx.In to create query")
	}
	// sqlx.In
	query, args, err := sqlx.In("select id from place where telcode in (?);", levels)
	if err != nil {
		fmt.Println("sqlx in has err...")
	}
	query = db.Rebind(query) // fuck!!!
	fmt.Println(query)
	fmt.Println(args)
	rows, err = db.Query(query, args...) // fuck!!! fuck  ...
	if err != nil {
		fmt.Println("query has err...")
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println("scan has err...")
			panic(err)
		} else {
			fmt.Println(id)
		}
	}
	// fuck NamedQuery
	fmt.Println("NamedQuery")
	type Place struct {
		Id            int64
		Country       string
		City          sql.NullString
		TelephoneCode int       `db:"telcode"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
	}

	p := Place{
		Country: "China",
		City: sql.NullString{
			"Beijing",
			false,
		},
		TelephoneCode: 5201314,
	}
	rowsNamedQuery, err := db.NamedQuery("select id from place where country=:country", p)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for rowsNamedQuery.Next() {
		var id int
		rowsNamedQuery.Scan(&id)
		fmt.Println(id)
		// var pp Place
		// err := rowsNamedQuery.Scan(&pp)
		if err != nil {
			fmt.Println("scan error...")
			fmt.Println(err)
			panic(err)
		}
	}
	m := map[string]interface{}{"city": "TaiWan"}
	// m := map[string]interface{}{"city": "Beijing"}
	result, err := db.NamedQuery("select id from place where city=:city", m)
	if err != nil {
		fmt.Println("query err...")
		fmt.Println(err)
	}

	fin := result.Next()
	if fin {
		fmt.Println("hint next")
		var id int
		err = result.Scan(&id)
		if err != nil {
			fmt.Println("scan err...")
			fmt.Println(err)
			panic(err)
		}
		fmt.Println("-----------")
		fmt.Println(id)
	}

	fmt.Println("PreNamed")
	p = Place{TelephoneCode: 50}
	pp := []Place{}
	nstmt, err := db.PrepareNamed("select * from place where telcode > :telcode")
	if err != nil {
		fmt.Println("prenamed error...")
		fmt.Println(err)
		panic(err)
	}
	err = nstmt.Select(&pp, p)
	if err != nil {
		fmt.Println("select err...")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(pp)

	// Named, In...
	fmt.Println("Named, In...")
	arg := map[string]interface{}{
		"country": "AAA",
		"telcode": []int{0, 1, 2},
	}
	query, args, err = sqlx.Named("select id from place where country=:country and telcode in (:telcode)", arg)
	query, args, err = sqlx.In(query, args...)
	query = db.Rebind(query)
	rowws, err := db.Query(query, args...)
	if err != nil {
		return
	}
	for rowws.Next() {
		fmt.Println("id...")
		var id int
		rowws.Scan(&id)
		fmt.Println(id)
	}

}
