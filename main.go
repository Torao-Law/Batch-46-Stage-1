package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/about", func(c echo.Context) error {
		return c.String(200, "Ini tentang saya yang di tipu golang")
	})

	fmt.Println("server started on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
