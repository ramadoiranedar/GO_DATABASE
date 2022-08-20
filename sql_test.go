package go_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	// transaction (tx)
	for i := 0; i < 10; i++ {
		email := "damar" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment-" + strconv.Itoa(i) + " from " + email

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID", id)
	}
	err = tx.Commit()
	// err = tx.Rollback() // rollback whatever u want in some condition
	if err != nil {
		panic(err)
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "damar" + strconv.Itoa(i) + "@gmail.com"
		comment := "comment-" + strconv.Itoa(i) + "from " + email

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment ID", id)
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "root@gmail.com"
	comment := "comment from root :D"
	script := "INSERT INTO comments(email, comment) VALUES(?, ?);"
	result, err := db.ExecContext(ctx, script, email, comment) // ExecContext didn't return any data, so it can't be get result data, for execute sql didn't return any data AS not query data
	if err != nil {
		panic(err)
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert the new COMMENT with ID:", insertID)
}

func TestExecContextSqlParemeter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "root"
	pw := "root"
	script := "INSERT INTO user(username, pw) VALUES(?, ?);"
	_, err := db.ExecContext(ctx, script, username, pw) // ExecContext didn't return any data, so it can't be get result data, for execute sql didn't return any data AS not query data
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Success insert the new USER")
}

func TestSqlInjectionSave(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	pw := "admin"

	script := "SELECT username FROM user WHERE username = ? AND pw = ? LIMIT 1"
	fmt.Println("script", script)
	rows, err := db.QueryContext(ctx, script, username, pw) // QueryContext will return data rows result of iterate all data from table
	if err != nil {
		panic(err)
	}
	defer rows.Close() // and row must be close

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("Success to login", username)
	} else {
		fmt.Println("Failed to login")
		return
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	pw := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND pw = '" + pw + "' LIMIT 1"
	fmt.Println("script", script)
	rows, err := db.QueryContext(ctx, script) // QueryContext will return data rows result of iterate all data from table
	if err != nil {
		panic(err)
	}
	defer rows.Close() // and row must be close

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("Success to login", username)
	} else {
		fmt.Println("Failed to login")
		return
	}
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, script) // QueryContext will return data rows result of iterate all data from table
	if err != nil {
		panic(err)
		return
	}
	defer rows.Close() // and row must be close

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("id:", id, " | ", "name:", name, " | ", "email:", email, " | ", "balance:", balance, " | ", "rating:", rating, " | ", "createdAt:", createdAt, " | ", "birthDate:", birthDate, " | ", "married:", married)
	}
}

func TestExecContextSqlInsert(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer(id, name) VALUES('3', 'dendi');"
	_, err := db.ExecContext(ctx, script) // ExecContext didn't return any data, so it can't be get result data, for execute sql didn't return any data AS not query data
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Success insert the new Customer")
}

func TestExecContextSqlUpdate(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "UPDATE customer SET name = 'ASEP' WHERE id = '2';"
	_, err := db.ExecContext(ctx, script) // ExecContext didn't return any data, so it can't be get result data, for execute sql didn't return any data AS not query data
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("Success insert the new Customer")
}
func TestQueryContextSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script) // QueryContext will return data rows result of iterate all data from table
	if err != nil {
		panic(err)
		return
	}
	defer rows.Close() // and row must be close

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println("Id:", id, "Name:", name)
	}
}

// MAPPING TIPE DATA GOLANG AND TIPE DATA DATABASE
// VARCHAR, CHAR -> string
// INT, BIGINT -> int32, int64
// FLOAT, DOUBLE -> float32, float64,
// BOOLEAN -> bool
// DATE, DATETIME, TIME, TIMESTAMP -> time.Time

// ALTER TABLE
// alter table customer add column email varchar(100), add column balance integer default 0, add column rating double default 0.0, add column created_at timestamp default current_timestamp, add column birth_date date, add column married boolean default false;

// create table user (
// 	username varchar(100) not null,
// 	pw varchar(100) not null,
// 	primary key (username)
// );

// INSERT INTO user(username, pw) values ('admin', 'admin');

// INSTER TABLE
// INSERT INTO customer (id, name, email, balance, rating, birth_date, married) VALUES
// ('1', 'Ferri', 'ferri@gmail.com', 1000000, 90.0, '1999-10-07', true),
// ('2', 'Deni', 'deni@gmail.com', 2000000, 85.5, '1998-12-10', true),
// ('3', 'Joko', 'joko@gmail.com', 7500000, 88.5, '1997-06-10', false);

// UDPDATE TABLE
// update customer set email = null, birth_date = null where id = '1';
