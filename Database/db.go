package Database

import (
	"Natural/Controllers"
	"database/sql"
	"fmt"
)

func StartMySQL() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/productdb")
	if err != nil {
		fmt.Println(err)
	}
	Controllers.Database = db
}
