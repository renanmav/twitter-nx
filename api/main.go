package main

import (
	"log"
	"net/http"
	"twitter-nx/api/schema"

	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	log.Printf("Listening requests on http://localhost:9000/graphql")
	http.Handle("/graphql", h)
	http.ListenAndServe(":9000", nil)
}
