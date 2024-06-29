package comment

import "ozonTech/internal/models"

type CommentUsecase interface {
}

type CommentRepository interface {
	GetByPostID(postID int) ([]*models.Comment, error)
	Create(comment *models.CommentCreateData) (*models.Comment, error)
}
