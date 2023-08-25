package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type ProjectConfig struct {
	Id     string `mapstructure:"id"`
	Region string `mapstructure:"region"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type RedisConnectionConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type PsqlConnectionConfig struct {
	Uri string `mapstructure:"uri"`
}

type Config struct {
	Server  ServerConfig          `mapstructure:"server"`
	Logging LoggerConfig          `mapstructure:"logging"`
	Redis   RedisConnectionConfig `mapstructure:"redis"`
	Psql    PsqlConnectionConfig  `mapstructure:"psql"`
	Project ProjectConfig         `mapstructure:"project"`
	Env     string
}

func NewConfig(configFilePath string, env string) *Config {
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Unable to decode Config: %s \n", err))
	}

	config.Env = env

	return &config
}
