package forum

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // Import du pilote SQLite
)

// ...

func AddCommentToPost(postID int, commentContent string) error {
	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la base de données : %w", err)
	}
	defer db.Close()

	// Préparer la requête d'insertion du commentaire
	query := "INSERT INTO comments (postID,  commentContent,creationDate) VALUES (?, ?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("erreur lors de la préparation de la requête : %w", err)
	}
	defer stmt.Close()

	// Obtenir la date et l'heure actuelles
	creationDate := time.Now().Format("2006-01-02 15:04:05")

	// Exécuter la requête d'insertion du commentaire dans la base de données
	_, err = stmt.Exec(postID, commentContent, creationDate)
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de la requête : %w", err)
	}

	fmt.Println("Comment added to post successfully.")
	return nil
}

// ...
