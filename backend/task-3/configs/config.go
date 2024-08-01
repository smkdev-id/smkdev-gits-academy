package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:"5000"`
	AppEnv  string `env:"APP_ENV" envDefault:"development"`
}

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
	c.AppPort = v.GetString("APP_PORT")
	c.AppEnv = v.GetString("APP_ENV")
	return nil
}
