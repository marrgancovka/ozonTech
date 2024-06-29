package usecase

import (
	"errors"
	"ozonTech/internal/models"
	"ozonTech/internal/pkg/comment"
)

type CommentUsecase struct {
	repo comment.CommentRepository
}

func NewCommentUsecase(repo comment.CommentRepository) *CommentUsecase {
	return &CommentUsecase{repo: repo}
}

func (u *CommentUsecase) GetCommentsByPostID(postID int) ([]*models.Comment, error) {
	return u.repo.GetByPostID(postID)
}

func (u *CommentUsecase) CreateComment(comment *models.CommentCreateData) (*models.Comment, error) {
	if len(comment.Text) > 2000 {
		return nil, errors.New("comment content exceeds 2000 characters")
	}
	return u.repo.Create(comment)
}
