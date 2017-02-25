package cfg

import (
	"crypto/rsa"
	"os"
	"log"
	"github.com/dgrijalva/jwt-go"
	"encoding/base64"
	"errors"
)

type JwtConfig struct {
	RsaPrivateKey *rsa.PrivateKey
}

func jwtConfigFromEnv() JwtConfig {
	privateKeyBase64 := os.Getenv("edustor.jwt.privatekey")

	if len(privateKeyBase64) == 0 {
		panic(errors.New("JWT key (edustor.jwt.privatekey env) cannot be empty"))
	}

	keyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		log.Panic(err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		log.Panic(err)
	}

	err = key.Validate()
	if err != nil {
		log.Panic(err)
	}

	return JwtConfig{
		RsaPrivateKey: key,
	}
}