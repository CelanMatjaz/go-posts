package routes

import (
	"html/template"
	"net/http"

	"github.com/CelanMatjaz/go-posts/internal/middleware"
	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type AccountContext struct {
	Username   string
	Posts      []types.PostWithUsername
	IsLoggedIn bool
	UserId     string
}

func AddAccountRoutes(r *mux.Router) {
	router := r.PathPrefix("/").Subrouter()

	router.Use(middleware.EnsureAuthenticatedMiddleware)

	r.HandleFunc("/account", AccountPage).Methods("GET")
}

func AccountPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"formatDate": utils.FormatDate,
		"getUserId": func() uuid.UUID  { return ctx.GetId() },
        "isPostPage": func() bool { return false },
	}).ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/account_page.html", "views/partials/post.html"))
	var username string = ctx.GetUsername()
	var posts []types.PostWithUsername = services.GetPostsOfUser(0, 0, ctx.GetId())
    err :=tmpl.Execute(w, AccountContext{Username: username, Posts: posts, IsLoggedIn: ctx.IsAuthenticated(), UserId: ctx.GetId().String()})
    utils.CheckError(err)
}
