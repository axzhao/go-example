package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/graph-gophers/dataloader"

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

type options struct {
	Logger int
	DB     int
}
type RootResolver struct {
	Global options
}

type UserArgs struct {
	ID graphql.ID
}

func (r *RootResolver) User(ctx context.Context, args *UserArgs) (*UserResolver, error) {
	var user *User
	for _, u := range mockUsers {
		if string(args.ID) == u.id {
			user = u
		}
	}
	return &UserResolver{u: user}, nil
}

var CallTimes1, CallTimes2 int

// mock
var mockUsers = []*User{
	&User{id: "1", name: "111"},
	&User{id: "2", name: "222"},
	&User{id: "3", name: "333"},
	&User{id: "4", name: "444"},
	&User{id: "5", name: "555"},
	&User{id: "6", name: "555"},
	&User{id: "7", name: "555"},
	&User{id: "8", name: "555"},
	&User{id: "9", name: "555"},
	&User{id: "10", name: "555"},
	&User{id: "11", name: "555"},
	&User{id: "12", name: "555"},
	&User{id: "13", name: "555"},
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
			dls := DataLoaders{}
			cache := NewCache()
			for _, v := range defaultLoaders {
				loader := dataloader.NewBatchedLoader(v.Batch, dataloader.WithCache(cache))
				dls[v.Key()] = loader
			}
			c := &Context{Values: make(map[string]interface{}), Loaders: dls}
			ctx := r.Context()
			ctx = context.WithValue(ctx, GraphqlContextKey, c)
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
