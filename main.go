package main

import (
	"github.com/jacksonflores/go-todo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", handlers.Home)
	e.POST("/", handlers.SubmitTask)

	e.Logger.Fatal(e.Start(":8080"))
}
