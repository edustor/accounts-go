package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config, ok := r.Context().Value(ConfigKey).(Config)
	if !ok {
		log.Panic("Can't get token from context")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"foo": "bar",
	})

	key := config.Jwt.RsaPrivateKey
	signedToken, err := token.SignedString(key)

	if err != nil {
		log.Panic(err)
	}

	fmt.Fprint(w, signedToken)

}

func ConfigMiddleware(h http.HandlerFunc, config interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ConfigKey, config)
		r = r.WithContext(ctx)
		h(w, r)
	}
}

func Router(config Config) http.Handler {
	router := httprouter.New()
	router.GET("/", index)
	return ConfigMiddleware(http.HandlerFunc(router.ServeHTTP), config)
}
