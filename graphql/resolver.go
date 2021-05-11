package graphql

import (
	"app/model"
	"context"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	peers sync.Map // map[chan]*PeerInfo
}

type PeerInfo struct {
	ctx context.Context
}

func (r *Resolver) AddPeer(ch chan *model.Time, payload *PeerInfo) {
	r.peers.Store(ch, payload)
}

func (r *Resolver) RemovePeer(ch chan *model.Time) {
	r.peers.Delete(ch)
}

func (r *Resolver) ForEachPeer(cb func(ch chan *model.Time, value *PeerInfo)) {
	r.peers.Range(func(key, value interface{}) bool {
		ch := key.(chan *model.Time)
		v := value.(*PeerInfo)
		cb(ch, v)
		return true
	})
}

func (r *Resolver) PeerCount() int {
	count := 0
	r.ForEachPeer(func(ch chan *model.Time, value *PeerInfo) {
		count++
	})
	return count
}

func (r *Resolver) BroadcastTime(t *model.Time) {
	r.ForEachPeer(func(ch chan *model.Time, value *PeerInfo) {
		ch <- t
	})
}
