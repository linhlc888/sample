package user

import (
	"database/sql"

	"os"
	//init mysql driver
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//User struct
type User struct {
	Fullname string
	Email    string
	Password string
}

//Insert User to DB
func Insert(u *User) error {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECT_STR"))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO user(fullname, email, password) values(?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Fullname, u.Email, hash)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//CheckLogin check credential with email and password
func CheckLogin(email string, password string) (bool, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECT_STR"))
	if err != nil {
		return false, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT fullname, email, password FROM user WHERE email=? LIMIT 1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var fullname, emailFromDB, passwordFromDB []byte
	err = stmt.QueryRow(email).Scan(&fullname, &emailFromDB, &passwordFromDB)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword(passwordFromDB, []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

//CheckExistingEmail check given email is existing or not
func CheckExistingEmail(email string) (bool, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECT_STR"))
	if err != nil {
		return false, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT email FROM user WHERE email=? LIMIT 1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var emailFromDB []byte
	err = stmt.QueryRow(email).Scan(&emailFromDB)
	if err != nil {
		return false, err
	}
	return true, nil
}
