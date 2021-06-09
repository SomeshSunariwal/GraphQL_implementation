package modal

import (
	"github.com/graphql-go/graphql"
)

var Book = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Books",
		Fields: graphql.Fields{
			"id":           &graphql.Field{Type: graphql.ID},
			"book":         &graphql.Field{Type: graphql.String},
			"author":       &graphql.Field{Type: Author},
			"availability": &graphql.Field{Type: graphql.NewList(Availability)},
			"available":    &graphql.Field{Type: graphql.Boolean},
		},
		Description: "Users data",
	},
)

var Author = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"author": &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)

var Availability = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Availability",
		Fields: graphql.Fields{
			"author": &graphql.Field{Type: graphql.String},
		},
		Description: "Users data",
	},
)
