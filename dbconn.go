package authsystem

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var (
	Admin, Pasd, DbName, Addr string
	err                       error
	connErr                   = errors.New("unable to connect with database")
	readErr                   = errors.New("could not read request from file")
)

type UserDB struct {
	Db *sql.DB
}

func init() {
	var dbase UserDB

	//Save Connection config.
	cfg := mysql.Config{
		User:                 Admin,
		Passwd:               Pasd,
		Net:                  "tcp",
		Addr:                 Addr,
		DBName:               DbName,
		AllowNativePasswords: true,
	}

	// Get a DB handle.
	//cfg.FormatDSN is driver specific parameter
	dbase.Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(connErr)
	}
}

//Get Db User
func GetUser(f string, n int64) string {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, n)
	_, err = file.Read(buf)

	if err != nil {
		panic(readErr)
	}
	return string(buf[:n])
}

func GetPasswd(f string, n int64) string {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, n)
	_, err = file.Read(buf)

	if err != nil {
		panic(readErr)
	}
	return string(buf[:n])
}
