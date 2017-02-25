package cfg

type contextKey int

const ConfigKey contextKey = 0

//noinspection GoNameStartsWithPackageName
type Config struct {
	Jwt JwtConfig
}

func FromEnv() Config {
	return Config{
		Jwt: jwtConfigFromEnv(),
	}
}