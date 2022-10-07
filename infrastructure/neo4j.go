package infrastructure

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"os"
)

func NewNeo4jSession() neo4j.Session {
	host := os.Getenv("NEO4J_HOST")
	user := os.Getenv("NEO4J_USER")
	password := os.Getenv("NEO4J_PASSWORD")
	driver, err := neo4j.NewDriver(host, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		panic("CANT CONNECT TO NEO4J")
	}
	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	return session
}
