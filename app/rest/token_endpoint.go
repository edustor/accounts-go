package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
)

func tokenEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseMultipartForm(10 * 1024 ^ 2)
	if err != nil {
		http.Error(w, fmt.Sprintf("Faild to parse form data %v", err), http.StatusBadRequest)
		return
	}
	grantType := r.Form.Get("grant_type")
	log.Println(r.Form)
	switch grantType {
	case "password":
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if username == "" || password == "" {
			http.Error(w, "'username' and 'password fields are required'", http.StatusBadRequest)
		}

		processPasswordGrant(username, password)
	default:
		http.Error(w, "Unsupported grant type", http.StatusBadRequest)
		return
	}
}

func processPasswordGrant(username string, password string) {

}