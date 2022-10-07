package main

import (
	"followers-ms/events"
	"followers-ms/graph/generated"
	graph "followers-ms/graph/resolvers"
	"followers-ms/infrastructure"
	"followers-ms/persistence"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	neo4jSession := infrastructure.NewNeo4jSession()

	log.Println(neo4jSession)

	repository := persistence.Repository{
		Driver: neo4jSession,
	}
	consumer := events.KafkaConsumer{
		Repository: repository,
	}

	go consumer.StartConsumer()

	producer := infrastructure.NewKafkaProducer()

	resolver := graph.Resolver{
		Repository: repository,
		Producer: events.KafkaProducer{
			Client: producer,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
