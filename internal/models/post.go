package models

type Post struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Content         string `json:"content"`
	CommentsAllowed bool   `json:"comments_allowed"`
}
