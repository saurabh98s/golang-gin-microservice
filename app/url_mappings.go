package app

import (
	"golang-gin-microservice/controllers/ping"
	"golang-gin-microservice/controllers/users"
)

func mapURLs() {

	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.FindUser)
	router.GET("/user/search", users.SearchUser)

}
