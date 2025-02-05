package config

import (
	"log"
	"os"

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

func LoadConfig() (*Config, error) {
	log.Print("Read config")
	var cfg Config

	os.Setenv("POSTGRES_USER", "admin")
	os.Setenv("POSTGRES_PASSWORD", "admin")
	os.Setenv("POSTGRES_DB", "task_manager")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
