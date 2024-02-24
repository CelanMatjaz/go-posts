package utils

import (
	"github.com/CelanMatjaz/go-posts/internal/types"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("----------------")
	store = sessions.NewCookieStore(key)
)

func (context *CustomContext) IsAuthenticated() bool {
	session, _ := store.Get(context.R, "user")
	return session.Values["authenticated"].(bool) == true
}

func (context *CustomContext) GetUsername() string {
	session, _ := store.Get(context.R, "user")
	return session.Values["username"].(string)
}

func (context *CustomContext) GetId() uuid.UUID {
	session, _ := store.Get(context.R, "user")
	id, _ := uuid.Parse(session.Values["id"].(string))
	return id
}

func (context *CustomContext) Login(user types.User) {
	session, _ := store.Get(context.R, "user")
	session.Values["authenticated"] = true
	session.Values["id"] = user.Id.String()
	session.Values["username"] = user.Username
	session.Save(context.R, context.W)
}

func (context *CustomContext) Logout() {
	session, _ := store.Get(context.R, "user")
	session.Values["authenticated"] = false
	session.Save(context.R, context.W)
}
