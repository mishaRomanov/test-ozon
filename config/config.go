package config

import (
	"github.com/caarlos0/env/v10"
)

// seting up a config for custom values for postgres
type Config struct {
	Storage      string `env:"STORAGE"`
	User         string `env:"USER"`
	Address      string `env:"ADDRESS"`
	Port         string `env:"PORT"`
	DatabaseName string `env:"DB_NAME"`
	Password     string `env:"PASSWORD"`
}

func InitConfig() (Config, error) {
	var config Config
	//here we parse all environmental variables into struct object
	err := env.Parse(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
