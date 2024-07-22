package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	DbConfig struct {
		Driver   string
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
		TimeZone string
		// Additional configurations⤵️
	}

	ApiConfig struct {
		ApiHost string
		ApiPort string
		// Additional configurations⤵️
	}

	FileLogConfig struct {
		FilePath string
		// Additional configurations⤵️
	}
)

type Config struct {
	DbConfig
	ApiConfig
	FileLogConfig
}

func (c *Config) LoadConfig() error {
	v := viper.New()
	v.AutomaticEnv()

	c.DbConfig = DbConfig{
		Driver:   v.GetString("APP_DB_DRIVER"),
		Host:     v.GetString("APP_DB_HOST"),
		Port:     v.GetString("APP_DB_PORT"),
		User:     v.GetString("APP_DB_USER"),
		Password: v.GetString("APP_DB_PASSWORD"),
		Name:     v.GetString("APP_DB_NAME"),
		SSLMode:  v.GetString("APP_SSL_MODE"),
		TimeZone: v.GetString("APP_TIME_ZONE"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: v.GetString("APP_API_HOST"),
		ApiPort: v.GetString("APP_API_PORT"),
	}

	c.FileLogConfig = FileLogConfig{
		FilePath: v.GetString("APP_FILE_PATH"),
	}

	if c.DbConfig.Driver == "" || c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.User == "" || c.DbConfig.Password == "" || c.DbConfig.Name == "" || c.DbConfig.SSLMode == "" || c.DbConfig.TimeZone == "" ||
		c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" ||
		c.FileLogConfig.FilePath == "" {
		return fmt.Errorf("missing required env var")
	}
	return nil
}

// NewConfig Constructor
func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.LoadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
