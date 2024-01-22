package handlers

import (
	"github.com/jacksonflores/go-todo/templates"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	page := templates.Index()

	return page.Render(c.Request().Context(), c.Response().Writer)
}
