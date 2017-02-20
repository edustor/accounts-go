package app

import "crypto/rsa"

type configKey int

const ConfigKey configKey = 0

type Config struct {
	Jwt JwtConfig
}

type JwtConfig struct {
	RsaPrivateKey *rsa.PrivateKey
}
