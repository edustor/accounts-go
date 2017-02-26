package conf

import (
	"time"
	"gopkg.in/gin-gonic/gin.v1"
	"errors"
)

type contextKey int

//noinspection GoNameStartsWithPackageName
type Config struct {
	Jwt JwtConfig
	TokenExpirationTime time.Duration
}

func Default() Config {
	return Config{
		Jwt: jwtConfigFromEnv(),
		TokenExpirationTime: 10 * time.Minute,
	}
}

func FromGinContext(ctx *gin.Context) (Config, error) {
	_config, ok := ctx.Get("config")
	if ok != true {
		return Config{}, errors.New("Cannot get config from GIN context")
	}

	config, ok := _config.(Config)
	if ok != true {
		return Config{}, errors.New("Cannot convert config from GIN context to conf.Config")
	}

	return config, nil
}