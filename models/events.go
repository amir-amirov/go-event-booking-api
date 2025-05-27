package models

import (
	"fmt"
	"time"

	"github.com/amir-amirov/go-event-booking-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      int64     `json:"userId"`
}

func (event *Event) Save() error {
	query := `
		INSERT INTO events (name, description, location, dateTime, userId)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := db.DB.QueryRow(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId).Scan(&event.ID)
	if err != nil {
		return fmt.Errorf("failed to save event: %s", err.Error())
	}

	fmt.Println("Event saved successfully")

	return nil
}

func GetEvents() ([]Event, error) {

	var events []Event = []Event{}

	query := `
		SELECT * FROM events
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve events: %s", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event: %s", err.Error())
		}

		events = append(events, event)
	}

	return events, nil
}

func Delete(id int64) error {
	query := `
		DELETE FROM events
		WHERE id = $1
	`

	_, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete event: %s", err.Error())
	}

	fmt.Println("Event deleted successfully")

	return nil
}

func (event *Event) Update(id int64) error {

	query := `
		UPDATE events
		SET name = $1, description = $2, location = $3, dateTime = $4, userId = $5
		WHERE id = $6
	`

	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, id)
	if err != nil {
		return fmt.Errorf("failed to update event: %s", err.Error())
	}

	event.ID = id

	fmt.Println("Event updated successfully")
	return nil
}
