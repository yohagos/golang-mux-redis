package main

import (
	"golang-mux-redis/mux"
	"golang-mux-redis/objects"
	"golang-mux-redis/redis"

	"net/http"
)

func main() {
	objects.UserInit()

	redis.RedisInit()

	redis.GetUsers()

	mux.LoadTemplates("dynamicTemplates/*.html")
	mux := mux.NewRouter()
	
	if err := http.ListenAndServe(":8888", mux); err != nil {
		panic(err)
	}
}