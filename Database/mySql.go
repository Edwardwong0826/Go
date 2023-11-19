package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	username string
	password string
}

func queryOneRow() {
	db, err := sql.Open("mysql", "root:61376554@/go")

	sql := "select * from user where id = ?"
	var u User
	db.QueryRow(sql, 1).Scan(&u.id, &u.username, &u.password)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("u: %v\n", u)
	}

}

func queryManyRow() {
	db, err := sql.Open("mysql", "root:61376554@/go")

	sql := "select * from user"
	rows, err := db.Query(sql)
	var u User

	defer rows.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for rows.Next() {
			rows.Scan(&u.id, &u.username, &u.password)
			fmt.Printf("u: %v\n", u)
		}
	}

}

func insert(username string, password string) {
	db, err := sql.Open("mysql", "root:61376554@/go")

	sql := "insert into user(username, PASSWORD) values(?,?)"
	result, err := db.Exec(sql, username, password)

	if err != nil {
		fmt.Printf("insert failed, err%v\n", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert ID failed, err%v\n", err)
		return
	}

	fmt.Printf("insert success, ths id is: %d\n", id)
}

func update(username string, password string, id int) {
	db, err := sql.Open("mysql", "root:61376554@/go")

	sql := "update user set username=?, password=? where id=?"
	result, err := db.Exec(sql, username, password, id)

	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("update rows failed , err: %v\n", err)
		return
	}

	fmt.Printf("update success, update row : %d \n", rows)

}

func delete(id int) {

	db, err := sql.Open("mysql", "root:61376554@/go")

	sql := "delete from user where id = ?"
	result, err := db.Exec(sql, id)

	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("delete rows failed , err: %v\n", err)
		return
	}

	fmt.Printf("delete success, delete rows : %d \n", rows)

}

// https://pkg.go.dev/github.com/go-sql-driver/mysql#section-readme
// run go get -u github.com/go-sql-driver/mysql in terminal
func main() {

	// db, err := sql.Open("mysql", "root:61376554@/go")
	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	insert("ed", "123456")
	queryOneRow()
	queryManyRow()
	update("ed", "654321", 5)
	delete(5)
}
