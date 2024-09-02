package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ProducerTopic string `mapstructure:"PRODUCER_TOPIC"`
	ConsumerTopic string `mapstructure:"CONSUMER_TOPIC"`
	Brokers       string `mapstructure:"BROKERS"`
	Group         string `mapstructure:"GROUP"`
	GroupId       string `mapstructure:"GROUP_ID"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("main_config")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("fatal unable to decode struct: %w", err))
	}

	return cfg, nil
}
