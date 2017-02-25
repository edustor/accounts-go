package main

import (
	"github.com/edustor/accounts-go/app"
	"log"
	"net/http"
	"github.com/edustor/accounts-go/app/cfg"
)

func main() {
	router := app.Router(cfg.FromEnv())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
