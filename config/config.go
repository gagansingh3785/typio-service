package config

import (
	"errors"

	"github.com/gagansingh3785/typio-service/utils"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"SERVER"`
	Logger LoggerConfig `mapstructure:"LOGGER"`
}

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

type LoggerConfig struct {
	Level string `mapstructure:"LEVEL"`
}

func SetupConfig() (*Config, error) {
	var config Config

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// validate config
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) Validate() error {
	if utils.IsZero(c.Server.Port) {
		return errors.New("server port is required")
	}

	if utils.IsZero(c.Server.Host) {
		return errors.New("server host is required")
	}

	return nil
}
