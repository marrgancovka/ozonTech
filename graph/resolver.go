package graph

import (
	authUsecase "ozonTech/internal/pkg/auth/usecase"
	commentUsecase "ozonTech/internal/pkg/comment/usecase"
	postUsecase "ozonTech/internal/pkg/post/usecase"
	"sync"
)

type Resolver struct {
	PostUsecase    *postUsecase.PostUsecase
	CommentUsecase *commentUsecase.CommentUsecase
	AuthUsecase    *authUsecase.AuthUsecase
	mu             sync.RWMutex
	//subscribers map[string]chan *entity.Comment
}

func NewResolver(postUseCase *postUsecase.PostUsecase, commentUseCase *commentUsecase.CommentUsecase) *Resolver {
	return &Resolver{
		PostUsecase:    postUseCase,
		CommentUsecase: commentUseCase,
		//subscribers:    make(map[string]chan *entity.Comment),
	}
}
