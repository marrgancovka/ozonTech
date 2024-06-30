package postgres

import (
	"database/sql"
	"fmt"
	"ozonTech/graph/model"
	"ozonTech/internal/models"
	"strconv"
	"strings"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepositoryImpl {
	return &CommentRepositoryImpl{DB: db}
}

func (r *CommentRepositoryImpl) GetByPostID(postID int) ([]*model.Comment, error) {
	rows, err := r.DB.Query(`
        SELECT id, post_id, user_id, text, parent_comment_id, child_comments
        FROM comment
        WHERE post_id = $1 AND parent_comment_id = 0
    `, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rootComments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		var childComments []uint8
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Text, &comment.ParentCommentID, &childComments); err != nil {
			return nil, err
		}
		fmt.Println(string(childComments), "---")
		comment.ChildComments = decodeChildComments(string(childComments))
		fmt.Println("111", comment.ChildComments)
		fmt.Println(comment.ChildComments)
		rootComments = append(rootComments, &comment)
	}

	var comments []*model.Comment
	for _, rootComment := range rootComments {
		comments = append(comments, r.convertToModelComment(rootComment))
	}

	return comments, nil
}

func decodeChildComments(data string) []int {
	str := strings.Trim(data, "{}")

	// Разбиваем строку по запятым
	strValues := strings.Split(str, ",")

	// Создаем массив для хранения чисел
	var nums []int

	// Преобразуем строки в числа и добавляем в массив
	for _, strVal := range strValues {
		num, err := strconv.Atoi(strVal)
		if err != nil {
			return nil
		}
		nums = append(nums, num)
	}
	return nums
}

func (r *CommentRepositoryImpl) convertToModelComment(comment *models.Comment) *model.Comment {
	parentID := ""
	if comment.ParentCommentID != 0 {
		parentID = strconv.Itoa(comment.ParentCommentID)
	}
	modelComment := &model.Comment{
		ID:              strconv.Itoa(comment.ID),
		PostID:          strconv.Itoa(comment.PostID),
		UserID:          strconv.Itoa(comment.UserID),
		Content:         comment.Text,
		ParentCommentID: &parentID,
	}

	var childComments []*model.Comment
	for _, childID := range comment.ChildComments {
		childComment, err := r.GetCommentByID(childID)
		if err != nil {
			continue
		}
		childComments = append(childComments, r.convertToModelComment(childComment))
	}
	modelComment.ChildComments = childComments

	return modelComment
}

func (r *CommentRepositoryImpl) GetCommentByID(id int) (*models.Comment, error) {
	var comment models.Comment
	var childComments []byte
	query := `SELECT id, post_id, user_id, text, parent_comment_id, child_comments FROM comment WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Text, &comment.ParentCommentID, &childComments)
	if err != nil {
		return nil, err
	}
	fmt.Println(comment.ChildComments, id)

	comment.ChildComments = decodeChildComments(string(childComments))
	return &comment, nil
}

func (r *CommentRepositoryImpl) Create(commentData *models.CommentCreateData) (*models.Comment, error) {
	var comment models.Comment
	query := `INSERT INTO comment (post_id, user_id, text, parent_comment_id) VALUES ($1, $2, $3, $4) RETURNING id, post_id, user_id, text, parent_comment_id`
	err := r.DB.QueryRow(query, commentData.PostID, commentData.UserID, commentData.Text, commentData.ParentCommentID).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Text, &comment.ParentCommentID)
	if err != nil {
		return nil, err
	}

	if comment.ParentCommentID != 0 {
		updateQuery := `
            UPDATE comment 
            SET child_comments = array_append(COALESCE(child_comments, ARRAY[]::INTEGER[]), $1)
            WHERE id = $2 
            RETURNING child_comments`
		var rawChildComments []byte
		if err = r.DB.QueryRow(updateQuery, comment.ID, comment.ParentCommentID).Scan(&rawChildComments); err != nil {
			return nil, err
		}
	}

	return &comment, nil
}
