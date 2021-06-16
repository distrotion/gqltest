package users

import (
	"context"
	"fmt"
	"strconv"

	"log"

	"github.com/distrotion/gqltest/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

var ctx = context.TODO()

func (user *User) Create() {
	statement, err := db.Getcol().InsertOne(ctx, user)
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
	// hashedPassword, err := HashPassword(user.Password)
	// _, err = statement.Exec(user.Username, hashedPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func GetUserIdByUsername(username string) (int, error) {
	res, err := db.Getcol().Find(ctx, username)
	if err != nil {
		log.Fatal(err)
	}

	var msg User
	if err = res.All(ctx, &msg); err != nil {
		panic(err)
	}

	var Id int
	Id, err = strconv.Atoi(msg.ID)

	return Id, nil
}

func (user *User) Authenticate() bool {
	// statement, err := database.Db.Prepare("select Password from Users WHERE Username = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// row := statement.QueryRow(user.Username)

	res, err := db.Getcol().Find(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	var msg []User
	if err = res.All(ctx, &msg); err != nil {
		panic(err)
	}

	fmt.Println(msg)

	//var hashedPassword string

	//hashedPassword = msg[0].Password
	var result bool
	// //fmt.Print(msg.Password)
	// return CheckPasswordHash(user.Password, hashedPassword)
	if msg[0].Password == user.Password {
		result = true
	} else {
		result = false
	}

	return result
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
