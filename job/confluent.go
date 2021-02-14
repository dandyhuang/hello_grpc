package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	broker := "testip"
	group := "test"
	topics := []string{"test"}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     broker,
		"broker.address.family": "v4",
		"group.id":              group,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	fmt.Println(err)
	err = c.SubscribeTopics(topics, nil)
	fmt.Println(err)
	for {
	}
}
