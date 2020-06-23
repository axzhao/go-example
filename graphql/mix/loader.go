package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/graph-gophers/dataloader"
)

// context
const GraphqlContextKey string = "graphqlKey"

type Context struct {
	Values  map[string]interface{}
	Loaders DataLoaders
	lock    sync.RWMutex
}

func (c *Context) WithValue(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.Values[key] = value
}

func (c *Context) Value(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.Values[key]
}

// dataloader
type DataLoaders map[string]*dataloader.Loader

func (dls DataLoaders) Key(key string) (*dataloader.Loader, error) {
	dl, ok := dls[key]
	if !ok {
		return nil, fmt.Errorf("not found key %s", key)
	}
	return dl, nil
}

// loader
type Loader interface {
	Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result
	Key() string
}

var defaultLoaders = map[string]Loader{
	"getUserById":        &getUserByIdLoader{},
	"getFriendsByUserId": &getFriendsByUserIdLoader{},
}

// getUserById
type GetUserByIdKey struct{ Id string }

func (k GetUserByIdKey) String() string   { return k.Id }
func (k GetUserByIdKey) Raw() interface{} { return k.Id }

type getUserByIdLoader struct{}

func (l *getUserByIdLoader) Key() string { return "getUserById" }
func (l *getUserByIdLoader) Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var ids []string
	for _, key := range keys {
		k, ok := key.(GetUserByIdKey)
		if ok {
			ids = append(ids, k.Id)
		}
	}
	users := getUserByIds(ids)
	var results []*dataloader.Result
	for _, u := range users {
		results = append(results, &dataloader.Result{Data: u})
	}
	return results
}

// getFriendsByUserId
type GetFriendsByUserIdKey struct{ UserId string }

func (k GetFriendsByUserIdKey) String() string   { return k.UserId }
func (k GetFriendsByUserIdKey) Raw() interface{} { return k.UserId }

type getFriendsByUserIdLoader struct{}

func (l *getFriendsByUserIdLoader) Key() string { return "getFriendsByUserId" }
func (l *getFriendsByUserIdLoader) Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var ids []string
	for _, key := range keys {
		k, ok := key.(GetFriendsByUserIdKey)
		if ok {
			ids = append(ids, k.UserId)
		}
	}
	friends := getFriendsByUserIds(ids)
	sortFriends := map[string][]*Friend{}
	for _, f := range friends {
		if v, ok := sortFriends[f.userId]; ok {
			sortFriends[f.userId] = append(v, f)
		} else {
			sortFriends[f.userId] = []*Friend{f}
		}
	}
	var results []*dataloader.Result
	for _, id := range ids {
		if v, ok := sortFriends[id]; ok {
			results = append(results, &dataloader.Result{Data: v})
		} else {
			results = append(results, &dataloader.Result{Data: []*Friend{}})
		}
	}
	return results
}
