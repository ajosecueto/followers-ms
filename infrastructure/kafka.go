package infrastructure

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func NewKafkaConsumer() *kafka.Consumer {
	server := os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	group := os.Getenv("KAFKA_GROUP")
	client := os.Getenv("KAFKA_CLIENT")

	log.Print("KAFKA CONSUMER ", server, group, client)

	data, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"client.id":         client,
		"group.id":          group,
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		panic(err)
	}

	return data
}

func NewKafkaProducer() *kafka.Producer {

	server := os.Getenv("KAFKA_BOOTSTRAP_SERVER")
	group := os.Getenv("KAFKA_GROUP")
	client := os.Getenv("KAFKA_CLIENT")

	data, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"client.id":         client,
		"group.id":          group,
	})
	if err != nil {
		panic(err)
	}

	return data
}
