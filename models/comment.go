package models

import "github.com/golang-events-planning-backend/db"

type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	EventId string `json:"event_id"`
	Text    string `binding:"required"`
}

type CommentResponse struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	EventId   string `json:"event_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

func GetAllComments() ([]CommentResponse, error) {
	var comments []CommentResponse
	query := `SELECT * FROM comments`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment CommentResponse
		err = rows.Scan(&comment.Id, &comment.UserId, &comment.EventId, &comment.Text, &comment.CreatedAt)
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
