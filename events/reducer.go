package events

import (
	"fmt"
	"followers-ms/infrastructure"
	"followers-ms/persistence"
)

type KafkaConsumer struct {
	Repository persistence.Repository
}

type AppConsumer interface {
	StartConsumer()
	FollowTopic()
	UnfollowTopic()
}

func (k *KafkaConsumer) StartConsumer() {
	c := infrastructure.NewKafkaConsumer()

	c.SubscribeTopics([]string{"profile", "edit", "follow", "unfollow", "block", "request"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s \n", msg.TopicPartition, string(msg.Value))
			switch *msg.TopicPartition.Topic {
			case "follow":
				// decoding capn'p
				// create struct
				k.FollowTopic()
				break
			case "unfollow":
				k.UnfollowTopic()
				break

			}

		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}

func (k *KafkaConsumer) FollowTopic() {
	request := persistence.RelationshipRequest{
		FromUserId: "A",
		ToUserId:   "B",
	}
	check, err := k.Repository.CheckRelationship(request)
	if err != nil {
		return
	}
	if check {
		k.Repository.CreateRelationship(request)
	}
}

func (k *KafkaConsumer) UnfollowTopic() {
	request := persistence.RelationshipRequest{
		FromUserId: "A",
		ToUserId:   "B",
	}
	check, err := k.Repository.CheckRelationship(request)
	if err != nil {
		return
	}
	if check {
		k.Repository.CreateRelationship(request)
	}
}
