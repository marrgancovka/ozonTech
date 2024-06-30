package in_memory

import (
	"ozonTech/graph/model"
	"ozonTech/internal/models"
	"strconv"
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

func (r *InMemoryCommentRepo) GetByPostID(postID int) ([]*model.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var rootComments []*models.Comment
	for _, comment := range r.comments {
		if comment.PostID == postID && comment.ParentCommentID == 0 {
			rootComments = append(rootComments, comment)
		}
	}

	var comments []*model.Comment
	for _, rootComment := range rootComments {
		comments = append(comments, convertToModelComment(rootComment, r))
	}

	return comments, nil
}

func convertToModelComment(comment *models.Comment, r *InMemoryCommentRepo) *model.Comment {
	parentID := ""
	if comment.ParentCommentID != 0 {
		parentID = strconv.Itoa(comment.ParentCommentID)
	}
	modelComment := &model.Comment{
		ID:              strconv.Itoa(comment.ID),
		PostID:          strconv.Itoa(comment.PostID),
		UserID:          strconv.Itoa(comment.UserID),
		Content:         comment.Text,
		ParentCommentID: &parentID,
	}

	var childComments []*model.Comment
	for _, childID := range comment.ChildComments {
		for _, c := range r.comments {
			if c.ID == childID {
				childComments = append(childComments, convertToModelComment(c, r))
			}
		}
	}
	modelComment.ChildComments = childComments

	return modelComment
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

	r.comments[newComment.ID] = newComment
	if newComment.ParentCommentID != 0 {
		r.comments[newComment.ParentCommentID].ChildComments = append(r.comments[newComment.ParentCommentID].ChildComments, newComment.ID)
	}
	return newComment, nil
}
