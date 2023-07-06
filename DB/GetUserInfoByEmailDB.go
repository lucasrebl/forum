package forum

import (
	"database/sql"
)

func GetUserInfoByEmailDB(email string) User {
	db, _ := sql.Open("sqlite3", "database.db")
	defer db.Close()

	var userInfo User
	_ = db.QueryRow("SELECT name FROM users WHERE email = ?", email).Scan(&userInfo.Name)
	_ = db.QueryRow("SELECT userId FROM users WHERE email = ?", email).Scan(&userInfo.Id)

	return userInfo
}
