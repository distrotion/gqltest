package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/distrotion/gqltest/db"
	dblog "github.com/distrotion/gqltest/dbinlog"
	"github.com/distrotion/gqltest/graph"
	"github.com/distrotion/gqltest/graph/generated"
	"github.com/distrotion/gqltest/internal/auth"
	"github.com/go-chi/chi"
	_ "github.com/urfave/cli/v2"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	db.Getcol()
	dblog.Getcolin()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
