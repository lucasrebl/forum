package forum

import (
	"database/sql"
	"fmt"
	"time"
)

func AddPostDB(title string, userID int, content string, creationDate time.Time, theme string) error {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la base de données : %w", err)
	}
	defer db.Close()

	// Ajout du post
	_, err = db.Exec("INSERT INTO posts (title, userID, content, creationDate, theme) VALUES (?, ?, ?, ?, ?)",
		title, userID, content, creationDate.Format("2006-01-02"), theme)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion du post : %w", err)
	}

	return nil
}
