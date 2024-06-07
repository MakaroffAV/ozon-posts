package db

import (
	"database/sql"
	"ozon-posts/internal/models"
	"sync"
)

type CommentDbRepository struct {
	db *sql.DB
	mu sync.RWMutex
}

func (r *CommentDbRepository) Comment(id string) (*models.Comment, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var c = &models.Comment{}

	if sErr := r.db.QueryRow(
		`
		SELECT
			id,
			post_id,
			parent_id,
			author,
			content,
			created_at
		FROM
			comment
		WHERE
			id = $1
		`,
		id,
	).Scan(
		&c.ID,
		&c.PostID,
		&c.ParentID,
		&c.Author,
		&c.Content,
		&c.CreatedAt,
	); sErr != nil {
		return nil, sErr
	}

	return c, nil

}

func (r *CommentDbRepository) CreateComment(comment *models.Comment) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	_, eErr := r.db.Exec(
		`
		INSERT INTO comment (
			id,
			post_id,
			parent_id,
			author,
			content,
			created_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
		`,
		comment.ID,
		comment.PostID,
		comment.ParentID,
		comment.Author,
		comment.Content,
		comment.CreatedAt,
	)

	return eErr

}

func (r *CommentDbRepository) CommentChildren(parentID string) ([]*models.Comment, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var c = []*models.Comment{}

	q, qErr := r.db.Query(
		`
		SELECT
			id,
			post_id,
			parent_id,
			author,
			content,
			created_at
		FROM
			comment
		WHERE
			comment.parent_id = $1
		`,
		parentID,
	)
	if qErr != nil {
		return nil, qErr
	}
	defer q.Close()

	for q.Next() {
		r := &models.Comment{}
		if sErr := q.Scan(
			&r.ID,
			&r.PostID,
			&r.ParentID,
			&r.Author,
			&r.Content,
			&r.CreatedAt,
		); sErr != nil {
			return nil, sErr
		}
		c = append(c, r)
	}

	return c, nil

}

func (r *CommentDbRepository) CommentsAll(postID string) ([]*models.Comment, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var (
		s = `
		SELECT
			id,
			post_id,
			parent_id,
			author,
			content,
			created_at
		FROM
			comment
		WHERE
			comment.post_id = $1
		`
		c = []*models.Comment{}
	)

	q, qErr := r.db.Query(s, postID)
	if qErr != nil {
		return nil, qErr
	}
	defer q.Close()

	for q.Next() {
		r := &models.Comment{}
		if sErr := q.Scan(
			&r.ID,
			&r.PostID,
			&r.ParentID,
			&r.Author,
			&r.Content,
			&r.CreatedAt,
		); sErr != nil {
			return nil, sErr
		}
		c = append(c, r)
	}

	return c, nil

}

func (r *CommentDbRepository) Comments(postID string, limit, offset int) ([]*models.Comment, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	var c = []*models.Comment{}

	q, qErr := r.db.Query(
		`
		SELECT
			id,
			post_id,
			parent_id,
			author,
			content,
			created_at
		FROM
			comment
		WHERE
			comment.post_id = $1
		LIMIT
			$2
		OFFSET
			$3
		`,
		postID,
		limit,
		offset,
	)
	if qErr != nil {
		return nil, qErr
	}
	defer q.Close()

	for q.Next() {
		r := &models.Comment{}
		if sErr := q.Scan(
			&r.ID,
			&r.PostID,
			&r.ParentID,
			&r.Author,
			&r.Content,
			&r.CreatedAt,
		); sErr != nil {
			return nil, sErr
		}
		c = append(c, r)
	}

	return c, nil

}

func NewCommentDbRepository(dbConn *sql.DB) *CommentDbRepository {

	return &CommentDbRepository{
		mu: sync.RWMutex{},
		db: dbConn,
	}

}
