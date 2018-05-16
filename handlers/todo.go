package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "todos")
	}
}

func PutTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"created": 123,
		})
	}
}

func DeleteTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
