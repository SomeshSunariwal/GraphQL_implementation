package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {

	handler := Handler{}
	e.POST("/graphql", handler.NewGraphQLHandler)
	e.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"Response": "OK"})
	})
	e.GET("/health", DB_Health)

	e.Use(middleware.Logger())
}
