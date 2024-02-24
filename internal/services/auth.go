package services

import (
	"log"
	"os"

	"github.com/CelanMatjaz/go-posts/database"
	"github.com/CelanMatjaz/go-posts/internal/types"
	"github.com/google/uuid"
)

func FindUserById(id uuid.UUID) (types.User, error) {
	var user types.User
    err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}

func FindUserByUsername(username string) (types.User, error) {
	var user types.User
    err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.Id, &user.Username, &user.Password)
	return user, err
}

func CreateUser(user types.User) {
	_, err := database.DB.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", user.Id, user.Username, user.Password)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
