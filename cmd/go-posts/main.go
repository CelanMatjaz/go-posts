package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CelanMatjaz/go-posts/database"
	"github.com/CelanMatjaz/go-posts/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.SetupEnv()
    database.CreateConnection()
	port := os.Getenv("PORT")

	r := mux.NewRouter()

    routes.AddHomeRoutes(r)
	routes.AddAuthRoutes(r)
	routes.AddAccountRoutes(r)
    routes.AddPostsRoutes(r)

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}



	srv.ListenAndServe()
}
