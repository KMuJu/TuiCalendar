package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kmuju/TuiCalendar/cmd/model"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./events.db")
	if err != nil {
		return nil, err
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
            CREATE TABLE IF NOT EXISTS Event (
                Id TEXT PRIMARY KEY,
                Date INTEGER NOT NULL,
                Name TEXT NOT NULL,
                Description TEXT,
                Status TEXT NOT NULL,
                Start TIMESTAMP NOT NULL,
                End TIMESTAMP NOT NULL
            );
	   `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetEvents(db *sql.DB) ([]model.Event, error) {
	query := `SELECT Id, Date, Name, Description, Status, Start, End FROM Event`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var event model.Event
		var start, end string
		if err := rows.Scan(
			&event.Id,
			&event.Date,
			&event.Name,
			&event.Description,
			&event.Status,
			&start,
			&end,
		); err != nil {
			fmt.Println(err)
			return nil, err
		}

		// Parse the Start and End time strings into time.Time
		event.Start, err = time.Parse(time.RFC3339, start)
		if err != nil {
			return nil, err
		}
		event.End, err = time.Parse(time.RFC3339, end)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func InsertEvent(db *sql.DB, event model.Event) error {
	query := `
    INSERT INTO Event
    (Id, Date, Name, Description, Status, Start, End)
    VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(
		query,
		event.Id,
		event.Date,
		event.Name,
		event.Description,
		event.Status,
		event.Start.Format(time.RFC3339),
		event.End.Format(time.RFC3339),
	)
	return err
}

func DeleteEvent(db *sql.DB, eventID string) error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventID)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}

	return nil
}
