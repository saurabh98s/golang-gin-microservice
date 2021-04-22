package users

import (
	"fmt"
	"golang-gin-microservice/domain/users"
	"golang-gin-microservice/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil { //using shouldBIndJSOn() in place of reading the data from body and
		//   Unmarshalling it.
		// TODO: Handle Error
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		// TODO: Handle Error
		return
	}

	c.JSON(http.StatusCreated, result)

}

func FindUser(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me!")

}

func SearchUser(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me!")

}
