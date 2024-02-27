package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CelanMatjaz/go-posts/internal/services"
	"github.com/CelanMatjaz/go-posts/internal/types"
	utils "github.com/CelanMatjaz/go-posts/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func AddCommentsRoutes(r *mux.Router) {
	// Api
	api_router := r.PathPrefix("/api/comments").Subrouter()

	api_router.HandleFunc("", getComments).Methods("GET")
	api_router.HandleFunc("/{comment_id}", getPost).Methods("GET")
	api_router.HandleFunc("", createComment).Methods("POST")
	api_router.HandleFunc("/update/{post_id}", updateComment).Methods("POST")
	api_router.HandleFunc("/{post_id}", deleteComment).Methods("DELETE")
}

func getComments(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	ctx.WriteJson(services.GetCommentsOfPost(ctx.ParseQueryAsUUID("post_id") ,ctx.ParseQueryInt("offset"), ctx.ParseQueryInt("limit")))
}

func getComment(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}

	uuid, err := ctx.ParseVarAsUUID("comment_id")

	if err != nil {
		log.Fatal(err)
		ctx.WriteBadRequest()
		return
	}

	comment, err := services.GetComment(uuid)

	if err != nil {
		ctx.WriteNotFound()
		return
	}

	ctx.WriteJson(comment)
}

func createComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("djwaiohdwoiahd")
	ctx := &utils.CustomContext{r, w}

	newComment := types.NewComment{
		UserId:  ctx.GetId(),
		PostId: uuid.MustParse(r.FormValue("post_id")),
		Content: r.FormValue("content"),
	}

	if length := len(newComment.Content); length == 0 || length > 1024 {
		ctx.Redirect("/posts/create")
		return
	}

	_, err := services.CreateComment(newComment)

	if err != nil {
		log.Fatal(err)
		ctx.Redirect("/")
		return
	}

	ctx.Redirect("/posts/" + newComment.PostId.String())
}

func updateComment(w http.ResponseWriter, r *http.Request) {

	/* ctx := &utils.CustomContext{r, w}

	id, err := ctx.ParseVarAsUUID("comment_id")

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
	ctx.Redirect("/posts/" + updatedPost.Id.String()) */
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := &utils.CustomContext{r, w}
	id, _ := ctx.ParseVarAsUUID("comment_id")
	services.DeleteComment(id)
}
