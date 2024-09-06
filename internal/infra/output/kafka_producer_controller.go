package output

import (
	"fmt"
	"kafkago/configs"
	"kafkago/internal/common/logger"
	kafka_producer "kafkago/internal/kafka/producer"
)

const (
	NUM_WORKERS = 100
)

func worker(cfg *configs.Config, kProducer *kafka_producer.KafkaProducer, logger logger.ILogger, countChan <-chan int) {
	for {
		key := fmt.Sprintf("%s:%d", "chave", <-countChan)
		err := kProducer.Produce(cfg.Topic_1, "messagem", key)
		if err != nil {
			logger.Error(fmt.Sprintf("erro writing to kafka topic: %w", err))
		} else {
			logger.Info(fmt.Sprintf("produced message %d", <-countChan))
		}
	}
}

func KafkaProducerController(kProducer *kafka_producer.KafkaProducer, cfg *configs.Config, logger logger.ILogger) {
	countChan := make(chan int, NUM_WORKERS)

	for i := 0; i < NUM_WORKERS; i++ {
		go worker(cfg, kProducer, logger, countChan)
	}

	count := 1
	for {
		countChan <- count
		count ++
	}

}
