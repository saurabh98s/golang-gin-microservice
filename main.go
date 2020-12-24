package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestModel is used for test data
type TestModel struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	r := gin.Default()

	// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
	v1 := r.Group("v1")
	// Adding middleware to the router.
	v1.Use(func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authentication")

		// if no authHeader is provided

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No token",
			})
			return
		}

	})

}
