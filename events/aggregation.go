package events

import (
	"followers-ms/persistence"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Client *kafka.Producer
}

type AppProducer interface {
	FollowProducer(request persistence.RelationshipRequest, topic string) error
}

func (k *KafkaProducer) FollowProducer(request persistence.RelationshipRequest, topic string) error {

	var err error
	deliveryChan := make(chan kafka.Event, 10000)
	// send Schema capn'p
	err = k.Client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(request.ToUserId)},
		deliveryChan,
	)
	if err != nil {
		return err
	}
	return nil

}
