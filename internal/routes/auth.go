package routes

import (
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
)

type Context struct {
	IsLoggedIn bool
	Path       string
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddAuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", RegisterPage).Methods("GET")
	r.HandleFunc("/login", LoginPage).Methods("GET")

	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/register", RegisterUser).Methods("POST")
	router.HandleFunc("/login", LoginUser).Methods("POST")
	router.HandleFunc("/logout", LogoutUser).Methods("POST")
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/register_page.html", "views/partials/navbar.html", "views/partials/auth_form_fields.html"))
	tmpl.Execute(w, Context{IsLoggedIn: ctx.IsAuthenticated(), Path: r.URL.Path})
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	tmpl := template.Must(template.ParseFiles("views/layout/layout.html", "views/login_page.html", "views/partials/navbar.html", "views/partials/auth_form_fields.html"))
	tmpl.Execute(w, Context{IsLoggedIn: ctx.IsAuthenticated(), Path: r.URL.Path})
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	type UserRegister struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		PasswordRepeat string `json:"password_repeat"`
	}

    user := UserRegister {
        Username: r.FormValue("username"),
        Password: r.FormValue("password"),
        PasswordRepeat: r.FormValue("password_repeat"),
    }
    
	if len(user.Username) == 0 || len(user.Password) == 0 || len(user.PasswordRepeat) == 0 || user.Password != user.PasswordRepeat {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userUUID, _ := uuid.NewV7()

	newUser := types.User{
		Id:       userUUID,
		Username: user.Username,
		Password: HashPassword(user.Password),
	}

	if _, err := services.FindUserByUsername(user.Username); err != nil {
		services.CreateUser(newUser)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

    user := User {
        Username: r.FormValue("username"),
        Password: r.FormValue("password"),
    }

    foundUser, err := services.FindUserByUsername(user.Username)

	if err != nil || !CheckPasswordHash(user.Password, foundUser.Password) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	ctx := &utils.CustomContext{r, w}
	ctx.Login(foundUser)
	http.Redirect(w, r, "/account", http.StatusSeeOther)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	ctx.Logout()
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
