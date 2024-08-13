package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-events-planning-backend/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (receiver Event) Save() error {
	query := `
INSERT INTO events(name,description,location,dateTime,userId) 
VALUES (? , ? , ? , ? , ?) 
`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(receiver.Name, receiver.Description, receiver.Location, receiver.DateTime, receiver.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	receiver.Id = id
	return err
}

func GetEvents() ([]Event, error) {
	var events []Event
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	query := `
	SELECT * FROM events WHERE id = ?
	`
	row := db.DB.QueryRow(query, eventId)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no event found with id %d", eventId)
		}
		return nil, fmt.Errorf("error fetching event: %v", err)
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
UPDATE events
SET name = ?, description = ?, location =?,dateTime = ?
WHERE id = ?
`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return err
}
