package in_memory

import (
	"fmt"
	"ozonTech/internal/models"
	"sync"
)

type InMemoryCommentRepo struct {
	mu       sync.RWMutex
	comments map[int]*models.Comment
}

func NewInMemoryCommentRepo() *InMemoryCommentRepo {
	return &InMemoryCommentRepo{
		comments: make(map[int]*models.Comment),
	}
}

func (r *InMemoryCommentRepo) GetByPostID(postID int) ([]*models.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var comments []*models.Comment
	for _, comment := range r.comments {
		if comment.PostID == postID {
			comments = append(comments, comment)
		}
	}
	return comments, nil
}

func (r *InMemoryCommentRepo) Create(comment *models.CommentCreateData) (*models.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var newComment = &models.Comment{
		ID:              len(r.comments) + 1,
		UserID:          comment.UserID,
		PostID:          comment.PostID,
		ChildComments:   make([]int, 0),
		ParentCommentID: comment.ParentCommentID,
		Text:            comment.Text,
	}

	fmt.Println(newComment)

	r.comments[newComment.ID] = newComment
	fmt.Println(newComment.ID)
	if newComment.ParentCommentID != 0 {
		r.comments[newComment.ParentCommentID].ChildComments = append(r.comments[newComment.ParentCommentID].ChildComments, newComment.ID)
	}
	fmt.Println(newComment.ParentCommentID)
	return newComment, nil
}
