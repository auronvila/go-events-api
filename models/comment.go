package models

import "github.com/golang-events-planning-backend/db"

type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	EventId string `json:"event_id"`
	Text    string `binding:"required" json:"text"`
}

type CommentResponse struct {
	Username    string `json:"username"`
	EventName   string `json:"event_name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	EventId     string `json:"event_id"`
	Text        string `json:"text"`
}

func GetAllComments() ([]CommentResponse, error) {
	var comments []CommentResponse
	query := `SELECT username, name AS event_name, email, description,eventId,text
FROM comments
         JOIN events ON comments.eventId = events.id
         JOIN users ON comments.userId = users.id;`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment CommentResponse
		err = rows.Scan(&comment.Username, &comment.EventName, &comment.Email, &comment.Description, &comment.EventId, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func GetAllCommentsByEventId(eventId string) ([]CommentResponse, error) {
	query := `
SELECT username, name AS event_name, email, description,eventId,text
FROM comments
         JOIN events ON comments.eventId = events.id
         JOIN users ON comments.userId = users.id WHERE eventId = ?
`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(eventId)
	if err != nil {
		return nil, err
	}
	var comments []CommentResponse

	for rows.Next() {
		var comment CommentResponse
		rows.Scan(&comment.Username, &comment.EventName, &comment.Email, &comment.Description, &comment.EventId, &comment.Text)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (c Comment) Save() error {
	query := `INSERT INTO comments(id,userId,eventId,text) VALUES (?,?,?,?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(c.Id, c.UserId, c.EventId, c.Text)
	if err != nil {
		return err
	}

	return err
}
