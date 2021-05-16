package users

import (
	"fmt"
	"golang-gin-microservice/utils/errors"
	"log"
	"time"

	"golang-gin-microservice/db"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil

}

// Save saves user info in the DB
func (user *User) Save() *errors.RestErr {

	if err := db.Client.Ping(); err != nil {
		return errors.NewInternalServerError("prepare statement error: " + err.Error())
	}
	log.Println("after ping")
	stmt, err := db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("prepare statement error: " + err.Error())
	}
	defer stmt.Close()
	user.DateCreated = time.Now().Format("2006-01-02 15:04:05")
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when saving user: %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when fetching last insert id: %s", err.Error()))
	}

	user.Id = userID
	return nil
}
