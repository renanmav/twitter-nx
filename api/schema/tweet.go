package schema

import (
	"os"
	"path/filepath"

	"github.com/graphql-go/graphql"
)

type Tweet struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}

var TweetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tweet",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"likes": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var TweetList []Tweet

var pwd, _ = os.Getwd()
var _ = GetJSONDataFromFile(filepath.Join(pwd, "schema", "tweets.json"), &TweetList)
