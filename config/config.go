package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURI string
	RedisURI string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile("config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	AppConfig = Config{
		MongoURI: viper.GetString("MONGO_URI"),
		RedisURI: viper.GetString("REDIS_URI"),
	}
}
