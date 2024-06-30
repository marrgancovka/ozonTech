package usecase

import (
	"errors"
	"ozonTech/graph/model"
	"ozonTech/internal/models"
	"ozonTech/internal/pkg/comment"
	"ozonTech/internal/utils"
)

type CommentUsecase struct {
	repo comment.CommentRepository
}

func NewCommentUsecase(repo comment.CommentRepository) *CommentUsecase {
	return &CommentUsecase{repo: repo}
}

func (u *CommentUsecase) GetCommentsByPostID(postID int) ([]*model.Comment, error) {
	return u.repo.GetByPostID(postID)
}

func (u *CommentUsecase) CreateComment(comment *models.CommentCreateData) (*model.Comment, error) {
	if len(comment.Text) > 2000 {
		return nil, errors.New("comment content exceeds 2000 characters")
	}
	created, err := u.repo.Create(comment)
	if err != nil {
		return nil, err
	}
	result := utils.ConvertToGraphQLComment(created)
	return result, nil
}
