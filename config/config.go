package config

import (
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV *Config

func LoadConfig() {
  viper.AddConfigPath(".")
  viper.SetConfigName(".env")
  viper.SetConfigType("env")

  if err := viper.ReadInConfig(); err != nil {
    log.Panic(err)
  }

  if err := viper.Unmarshal(&ENV); err != nil {
    log.Panic(err)
  }
}
