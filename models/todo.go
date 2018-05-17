package models

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

// Task is a struct containing Task data
type Job struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TodoCollection is collection of Tasks
type JobCollection struct {
	Jobs []Job `json:"items"`
}

func GetJobs(db *sql.DB) JobCollection {
	sql := "SELECT * FROM jobs"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := JobCollection{}
	for rows.Next() {
		job := Job{}
		err2 := rows.Scan(&job.ID, &job.Name)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Jobs = append(result.Jobs, job)
	}
	return result
}

func PutJob(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO jobs(name) VALUES(?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteJob(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM jobs WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}