package config

import (
	"auth_service/pkg/db/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCServerPort          int    `env:"GRPC_SERVER_PORT" env-default:"9090"`
	JWTSecretKey            string `env:"JWT_SECRET_KEY"`
	JWTAccessTokenLifetime  int    `env:"JWT_ACCESS_TOKEN_LIFETIME"`
	JWTRefreshTokenLifetime int    `env:"JWT_REFRESH_TOKEN_LIFETIME"`
	postgres.Config
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("../../.env", &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
