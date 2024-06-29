package post

import (
	"ozonTech/graph/model"
	"ozonTech/internal/models"
)

type PostUsecase interface {
	GetAllPosts() ([]*model.Post, error)
	GetPostByID(id int) (*model.Post, error)
	CreatePost(post *models.PostCreateData) (*model.Post, error)
}

type PostRepository interface {
	GetAll() ([]*models.Post, error)
	GetByID(id int) (*models.Post, error)
	Create(post *models.PostCreateData) (*models.Post, error)
}
