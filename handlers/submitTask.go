package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func SubmitTask(c echo.Context) error {
	fmt.Println(c.Request().Header.Get("HX-Request"))
	task := c.FormValue("task")
	fmt.Println(task)
	return c.String(200, "<li>"+task+"</li>")
}
