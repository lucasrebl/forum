package forum

import (
	"database/sql"
	"fmt"
)

func AddUserDB(name, email, password string) error {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la base de données : %w", err)
	}
	defer db.Close()

	// Vérification si un utilisateur est deja présent
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ? OR email = ?", name, email).Scan(&count)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'unicité : %w", err)
	}

	if count > 0 {
		return fmt.Errorf("un utilisateur avec le même nom ou la même adresse e-mail existe déjà")
	}

	// Ajout de l'utilisateur
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		return fmt.Errorf("erreur lors de l'insertion de l'utilisateur : %w", err)
	}

	return nil
}
