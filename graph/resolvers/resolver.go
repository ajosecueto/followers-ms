package graph

import (
	"followers-ms/events"
	"followers-ms/persistence"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Producer   events.KafkaProducer
	Repository persistence.Repository
}
