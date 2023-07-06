package forum

import (
	"database/sql"
	"fmt"
)

func GetUserLogins(email, password string) error {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la base de données : %w", err)
	}
	defer db.Close()

	// Vérification si un utilisateur est deja présent
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return fmt.Errorf("erreur lors de la vérification de l'unicité : %w", err)
	}

	if count == 0 {
		return fmt.Errorf("l'email ne corespond à aucun compte")
	}

	// Vérification du mot de passe
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération du mot de passe : %w", err)
	}

	// Comparaison du mot de passe
	if storedPassword != password {
		return fmt.Errorf("le mot de passe est incorrect")
	}

	return nil
}
