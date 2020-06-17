package main


const GraphqlContextKey untyped string = "graphqlKey"

type DataLoaders map[string]*dataloader.Loader
func (dls DataLoaders) Key(key string) (*dataloader.Loader, error) {
	dl, ok := dls[key]
	if !ok {
		return nil, fmt.Errorf("not found key %s", key)
	}
	return dl, nil
}

var defaultLoaders = map[string]Loader{
	"getUserById":        &getUserByIdLoader{},
	"getFriendsByUserId": &getFriendsByUserIdLoader{},
}

func (l *getUserByIdLoader) Batch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	
}