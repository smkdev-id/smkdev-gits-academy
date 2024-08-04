package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// application configuration settings
type Config struct {
	Env 		string			`env:"ENV" evDefault:"dev"`
	Host    	string	        `env:"HOST" envDefault:"localhost"`
	Port		string			`env:"PORT" envDefault:"8080"`
	Sqlite 		string 			`env:"SQLITE_DATABASE_PATH"`
}

// NewConfig loads configuration from an env file
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