package service

import (
	"github.com/SomeshSunariwal/GraphQL_implementation/database"
	"github.com/graphql-go/graphql"
)

type Service struct {
	database database.Database
}

func (service *Service) AddItem() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Users": service.database.AddItem(),
		},
	})
}
