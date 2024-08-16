package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-events-planning-backend/db"
	"time"
)

type Event struct {
	Id          string
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      string
}

type UserAssignedEvents struct {
	UserID           string `json:"userId"`
	UserEmail        string `json:"userEmail"`
	EventID          string `json:"eventId"`
	EventName        string `json:"eventName"`
	EventLocation    string `json:"eventLocation"`
	EventDescription string `json:"eventDescription"`
	EventDate        string `json:"eventDate"`
}

func (receiver Event) Save() error {
	query := `
INSERT INTO events(id,name,description,location,dateTime,userId) 
VALUES (?, ? , ? , ? , ? , ?) 
`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(receiver.Id, receiver.Name, receiver.Description, receiver.Location, receiver.DateTime, receiver.UserId)
	if err != nil {
		return err
	}

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

func GetEventById(eventId string) (*Event, error) {
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

func (event Event) Delete() error {
	query := `
DELETE FROM events WHERE id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Id)

	if err != nil {
		return err
	}
	return nil
}

func (e Event) Register(userId string, id string) error {
	query := `
INSERT INTO user_events(id,event_id, user_id) VALUES (?,?,?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, e.Id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) CancelRegistration(userId string) error {
	query := `
DELETE FROM user_events WHERE event_id = ? AND user_id = ?
`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Id, userId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no registration found for event_id %d and user_id %d", e.Id, userId)
	}

	return nil
}

func GetUserAssignedRegistrations() ([]UserAssignedEvents, error) {
	query := ` 
SELECT 
    u.id AS user_id, 
    u.email, 
    e.id AS event_id, 
    e.name, 
    e.location,
    e.description,
    e.dateTime
FROM 
    user_events ue
JOIN 
    users u ON ue.user_id = u.id
JOIN 
    events e ON ue.event_id = e.id;
`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []UserAssignedEvents

	for rows.Next() {
		var ue UserAssignedEvents
		err := rows.Scan(&ue.UserID, &ue.UserEmail, &ue.EventID, &ue.EventName, &ue.EventLocation, &ue.EventDescription, &ue.EventDate)
		if err != nil {
			return nil, err
		}
		results = append(results, ue)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func GetSpecificEventUserAssignee(eventId string) ([]UserResponse, error) {
	query := ` 
SELECT 
    u.id AS user_id, 
    u.email
FROM 
    user_events ue
JOIN 
    users u ON ue.user_id = u.id
WHERE 
    ue.event_id = ?;
`

	rows, err := db.DB.Query(query, eventId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserResponse

	for rows.Next() {
		var user UserResponse
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors after the iteration is done
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
