package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() *sql.DB {
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_pbp_tools?parseTime=true&loc=Asia%2FJakarta")
	db, err := sql.Open("mysql", "root:M4D3ENA@tcp(localhost:3306)/db_latihan_pbp_tools")
	if err != nil {
		log.Fatal(err)
	}
	return db
}