package rest

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"github.com/edustor/accounts-go/app/conf"
	"time"
	"gopkg.in/gin-gonic/gin.v1"
)

func index(c *gin.Context) {
	_config, ok := c.Get("config")
	config, ok := _config.(conf.Config)
	if !ok {
		log.Panic("Can't get token from context")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(config.TokenExpirationTime).Unix(),
		"scope": "test",
		"sub": "test",
	})

	key := config.Jwt.RsaPrivateKey
	signedToken, err := token.SignedString(key)

	if err != nil {
		log.Panic(err)
	}

	c.String(200, signedToken)
}

func ConfigMiddleware(config conf.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) != 0 {
			c.Header("Content-Type", "application/json")
			c.JSON(0, map[string]interface{}{
				"success": "false",
				"errors": c.Errors.Errors(),
			})
		}
	}
}

func Router(config conf.Config) *gin.Engine {
	router := gin.Default()
	router.Use(ConfigMiddleware(config))
	router.Use(ErrorMiddleware())

	router.GET("/", index)
	router.POST("/oauth2/token", tokenEndpoint)

	return router
}
