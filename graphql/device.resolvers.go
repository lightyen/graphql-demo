package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/model"
	"context"
	"fmt"
	"net"
)

func (r *deviceResolver) IP(ctx context.Context, obj *model.Device) (*model.IP, error) {
	return &model.IP{IP: net.ParseIP("127.0.0.1")}, nil
}

func (r *deviceResolver) Now(ctx context.Context, obj *model.Device) (*model.Time, error) {
	t := model.Now()
	return &t, nil
}

func (r *deviceResolver) Count(ctx context.Context, obj *model.Device, param string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// Device returns generated.DeviceResolver implementation.
func (r *Resolver) Device() generated.DeviceResolver { return &deviceResolver{r} }

type deviceResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *deviceResolver) Description(ctx context.Context, obj *model.Device) (*string, error) {
	s := "helloworld"
	return &s, nil
}
