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

	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.hub[ch] = struct{}{}
	fmt.Println("websocket create: len = ", len(r.hub))

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		defer r.mutex.Unlock()
		delete(r.hub, ch)
		fmt.Println("websocket close: len = ", len(r.hub))
	}()

	return ch, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
