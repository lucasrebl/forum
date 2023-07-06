package forum

import (
	"database/sql"
)

func GetUserInfoByIdDB(userId int) User {
	db, _ := sql.Open("sqlite3", "database.db")
	defer db.Close()

	var userInfo User
	_ = db.QueryRow("SELECT name FROM users WHERE userId = ?", userId).Scan(&userInfo.Name)
	_ = db.QueryRow("SELECT email FROM users WHERE userId = ?", userId).Scan(&userInfo.Email)

	return userInfo
}
