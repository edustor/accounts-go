package main

import (
	"log"
	"github.com/edustor/accounts-go/app/rest"
	"github.com/edustor/accounts-go/app/conf"
)

func main() {
	router := rest.Router(conf.Default())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	router.Run(":8081")
}
