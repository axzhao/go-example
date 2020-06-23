package main

import (
	"context"
	"errors"
	"fmt"
)

// Friend
type Friend struct {
	userId   string
	friendId string
}
type FriendResolver struct {
	f    *Friend
	root *RootResolver
}

func (r *FriendResolver) FriendId() *string { return &r.f.friendId }
func (r *FriendResolver) User(ctx context.Context) (*UserResolver, error) {
	c, found := ctx.Value(GraphqlContextKey).(*Context)
	if !found {
		return nil, errors.New("unable to find the custom context")
	}
	loader, err := c.Loaders.Key("getUserById")
	if err != nil {
		return nil, err
	}
	thunk := loader.Load(ctx, GetUserByIdKey{Id: r.f.friendId})
	data, err := thunk()
	if err != nil {
		return nil, fmt.Errorf("getUserById: %v", err)
	}
	user, ok := data.(*User)
	if !ok {
		return nil, fmt.Errorf("User: loaded the wrong type of data: %#v", data)
	}
	return &UserResolver{user}, nil
}

func getFriendsByUserIds(ids []string) []*Friend {
	CallTimes2++
	var friends []*Friend
	for _, f := range mockUserFriends {
		for _, id := range ids {
			if f.userId == id {
				friends = append(friends, f)
				break
			}
		}
	}
	return friends
}
