package conn

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type UserDB struct {
	db *sql.DB
}

func init() {
	var dbase UserDB
	// Save connection properties.
	cfg := mysql.Config{
		User:                 "ARTHEMIS",
		Passwd:               "moleculr22",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "user",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	dbase.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
}
