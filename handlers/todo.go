package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"github.com/pamost/todo/models"
	"strconv"
)

type H map[string]interface{}

func GetJobs(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch tasks using our new model
		return c.JSON(http.StatusOK, models.GetJobs(db))
	}
}

func PutJob(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var job models.Job
		// Map imcoming JSON body to the new Task
		c.Bind(&job)
		// Add a task using our new model
		id, err := models.PutJob(db, job.Name)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

func DeleteJob(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteJob(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
