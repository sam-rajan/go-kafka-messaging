package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type KafkaProperties struct {
	bootstrapServer  string
	securityProtocol string
	sasl             SecurityLayer
}

type SecurityLayer struct {
	mechanism string
	userName  string
	password  string
}

func (self *KafkaProperties) LoadProperties() {
	configFile, err := os.Open("go-kafka-messaging.properties")

	if err != nil {
		fmt.Printf("Failed to read config file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(configFile)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			keyval := strings.Split(line, "=")
			self.sasl = SecurityLayer{}

			switch keyval[0] {
			case "bootstrap.servers":
				self.bootstrapServer = keyval[1]
			case "security.protocol":
				self.securityProtocol = keyval[1]
			case "sasl.mechanisms":
				self.sasl.mechanism = keyval[1]
			case "sasl.password":
				self.sasl.userName = keyval[1]
			case "sasl.username":
				self.sasl.userName = keyval[1]
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Failed to read file: %s", err)
			os.Exit(1)
		}
	}

}
