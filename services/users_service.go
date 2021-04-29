package services

import (
	"golang-gin-microservice/domain/users"
	"golang-gin-microservice/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
