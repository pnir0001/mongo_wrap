package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pnir0001/mongo_wrap/wrap"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// test mongo library wrap

	fmt.Println("start !!!")

	// nomal mongo client
	// fmt.Println("start mongo ----------------------")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// id := mongoCore(client)
	// fmt.Println(id)
	// fmt.Println("end mongo ----------------------")
	// fmt.Println("")

	// wrap mongo client(mongo)
	fmt.Println("start mongo wrap ---------------")
	wClient := &wrap.WrapClient{Client: client}
	wID := mongoWrapCore(wClient)
	fmt.Println(wID)
	fmt.Println("end mongo wrap -----------------")
	fmt.Println("")

	// wrap mongo client(wrap mock)
	fmt.Println("start mongo mock ---------------")
	wClient = &wrap.WrapClient{}
	wID = mongoWrapCore(wClient)
	fmt.Println(wID)
	fmt.Println("end mongo mock -----------------")
	fmt.Println("")

	fmt.Println("all end !!!")
}

func mongoCore(client *mongo.Client) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("test").Collection("test_collection")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	return id
}
func mongoWrapCore(client *wrap.WrapClient) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("test").Collection("test_collection")

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	return id
}
