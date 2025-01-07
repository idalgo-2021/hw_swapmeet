package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PGConfig struct {
	User     string `env:"POSTGRES_USER" env-default:"pgadmin"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"12345"`
	Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	Database string `env:"POSTGRES_DB" env-default:"yandex"`
}

type DB struct {
	Db *sqlx.DB
}

func New(config PGConfig) (*DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", config.User, config.Password, config.Database, config.Host, config.Port)

	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if _, err := db.Conn(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to establish a database connection: %w", err)
	}

	return &DB{Db: db}, nil
}
