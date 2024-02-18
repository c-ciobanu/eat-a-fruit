package main

import (
	"github.com/c-ciobanu/eat-a-fruit/models"
	"github.com/c-ciobanu/eat-a-fruit/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var fruits = []models.Fruit{
	{Id: "apple", Name: "Apple", Months: []string{"October", "November", "December", "January", "February"}},
	{Id: "blackberry", Name: "Blackberry", Months: []string{"July", "August", "September", "October"}},
	{Id: "blackcurrant", Name: "Blackcurrant", Months: []string{"June", "July", "August"}},
	{Id: "blueberry", Name: "Blueberry", Months: []string{"July"}},
	{Id: "cherry", Name: "Cherry", Months: []string{"June", "July", "August"}},
	{Id: "cranberry", Name: "Cranberry", Months: []string{"November", "December"}},
	{Id: "damson", Name: "Damson", Months: []string{"August", "September"}},
	{Id: "elderberry", Name: "Elderberry", Months: []string{"October", "November"}},
	{Id: "gooseberry", Name: "Gooseberry", Months: []string{"June", "July"}},
	{Id: "greengage", Name: "Greengage", Months: []string{"July", "August"}},
	{Id: "loganberry", Name: "Loganberry", Months: []string{"July", "August"}},
	{Id: "pear", Name: "Pear", Months: []string{"September", "October", "November", "December", "January", "February"}},
	{Id: "plum", Name: "Plum", Months: []string{"August", "September"}},
	{Id: "raspberry", Name: "Raspberry", Months: []string{"June", "July", "August", "September"}},
	{Id: "redcurrant", Name: "Redcurrant", Months: []string{"June", "July", "August"}},
	{Id: "strawberry", Name: "Strawberry", Months: []string{"May", "June", "July", "August", "September"}},
	{Id: "tayberry", Name: "Tayberry", Months: []string{"June"}},
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
