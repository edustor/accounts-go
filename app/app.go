package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"github.com/edustor/accounts-go/app/cfg"
	"time"
	"github.com/urfave/negroni"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	config, ok := r.Context().Value(cfg.ConfigKey).(cfg.Config)
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

	fmt.Fprint(w, signedToken)
}

func ConfigMiddleware(config cfg.Config) negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, cfg.ConfigKey, config)
		r = r.WithContext(ctx)
		next(rw, r)
	}
}

func Router(config cfg.Config) http.Handler {
	router := httprouter.New()
	router.GET("/", index)

	n := negroni.Classic()
	negroni.NewLogger()
	n.Use(ConfigMiddleware(config))
	n.UseHandler(router)

	return n
}
