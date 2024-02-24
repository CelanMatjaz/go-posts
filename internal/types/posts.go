package types

import (
	"time"

	"github.com/google/uuid"
)

type NewPost struct {
    UserId uuid.UUID
    Content string
}

type Post struct {
	Id uuid.UUID
    NewPost
    CreatedAt time.Time
    UpdatedAt time.Time
}

type PostWithUsername struct {
    Post
    Username string
}