package forum

import (
	"database/sql"
	"fmt"
	"time"
)

func GetPostsFromDB() ([]Post, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture de la base de données : %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT postID, title, userId, content, creationDate, theme FROM posts")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des posts : %w", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		var creationDateBDD string
		err := rows.Scan(&post.PostId, &post.Title, &post.UserId, &post.Content, &creationDateBDD, &post.Theme)
		if err != nil {
			return nil, fmt.Errorf("erreur lors de la lecture des données des posts : %w", err)
		}

		fmt.Println("Creation Date BDD -> ", creationDateBDD)

		parsedCreationDate, err := time.Parse(time.RFC3339, creationDateBDD)
		if err != nil {
			return nil, fmt.Errorf("erreur lors de la conversion de la date : %w", err)
		}
		fmt.Println("Creation Date after parse -> ", parsedCreationDate)

		fmt.Println("Creation Date after Format -> ", parsedCreationDate.Format("2006-01-02"))

		post.CreationDate = parsedCreationDate

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture des données des posts : %w", err)
	}

	return posts, nil
}
