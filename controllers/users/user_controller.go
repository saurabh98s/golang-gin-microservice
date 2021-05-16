package users

import (
	"fmt"
	"golang-gin-microservice/domain/users"
	"golang-gin-microservice/services"
	"golang-gin-microservice/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil { //using shouldBIndJSOn() in place of reading the data from body and
		//   Unmarshalling it.
		restError := errors.NewBadRequestError("invalid json body")
		c.JSON(restError.Status, restError)
		return
	}
	fmt.Printf("First Name %s , Last Name %s ,Date Created %s", user.FirstName, user.LastName, user.DateCreated)
	result, err := services.CreateUser(user)
	if result == nil {
		return
	}
	if err != nil {
		c.JSON(err.Status, err.Error)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func FindUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid User ID, should be a number")
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, result)

}

func SearchUser(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me!")

}
