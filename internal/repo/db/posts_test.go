package db

import (
	"ozon-posts/internal/models"
	"reflect"
	"sync"
	"testing"

	"github.com/google/uuid"
)

var testPost = &models.Post{
	ID:              uuid.New().String(),
	Title:           "test_post_t",
	Author:          "test_post_a",
	Content:         "test_post_c",
	CommentsAllowed: true,
}

func TestPosts(t *testing.T) {

	c, cErr := Connection()
	if cErr != nil {
		t.Fatalf(
			"Test failed: TestPosts (%s)", cErr.Error(),
		)
	}
	defer c.Close()

	r := PostDbRepository{
		db: c,
		mu: sync.RWMutex{},
	}

	_, pErr := r.Posts()
	if pErr != nil {
		t.Fatalf(
			"Test failed: TestPosts (%s)", pErr.Error(),
		)
	}

}

func TestCreatePost(t *testing.T) {

	c, cErr := Connection()
	if cErr != nil {
		t.Fatalf(
			"Test failed: TestCreatePost (%s)", cErr.Error(),
		)
	}
	defer c.Close()

	r := PostDbRepository{
		db: c,
		mu: sync.RWMutex{},
	}

	if iErr := r.CreatePost(testPost); iErr != nil {
		t.Fatalf(
			"Test failed: TestCreatePost (%s)", cErr.Error(),
		)
	}

}

func TestPost(t *testing.T) {

	c, cErr := Connection()
	if cErr != nil {
		t.Fatalf(
			"Test failed: TestPost (%s)", cErr.Error(),
		)
	}
	defer c.Close()

	r := PostDbRepository{
		db: c,
		mu: sync.RWMutex{},
	}

	p, pErr := r.Post(testPost.ID)
	if pErr != nil {
		t.Fatalf(
			"Test failed: TestPost (%s)", pErr.Error(),
		)
	}

	if !reflect.DeepEqual(p, testPost) {
		t.Fatalf(
			"Test failed: TestPost; got (%v), want (%v)", p, testPost,
		)
	}

}
