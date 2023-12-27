package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}
	topic := "coordinators"

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatal("error create new Producer")
	}
	//send messages
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message-%d", i)
		produceMsg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(value),
		}
		err = producer.Produce(produceMsg, nil)
		time.Sleep(time.Second * 1)
		if err != nil {
			log.Fatalf("error send message: %s-  topic: %s \n", value, topic)
		} else {
			fmt.Printf("Produced message: %s -  topic: %s \n", value, topic)
		}
	}

	producer.Flush(15 * 1000)
	producer.Close()
}
