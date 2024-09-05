package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"kafkago/configs"
	"kafkago/internal/common/logger"
	"kafkago/internal/kafka"
	make_controller "kafkago/internal/make"
	"kafkago/internal/utils"
)

func Start(cfg *configs.Config, logger logger.ILogger) {
	done := make(chan bool)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", getHealth)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: mux,
	}

	kclient, err := kafka.NewKafkaClient(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("%w", err))
		os.Exit(utils.EXIT_FAILURE)
	}

	go func() {

		go make_controller.MakeKafkaProducerController(kclient, cfg, logger)
		logger.Info(fmt.Sprintf("Start server at port 0.0.0.0:%s", cfg.ServerPort))
		err = server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server closed")
		} else {
			logger.Error(fmt.Sprintf("Error starting server: %s", err.Error()))
			os.Exit(utils.EXIT_FAILURE)
		}

		done <- true
	}()

	<-done

	logger.Info("server finish")
}
