package go_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_golang")
	if err != nil { // catch error
		panic(err)
	}
	defer db.Close() // dont forget to close after using connection
}
