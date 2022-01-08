package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Env    string
	Db     DbConfig
	Token  TokenConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type DbConfig struct {
	Driver string
	Source string
}

type TokenConfig struct {
	SymmetricKey   string `mapstructure:"symmetric_key"`
	AccessDuration time.Duration
}

func LoadConfig(path string) (config Config, err error) {

	// Set config path, name, and type
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	// Automatically read values from environment variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
