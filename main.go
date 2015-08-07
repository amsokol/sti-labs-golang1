package main

import (
	"github.com/amsokol/sti-labs-golang1/controllers"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/thoas/stats"
	"github.com/tylerb/graceful"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()

	// Enable colored log
	e.ColoredLog(true)

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	e.Use(mw.Gzip())

	// https://github.com/thoas/stats
	s := stats.New()
	e.Use(s.Handler)
	// Route
	e.Get("/stats", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, s.Data())
	})

	// Serve index file
	e.Index("public/index.html")

	// Serve favicon
	e.Favicon("public/favicon.ico")

	// Serve static files
	e.Static("/scripts", "public/scripts")

	//-----------
	// Templates
	//-----------
	e.SetRenderer(controllers.CreateRenderer())
	e.Get("/page1", controllers.Page1Handler)
	e.Get("/page2", controllers.Page2Handler)

	// Start server
	log.Println("Server is starting for port 8080...")
	graceful.ListenAndServe(e.Server(":8080"), 2*time.Second)
	log.Println("Server stoped")
}
