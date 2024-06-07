package models

type Comment struct {
	ID        string  `json:"id"`
	Author    string  `json:"author"`
	PostID    string  `json:"post_id"`
	Content   string  `json:"content"`
	ParentID  *string `json:"parent_id"`
	CreatedAt int64   `json:"created_at"`
}
