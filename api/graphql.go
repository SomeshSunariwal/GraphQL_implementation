package api

import (
	"net/http"

	"github.com/SomeshSunariwal/GraphQL_implementation/database"
	"github.com/SomeshSunariwal/GraphQL_implementation/service"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
			Query:    GETTING(handler),
			Mutation: ADDING(handler),
		},
	)

	if err != nil {
		return nil
	}

	GraphQLHandler := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: requestString,
	})

	if GraphQLHandler.HasErrors() {
		log.Info("Error", GraphQLHandler.Errors)
		return context.JSON(http.StatusOK, GraphQLHandler.Errors)
	}

	return context.JSON(http.StatusOK, GraphQLHandler.Data)
}

// POST, PUT, PATCH,DELETE,
func ADDING(handler *Handler) *graphql.Object {
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
func GETTING(handler *Handler) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"GetItems":    handler.service.GetItems(),
			"GetItemByID": handler.service.GetItemByID(),
		},
	})
}

func DB_Health(c echo.Context) error {
	client := database.Client()
	err := client.Health()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"DB": "Down"})
	}
	return c.JSON(http.StatusOK, map[string]string{"DB": "ok"})
}
