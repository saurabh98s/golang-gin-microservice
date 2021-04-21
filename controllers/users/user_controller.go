package users

import (
	"encoding/json"
	"fmt"
	"golang-gin-microservice/domain/users"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var users users.User
	fmt.Println(users)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handle Error
	}
	if err := json.Unmarshal(bytes, &users); err != nil {

	}
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.JSON(http.StatusNotImplemented, "implement me!")

}

func FindUser(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me!")

}

func SearchUser(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, "implement me!")

}
