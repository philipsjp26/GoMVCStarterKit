package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func NewEnv() *EnvConfig {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
	}
	var configurations EnvConfig
	if err := viper.Unmarshal(&configurations); err != nil {
		log.Errorf("failed unmarshall viper got err :%v", err)
	}
	return &configurations
}
