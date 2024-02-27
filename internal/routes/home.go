package routes

import (
	"html/template"
	"net/http"

	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddHomeRoutes(r *mux.Router) {
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/home", homePage).Methods("GET")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	type HomeContext struct {
		Posts      []types.PostWithUsername
		IsLoggedIn bool
	}

	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.New("layout.html").Funcs(template.FuncMap{
		"formatDate": utils.FormatDate,
		"getUserId": func() uuid.UUID  { return ctx.GetId() },
        "isPostPage": func() bool { return false },
	}).ParseFiles("views/layout/layout.html", "views/partials/navbar.html", "views/home_page.html", "views/partials/post.html"))
	var posts []types.PostWithUsername = services.GetPosts(0, 10)
	tmpl.Execute(w, HomeContext{Posts: posts, IsLoggedIn: ctx.IsAuthenticated()})
}
