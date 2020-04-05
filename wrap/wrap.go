package wrap

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WrapClient struct {
	Client *mongo.Client
}
type WrapDatabase struct {
	Database *mongo.Database
}
type WrapCollection struct {
	Collection *mongo.Collection
}
type WrapInsertOneResult struct {
	InsertOneResult *mongo.InsertOneResult
	InsertedID      interface{}
}

func (w *WrapClient) Connect(ctx context.Context) error {
	fmt.Println("Wrap Connect")
	if w.Client != nil {
		fmt.Println("but mongo Connect")
		return w.Client.Connect(ctx)
	}
	return nil
}

func (w *WrapClient) Database(name string, opts ...*options.DatabaseOptions) *WrapDatabase {
	fmt.Println("Wrap Databse")
	if w.Client != nil {
		fmt.Println("but mongo Databse")
		return &WrapDatabase{Database: w.Client.Database(name, opts...)}
	}
	return &WrapDatabase{}
}

func (w *WrapDatabase) Collection(name string, opts ...*options.CollectionOptions) *WrapCollection {
	fmt.Println("Wrap Collection")
	if w.Database != nil {
		fmt.Println("but mongo Collection")
		return &WrapCollection{Collection: w.Database.Collection(name, opts...)}
	}
	return &WrapCollection{}
}

func (coll *WrapCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*WrapInsertOneResult, error) {
	fmt.Println("Wrap InsertOne")
	if coll.Collection != nil {
		fmt.Println("but mongo InsertOne")
		r, err := coll.Collection.InsertOne(ctx, document, opts...)
		return &WrapInsertOneResult{InsertedID: r.InsertedID}, err
	}
	return &WrapInsertOneResult{InsertedID: "WrapObjectID..."}, nil
}
