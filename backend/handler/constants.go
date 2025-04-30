package handler

import "database/sql"

var db *sql.DB

func SetDB(DB *sql.DB) {
	db = DB
}
