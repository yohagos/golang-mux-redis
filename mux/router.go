package mux

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	content = "Content-Type"
	cType   = "application/json"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, cType)
	// redis safe
}
