package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env  string `env:"ENV" envDefault:"dev"`
	Port string `port:"PORT" envDefault:"8080"`
}

func NewConfig(envPath string) (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
