package routes

import (
	"html/template"
	"net/http"

	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/CelanMatjaz/go-posts/internal/middleware"
	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
	"github.com/gorilla/mux"
)

type AccountContext struct {
	Username   string
	Posts      []types.PostWithUsername
	IsLoggedIn bool
}

func AddAccountRoutes(r *mux.Router) {
	router := r.PathPrefix("/").Subrouter()

	router.Use(middleware.EnsureAuthenticatedMiddleware)

	r.HandleFunc("/account", AccountPage).Methods("GET")
}

func AccountPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"formatDate": utils.FormatDate,
	}).ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/account_page.html", "views/partials/post.html"))
	ctx := &utils.CustomContext{r, w}
	var username string = ctx.GetUsername()
	var posts []types.PostWithUsername = services.GetPosts(0, 0)
	tmpl.Execute(w, AccountContext{Username: username, Posts: posts, IsLoggedIn: ctx.IsAuthenticated()})
}
