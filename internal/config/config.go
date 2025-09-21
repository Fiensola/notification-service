package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpAddr string `mapstructure:"HTTP_ADDR"`
}

func LoadConfig() (c *Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, nil
}
