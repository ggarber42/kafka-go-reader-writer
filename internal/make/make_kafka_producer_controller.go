package make_controller

import (
	"kafkago/configs"
	"kafkago/internal/infra/output"
	kafka_producer "kafkago/internal/kafka/producer"

	"github.com/twmb/franz-go/pkg/kgo"
)

func MakeKafkaProducerController(kclient *kgo.Client, cfg *configs.Config) {
	kProducer := kafka_producer.NewProducerHandler(kclient)
	output.KafkaProducerController(kProducer)
}
