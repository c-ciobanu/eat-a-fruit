package main

import (
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
	{Id: "apple", Name: "Apple", Months: []time.Month{time.September, time.October, time.November, time.December, time.January, time.February}},
	{Id: "apricot", Name: "Apricot", Months: []time.Month{time.June, time.July, time.August}},
	{Id: "blackberry", Name: "Blackberry", Months: []time.Month{time.July, time.August, time.September, time.October}},
	{Id: "blackcurrant", Name: "Blackcurrant", Months: []time.Month{time.June, time.July, time.August}},
	{Id: "blueberry", Name: "Blueberry", Months: []time.Month{time.July}},
	{Id: "cherry", Name: "Cherry", Months: []time.Month{time.June, time.July, time.August}},
	{Id: "cranberry", Name: "Cranberry", Months: []time.Month{time.November, time.December}},
	{Id: "damson", Name: "Damson", Months: []time.Month{time.August, time.September}},
	{Id: "elderberry", Name: "Elderberry", Months: []time.Month{time.October, time.November}},
	{Id: "gooseberry", Name: "Gooseberry", Months: []time.Month{time.June, time.July}},
	{Id: "greengage", Name: "Greengage", Months: []time.Month{time.July, time.August}},
	{Id: "loganberry", Name: "Loganberry", Months: []time.Month{time.July, time.August}},
	{Id: "pear", Name: "Pear", Months: []time.Month{time.September, time.October, time.November, time.December, time.January, time.February}},
	{Id: "plum", Name: "Plum", Months: []time.Month{time.August, time.September}},
	{Id: "quince", Name: "Quince", Months: []time.Month{time.October, time.November, time.December}},
	{Id: "raspberry", Name: "Raspberry", Months: []time.Month{time.June, time.July, time.August, time.September}},
	{Id: "redcurrant", Name: "Redcurrant", Months: []time.Month{time.June, time.July, time.August}},
	{Id: "strawberry", Name: "Strawberry", Months: []time.Month{time.May, time.June, time.July, time.August, time.September}},
	{Id: "tayberry", Name: "Tayberry", Months: []time.Month{time.June}},
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", indexHandler)

	e.Logger.Fatal(e.Start(":4000"))
}

func indexHandler(c echo.Context) error {
	currentMonth := time.Now().Month()

	var fruitsInSeasonNow []models.Fruit
	for _, fruit := range ukSeasonalFruits {
		for _, month := range fruit.Months {
			if month >= currentMonth-1 && month <= currentMonth+1 {
				fruitsInSeasonNow = append(fruitsInSeasonNow, fruit)
				break
			}
		}
	}

	return views.Index(fruitsInSeasonNow).Render(c.Request().Context(), c.Response())
}
