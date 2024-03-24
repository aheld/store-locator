package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/aheld/market-locator/templates"
	"github.com/labstack/echo/v4"
)

// add market struct:
func main() {
	settings := templates.Settings{
		UserName: "Aaron D",
	}
	loadedMarkets, err := Import()
	if err != nil {
		panic(err)
	}
	markets := templates.Markets{Markets: []templates.Market{}}
	for _, market := range loadedMarkets {
		m := templates.Market{
			Name:        market.Name,
			Address:     market.Address,
			Description: market.Description,
			Website:     market.Website,
			Latitude:    market.Latitude,
			Longitude:   market.Longitude,
			Image:       fmt.Sprint("https://www.usdalocalfoodportal.com/mywp/uploadimages/", market.Image),
			Products:    market.Products,
		}
		m.Slug = m.GetSlug()
		markets.Markets = append(markets.Markets, m)
	}
	settings.Markets = markets

	e := echo.New()
	e.Static("/assets", "assets")
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/markets/pa/")
	})

	e.GET("/markets/pa/", func(c echo.Context) error {
		hero := templates.StateHero("Pennsylvania")
		component := templates.Layout(settings, hero)
		return component.Render(c.Request().Context(), c.Response())
	})

	e.GET("/markets/pa/:slug", func(c echo.Context) error {
		htx := c.Request().Header.Get("Hx-Request")
		name := c.Param("slug")
		name, err := url.QueryUnescape(name)
		if err != nil {
			e.Logger.Fatal(err)
		}
		market := markets.FindSlug(name)
		if htx != "true" {
			marketHero := templates.MarketHero(market)
			component := templates.Layout(settings, marketHero)
			return component.Render(c.Request().Context(), c.Response())
		}

		component := templates.MarketHero(market)
		return component.Render(c.Request().Context(), c.Response())
	})

	e.POST("/markets/search", func(c echo.Context) error {
		marketList := markets
		searchTerm := c.FormValue("q")
		if searchTerm != "" {
			marketList = markets.Search(searchTerm)
		}
		component := templates.MarketList(marketList)
		return component.Render(c.Request().Context(), c.Response())
	})

	host := ""
	if os.Getenv("CONTAINER") == "" {
		host = "localhost"
	}

	e.Logger.Fatal(e.Start(host + ":8000"))

	fmt.Println("Listening on :8000")
}

// a function that returns the routes for kubernets health checks
func healthCheckRoutes() *echo.Echo {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	return e
}
