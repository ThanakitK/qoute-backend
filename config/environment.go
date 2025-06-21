package config

import (
	"log"

	"github.com/spf13/viper"
)

var Env = struct {
	DBURI      string `mapstructure:"DB_URI" validate:"required"`
	DBName     string `mapstructure:"DB_NAME" validate:"required"`
	Cors       string `mapstructure:"CORS"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}{
	Cors:       "*",
	JWT_SECRET: "secret",
}

func NewAppInitEnvironment() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println(".env file not found, loading from environment variables only.")
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatalf("Unable to unmarshal environment variables: %s", err)
	}

	log.Println("Environment variables loaded successfully.")
}
