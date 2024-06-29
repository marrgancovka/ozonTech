package post

import "ozonTech/internal/models"

type PostUsecase interface {
}

type PostRepository interface {
	GetAll() ([]*models.Post, error)
	GetByID(id int) (*models.Post, error)
	Create(post *models.PostCreateData) (*models.Post, error)
}
