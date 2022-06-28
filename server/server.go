package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/huyhuy1020/contactFriend"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
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
