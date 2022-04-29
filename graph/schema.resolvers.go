package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/haulerkonj/gqlgen_todos/graph/generated"
	"github.com/haulerkonj/gqlgen_todos/graph/model"
)

func (r *mutationResolver) Fetch(ctx context.Context, input string) (*model.Record, error) {
	record := &model.Record{
		Num:           input,
		Names:         input,
		SumOfAllForks: input,
		// User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.records = append(r.records, record)
	return record, nil
}

func (r *queryResolver) Record(ctx context.Context) (*model.Record, error) {
	return r.records[0], nil
}

func (r *queryResolver) Records(ctx context.Context) ([]*model.Record, error) {
	return r.records, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
