package comment

import (
	"ozonTech/graph/model"
	"ozonTech/internal/models"
)

type CommentUsecase interface {
	GetCommentsByPostID(postID int) ([]*model.Comment, error)
	CreateComment(comment *models.CommentCreateData) (*model.Comment, error)
}

type CommentRepository interface {
	GetByPostID(postID int) ([]*model.Comment, error)
	Create(comment *models.CommentCreateData) (*models.Comment, error)
}
