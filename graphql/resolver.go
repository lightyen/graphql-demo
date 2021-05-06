package graphql

import (
	"app/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	mutex sync.RWMutex
	hub   map[chan *model.Time]struct{}
}

func (r *Resolver) NotifyTime(t *model.Time) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	for ch := range r.hub {
		ch <- t
	}
}
