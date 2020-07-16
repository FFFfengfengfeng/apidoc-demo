package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Cate struct {
	id   int
	name string
}

func main() {
	DB, _ := sql.Open("mysql", "feng:123456@tcp(127.0.0.1:3306)/doc")

	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)

	err := DB.Ping()

	if err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	rows, err := DB.Query("select * from cate")

	if err != nil {
		fmt.Println("查询失败")
		return
	}

	for rows.Next() {
		var name string
		var id int
		var createAt string
		if err := rows.Scan(&name, &id, &createAt); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, name: %s, create_at: %s\n", id, name, createAt)
	}
	rows.Close()
}
