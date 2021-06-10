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
		Args: graphql.FieldConfigArgument{
			"bookName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"author": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"available": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"seller": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			userPostRequest, err := validatePostRequest(p.Args)
			if err != nil {
				return nil, err
			}
			// Database Function Call Here
			client := database.Client()
			result, err := client.UpdateItem(userPostRequest)
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
		Args: graphql.FieldConfigArgument{
			"bookName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			bookName := p.Args["bookName"].(string)
			// Database Function Call Here
			client := database.Client()
			result, err := client.DeleteItem(bookName)
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
		Args: graphql.FieldConfigArgument{
			"bookName": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

			bookName := p.Args["bookName"].(string)
			// Database Function Call Here
			client := database.Client()
			result, err := client.GetItemByID(bookName)
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

	bookName := Args["bookName"]
	if bookName != nil {
		bookNewName := bookName.(string)
		validRequest.BookName = &bookNewName
	}

	author := Args["author"]
	if author != nil {
		authorNew := author.(string)
		validRequest.Author = &authorNew
	}

	seller := Args["seller"]
	if seller != nil {
		sellerNew := seller.(string)
		validRequest.Seller = &sellerNew
	}

	available := Args["available"]
	if available != nil {
		availableNew := available.(bool)
		validRequest.Available = &availableNew
	}

	location := Args["location"]
	if location != nil {
		locationNew := location.(string)
		validRequest.Location = &locationNew
	}

	return validRequest, nil
}
