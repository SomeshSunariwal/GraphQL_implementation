package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {

	handler := Handler{}
	group := e.Group("/graph") // For Kubernetes

	group.POST("/graphql", handler.NewGraphQLHandler)
	group.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"Response": "OK"})
	})
	group.GET("/health", DB_Health)

	e.Use(middleware.Logger())
}
