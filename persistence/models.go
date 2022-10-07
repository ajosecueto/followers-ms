package persistence

type RelationshipRequest struct {
	FromUserId string
	ToUserId   string
	Relation   string
}
