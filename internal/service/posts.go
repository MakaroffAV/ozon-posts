package service

import (
	"ozon-posts/internal/models"
	"ozon-posts/internal/repo"

	"github.com/google/uuid"
)

type PostService struct {
	PostStorage repo.PostRepo
}

func (s PostService) Posts() ([]*models.Post, error) {
	return s.PostStorage.Posts()
}

func (s PostService) Post(postID string) (*models.Post, error) {
	return s.PostStorage.Post(postID)
}

func (s PostService) CreatePost(title string, content string, author string, commentsAllowed bool) (*models.Post, error) {

	p := &models.Post{
		ID:              uuid.New().String(),
		Title:           title,
		Author:          author,
		Content:         content,
		CommentsAllowed: commentsAllowed,
	}
	return p, s.PostStorage.CreatePost(p)

}

func NewPostService(storage repo.PostRepo) *PostService {

	return &PostService{
		PostStorage: storage,
	}

}
