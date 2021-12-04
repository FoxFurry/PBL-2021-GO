package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig(configPath string) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Could not read config file: %v", err)
	}
}
