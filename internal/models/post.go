package models

type Post struct {
	ID              int        `json:"id"`
	UserID          int        `json:"id_user"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	CommentsAllowed bool       `json:"comments_allowed"`
	Comments        []*Comment `json:"comments,omitempty"`
}

type PostCreateData struct {
	UserID          int    `json:"id"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	CommentsAllowed bool   `json:"comments_allowed"`
}
