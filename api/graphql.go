package api

import (
	"net/http"

	"github.com/SomeshSunariwal/GraphQL_implementation/service"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service service.Service
}

func (handler *Handler) NewGraphQLHandler(context echo.Context) error {

	var requestBody map[string]interface{}
	context.Bind(&requestBody)

	// Request String
	requestString := requestBody["query"].(string)

	// Input Schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    GetItem(handler),
			Mutation: AddItem(handler),
		},
	)

	if err != nil {
		return nil
	}

	GraphQLHandler := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestString,
	})

	return context.JSON(http.StatusOK, GraphQLHandler.Data)
}

// POST, PUT, PATCH,DELETE,
func AddItem(handler *Handler) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"AddItem":    handler.service.AddItem(),
			"UpdateItem": handler.service.UpdateItem(),
			"DeleteItem": handler.service.DeleteItem(),
		},
	})
}

// GET
func GetItem(handler *Handler) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetItems":    handler.service.GetItems(),
			"GetItemByID": handler.service.GetItemByID(),
		},
	})
}
