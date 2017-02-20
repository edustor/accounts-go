package main

import (
	"github.com/edustor/accounts-go/app"
	"log"
	"net/http"
	"encoding/base64"
	"os"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func main() {
	privateKeyBase64 := os.Getenv("edustor.jwt.privatekey")

	log.Print("Key len: ", strings.Count(privateKeyBase64, ""))

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

	cfg := app.Config{
		Jwt: app.JwtConfig{
			RsaPrivateKey: key,
		},
	}
	router := app.Router(cfg)
	log.Fatal(http.ListenAndServe(":8080", router))
}
