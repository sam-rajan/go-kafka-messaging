package client

import (
	"errors"
	configreader "go-kafka-messaging/internal/pkg/config-reader"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var config kafka.ConfigMap

func init() {
	configFile := "configs/kafka-consumer.properties"
	config = configreader.ReadConfig(configFile)
}

func GetConfig() (kafka.ConfigMap, error) {
	if config == nil {
		return nil, errors.New("Config props not loaded yet")
	}

	return config, nil
}
