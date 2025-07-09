package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Host string
		Port int
	}
	Storage struct {
		Port int
		Host string
	}
	Database struct {
		Host            string
		Port            int
		User            string
		Password        string
		Name            string
		MaxConns        int32
		MinConns        int32
		MaxConnIdleTime int32
		MaxConnLifetime int32
	}
	Redis struct {
		Host string
		Port int
		DB   int
	}
	JWTSecret struct {
		AccessSecret  string
		RefreshSecret string
	}
	Google struct {
		ClientID     string
		ClientSecret string
		RedirectURI  string
	}
	Quotes struct {
		ApiKey string
		ApiUrl string
	}
}

var AppConfig Config

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../app")
	viper.AddConfigPath("../todo-app-server")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return &AppConfig
}
