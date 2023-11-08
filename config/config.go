package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// seting up a config for custom values for postgres
type Config struct {
	User         string `mapstructure:"USER"`
	Adress       string `mapstructure:"ADDRESS"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	Password     string `mapstructure:"PASSWORD"`
	StorageType  string `mapstructure:"STORAGE_TYPE"`
}

func LoadConfig(path string) (Config, error) {
	var config Config
	//viper sets config path
	viper.AddConfigPath(path)
	//setting name and extension
	viper.SetConfigName("conf")
	viper.SetConfigType("env")

	//parse config file
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Error while reading config file %v", err)
		return Config{}, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Errorf("Error while reading parsing config data %v", err)
		return Config{}, err
	}
	return config, nil
}
