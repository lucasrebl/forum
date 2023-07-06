package forum

import (
	"database/sql"
	"fmt"
)

func GetPostByIDFromDB(postID int) (*Post, error) {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM posts WHERE postId = ?", postID).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("impossible d'acceder au nombre de post : %w", err)
	}

	if count == 0 {
		return nil, fmt.Errorf("aucun post dans la bdd")
	}

	var post Post

	err = db.QueryRow("SELECT * FROM posts WHERE postId = ?", postID).Scan(&post.PostId, &post.Title, &post.UserId, &post.Content, &post.CreationDate, &post.Theme)
	if err != nil {
		return nil, fmt.Errorf("GetPostIdFromDB -> Erreur l'ors de la lecture du post dans la bdd")
	}

	return &post, err
}
