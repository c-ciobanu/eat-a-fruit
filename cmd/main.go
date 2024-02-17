package main

import (
	"github.com/c-ciobanu/eat-a-fruit/models"
	"github.com/c-ciobanu/eat-a-fruit/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var fruits = []models.Fruit{
	{Id: "apple", Name: "Apple"},
	{Id: "banana", Name: "Banana"},
	{Id: "pear", Name: "Pear"},
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", indexHandler)

	e.Logger.Fatal(e.Start(":4000"))
}

func indexHandler(c echo.Context) error {
	return views.Index(fruits).Render(c.Request().Context(), c.Response())
}
