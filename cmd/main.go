package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hackstock/webmock/pkg/api"
	"github.com/hackstock/webmock/pkg/parsing"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed initializing logger : %v", err)
	}

	endpoints, err := parsing.ParseEndpoints(viper.GetString("WEBMOCK_ENDPOINT_DEFINITIONS"))

	if err != nil {
		logger.Fatal("failed parsing endpoints definitions", zap.Error(err))
	}

	router := gin.Default()
	server := api.NewServer(router, endpoints, logger)

	port := viper.GetInt("WEBMOCK_SERVER_PORT")
	err = server.Run(port)
	if err != nil {
		logger.Fatal("failed starting server", zap.Int("port", port))
	}
}
