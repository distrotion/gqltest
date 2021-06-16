package datalog

import (
	"context"
	"log"
	_ "strconv"

	dblog "github.com/distrotion/gqltest/dbinlog"
)

var ctx = context.TODO()

type data struct {
	Data string `json:"data"`
}

func (input *data) fromin() {
	statement, err := dblog.Getcolin().InsertOne(ctx, input)
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	// hashedPassword, err := HashPassword(user.Password)
	// _, err = statement.Exec(user.Data, hashedPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
