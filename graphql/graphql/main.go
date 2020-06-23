package main

import (
	"bytes"
	"context"
	"log"
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// schema
const schemaString = `
		schema {
			query: Query
		}
		type Query {
			user(id: ID!): User
		}
		type User {
			id: ID
			name: String
			friends(): [Friend]!
		}
		type Friend {
			friendId: String
			user(): User
		}
	`

// resolver
type RootResolver struct{}

type UserArgs struct {
	ID graphql.ID
}

func (r *RootResolver) User(ctx context.Context, args *UserArgs) (*UserResolver, error) {
	var user *User
	for _, u := range mockUser {
		if string(args.ID) == u.id {
			user = u
		}
	}
	return &UserResolver{u: user}, nil
}

// xxx

var CallTimes1, CallTimes2 int


// User
type User struct {
	id   string
	name string
}
type UserResolver struct {
	u *User
}

func (r *UserResolver) ID() *graphql.ID {
	id := graphql.ID(r.u.id)
	return &id
}
func (r *UserResolver) Name() *string { return &r.u.name }
func (r *UserResolver) Friends(ctx context.Context) ([]*FriendResolver, error) {
	CallTimes1++
	var friends []*Friend
	for _, f := range mockUserFriends {
		if f.userId == r.u.id {
			friends = append(friends, f)
		}
	}
	var resolvers []*FriendResolver
	for _, friend := range friends {
		resolvers = append(resolvers, &FriendResolver{f: friend})
	}
	return resolvers, nil
}

// Friend
type Friend struct {
	userId   string
	friendId string
}
type FriendResolver struct {
	f *Friend
}

func (r *FriendResolver) FriendId() *string { return &r.f.friendId }
func (r *FriendResolver) User(ctx context.Context) (*UserResolver, error) {
	CallTimes2++
	var user *User
	for _, u := range mockUser {
		if r.f.friendId == u.id {
			user = u
		}
	}
	return &UserResolver{u: user}, nil
}

// mock
var mockUser = []*User{
	&User{id: "1", name: "111"},
	&User{id: "2", name: "222"},
	&User{id: "3", name: "333"},
	&User{id: "4", name: "444"},
	&User{id: "5", name: "555"},
}
var mockUserFriends = []*Friend{
	&Friend{userId: "1", friendId: "2"},
	&Friend{userId: "1", friendId: "3"},
	&Friend{userId: "2", friendId: "1"},
	&Friend{userId: "2", friendId: "4"},
	&Friend{userId: "3", friendId: "1"},
	&Friend{userId: "4", friendId: "1"},
	&Friend{userId: "4", friendId: "5"},
	&Friend{userId: "5", friendId: "4"},
}

// main
func main() {
	withContext := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			h.ServeHTTP(w, r.WithContext(ctx))
			fmt.Println("1: call user get friends ", CallTimes1)
			fmt.Println("2: call friend get user", CallTimes2)
		})
	}

	schema := graphql.MustParseSchema(schemaString, &RootResolver{})
	http.Handle("/graphql", withContext(&relay.Handler{Schema: schema}))

	// debug
	debugPage := bytes.Replace(GraphiQLPage, []byte("fetch('/'"), []byte("fetch('/graphql'"), -1)
	http.HandleFunc("/debug.html", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(debugPage)
	})
	// http
	log.Println("run graphql server, :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}



// query {
// 	user(id: 1) {
// 	  id,
// 	  name,
// 	  friends {
// 		friendId
// 		user {
// 		  id,
// 		  name
// 		}
// 	  }
// 	}
//   }
  
// 2020/06/18 11:39:15 run graphql server, :9000
// 1: call user get friends  1
// 2: call friend get user 2
// 1: call user get friends  2
// 2: call friend get user 4