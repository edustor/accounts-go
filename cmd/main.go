package main

import (
	"github.com/edustor/accounts-go/app"
	"log"
	"net/http"
)

func main() {
	router := app.Router()
	log.Fatal(http.ListenAndServe(":8080", router))
}

