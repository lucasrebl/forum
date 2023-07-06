package forum

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func SplitToken(toSplit string) (int, string) {

	parts := strings.Split(toSplit, "#")

	userId := parts[0] // '1234'
	token := parts[1]  // 'abcd'

	fmt.Println("User ID: ", userId)
	fmt.Println("Token: ", token)

	i, _ := strconv.Atoi(userId)

	return i, token
}

func UserTokenIsValid(tokenToVerify string) bool {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println("impossible d'ouvrir la bdd")
		return false
	}
	defer db.Close()

	userId, tokenToVerify := SplitToken(tokenToVerify)

	// Vérification si l'utilisateur est present
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE userId = ?", userId).Scan(&count)
	if err != nil {
		fmt.Println("aucun utilisateur pour cette id")
		return false
	}

	if count == 0 {
		return false
	}

	// Vérification du mot de passe
	var storedToken string
	err = db.QueryRow("SELECT token FROM users WHERE userId = ?", userId).Scan(&storedToken)
	if err != nil {
		fmt.Println("aucun token pour cette utilisateur")
		return false
	}

	// Comparaison du token
	if storedToken != tokenToVerify {
		fmt.Println(storedToken + " != " + tokenToVerify)
		return false
	}

	return true
}
