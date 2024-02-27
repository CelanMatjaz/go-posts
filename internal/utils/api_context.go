package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CustomContext struct {
	R *http.Request
	W http.ResponseWriter
}

func (context *CustomContext) ParseQueryInt(key string) int {
	value, _ := strconv.Atoi("0" + context.R.URL.Query().Get(key))
	return value
}

func (context *CustomContext) ParseQueryBool(key string) bool {
	if val := context.R.URL.Query().Get(key); val == "1" || val == "true" {
		return true
	}
	return false
}

func (context *CustomContext) ParseQueryAsUUID(key string) (uuid.UUID) {
	id, _ := uuid.Parse(context.R.URL.Query().Get(key))
	return id
}

func (context *CustomContext) WriteJson(val any) {
	context.W.Header().Add("Content-type", "application/json")
	json.NewEncoder(context.W).Encode(val)
}

func (context *CustomContext) WriteBadRequest() {
	context.W.WriteHeader(http.StatusBadRequest)
}

func (context *CustomContext) WriteNotFound() {
	context.W.WriteHeader(http.StatusNotFound)
}

func (context *CustomContext) Vars() map[string]string {
	return mux.Vars(context.R)
}

func (context *CustomContext) ParseVarAsUUID(key string) (uuid.UUID, error) {
	id, err := uuid.Parse(mux.Vars(context.R)[key])
	return id, err
}

func (context *CustomContext) Redirect(path string) {
	http.Redirect(context.W, context.R, path, http.StatusSeeOther)
}
