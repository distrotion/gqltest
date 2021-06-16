package links

import (
	"context"
	"fmt"
	"log"

	"github.com/distrotion/gqltest/db"
	"github.com/distrotion/gqltest/internal/users"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

var ctx = context.TODO()

//#2
func (link Link) Save() int64 {
	//#3

	//db.Getcol()

	// stmt, err := database.Getcol()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //#4
	// res, err := stmt.Exec(link.Title, link.Address)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //#5
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal("Error:", err.Error())
	// }
	// log.Print("Row inserted!")

	res, insertErr := db.Getcol().InsertOne(ctx, link)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)

	return 0
}
