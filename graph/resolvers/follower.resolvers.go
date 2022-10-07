package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"followers-ms/graph/customTypes"
	"followers-ms/graph/generated"
	"followers-ms/persistence"
	"github.com/99designs/gqlgen/graphql"
)

// Follow is the resolver for the follow field.
func (r *mutationResolver) Follow(ctx context.Context, userID string) (bool, error) {
	request := persistence.RelationshipRequest{FromUserId: "UserA", ToUserId: userID, Relation: "FOLLOW"}
	check, err := r.Repository.CheckRelationship(request)

	if err != nil {
		graphql.AddErrorf(ctx, "Cant check follow %v", err)
	}

	if !check {
		graphql.AddErrorf(ctx, "User Can't Follow this user")
	}

	r.Producer.FollowProducer(request, "follow")

	return true, nil
}

// Unfollow is the resolver for the unfollow field.
func (r *mutationResolver) Unfollow(ctx context.Context, userID string) (bool, error) {
	return true, nil
}

// GetFollowers is the resolver for the getFollowers field.
func (r *queryResolver) GetFollowers(ctx context.Context, userID string) ([]*customTypes.Profile, error) {
	panic(fmt.Errorf("not implemented: GetFollowers - getFollowers"))
}

// GetFollowees is the resolver for the getFollowees field.
func (r *queryResolver) GetFollowees(ctx context.Context, userID string) ([]*customTypes.Profile, error) {
	panic(fmt.Errorf("not implemented: GetFollowees - getFollowees"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
