package cfg

import "time"

type contextKey int

const ConfigKey contextKey = 0

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