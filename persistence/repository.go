package persistence

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Repository struct {
	Driver neo4j.Session
}

type RepositoryInterface interface {
	CheckRelationship(relation RelationshipRequest) (bool, error)
	CreateRelationship(relation RelationshipRequest) (bool, error)
}

func (r Repository) CheckRelationship(relation RelationshipRequest) (bool, error) {
	fmt.Printf("Repository CheckRelationship %v, %v \n", relation.FromUserId, relation.ToUserId)
	return true, nil
}

func (r Repository) CreateRelationship(relation RelationshipRequest) (bool, error) {
	// FromUserId -Relation-> ToUserId
	fmt.Printf("Repository CreateRelationship %v, %v \n", relation.FromUserId, relation.ToUserId)
	return true, nil
}
