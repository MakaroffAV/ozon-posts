package graph

import (
	"ozon-posts/internal/models"
	"ozon-posts/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostService    *service.PostService
	CommentService *service.CommentService
	Subscribers    map[string]chan *models.Comment
}

func NewResolver(postService *service.PostService, commentService *service.CommentService) *Resolver {

	return &Resolver{
		PostService:    postService,
		CommentService: commentService,
		Subscribers:    make(map[string]chan *models.Comment),
	}
}
