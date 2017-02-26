package conf

import "time"

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