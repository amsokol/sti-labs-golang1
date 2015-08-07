package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	Page1Data struct {
		Name string
		Title string
	}
)

var Page1File string = "public/views/page1.html"

func Page1Handler(c *echo.Context) error {
	return c.Render(http.StatusOK, "page1", Page1Data{Name: "Sasha", Title:"CIO )))"})
}
