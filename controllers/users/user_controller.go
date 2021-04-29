package users

import (
	"fmt"
	"golang-gin-microservice/domain/users"
	"golang-gin-microservice/services"
	"golang-gin-microservice/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil { //using shouldBIndJSOn() in place of reading the data from body and
		//   Unmarshalling it.
		restError := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad request",
		}
		c.JSON(restError.Status, restError)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
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
