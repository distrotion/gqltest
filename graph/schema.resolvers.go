package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	dblog "github.com/distrotion/gqltest/dbinlog"
	"github.com/distrotion/gqltest/graph/generated"
	"github.com/distrotion/gqltest/graph/model"
	"github.com/distrotion/gqltest/internal/links"
	"github.com/distrotion/gqltest/internal/users"
	"github.com/distrotion/gqltest/pkg/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

// func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
// 	var link model.Link
// 	var user model.User
// 	link.Address = input.Address
// 	link.Title = input.Title
// 	user.Name = "test"
// 	link.User = &user
// 	return &link, nil
// }

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {

	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	linkID := link.Save()
	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {

	inlog, err := dblog.Getcolin().InsertOne(ctx, bson.M{"From": "CreateUser", "data_newuser_in": input})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inlog)

	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	outlog, err := dblog.Getcolin().InsertOne(ctx, bson.M{"From": "CreateUser", "data_newuser_out": token})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outlog)
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {

	//tokenin, err := jwt.GenerateToken(input.Username)

	//var inlog interface{}
	//r.ShouldBind(&inlog)

	inlog, err := dblog.Getcolin().InsertOne(ctx, bson.M{"From": "login", "data_login_in": input})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inlog)

	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	//correct := true
	if !correct {
		// 1
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	outlog, err := dblog.Getcolin().InsertOne(ctx, bson.M{"From": "login", "data_login_out": token})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outlog)

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {

	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {

	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address})
	}

	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
