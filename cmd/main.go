package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	// defines expected config file name and type
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	// defines default values for important configs
	viper.SetDefault("WEBMOCK_SERVER_PORT", "9000")
	viper.SetDefault("WEBMOCK_ENDPOINT_DEFINITIONS", "endpoints.json")
	// defines paths from which configs could be loaded
	viper.AddConfigPath("$HOME/.webmock")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("config file not found : %v", err)
		} else {
			log.Fatalf("failed processing config file : %v", err)
		}
	}
}
