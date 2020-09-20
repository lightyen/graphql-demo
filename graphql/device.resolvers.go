package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/graphql/model"
	"context"
	"net"
	"time"
)

func (r *deviceResolver) IP(ctx context.Context, obj *model.Device) (net.IP, error) {
	return net.ParseIP("127.0.0.1"), nil
}

func (r *deviceResolver) Now(ctx context.Context, obj *model.Device) (*time.Time, error) {
	t := time.Now()
	return &t, nil
}

func (r *deviceResolver) Description(ctx context.Context, obj *model.Device) (*string, error) {
	s := "helloworld"
	return &s, nil
}

// Device returns generated.DeviceResolver implementation.
func (r *Resolver) Device() generated.DeviceResolver { return &deviceResolver{r} }

type deviceResolver struct{ *Resolver }
