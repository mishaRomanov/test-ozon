package config

import (
	"github.com/caarlos0/env/v10"
)

// seting up a config for custom values for postgres
type Config struct {
	User         string `env:"USER"`
	Adress       string `env:"ADDRESS"`
	DatabaseName string `env:"DB_NAME"`
	Password     string `env:"PASSWORD"`
}

func InitConfig() (Config, error) {
	var config Config
	//viper sets config path
	err := env.Parse(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
