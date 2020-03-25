package main

import (
	_ "echo-test/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
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
	e.GET("/persons", getAllPerson)
	e.GET("/person/:id", getPerson)
	e.GET("/person/edit/:id", editPerson)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// @Summary Get All person
// @Success 200 {array} Person
// @Router /persons [get]
func getAllPerson(c echo.Context) error {
	return c.JSON(http.StatusOK, persons)
}

func editPerson(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, "Edit: "+id)
}

// @Summary Get person
// @Description get person
// @Param id path int true "person search by id"
// @Success 200 {string} string
// @Router /person/{id} [get]
func getPerson(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, id)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

var persons = []Person{
	{
		Id:   0,
		Name: "Apple",
		Desc: "Apple is sweet",
	},
	{
		Id:   1,
		Name: "Flower",
		Desc: "Flower is colorful",
	},
}
