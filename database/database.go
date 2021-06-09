package database

import (
	"github.com/SomeshSunariwal/GraphQL_implementation/modal"
	"github.com/graphql-go/graphql"
)

type Database struct {
}

func (database *Database) AddItem() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(modal.User),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			// Database Function Call Here
			result, err := AddItem()
			if err != nil {
				return nil, nil
			}

			return result, nil
		},
		Description: "user",
	}
}

func AddItem() (string, error) {
	return "", nil
}
