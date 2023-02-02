package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	// responseMessage := gin.H{
	// 	"message": "Pong",
	// }

	c.JSON(http.StatusOK, "Pong")
}
