package configreader

import "github.com/confluentinc/confluent-kafka-go/kafka"

type ConfigReader interface {
	ReadConfig() kafka.ConfigMap
}
