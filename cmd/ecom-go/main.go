package main

import (
	"log"
	"strconv"

	"github.com/AntonioMorales97/ecom-go/internal/config"
	"github.com/AntonioMorales97/ecom-go/internal/logger"
	"github.com/AntonioMorales97/ecom-go/pkg/api"
	"go.uber.org/zap"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to init Viper", zap.Error(err))
	}

	err = logger.InitializeZapCustomLogger(config.Env)
	if err != nil {
		log.Fatal("Failed to init logger")
	}

	server := api.NewServer()

	err = server.Start(config.Server.Host + ":" + strconv.Itoa(config.Server.Port))
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
