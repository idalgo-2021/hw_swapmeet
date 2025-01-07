package config

import (
	"app_server/pkg/db/cache"
	"app_server/pkg/db/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCServerPort int    `env:"GRPC_SERVER_PORT" env-default:"9092"`
	JWTSecretKey   string `env:"JWT_SECRET_KEY"`
	postgres.PGConfig
	cache.RedisConfig
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("../../.env", &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
