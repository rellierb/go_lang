package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "rellie:rellie/test?charset=utf8")

	// insert
	stmt, err := db.Prepare("INSERT userinfo SET username=")

}

