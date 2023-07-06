package forum

import (
	"database/sql"
)

type Comment struct {
	CommentID    int    `json:"commentid"`
	PostID       int    `json:"postid"`
	UserID       int    `json:"userid"`
	CreationDate string `json:"creationdate"`
	Content      string `json:"content"`
	UserName     string `json:"username"`
}

func GetCommentsByPostIDFromDB(postID string) ([]Comment, error) {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL query to retrieve comments by post ID
	query := "SELECT * FROM comments WHERE postid = ?"
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and create Comment objects
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(
			&comment.CommentID,
			&comment.PostID,
			&comment.UserID,
			&comment.CreationDate,
			&comment.Content,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
