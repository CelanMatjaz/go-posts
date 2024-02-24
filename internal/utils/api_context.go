package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

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