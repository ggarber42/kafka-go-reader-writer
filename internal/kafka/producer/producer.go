package kafka_producer

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaProducer struct {
	Client *kgo.Client
}

func NewProducerHandler(client *kgo.Client) *KafkaProducer {
	return &KafkaProducer{
		Client: client,
	}
}

func (kp *KafkaProducer) Produce(topicName string, value string, key string) error {
	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{Topic: topicName, Value: []byte(value), Key: []byte(key)}
	ctx := context.Background()
	kp.Client.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()
	return nil
}
