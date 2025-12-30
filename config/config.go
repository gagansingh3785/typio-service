package config

import (
	"errors"
	"fmt"

	"github.com/gagansingh3785/typio-service/utils"
	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"SERVER"`
	Logger     LoggerConfig     `mapstructure:"LOGGER"`
	DB         DBConfig         `mapstructure:"DB"`
	Migrations MigrationsConfig `mapstructure:"MIGRATIONS"`
}

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

type LoggerConfig struct {
	Level string `mapstructure:"LEVEL"`
}

type DBConfig struct {
	// TODO: Add more database configuration to tune
	// connection performance
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

type MigrationsConfig struct {
	Path string `mapstructure:"PATH"`
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

	if utils.IsZero(c.DB.Host) {
		return errors.New("database host is required")
	}

	if utils.IsZero(c.DB.Port) {
		return errors.New("database port is required")
	}

	if utils.IsZero(c.DB.User) {
		return errors.New("database user is required")
	}

	if utils.IsZero(c.DB.Password) {
		return errors.New("database password is required")
	}

	if utils.IsZero(c.DB.DBName) {
		return errors.New("database name is required")
	}

	if utils.IsZero(c.Migrations.Path) {
		return errors.New("migrations path is required")
	}

	return nil
}

func (c *Config) GetServerAddr() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}

func (c *Config) GetDBURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.DBName)
}
