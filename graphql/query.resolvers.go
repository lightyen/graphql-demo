package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/model"
	"context"
)

func (r *queryResolver) Device(ctx context.Context) (*model.Device, error) {
	return &model.Device{}, nil
}

func (r *queryResolver) Test(ctx context.Context, id model.UUID) (interface{}, error) {
	bs := make([]byte, 0)
	if bs[3] == 5 {
		//
	}
	return &model.Void{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
