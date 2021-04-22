package services

import "golang-gin-microservice/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
