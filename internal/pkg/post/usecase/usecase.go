package usecase

import (
	"ozonTech/graph/model"
	"ozonTech/internal/models"
	"ozonTech/internal/pkg/comment"
	"ozonTech/internal/pkg/post"
	"ozonTech/internal/utils"
)

type PostUsecase struct {
	repo        post.PostRepository
	repoComment comment.CommentRepository
}

func NewPostUsecase(repo post.PostRepository, repoComment comment.CommentRepository) *PostUsecase {
	return &PostUsecase{repo: repo, repoComment: repoComment}
}

func (u *PostUsecase) GetAllPosts() ([]*model.Post, error) {
	posts, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var result []*model.Post
	for _, onePost := range posts {
		result = append(result, utils.ConvertToGraphQLPost(onePost))
	}
	return result, nil
}

func (u *PostUsecase) GetPostByID(id int) (*model.Post, error) {
	curPost, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	convertCurPost := utils.ConvertToGraphQLPost(curPost)

	comments, err := u.repoComment.GetByPostID(curPost.ID)
	if err != nil {
		return nil, err
	}

	convertCurPost.Comments = comments

	return convertCurPost, nil
}

func (u *PostUsecase) CreatePost(post *models.PostCreateData) (*model.Post, error) {
	createdPost, err := u.repo.Create(post)
	if err != nil {
		return nil, err
	}
	resultPost := utils.ConvertToGraphQLPost(createdPost)
	return resultPost, nil
}
