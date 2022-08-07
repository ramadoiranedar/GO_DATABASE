package go_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(host:3306)/my_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
