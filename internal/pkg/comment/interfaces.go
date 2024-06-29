package comment

import (
	"ozonTech/graph/model"
	"ozonTech/internal/models"
)

type CommentUsecase interface {
}

type CommentRepository interface {
	GetByPostID(postID int) ([]*model.Comment, error)
	Create(comment *models.CommentCreateData) (*models.Comment, error)
}
