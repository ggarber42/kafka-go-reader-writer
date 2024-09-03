package output

import (
	"fmt"
	kafka_producer "kafkago/internal/kafka/producer"
	"kafkago/internal/utils"
	"os"
)

func KafkaProducerController(kProducer *kafka_producer.KafkaProducer) {
	for char := 'A'; char <= 'Z'; char++ {
		err := kProducer.Produce("aaaa", fmt.Sprintf("%c", char))
		if err != nil {
			fmt.Println("corram para as montanhas", err)
			os.Exit(utils.EXIT_FAILURE)
		}
	}
}
