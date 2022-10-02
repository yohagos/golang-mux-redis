package mux

import (
	"golang-mux-redis/objects"
	"golang-mux-redis/redis"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	content = "Content-Type"
	cType   = "application/json"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", getUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/addUser", createUserHandler).Methods(http.MethodPost)

	fs := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(http.StripPrefix(".", fs))

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplates(w, "index.html", nil)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, cType)
	users := redis.GetUsers()
	ExecuteTemplates(w, "index.html", struct {
		Users *[]objects.User
	}{
		Users: users,
	})
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(content, cType)
	// redis safe
}