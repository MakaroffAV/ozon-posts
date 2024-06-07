package service

import (
	"errors"
	"ozon-posts/internal/models"
	"ozon-posts/internal/repo"
	"time"

	"github.com/google/uuid"
)

var errEmptyPost = errors.New("post is not found")
var errEmptyParent = errors.New("empty comment parent")
var errMaxCommentLen = errors.New("comment is too long")
var errCommentsNotAllowed = errors.New("comments are not allowed")

type CommentService struct {
	CommentStorage repo.CommentRepo
}

func (s CommentService) CreateComment(postID string, parentID *string, author string, content string, postService *PostService) (*models.Comment, error) {

	if len(content) > 2000 {
		return nil, errMaxCommentLen
	}

	// не знаю надо  проверять  тут  есть
	// ли пост, или комментарий-родитель,
	// по идее, если  работает в  проде -
	// то там база  данных выкинет ошибку
	// из-за      настроенных      ключей
	//
	// но  напишем  явную  проверку, в  любом
	// случае не пусть лучше будет тормизить,
	// чем      сломается      логика,   имхо

	p, pErr := postService.Post(postID)
	if p == nil || pErr != nil {
		return nil, errEmptyPost
	}

	if !p.CommentsAllowed {
		return nil, errCommentsNotAllowed
	}

	if parentID != nil {
		c, cErr := s.CommentStorage.Comment(*parentID)
		if cErr != nil || c == nil {
			return nil, errEmptyParent
		}
	}

	c := &models.Comment{
		ID:        uuid.New().String(),
		Author:    author,
		PostID:    postID,
		Content:   content,
		ParentID:  parentID,
		CreatedAt: time.Now().Unix(),
	}
	return c, s.CommentStorage.CreateComment(c)

}

func (s CommentService) CommentChildren(parentID string) ([]*models.Comment, error) {

	return s.CommentStorage.CommentChildren(parentID)

}

func (s CommentService) Comments(postID string, limit, offset int) ([]*models.Comment, error) {

	if limit < 0 && offset < 0 {
		return s.CommentStorage.CommentsAll(postID)
	} else {
		return s.CommentStorage.Comments(postID, limit, offset)
	}

}

func NewCommentService(storage repo.CommentRepo) *CommentService {

	return &CommentService{
		CommentStorage: storage,
	}

}
