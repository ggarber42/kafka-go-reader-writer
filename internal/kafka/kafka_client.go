package kafka

import (
	"strings"
	"time"

	"kafkago/configs"

	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewKafkaClient(cfg *configs.Config) (*kgo.Client, error) {
	instanceID := uuid.New()
	client, err := kgo.NewClient(
		kgo.SeedBrokers(strings.Split(cfg.Brokers, ",")...),
		kgo.ConsumerGroup(cfg.Group),
		kgo.InstanceID(cfg.GroupId+instanceID.String()),
		kgo.FetchMaxWait(500*time.Millisecond),
		kgo.MaxBufferedRecords(100),
		kgo.RebalanceTimeout(80000*time.Millisecond),
		kgo.SessionTimeout(10000*time.Millisecond),
		kgo.HeartbeatInterval(1000*time.Millisecond),
		kgo.RetryTimeout(500*time.Millisecond),
		kgo.DialTimeout(10*time.Minute),
		kgo.DisableAutoCommit(),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
