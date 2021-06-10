package service

import (
	"github.com/SomeshSunariwal/GraphQL_implementation/database"
	"github.com/SomeshSunariwal/GraphQL_implementation/modal"
	"github.com/graphql-go/graphql"
)

type Service struct {
}

func (service *Service) AddItem() *graphql.Field {

	return &graphql.Field{
		Type: modal.Book,
		Args: graphql.FieldConfigArgument{
			"bookName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"author": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"available": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"location": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"seller": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			userPostRequest, err := validatePostRequest(p.Args)
			if err != nil {
				return nil, err
			}
			// Database Function Call Here
			client := database.Client()
			result, err := client.AddItem(userPostRequest)
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
			client := database.Client()
			result, err := client.UpdateItem()
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
			client := database.Client()
			result, err := client.DeleteItem()
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
			client := database.Client()
			result, err := client.GetItems()
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
			client := database.Client()
			result, err := client.GetItemByID()
			if err != nil {
				return nil, err
			}

			return result, nil
		},
		Description: "GetItemByID",
	}
}

func validatePostRequest(Args map[string]interface{}) (*modal.PostBook, error) {
	validRequest := &modal.PostBook{}

	bookName := Args["bookName"].(string)
	validRequest.BookName = &bookName

	author := Args["author"].(string)
	validRequest.Author = &author

	seller := Args["seller"].(string)
	validRequest.Seller = &seller

	available := Args["available"].(bool)
	validRequest.Available = &available

	location := Args["location"].(string)
	validRequest.Location = &location

	return validRequest, nil
}
