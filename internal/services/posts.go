package services

import (
	"errors"
	"time"

	"github.com/CelanMatjaz/go-posts/database"
	"github.com/CelanMatjaz/go-posts/internal/types"
	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/google/uuid"
)

func GetPosts(offset int, limit int) []types.PostWithUsername {
	if limit <= 0 {
		limit = 10
	}

	rows, err := database.DB.Query("SELECT posts.id, user_id, content, created_at, updated_at, users.username as username FROM posts LEFT JOIN users ON users.id = posts.user_id ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)

	utils.CheckError(err)

	var posts []types.PostWithUsername
	for rows.Next() {
		var post types.PostWithUsername
		err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Username)
		utils.CheckError(err)
		posts = append(posts, post)
	}

	return posts
}

func GetPostsOfUser(offset int, limit int, userId uuid.UUID) []types.PostWithUsername {
	if limit <= 0 {
		limit = 10
	}

	rows, err := database.DB.Query("SELECT posts.id, user_id, content, created_at, updated_at, users.username as username FROM posts LEFT JOIN users ON users.id = posts.user_id WHERE user_id = $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset, userId)

	utils.CheckError(err)

	var posts []types.PostWithUsername
	for rows.Next() {
		var post types.PostWithUsername
		err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Username)
		utils.CheckError(err)
		posts = append(posts, post)
	}

	return posts
}

func GetPost(uuid uuid.UUID) (types.PostWithUsername, error) {
	row := database.DB.QueryRow("SELECT posts.id, user_id, content, created_at, updated_at, users.username as username FROM posts LEFT JOIN users ON users.id = posts.user_id WHERE posts.id = $1", uuid)

	var post types.PostWithUsername
	err := row.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Username)

	if row.Err() != nil || err != nil {
		return types.PostWithUsername{}, errors.New("no post returned")
	} else {
		return post, nil
	}
}

func CreatePost(newPost types.NewPost) (uuid.UUID, error) {
	res := database.DB.QueryRow("INSERT INTO posts (user_id, content) VALUES ($1, $2) returning id", newPost.UserId, newPost.Content)
	var id uuid.UUID
	err := res.Scan(&id)
	return id, err
}

func UpdatePost(post types.Post) error {
	_, err := database.DB.Exec("UPDATE posts SET content = $1, updated_at = $2 WHERE id = $3", post.Content, time.Now(), post.Id)
	return err
}

func DeletePost(postId uuid.UUID) error {
	_, err := database.DB.Exec("DELETE FROM posts WHERE id = $1", postId)
	return err
}
