package db

import (
	"ozon-posts/internal/models"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
)

var newTestPost = &models.Post{
	ID:              uuid.New().String(),
	Title:           "test_post_t",
	Author:          "test_post_a",
	Content:         "test_post_c",
	CommentsAllowed: true,
}

var comment1Id = uuid.New().String()
var comment2Id = uuid.New().String()

func TestCommentsInit(t *testing.T) {

	c, cErr := Connection()
	if cErr != nil {
		t.Fatalf(
			"Test failed: TestCommentsInit (%s)", cErr.Error(),
		)
	}
	defer c.Close()

	r := PostDbRepository{
		db: c,
		mu: sync.RWMutex{},
	}

	if iErr := r.CreatePost(newTestPost); iErr != nil {
		t.Fatalf(
			"Test failed: TestCommentsInit (%s)", cErr.Error(),
		)
	}

}

func TestCreateComment(t *testing.T) {

	c, cErr := Connection()
	if cErr != nil {
		t.Fatalf(
			"Test failed: TestCreateComment (%s)", cErr.Error(),
		)
	}
	defer c.Close()

	r := &CommentDbRepository{
		db: c,
		mu: sync.RWMutex{},
	}

	c1 := &models.Comment{
		ID:        comment1Id,
		PostID:    newTestPost.ID,
		ParentID:  nil,
		Author:    "comment_test_a",
		Content:   "comment_test_c",
		CreatedAt: time.Now().Unix(),
	}

	c2 := &models.Comment{
		ID:        comment2Id,
		PostID:    newTestPost.ID,
		ParentID:  &c1.ID,
		Author:    "comment_test_a_n",
		Content:   "comment_test_c_n",
		CreatedAt: time.Now().Unix(),
	}

	testCases := []*models.Comment{
		c1,
		c2,
	}

	for i, comment := range testCases {
		if dErr := r.CreateComment(comment); dErr != nil {
			t.Fatalf(
				"Test failed: TestCreateComment (%d, %s)", i, cErr.Error(),
			)
		}
	}

}

func TestCommentChildren(t *testing.T) {

	c, cErr := Connection()
	if cErr != nil {
		t.Fatalf(
			"Test failed: TestCommentChildren (%s)", cErr.Error(),
		)
	}
	defer c.Close()

	r := &CommentDbRepository{
		db: c,
		mu: sync.RWMutex{},
	}

	q, qErr := r.CommentChildren(comment1Id)
	if qErr != nil {
		t.Fatalf(
			"Test failed: TestCommentChildren (%s)", cErr.Error(),
		)
	}

	if len(q) != 1 {
		t.Fatalf(
			"Test failed: TestCommentChildren (%s), got (%d), want (1)", cErr.Error(), len(q),
		)
	}

}
