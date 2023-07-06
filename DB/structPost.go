package forum

import "time"

type Post struct {
	PostId       int       `json:"postid"`
	Title        string    `json:"title"`
	UserId       int       `json:"userid"`
	Content      string    `json:"content"`
	CreationDate time.Time `json:"creationdate"`
	Theme        string    `json:"theme"`
	Comments     []Comment `json:"comments"`
}

type FormatedPost struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Username string    `json:"username"`
	Theme    string    `json:"theme"`
	Comment  []Comment `json:"comment"`
}
