package modal

import (
	"github.com/graphql-go/graphql"
)

var Book = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Books",
		Fields: graphql.Fields{
			"id":           &graphql.Field{Type: graphql.ID},
			"bookName":     &graphql.Field{Type: graphql.String},
			"details":      &graphql.Field{Type: Details},
			"availability": &graphql.Field{Type: graphql.NewList(Availability)},
			"available":    &graphql.Field{Type: graphql.Boolean},
		},
		Description: "Book",
	},
)

var Details = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Details",
		Fields: graphql.Fields{
			"author": &graphql.Field{Type: graphql.String},
			"seller": &graphql.Field{Type: graphql.String},
		},
		Description: "Details",
	},
)

var Availability = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Availability",
		Fields: graphql.Fields{
			"location": &graphql.Field{Type: graphql.String},
			"id":       &graphql.Field{Type: graphql.Int},
		},
		Description: "Users data",
	},
)
