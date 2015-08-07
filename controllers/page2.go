package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	Page2Data struct {
		Name string
		Title string
	}
)

var Page2File string = "public/views/page2.html"

func Page2Handler(c *echo.Context) error {
	return c.Render(http.StatusOK, "page2", Page2Data{Name: "Sasha", Title:"CTO )))"})
}
