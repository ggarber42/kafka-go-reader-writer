package main

import (
	"fmt"
	"kafkago/configs"
	"kafkago/internal/common/logger"
	"kafkago/internal/server"
	"kafkago/internal/utils"
	"os"
)

func main() {
	logger := logger.NewLogger()
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		logger.Error(fmt.Sprintf("erro loading config"))
		os.Exit(utils.EXIT_FAILURE)
	}
	server.StartConsumerServer(cfg, logger)
}