package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	USERNAME = "taufik"
	PASSWORD = "secret"
)

func Auth(c *gin.Context) bool {

	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Unauthorizhed",
		})
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"massage": "Unauthorizhed",
		})
		return false
	}
	return true
}
