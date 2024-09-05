package make_controller

import (
	"kafkago/configs"
	"kafkago/internal/common/logger"
	"kafkago/internal/infra/output"

	"github.com/twmb/franz-go/pkg/kgo"
)

func MakeConsumerController(kclient *kgo.Client, cfg *configs.Config, logger logger.ILogger) {
	output.KafkaConsumerController(kclient, logger)
}
