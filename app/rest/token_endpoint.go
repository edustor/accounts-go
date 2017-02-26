package rest

import (
	"gopkg.in/gin-gonic/gin.v1"
	"errors"
)

func tokenEndpoint(c *gin.Context) {
	grantType := c.PostForm("grant_type")
	switch grantType {
	case "password":
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.AbortWithError(400, errors.New("'username' and 'password fields are required'"))
		}

		processPasswordGrant(username, password)
	case "":
		c.Status(400)
		c.Error(errors.New("grant_type field is required"))
	default:
		c.Status(400)
		c.Error(errors.New("Unsupported grant type"))
		return
	}
}

func processPasswordGrant(username string, password string) {

}