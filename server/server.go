package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v10"
	"github.com/huyhuy1020/contactFriend"
	"github.com/huyhuy1020/contactFriend/postgres"
	
)

const defaultPort = "8080"

func main() {
	// we put some example
	DB := postgres.NewReceive(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "0942877351",
		Database: "contact_friend",
	})

	defer DB.Close()

	//after we close database but we need a quick look so we use addqueryhook
	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := contactFriend.Config{Resolvers: &contactFriend.Resolver{}}
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/", handler.GraphQL(contactFriend.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
