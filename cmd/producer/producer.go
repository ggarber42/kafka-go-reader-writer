package main

import (
	"fmt"

	"kafkago/configs"
	logger "kafkago/internal/common/logger"
	"kafkago/internal/server"
)

func main() {
	log := logger.NewLogger()
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
	log.Info("start application")
	server.Start(cfg, log)
}
