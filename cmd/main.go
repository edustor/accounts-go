package main

import (
	"log"
	"net/http"
	"github.com/edustor/accounts-go/app/cfg"
	"github.com/edustor/accounts-go/app/rest"
)

func main() {
	router := rest.Router(cfg.Default())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
