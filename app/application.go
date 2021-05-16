package app

import (
	"github.com/gin-gonic/gin"
)

var (
	//router is a gin.Engine
	router = gin.Default()
)

// StartApplication initialises the application
func StartApplication() {
	mapURLs()
	router.Run(":8080")
}
