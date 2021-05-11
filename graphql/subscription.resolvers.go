package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graphql/generated"
	"app/model"
	"context"
	"fmt"
)

func (r *subscriptionResolver) NotificationTime(ctx context.Context) (<-chan *model.Time, error) {
	ch := make(chan *model.Time, 1)
	r.AddPeer(ch, &PeerInfo{ctx})
	fmt.Println("websocket create: len = ", r.PeerCount())

	go func() {
		<-ctx.Done()
		r.RemovePeer(ch)
		fmt.Println("websocket close: len = ", r.PeerCount())
	}()

	return ch, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
