package links

import (
	"context"
	"fmt"
	"log"

	"github.com/distrotion/gqltest/db"
	"github.com/distrotion/gqltest/internal/users"
	"go.mongodb.org/mongo-driver/bson"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

var ctx = context.TODO()

func (link Link) Save() int64 {

	res, insertErr := db.Getcol().InsertOne(ctx, link)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)

	return 0
}

func GetAll() []Link {

	//opts := options.Find()
	//opts.SetSort(bson.D{{"_id", -1}})

	res, insertErr := db.Getcol().Find(ctx, bson.D{{}})
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)

	var msg []Link
	if insertErr = res.All(ctx, &msg); insertErr != nil {
		panic(insertErr)
	}

	return msg

}
