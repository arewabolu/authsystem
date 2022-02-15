package authsystem

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var (
	User, Pasd, DbName, Addr string
	err                      error
	connerr                  = errors.New("Unable to connect with database")
	readerr                  = errors.New("Could not read from file")
)

type UserDB struct {
	Db *sql.DB
}

func init() {
	var dbase UserDB

	//Save Connection config.
	cfg := mysql.Config{
		User:                 User,
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
		log.Fatal(connerr)
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
		panic(readerr)
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
		panic(readerr)
	}
	return string(buf[:n])
}
