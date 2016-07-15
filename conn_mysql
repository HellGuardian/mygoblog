package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:linux@tcp(192.168.191.2:3306)/user?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 查询数据库中的内容
	rows, err := db.Query("select * from user")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var ID int64
		var name string
		var age int64
		var sex string
		if err := rows.Scan(&ID, &name, &age, &sex); err != nil {
			log.Fatal(err)
		}
		fmt.Println(ID, name, age, sex)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

