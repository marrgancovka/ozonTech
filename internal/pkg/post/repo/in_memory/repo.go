package in_memory

import (
	"errors"
	"ozonTech/internal/models"
	"sync"
)

type InMemoryPostRepo struct {
	mu    sync.RWMutex
	posts map[int]*models.Post
}

func NewInMemoryPostRepo() *InMemoryPostRepo {
	return &InMemoryPostRepo{
		posts: make(map[int]*models.Post),
	}
}

func (r *InMemoryPostRepo) GetAll() ([]*models.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	posts := make([]*models.Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *InMemoryPostRepo) GetByID(id int) (*models.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	post, exists := r.posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	return post, nil
}

func (r *InMemoryPostRepo) Create(post *models.PostCreateData) (*models.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var newPost models.Post

	newPost.Comments = []*models.Comment{}
	newPost.Title = post.Title
	newPost.Content = post.Content
	newPost.ID = len(r.posts) + 1
	newPost.UserID = post.UserID
	newPost.CommentsAllowed = post.CommentsAllowed

	r.posts[newPost.ID] = &newPost

	return &newPost, nil
}
