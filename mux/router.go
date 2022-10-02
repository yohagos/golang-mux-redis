package mux

import (
	"fmt"
	"golang-mux-redis/objects"
	"golang-mux-redis/redis"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", getUsersHandler).Methods(http.MethodGet)

	router.HandleFunc("/addUser", createUserGETHandler).Methods(http.MethodGet)
	router.HandleFunc("/addUser", createUserPOSTHandler).Methods(http.MethodPost)

	fs := http.FileServer(http.Dir("."))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplates(w, "index.html", nil)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := redis.GetUsers()
	ExecuteTemplates(w, "users.html", struct {
		Users *[]objects.User
	}{
		Users: users,
	})
}

func createUserGETHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteTemplates(w, "createUser.html", nil)
}

func createUserPOSTHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	phone := r.PostForm.Get("phone")
	fmt.Println(name)
	
	user := objects.User{
		Name: name,
		Email: email,
		Phone: phone,
	}
	fmt.Println(user)
	redis.CreateNewUser(user)
	http.Redirect(w, r, "/users", 302)
}