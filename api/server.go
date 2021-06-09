package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {

	handler := Handler{}
	e.POST("/graphql", handler.NewGraphQLHandler)
	e.Use(middleware.Logger())
}
