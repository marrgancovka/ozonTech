package usecase

import (
	"fmt"
	"ozonTech/internal/models"
	"ozonTech/internal/pkg/comment"
	"ozonTech/internal/pkg/post"
)

type PostUsecase struct {
	repo        post.PostRepository
	repoComment comment.CommentRepository
}

func NewPostUsecase(repo post.PostRepository, repoComment comment.CommentRepository) *PostUsecase {
	return &PostUsecase{repo: repo, repoComment: repoComment}
}

func (u *PostUsecase) GetAllPosts() ([]*models.Post, error) {
	return u.repo.GetAll()
}

func (u *PostUsecase) GetPostByID(id int) (*models.Post, error) {

	curPost, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	comments, err := u.repoComment.GetByPostID(curPost.ID)
	if err != nil {
		return nil, err
	}
	curPost.Comments = comments
	return curPost, nil
}

func (u *PostUsecase) CreatePost(post *models.PostCreateData) (*models.Post, error) {
	fmt.Println("usecase post")
	return u.repo.Create(post)
}
