package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	config := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9095",
	}
	topic := "coordinators"

	producer, err := kafka.NewProducer(&config)
	if err != nil {
		log.Fatal("error create new Producer")
	}
	//send messages
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message-%d\n", i)
		produceMsg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(value),
		}
		err = producer.Produce(produceMsg, nil)
		if err != nil {
			log.Fatalf("error send message: %s-  topic: %s \n", value, topic)
		} else {
			fmt.Printf("Produces message: %s -  topic: %s ", value, topic)
		}
	}

	producer.Flush(15 * 1000)
	producer.Close()
}
