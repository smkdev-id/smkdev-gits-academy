package config

import "github.com/spf13/viper"

type (
	DbConfig struct {
		DB_DRIVER   string
		DB_HOST     string
		DB_PORT     string
		DB_USER     string
		DB_PASSWORD string
		DB_NAME     string
	}

	Config struct {
		DB DbConfig
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.GetConfig()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) GetConfig() error {
	v := viper.New()
	v.AutomaticEnv()

	c.DB = DbConfig{
		DB_DRIVER:   v.GetString("DB_DRIVER"),
		DB_HOST:     v.GetString("DB_HOST"),
		DB_PORT:     v.GetString("DB_PORT"),
		DB_USER:     v.GetString("DB_USER"),
		DB_PASSWORD: v.GetString("DB_PASSWORD"),
		DB_NAME:     v.GetString("DB_NAME"),
	}

	return nil
}
