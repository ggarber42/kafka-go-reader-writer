package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Topic_1 string `mapstructure:"TOPIC_1"`
	Brokers       string `mapstructure:"BROKERS"`
	Group         string `mapstructure:"GROUP"`
	GroupId       string `mapstructure:"GROUP_ID"`
	ProducerServerPort string `mapstructure:"PRODUCER_SERVER_PORT"`
	ConsumerServerPort string `mapstructure:"CONSUMER_SERVER_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("main_config")
	viper.SetConfigFile(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
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
