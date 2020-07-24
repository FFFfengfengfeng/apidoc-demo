package cate

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type Cate struct {
	id int
	name string
	create_at time.Time
}

func GetList(conn *sql.DB) []byte  {
	rows, err := conn.Query("select * from cate")

	if err != nil {
		fmt.Printf("mysql query fail ! [%s]", err)
	}

	var req []interface{}

	for rows.Next() {
		var c Cate
		var create_at time.Time
		err := rows.Scan(&c.id, &c.name, &create_at)
		c.create_at = create_at
		if err != nil {
			fmt.Printf("query mysql fail ! [%s]\n", err)
		}
		req = append(req, c)
	}
	reqJson , _ := json.Marshal(req)

	return reqJson
}
