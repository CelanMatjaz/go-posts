package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddPostsRoutes(r *mux.Router) {
	pages_router := r.PathPrefix("/posts").Subrouter()

	// Pages
	pages_router.HandleFunc("/create", createPostPage).Methods("GET")
	pages_router.HandleFunc("/update/{post_id}", updatePostPage).Methods("GET")
	pages_router.HandleFunc("/{post_id}", singlePostPage).Methods("GET")

	// Api
	api_router := r.PathPrefix("/api/posts").Subrouter()

	api_router.HandleFunc("", getPosts).Methods("GET")
	api_router.HandleFunc("/{post_id}", getPost).Methods("GET")

	// api_router := api_router.PathPrefix("/post").Subrouter()
	// authRouter.Use(middleware.EnsureAuthenticatedMiddleware)
	api_router.HandleFunc("", createPost).Methods("POST")
	api_router.HandleFunc("/update/{post_id}", updatePost).Methods("POST")
	api_router.HandleFunc("/{post_id}", deletePost).Methods("DELETE")
}

type PostsContext struct {
	IsLoggedIn bool
	Post       types.PostWithUsername
	Comments   []types.CommentWithUsername
}

func singlePostPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}

	id, err := ctx.ParseVarAsUUID("post_id")

	if err != nil {
		ctx.Redirect("/posts")
		return
	}

	post, err := services.GetPost(id)

	if err != nil {
		tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/post_not_found_page.html"))
		tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated()})
		return
	}

	tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"formatDate": utils.FormatDate,
		"getUserId": func() uuid.UUID { return ctx.GetId() },
        "isPostPage": func() bool { return true },
        "showCommentBox": func() bool { return ctx.ParseQueryBool("showCommentBox") },
	}).ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/partials/post.html", "views/partials/comment.html", "views/single_post_page.html"))
    tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated(), Post: post, Comments: services.GetCommentsOfPost(post.Id, 0, 0)})
}

func createPostPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/create_post_page.html"))
	tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated()})
}

func updatePostPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}

	id, err := ctx.ParseVarAsUUID("post_id")

	if err != nil {
		ctx.Redirect("/posts")
		return
	}

	post, err := services.GetPost(id)

	if err != nil {
		tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/post_not_found_page.html"))
		tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated()})
		return
	}

	tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/update_post_page.html"))
	err = tmpl.Execute(w, PostsContext{IsLoggedIn: ctx.IsAuthenticated(), Post: post})
	utils.CheckError(err)
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
		ctx.WriteBadRequest()
		ctx.Redirect("/posts/create")
		return
	}

	_, err := services.CreatePost(newPost)

	if err != nil {
		ctx.WriteBadRequest()
		ctx.Redirect("/")
		return
	}

	ctx.Redirect("/account")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}

	id, err := ctx.ParseVarAsUUID("post_id")

	if err != nil {
		ctx.Redirect("/posts")
		return
	}

	updatedPost := types.Post{}
	updatedPost.Id = id
	updatedPost.Content = r.FormValue("content")

	if length := len(updatedPost.Content); length == 0 || length > 1024 {
		ctx.Redirect("/posts/update")
		return
	}

	services.UpdatePost(updatedPost)
	ctx.Redirect("/posts/" + updatedPost.Id.String())
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	id, _ := uuid.Parse(ctx.Vars()["post_id"])
	services.DeletePost(id)
}
