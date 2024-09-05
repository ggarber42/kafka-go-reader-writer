package output

import (
	"context"
	"fmt"
	"kafkago/internal/common/logger"

	"github.com/twmb/franz-go/pkg/kgo"
)

func KafkaConsumerController(kclient *kgo.Client, logger logger.ILogger) {
	logger.Info("running consumer")
	ctx := context.Background()
	for {
		fetches := kclient.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			logger.Error(fmt.Sprintf("error consuming message: %w", errs))
			continue
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			logger.Info(fmt.Sprintf("read message: %s", string(record.Key)))
		}
	}
}
