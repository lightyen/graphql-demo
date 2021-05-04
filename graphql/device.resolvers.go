package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/common"
	"app/graphql/generated"
	"context"
	"fmt"
	"net"
)

func (r *deviceResolver) IP(ctx context.Context, obj *common.Device) (*common.IP, error) {
	return &common.IP{IP: net.ParseIP("127.0.0.1")}, nil
}

func (r *deviceResolver) Now(ctx context.Context, obj *common.Device) (*common.Time, error) {
	t := common.Now()
	return &t, nil
}

func (r *deviceResolver) Description(ctx context.Context, obj *common.Device) (*string, error) {
	s := "helloworld"
	return &s, nil
}

func (r *deviceResolver) Count(ctx context.Context, obj *common.Device, param string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// Device returns generated.DeviceResolver implementation.
func (r *Resolver) Device() generated.DeviceResolver { return &deviceResolver{r} }

type deviceResolver struct{ *Resolver }
