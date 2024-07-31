package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerMQTTHost         string
	ServerMQTTPort         int
	ServerMQTTPassword     string
	ServerMQTTProtocol     string
	ServerMQTTUsername     string
	ServerMQTTEnableTLS    bool
	ServerMQTTValidateCert bool
	ServerMQTTCAPath       string // Add this line
	DatabaseHost           string
	DatabasePort           int
	DatabaseName           string
	DatabaseUser           string
	DatabasePassword       string
}

func LoadConfig(env string) (Config, error) {
	var config Config
	viper.SetConfigFile(".env." + env)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func InitConfig() Config {
	env := "prod" // Default to prod if no environment specified
	config, err := LoadConfig(env)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}
	return config
}
