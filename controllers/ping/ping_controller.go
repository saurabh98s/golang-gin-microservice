package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	c.JSON(http.StatusOK, "everything good here!")

}
