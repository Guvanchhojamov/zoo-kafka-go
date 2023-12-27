package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "console-consumer-1",
		"auto.offset.reset": "earliest",
	}

	//create kafka consume

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("error create new consumer... %s", err.Error())
	}
	err = consumer.SubscribeTopics([]string{"coordinators"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Resived message: %s\n", string(message.Value))
		}
	}
	consumer.Close()

}
