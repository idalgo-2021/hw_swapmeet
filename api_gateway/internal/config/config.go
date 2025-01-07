package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServerAddress  string `env:"HTTP_SERVER_ADDRESS" env-default:"localhost"`
	HTTPServerPort     int    `env:"HTTP_SERVER_PORT" env-default:"8090"`
	AuthServiceURL     string `env:"AUTH_SERVICE_URL" env-default:"localhost:50061"`
	SwapmeetServiceURL string `env:"SWAPMEET_SERVICE_URL" env-default:"localhost:50062"`
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("../../.env", &cfg)
	if err != nil {
		return nil
	}
	return &cfg
}
