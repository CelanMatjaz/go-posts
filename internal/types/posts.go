package types

import (
	"time"

	"github.com/google/uuid"
)

type dates struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NewPost struct {
	UserId  uuid.UUID
	Content string
}

type Post struct {
	Id uuid.UUID
	NewPost
	dates
}

type PostWithUsername struct {
	Post
	Username string
}

type NewComment struct {
	UserId  uuid.UUID
	Content string
	PostId  uuid.UUID
}

type Comment struct {
	Id uuid.UUID
	NewComment
	dates
}

type CommentWithUsername struct {
	Comment
	Username string
}
