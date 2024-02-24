package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddPostsRoutes(r *mux.Router) {
	pages_router := r.PathPrefix("/posts").Subrouter()

	// Pages
	pages_router.HandleFunc("/create", createPostPage).Methods("GET")

	// Api
	api_router := r.PathPrefix("/api/posts").Subrouter()

	api_router.HandleFunc("", getPosts).Methods("GET")
	api_router.HandleFunc("/{post_id}", getPost).Methods("GET")

	// api_router := api_router.PathPrefix("/post").Subrouter()
	// authRouter.Use(middleware.EnsureAuthenticatedMiddleware)
	api_router.HandleFunc("", createPost).Methods("POST")
	api_router.HandleFunc("/{post_id}", updatePost).Methods("PUT")
	api_router.HandleFunc("/{post_id}", deletePost).Methods("DELETE")
}

type PostsContext struct {
	IsLoggedIn bool
}

func singlePostPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/single_post_page.html"))
	tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated()})
}

func createPostPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/create_post_page.html"))
	tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated()})
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	ctx.WriteJson(services.GetPosts(ctx.ParseQueryInt("offset"), ctx.ParseQueryInt("limit")))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}

    uuid, err := uuid.Parse(ctx.Vars()["post_id"])

    if err != nil {
        log.Fatal(err)
        ctx.WriteBadRequest()
        return
    }

    post, err := services.GetPost(uuid)

    if err != nil {
        ctx.WriteNotFound()
        return
    }

	ctx.WriteJson(post)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	newPost := types.NewPost{
		UserId:  ctx.GetId(),
		Content: r.FormValue("content"),
	}

	if length := len(newPost.Content); length == 0 || length > 1024 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := services.CreatePost(newPost)

	utils.CheckError(err)
	fmt.Println("/posts/", id.String())
	http.Redirect(w, r, "/posts/"+id.String(), http.StatusSeeOther)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
}

func deletePost(w http.ResponseWriter, r *http.Request) {
}
