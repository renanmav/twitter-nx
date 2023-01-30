package schema

import "github.com/graphql-go/graphql"

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type:        graphql.String,
			Description: "My first resolver field implemented in golang",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"tweets": &graphql.Field{
			Type:        graphql.NewList(TweetType),
			Description: "List of tweets",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return TweetList, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
