package configreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProperties struct {
	Value kafka.ConfigMap
}

func (self *KafkaProperties) LoadProperties(configFileName string) {
	configFile, err := os.Open(configFileName)

	if err != nil {
		fmt.Printf("Failed to read config file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(configFile)
	self.Value = make(map[string]kafka.ConfigValue)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			keyval := strings.Split(line, "=")

			parameter := strings.TrimSpace(keyval[0])
			value := strings.TrimSpace(keyval[1])

			self.Value[parameter] = value
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Failed to read file: %s", err)
			os.Exit(1)
		}
	}

}
