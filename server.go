package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	// Echo instance
	e := echo.New()

	// Middleware
	// Logger is for server log
	e.Use(middleware.Logger())

	// Recover from panics anywhere in the chain
	e.Use(middleware.Recover())

	// CORS gives web servers cross-domain access controls,
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
		fmt.Println(requested_id)
		return c.JSON(http.StatusOK, requested_id)
	})

	e.GET("/get", func(c echo.Context) error { return c.JSON(200, "GET") })

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS todo (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
