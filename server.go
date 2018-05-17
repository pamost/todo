package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pamost/todo/handlers"
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
	e.File("/", "public/index.html")
	e.GET("/jobs", handlers.GetJobs(db))
	e.PUT("/job", handlers.PutJob(db))
	e.DELETE("/job/:id", handlers.DeleteJob(db))

	e.Logger.Fatal(e.Start(":8080"))
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
    CREATE TABLE IF NOT EXISTS jobs (
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
