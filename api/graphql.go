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
			Query: handler.service.AddItem(),
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
