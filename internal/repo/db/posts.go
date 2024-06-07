package db

import (
	"database/sql"
	"ozon-posts/internal/models"
	"sync"
)

type PostDbRepository struct {
	db *sql.DB
	mu sync.RWMutex
}

func (r *PostDbRepository) Posts() ([]*models.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var p = []*models.Post{}

	q, qErr := r.db.Query(
		`
		SELECT
			id,
			title,
			author,
			content,
			comments_on
		FROM
			post
		`,
	)
	if qErr != nil {
		return nil, qErr
	}
	defer q.Close()

	for q.Next() {
		r := &models.Post{}
		if sErr := q.Scan(
			&r.ID,
			&r.Title,
			&r.Author,
			&r.Content,
			&r.CommentsAllowed,
		); sErr != nil {
			return nil, sErr
		}
		p = append(p, r)
	}

	return p, nil

}

func (r *PostDbRepository) CreatePost(post *models.Post) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	_, eErr := r.db.Exec(
		`
		INSERT INTO post (
			id,
			title,
			author,
			content,
			comments_on
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5 
		)
		`,
		post.ID,
		post.Title,
		post.Author,
		post.Content,
		post.CommentsAllowed,
	)

	return eErr

}

func (r *PostDbRepository) Post(postID string) (*models.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var p = &models.Post{}

	if sErr := r.db.QueryRow(
		`
		SELECT
			id,
			title,
			author,
			content,
			comments_on
		FROM
			post
		WHERE
			post.id = $1
		`,
		postID,
	).Scan(
		&p.ID,
		&p.Title,
		&p.Author,
		&p.Content,
		&p.CommentsAllowed,
	); sErr != nil {
		return nil, sErr
	}

	return p, nil

}

func NewPostDbRepository(dbConn *sql.DB) *PostDbRepository {

	return &PostDbRepository{
		mu: sync.RWMutex{},
		db: dbConn,
	}

}
