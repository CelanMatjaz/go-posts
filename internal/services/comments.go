package services

import (
	"errors"
	"time"

	"github.com/CelanMatjaz/go-posts/database"
	"github.com/CelanMatjaz/go-posts/internal/types"
	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/google/uuid"
)

func GetCommentsOfPost(commentId uuid.UUID, offset int, limit int) []types.CommentWithUsername {
	if limit <= 0 {
		limit = 10
	}

	rows, err := database.DB.Query("SELECT comments.id, users.id, post_id, content, created_at, updated_at, users.username as username FROM comments LEFT JOIN users ON users.id = comments.user_id WHERE post_id = $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset, commentId)

	utils.CheckError(err)

	var comments []types.CommentWithUsername
	for rows.Next() {
		var comment types.CommentWithUsername
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.Username)
		utils.CheckError(err)
		comments = append(comments, comment)
	}

	return comments
}

func GetComment(uuid uuid.UUID) (types.CommentWithUsername, error) {
	row := database.DB.QueryRow("SELECT comments.id, user_id, post_id, content, created_at, updated_at, users.username as username FROM comments LEFT JOIN users ON users.id = comments.user_id WHERE posts.id = $1", uuid)

	var comment types.CommentWithUsername
	err := row.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt, &comment.Username)

	if row.Err() != nil || err != nil {
		return types.CommentWithUsername{}, errors.New("no comment returned")
	} else {
		return comment, nil
	}
}

func CreateComment(newComment types.NewComment) (uuid.UUID, error) {
	res := database.DB.QueryRow("INSERT INTO comments (user_id, post_id, content) VALUES ($1, $2, $3) returning id", newComment.UserId, newComment.PostId, newComment.Content)
	var id uuid.UUID
	err := res.Scan(&id)
	return id, err
}

func UpdateComment(comment types.Comment) error {
	_, err := database.DB.Exec("UPDATE comment SET content = $1, updated_at = $2 WHERE id = $3", comment.Content, time.Now(), comment.Id)
	return err
}

func DeleteComment(commentId uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM comments WHERE id = $1", commentId)
	return err
}
