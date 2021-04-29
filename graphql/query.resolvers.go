package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/common"
	"app/graphql/generated"
	"context"
)

func (r *queryResolver) Device(ctx context.Context) (*common.Device, error) {
	return &common.Device{}, nil
}

func (r *queryResolver) Test(ctx context.Context, id common.UUID) (interface{}, error) {
	return &common.Void{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
