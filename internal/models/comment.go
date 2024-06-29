package models

type Comment struct {
	ID              int    `json:"id"`
	PostID          int    `json:"post_id"`
	UserID          int    `json:"user_id"`
	ParentCommentID int    `json:"parent_comment_id"`
	ChildComments   []int  `json:"child_comments"`
	Text            string `json:"text"`
}

type CommentCreateData struct {
	PostID          int    `json:"post_id"`
	ParentCommentID int    `json:"parent_comment_id"`
	Text            string `json:"text"`
	UserID          int    `json:"user_id"`
}
