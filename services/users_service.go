package services

import (
	"golang-gin-microservice/domain/users"
	"golang-gin-microservice/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {

	result := users.User{Id: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return &result, nil

}
