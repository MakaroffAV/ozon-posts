package repo

import (
	"ozon-posts/internal/models"
)

type PostRepo interface {
	Posts() ([]*models.Post, error)
	CreatePost(post *models.Post) error
	Post(postID string) (*models.Post, error)
}

type CommentRepo interface {
	Comment(id string) (*models.Comment, error)
	CreateComment(comment *models.Comment) error
	CommentsAll(postID string) ([]*models.Comment, error)
	CommentChildren(parentID string) ([]*models.Comment, error)
	Comments(postID string, limit, offset int) ([]*models.Comment, error)
}
