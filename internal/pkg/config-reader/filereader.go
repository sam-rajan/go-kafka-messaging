package configreader

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	if configFile == "" {
		log.Fatal("Failed to config file")
	}

	config, err := os.Open(configFile)

	if err != nil {
		log.Fatalf("Failed to read config file %s", err)
	}

	scanner := bufio.NewScanner(config)
	configMap := make(map[string]kafka.ConfigValue)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			keyval := strings.Split(line, "=")

			parameter := strings.TrimSpace(keyval[0])
			value := strings.TrimSpace(keyval[1])

			configMap[parameter] = value
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}
	}

	return configMap
}
