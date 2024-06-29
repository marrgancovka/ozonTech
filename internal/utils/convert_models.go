package utils

import (
	"ozonTech/graph/model"
	"ozonTech/internal/models"
	"strconv"
)

func ConvertToGraphQLPost(post *models.Post) *model.Post {
	return &model.Post{
		ID:              strconv.Itoa(post.ID),
		Title:           post.Title,
		Content:         post.Content,
		CommentsAllowed: post.CommentsAllowed,
		UserID:          strconv.Itoa(post.UserID),
		Comments:        ConvertToGraphQLComments(post.Comments),
	}
}

func ConvertToGraphQLComment(comment *models.Comment) *model.Comment {
	var parentCommentID *string
	if comment.ParentCommentID != 0 {
		idStr := strconv.Itoa(comment.ParentCommentID)
		parentCommentID = &idStr
	}
	var childComments []*model.Comment
	for _, childID := range comment.ChildComments {
		childComment := &model.Comment{
			ID: strconv.Itoa(childID),
		}
		childComments = append(childComments, childComment)
	}
	return &model.Comment{
		ID:              strconv.Itoa(comment.ID),
		Content:         comment.Text,
		PostID:          strconv.Itoa(comment.PostID),
		UserID:          strconv.Itoa(comment.UserID),
		ParentCommentID: parentCommentID,
		ChildComments:   childComments,
	}
}
func ConvertToGraphQLComments(comments []*models.Comment) []*model.Comment {
	var gqlComments []*model.Comment
	for _, comment := range comments {
		gqlComments = append(gqlComments, ConvertToGraphQLComment(comment))
	}
	return gqlComments
}
