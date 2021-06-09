package main

import (
	"net/http"
	"os"

	"github.com/SomeshSunariwal/GraphQL_implementation/api"
	"github.com/SomeshSunariwal/GraphQL_implementation/config"
	"github.com/labstack/echo/v4"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.PORT
	}
	e := echo.New()

	// To write APIs in seprate file
	api.Init(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Start(":" + PORT)
}
