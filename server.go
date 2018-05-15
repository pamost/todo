package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// Logger is for server log
	e.Use(middleware.Logger())

	// Recover from panics anywhere in the chain
	e.Use(middleware.Recover())

	// ORS gives web servers cross-domain access controls,
	// which enable secure cross-domain data transfers
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Route => handler
	e.GET("/", func(c echo.Context) error {

		return c.JSON(http.StatusOK, "Hi!")
	})

	e.GET("/id/:id", func(c echo.Context) error {
		requested_id := c.Param("id")
		fmt.Println(requested_id);
		return c.JSON(http.StatusOK, requested_id)
	})

	e.GET("/get", func(c echo.Context) error { return c.JSON(200, "GET") })

	e.Logger.Fatal(e.Start(":8000"))
}