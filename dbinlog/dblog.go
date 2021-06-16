package dblog

import (
	"context"
	_ "fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var ctx = context.TODO()

var server = "mongodb://localhost:9001/"
var db_mongo = "auth_log"
var colec_mongo = "income"

func Getcolin() *mongo.Collection {
	return collection
}

//var collec = "income"

func loginit() {
	clientOptions := options.Client().ApplyURI(server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(db_mongo).Collection(colec_mongo)

}
