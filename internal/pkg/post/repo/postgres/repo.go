package postgres

import (
	"database/sql"
	"ozonTech/internal/models"
)

type PostRepositoryImpl struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

func (r *PostRepositoryImpl) GetAll() ([]*models.Post, error) {
	rows, err := r.DB.Query(`SELECT id, user_id, title, content, comments_allowed FROM post`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CommentsAllowed); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (r *PostRepositoryImpl) GetByID(id int) (*models.Post, error) {
	var post models.Post
	query := `SELECT id, user_id, title, content, comments_allowed FROM post WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CommentsAllowed)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepositoryImpl) Create(postData *models.PostCreateData) (*models.Post, error) {
	var post models.Post
	query := `INSERT INTO post (user_id, title, content, comments_allowed) VALUES ($1, $2, $3, $4) RETURNING id, user_id, title, content, comments_allowed`
	err := r.DB.QueryRow(query, postData.UserID, postData.Title, postData.Content, postData.CommentsAllowed).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CommentsAllowed)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
