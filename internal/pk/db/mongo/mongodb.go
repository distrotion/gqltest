package mongodb

import (
	"context"
	"log"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var sever = "mongodb://127.0.0.1:27017/"

var db_name = "aut"
var col_name = "users"

var collections *mongo.Collection
var ctx = context.TODO()

func init() {

	clientOptions := options.Client().ApplyURI(sever)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collections = client.Database(db_name).Collection(col_name)
}
