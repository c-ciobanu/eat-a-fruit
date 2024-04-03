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
	{Id: "apple", Name: "Apple", Months: []time.Month{time.September, time.October, time.November, time.December, time.January, time.February}, Img: "/assets/img/apple.jpg"},
	{Id: "apricot", Name: "Apricot", Months: []time.Month{time.June, time.July, time.August}, Img: "/assets/img/apricot.jpg"},
	{Id: "blackberry", Name: "Blackberry", Months: []time.Month{time.July, time.August, time.September, time.October}, Img: "/assets/img/blackberry.jpg"},
	{Id: "blackcurrant", Name: "Blackcurrant", Months: []time.Month{time.June, time.July, time.August}, Img: "/assets/img/blackcurrant.jpg"},
	{Id: "blueberry", Name: "Blueberry", Months: []time.Month{time.July}, Img: "/assets/img/blueberry.jpg"},
	{Id: "cherry", Name: "Cherry", Months: []time.Month{time.June, time.July, time.August}, Img: "/assets/img/cherry.jpg"},
	{Id: "cranberry", Name: "Cranberry", Months: []time.Month{time.November, time.December}, Img: "/assets/img/cranberry.jpg"},
	{Id: "damson", Name: "Damson", Months: []time.Month{time.August, time.September}, Img: "/assets/img/damson.jpg"},
	{Id: "elderberry", Name: "Elderberry", Months: []time.Month{time.October, time.November}, Img: "/assets/img/elderberry.jpg"},
	{Id: "gooseberry", Name: "Gooseberry", Months: []time.Month{time.June, time.July}, Img: "/assets/img/gooseberry.jpg"},
	{Id: "greengage", Name: "Greengage", Months: []time.Month{time.July, time.August}, Img: "/assets/img/greengage.jpg"},
	{Id: "loganberry", Name: "Loganberry", Months: []time.Month{time.July, time.August}, Img: "/assets/img/loganberry.jpg"},
	{Id: "pear", Name: "Pear", Months: []time.Month{time.September, time.October, time.November, time.December, time.January, time.February}, Img: "/assets/img/pear.jpg"},
	{Id: "plum", Name: "Plum", Months: []time.Month{time.August, time.September}, Img: "/assets/img/plum.jpg"},
	{Id: "quince", Name: "Quince", Months: []time.Month{time.October, time.November, time.December}, Img: "/assets/img/quince.jpg"},
	{Id: "raspberry", Name: "Raspberry", Months: []time.Month{time.June, time.July, time.August, time.September}, Img: "/assets/img/raspberry.jpg"},
	{Id: "redcurrant", Name: "Redcurrant", Months: []time.Month{time.June, time.July, time.August}, Img: "/assets/img/redcurrant.jpg"},
	{Id: "strawberry", Name: "Strawberry", Months: []time.Month{time.May, time.June, time.July, time.August, time.September}, Img: "/assets/img/strawberry.jpg"},
	{Id: "tayberry", Name: "Tayberry", Months: []time.Month{time.June}, Img: "/assets/img/tayberry.jpg"},
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
