package forum

import (
	"database/sql"
	"fmt"
)

func SetUserTokenDB(userId int, token string) error {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la base de données : %w", err)
	}
	defer db.Close()

	fmt.Println(userId, token)

	// Requête pour mettre à jour le token
	_, err = db.Exec("UPDATE users SET token = ? WHERE userId = ?", token, userId)
	if err != nil {
		return fmt.Errorf("erreur lors de l'écriture du token : %w", err)
	}

	return nil
}
