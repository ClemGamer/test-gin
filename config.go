package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var AppConfig *appConfig

// load config.json to AppConfig
func LoadAppConfig() {
	fmt.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

}

type appConfig struct {
	DatabaseConfig databaseConfig `mapstructure:"database"`
	Port           string         `mapstructure:"port"`
}

type databaseConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
}
