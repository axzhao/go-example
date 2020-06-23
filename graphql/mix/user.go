package main

import (
	"context"
	"errors"
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
)

// User
type User struct {
	id   string
	name string
}
type UserResolver struct {
	u    *User
	root *RootResolver
}

func (r *UserResolver) ID() *graphql.ID {
	id := graphql.ID(r.u.id)
	return &id
}
func (r *UserResolver) Name() *string { return &r.u.name }
func (r *UserResolver) Friends(ctx context.Context) ([]*FriendResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}
	loader, err := c.Loaders.Key("getFriendsByUserId")
	if err != nil {
		return nil, err
	}
	thunk := loader.Load(ctx, GetFriendsByUserIdKey{UserId: r.u.id})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getFriendsByUserId: %v", err)
	}
	friends, ok := data.([]*Friend)
	if !ok {
		return nil, fmt.Errorf("Friends: loaded the wrong type of data: %#v", data)
	}
	var resolvers []*FriendResolver
	for _, friend := range friends {
		resolvers = append(resolvers, &FriendResolver{friend})
	}
	return resolvers, nil
}

func getUserByIds(ids []string) []*User {
	CallTimes1++
	var users []*User
	for _, id := range ids {
		for _, u := range mockUsers {
			if u.id == id {
				users = append(users, u)
				break
			}
		}
	}
	return users
}
