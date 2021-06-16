package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func Getcol() *mongo.Collection {
	return collection
}

var ctx = context.TODO()

var server = "mongodb://localhost:9000/"
var db_mongo = "aut"
var collec = "users"

func init() {
	clientOptions := options.Client().ApplyURI(server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(db_mongo).Collection(collec)
	fmt.Println("database ready to use ...")
}
