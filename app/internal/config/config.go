package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DatabaseConfig struct {
		User     string `env:"POSTGRES_USER" env-default: "admin"`
		Password string `env:"POSTGRES_PASSWORD" env-default: "admin"`
		Name     string `env:"POSTGRES_DB" env-default: "task_manager"`
		Host     string `env:"POSTGRES_HOST" env-default: "localhost"`
		Port     string `env:"POSTGRES_PORT" env-default: "5432"`
	}
}

func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseConfig.User,
		c.DatabaseConfig.Password,
		c.DatabaseConfig.Host,
		c.DatabaseConfig.Port,
		c.DatabaseConfig.Name,
	)
}

func LoadConfig() (*Config, error) {
	log.Print("Read config")
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
