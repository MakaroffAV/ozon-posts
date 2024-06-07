package memory

import (
	"ozon-posts/internal/models"
	"sync"
)

type CommentRepository struct {
	mu              sync.RWMutex
	comments        map[string]*models.Comment
	postComments    map[string][]*models.Comment
	commentChildren map[string][]*models.Comment
}

func (r *CommentRepository) Comment(id string) (*models.Comment, error) {

	c, cExists := r.comments[id]
	if cExists {
		return c, nil
	}

	return nil, nil

}

func (r *CommentRepository) CreateComment(comment *models.Comment) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.comments[comment.ID] = comment
	r.postComments[comment.PostID] = append(r.postComments[comment.PostID], comment)

	if comment.ParentID != nil {
		r.commentChildren[*comment.ParentID] = append(r.commentChildren[*comment.ParentID], comment)
	}

	return nil

}

func (r *CommentRepository) CommentsAll(postID string) ([]*models.Comment, error) {

	return r.postComments[postID], nil

}

func (r *CommentRepository) CommentChildren(parentID string) ([]*models.Comment, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	c, cExists := r.commentChildren[parentID]
	if cExists {
		return c, nil
	}

	return []*models.Comment{}, nil

}

func (r *CommentRepository) Comments(postID string, limit, offset int) ([]*models.Comment, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	c := r.postComments[postID]

	if offset >= len(c) {
		return []*models.Comment{}, nil
	}

	if offset+limit > len(c) {
		limit = len(c) - offset
	}

	return c[offset : offset+limit], nil

}

func NewCommentRepository() *CommentRepository {

	return &CommentRepository{
		mu:              sync.RWMutex{},
		comments:        make(map[string]*models.Comment),
		postComments:    make(map[string][]*models.Comment),
		commentChildren: make(map[string][]*models.Comment),
	}

}
