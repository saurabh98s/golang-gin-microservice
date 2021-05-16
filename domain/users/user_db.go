package users

import (
	"fmt"
	"golang-gin-microservice/utils/errors"
	"strings"
	"time"

	"golang-gin-microservice/db"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users where id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("id %d does not exists", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to GET userID %d: %s", user.Id, err.Error()))
	}
	return nil
}

// Save saves user info in the DB
func (user *User) Save() *errors.RestErr {
	stmt, err := db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = time.Now().Format("2006-01-02 15:04:05")
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when saving user: %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when fetching last insert id: %s", err.Error()))
	}

	user.Id = userID
	return nil
}
