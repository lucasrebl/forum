package forum

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	//création BD
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	//crée la table si elle n'existe pas
	tableUsers := `CREATE TABLE if not EXISTS users (
        userId integer not null primary key AUTOINCREMENT,
        name text UNIQUE,
        email text UNIQUE,
		password text,
		token text
        );`
	_, err = db.Exec(tableUsers)
	if err != nil {
		log.Printf("%q: %s\n", err, tableUsers)
		return
	}

	//crée la table si elle n'existe pas
	tablePosts := `CREATE TABLE if not EXISTS posts (
		postId integer not null primary key AUTOINCREMENT,
        title text UNIQUE,
        userId integer not null,
        content text not null,
		creationDate DATETIME,
		theme text
        );`
	_, err = db.Exec(tablePosts)
	if err != nil {
		log.Printf("%q: %s\n", err, tablePosts)
		return
	}

	//crée la table si elle n'existe pas
	tableLikes := `CREATE TABLE if not EXISTS likes (
        likeId integer not null primary key AUTOINCREMENT,
        userID integer not null,
		postId integer not null,
		creationDate text
        );`
	_, err = db.Exec(tableLikes)
	if err != nil {
		log.Printf("%q: %s\n", err, tableLikes)
		return
	}

	// Crée la table si elle n'existe pas
	tableComment := `CREATE TABLE IF NOT EXISTS comments (
    commentId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    postId INTEGER NOT NULL,
    userId INTEGER NOT NULL,
    content TEXT,
    creationDate TEXT
);`
	_, err = db.Exec(tableComment)
	if err != nil {
		log.Printf("%q: %s\n", err, tableComment)
		return
	}

}
