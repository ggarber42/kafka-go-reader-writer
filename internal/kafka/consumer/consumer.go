package consumer

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaConsumer struct {
	Client *kgo.Client
}

func NewConsumerHandler (client *kgo.Client) (*KafkaConsumer) {
	return &KafkaConsumer{
		Client: client,
	}
}

func (kc *KafkaConsumer) Consume() kgo.Fetches {
	ctx := context.Background()
	fetches := kc.Client.PollFetches(ctx)
	return fetches
}
