package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Client struct {
	c kgo.Client
}

func NewProducer(brokers []string) (*kgo.Client, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup("my-group"),
		kgo.ConsumeTopics("foo"),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Publish(p *kgo.Client, topicName string, value string) error {
	var wg sync.WaitGroup
	wg.Add(1)
	record := &kgo.Record{Topic: topicName, Value: []byte(value)}
	ctx := context.Background()
	p.Produce(ctx, record, func(_ *kgo.Record, err error) {
		defer wg.Done()
		if err != nil {
			fmt.Printf("record had a produce error: %v\n", err)
		}

	})
	wg.Wait()
	return nil
}
