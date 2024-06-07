package memory

import (
	"ozon-posts/internal/models"
	"sync"
)

type PostRepository struct {
	mu    sync.RWMutex
	posts map[string]*models.Post
}

func (r *PostRepository) Posts() ([]*models.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	p := []*models.Post{}
	for _, v := range r.posts {
		p = append(p, v)
	}

	return p, nil

}

func (r *PostRepository) CreatePost(post *models.Post) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.posts[post.ID] = post

	return nil

}

func (r *PostRepository) Post(postID string) (*models.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	p, pExist := r.posts[postID]
	if pExist {
		return p, nil
	}

	return nil, nil

}

func NewPostRepository() *PostRepository {

	return &PostRepository{
		posts: make(map[string]*models.Post),
	}

}
