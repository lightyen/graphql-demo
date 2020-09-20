package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/graphql/model"
	"context"
)

func (r *queryResolver) Device(ctx context.Context) (*model.Device, error) {
	return &model.Device{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
