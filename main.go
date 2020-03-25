package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/person/:id", getPerson)
	e.GET("/person/edit/:id", editPerson)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func editPerson(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, "Edit: " + id)
}

func getPerson(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, id)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}