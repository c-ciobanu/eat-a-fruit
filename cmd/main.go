package main

import (
	"slices"
	"time"

	"github.com/c-ciobanu/eat-a-fruit/models"
	"github.com/c-ciobanu/eat-a-fruit/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// https://www.bda.uk.com/food-health/your-health/sustainable-diets/seasonal-fruit-and-veg-a-handy-guide.html
// https://www.oddbox.co.uk/blog/whats-in-season-when-a-guide-to-uk-seasonal-eating
// https://www.bbcgoodfood.com/seasonal-calendar
var ukSeasonalFruits = []models.Fruit{
	{Id: "apple", Name: "Apple", Months: []string{"September", "October", "November", "December", "January", "February"}},
	{Id: "apricot", Name: "Apricot", Months: []string{"June", "July", "August"}},
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
	{Id: "quince", Name: "Quince", Months: []string{"October", "November", "December"}},
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
	month := time.Now().Month()

	var fruitsInSeasonNow []models.Fruit
	for _, fruit := range ukSeasonalFruits {
		if slices.Contains(fruit.Months, month.String()) {
			fruitsInSeasonNow = append(fruitsInSeasonNow, fruit)
		}
	}

	return views.Index(fruitsInSeasonNow).Render(c.Request().Context(), c.Response())
}
