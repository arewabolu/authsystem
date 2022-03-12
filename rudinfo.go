package authsystem

import (
	"auth_system/usersinfo"
	"errors"
	"log"

	logger "github.com/arewabolu/Errors"
	randomid "github.com/arewabolu/urlgenerator"
)

func DbComms() {

}

/*
1. Check if user email exists;
2. If true return user id & email;
//1 & 2 are 1 function
3. generate a token and send to Database
4. send token to user email;(1 Functiom)
5. Check if token entered by user exists using id;
6. If true return to user page
*/
//remember to add maximum time limit for token storage

//Executes Tasks 1 & 2 above
func Query4Id(uEmail string) (string, usersinfo.Users) {

	var (
		db          UserDB
		Id, dbEmail usersinfo.Users
	)

	que, err := db.Db.Prepare("SELECT id, email from userData WHERE email=?")
	if err != nil {
		logger.ErrLogger(errors.New("unable to query database"), "log.txt")
	}
	defer que.Close()

	row, err := que.Query(uEmail)
	if err != nil {
		logger.ErrLogger(err, "log.txt")
	}

	for row.Next() {
		row.Scan(&Id, &dbEmail)
	}

	err = row.Err()
	if err != nil {
		logger.ErrLogger(err, "log.txt")
	}

	token := randomid.Generator(6)

	return token, Id

}

// Executes task 3 as above
func InsToken(token string, Id usersinfo.Users) {
	var (
		db UserDB
	)
	que, err := db.Db.Prepare("INSERT INTO userData(tempLogin) VALUES(?)")
	if err != nil {
		logger.ErrLogger(err, "log.txt")
	}
	_, err = que.Exec(token)
	if err != nil {
		logger.ErrLogger(err, "log.txt")
	}
}

//Executes Task 5
func Query4Token(uToken string, Id usersinfo.Users) (bool, error) {
	var (
		db      UserDB
		dbToken string
		err     error
		opt     bool
	)
	stmt, err := db.Db.Prepare("SELECT token from users where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(Id).Scan(&dbToken)
	if err != nil {
		log.Fatal(err)
	}
	if uToken != dbToken {
		opt = false
		return opt, errors.New("incorrect token entered")

	}
	return opt, err
}
