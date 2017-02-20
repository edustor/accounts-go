package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world")
}

func Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	return router
}
