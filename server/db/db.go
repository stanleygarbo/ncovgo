package db

import (
	"database/sql"
)

var db *sql.DB
var err error

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:pw@tcp(server_mysql_1:3306)/covid_tracker")
	if err != nil{
		return nil, err
	}
	db.SetMaxIdleConns(100)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return
}

