package service

import (
	"github.com/SomeshSunariwal/GraphQL_implementation/database"
	"github.com/SomeshSunariwal/GraphQL_implementation/modal"
	"github.com/graphql-go/graphql"
)

type Service struct {
	database database.Database
}

func (service *Service) AddItem() *graphql.Field {
	return &graphql.Field{
		Type: modal.Book,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			// Database Function Call Here
			result, err := service.database.AddItem()
			if err != nil {
				return nil, err
			}
			return result, nil
		},
		Description: "AddItem",
	}
}

func (service *Service) UpdateItem() *graphql.Field {
	return &graphql.Field{
		Type: modal.Book,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			// Database Function Call Here
			result, err := service.database.UpdateItem()
			if err != nil {
				return nil, err
			}
			return result, nil
		},
		Description: "UpdateItem",
	}
}

func (service *Service) DeleteItem() *graphql.Field {
	return &graphql.Field{
		Type: modal.Book,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			// Database Function Call Here
			result, err := service.database.DeleteItem()
			if err != nil {
				return nil, err
			}
			return result, nil
		},
		Description: "DeleteItem",
	}
}

func (service *Service) GetItems() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(modal.Book),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			// Database Function Call Here
			result, err := service.database.GetItems()
			if err != nil {
				return nil, err
			}
			return result, nil
		},
		Description: "GetItems",
	}
}

func (service *Service) GetItemByID() *graphql.Field {
	return &graphql.Field{
		Type: modal.Book,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			// Database Function Call Here
			result, err := service.database.GetItemByID()
			if err != nil {
				return nil, err
			}

			return result, nil
		},
		Description: "GetItemByID",
	}
}
