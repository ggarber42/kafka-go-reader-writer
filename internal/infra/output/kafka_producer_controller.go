package output

import (
	"fmt"
	"kafkago/configs"
	"kafkago/internal/common/logger"
	kafka_producer "kafkago/internal/kafka/producer"
)

func KafkaProducerController(kProducer *kafka_producer.KafkaProducer, cfg *configs.Config, logger logger.ILogger) {
	count := 1
	for {
		key := fmt.Sprintf("%s:%d", "chave", count)
		err := kProducer.Produce(cfg.ProducerTopic, "messagem", key)
		if err != nil {
			logger.Error(fmt.Sprintf("erro writing to kafka topic: %w", err))
		} else {
			logger.Info(fmt.Sprintf("processed message %d", count))
		}
		count ++
	}
}
